package test

import (
	"math/rand"
	"time"

	. "github.com/portworx/pds-functional-test/pkg/common"
)

func (suite *PDSTestSuite) TestDeployDataServices() {

	var (
		deploymentTargetId, storageTemplateId string
		deploymentTargetComponent             = suite.components.DeploymentTarget
		nsComponent                           = suite.components.Namespace
		storagetemplateComponent              = suite.components.StorageSettingsTemplate
		resourceTemplateComponent             = suite.components.ResourceSettingsTemplate
		dataServiceComponent                  = suite.components.DataService
		versionComponent                      = suite.components.Version
		imageComponent                        = suite.components.Image
		appConfigComponent                    = suite.components.AppConfigTemplate
	)

	clusterId, err := GetClusterId(suite.env.TARGET_KUBECONFIG)
	if err != nil {
		log.Panicf("Unable to fetch the cluster Id")
	}

	log.Info("Get the Target cluster details")
	targetClusters, _ := deploymentTargetComponent.ListDeploymentTargetsBelongsToTenant(tenantId)
	for i := 0; i < len(targetClusters); i++ {
		if targetClusters[i].GetClusterId() == clusterId {
			deploymentTargetId = targetClusters[i].GetId()
			log.Infof("Cluster ID: %v, Name: %v,Status: %v", targetClusters[i].GetClusterId(), targetClusters[i].GetName(), targetClusters[i].GetStatus())
		}
	}

	log.Infof("Get the available namespaces in the Cluster having Id: %v", clusterId)
	namespaces, _ := nsComponent.ListNamespaces(deploymentTargetId)
	for i := 0; i < len(namespaces); i++ {
		if namespaces[i].GetStatus() == "available" {
			namespaceNameIdMap[namespaces[i].GetName()] = namespaces[i].GetId()
			log.Infof("Available namespace - Name: %v , Id: %v , Status: %v", namespaces[i].GetName(), namespaces[i].GetId(), namespaces[i].GetStatus())
		}
	}
	log.Infof("Get the storage template")
	storageTemplates, _ := storagetemplateComponent.ListTemplates(tenantId)
	for i := 0; i < len(storageTemplates); i++ {
		if storageTemplates[i].GetName() == storageTemplateName {
			log.Infof("Storage template details -----> Name %v,Repl %v , Fg %v , Fs %v",
				storageTemplates[i].GetName(),
				storageTemplates[i].GetRepl(),
				storageTemplates[i].GetFg(),
				storageTemplates[i].GetFs())
			storageTemplateId = storageTemplates[i].GetId()
			log.Infof("Storage Id: %v", storageTemplateId)
		}
	}

	log.Infof("Get the resource template for each data services")
	resourceTemplates, _ := resourceTemplateComponent.ListTemplates(tenantId)
	for i := 0; i < len(resourceTemplates); i++ {
		if resourceTemplates[i].GetName() == resourceTemplateName {
			dataService, _ := dataServiceComponent.GetDataService(resourceTemplates[i].GetDataServiceId())
			log.Infof("Data service name: %v", dataService.GetName())
			log.Infof("Resource template details ---> Name %v, Id : %v ,DataServiceId %v , StorageReq %v , Memoryrequest %v",
				resourceTemplates[i].GetName(),
				resourceTemplates[i].GetId(),
				resourceTemplates[i].GetDataServiceId(),
				resourceTemplates[i].GetStorageRequest(),
				resourceTemplates[i].GetMemoryRequest())
			dataServiceDefaultResourceTemplateIdMap[dataService.GetName()] =
				resourceTemplates[i].GetId()
			dataServiceNameIdMap[dataService.GetName()] = dataService.GetId()
		}
	}

	log.Infof("Get the Versions.")
	for key := range dataServiceNameIdMap {
		versions, _ := versionComponent.ListDataServiceVersions(dataServiceNameIdMap[key])
		for i := 0; i < len(versions); i++ {
			dataServiceNameVersionMap[key] = append(dataServiceNameVersionMap[key], versions[i].GetId())
		}
	}

	log.Infof("Get the Versions.")
	for key := range dataServiceNameVersionMap {
		images, _ := imageComponent.ListImages(dataServiceNameVersionMap[key][0])
		for i := 0; i < len(images); i++ {
			dataServiceNameImagesMap[key] = images[i].GetId()
		}
	}

	log.Infof("Get the Application configuration template")
	appConfigs, _ := appConfigComponent.ListTemplates(tenantId)
	for i := 0; i < len(appConfigs); i++ {
		if appConfigs[i].GetName() == appConfigTemplateName {
			for key := range dataServiceNameIdMap {
				if dataServiceNameIdMap[key] == appConfigs[i].GetDataServiceId() {
					dataServiceNameDefaultAppConfigMap[key] = appConfigs[i].GetId()
				}
			}
		}

	}

	log.Info("Create dataservices")
	log.Info("Create dataservice with no scheduled backup enabled.")
	for i := range supportedDataServices {
		log.Infof("Key: %v, Value %v", supportedDataServices[i], dataServiceNameDefaultAppConfigMap[supportedDataServices[i]])
		n := rand.Int() % len(pdsNamespaces)
		namespace := pdsNamespaces[n]
		namespaceId := namespaceNameIdMap[namespace]
		log.Infof("Created %v deployment  in the namespace %v with no scheduled back up.", supportedDataServices[i], namespace)
		deployment, _ :=
			suite.components.DataServiceDeployment.CreateDeployment(projectId,
				deploymentTargetId,
				dnsZone,
				deploymentName,
				namespaceId,
				dataServiceNameDefaultAppConfigMap[supportedDataServices[i]],
				dataServiceNameImagesMap[supportedDataServices[i]],
				3,
				serviceType,
				dataServiceDefaultResourceTemplateIdMap[supportedDataServices[i]],
				storageTemplateId,
			)
		deployementIdNameMap[deployment.GetId()] = deployment.GetName()
		status, _ := suite.components.DataServiceDeployment.GetDeploymentSatus(deployment.GetId())
		sleeptime := 0
		for status.GetHealth() != "Healthy" && sleeptime < duration {
			time.Sleep(10 * time.Second)
			sleeptime += 10
			status, _ = suite.components.DataServiceDeployment.GetDeploymentSatus(deployment.GetId())
			log.Infof("Health status -  %v", status.GetHealth())
		}
		log.Infof("Deployment details ---> Id:%v ,Health status -  %v,Replicas - %v, Ready replicas - %v", deployment.GetId(), status.GetHealth(), status.GetReplicas(), status.GetReadyReplicas())

	}

}
