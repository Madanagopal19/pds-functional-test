package test

import (
	"os"
	"time"
)

func (suite *PDSTestSuite) TestKafkaDeployment() {
	var (
		versionComponent = suite.components.Version
		imageComponent   = suite.components.Image
	)

	clusterId, err := suite.getClusterID()
	log.Infof(clusterId)
	if err != nil {
		log.Panicf("Unable to fetch the cluster Id")
	}

	deploymentTargetId := suite.getTargetClusters(clusterId)
	log.Infoln("targetId", deploymentTargetId)

	namespaces := suite.listNamespaces(deploymentTargetId)
	for i := 0; i < len(namespaces); i++ {
		if namespaces[i].GetStatus() == "available" {
			namespaceNameIdMap[namespaces[i].GetName()] = namespaces[i].GetId()
			log.Infof("Available namespace - Name: %v , Id: %v , Status: %v", namespaces[i].GetName(), namespaces[i].GetId(), namespaces[i].GetStatus())
		}
	}

	storageTemplateId := suite.getStorageTemplateId(tenantId)
	log.Info(storageTemplateId)

	resourceTemplate, dataServiceNameIdMap, dataServiceDefaultResourceTemplateIdMap := suite.getResourceTemplate(tenantId)

	log.Info(resourceTemplate)
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
	appConfigs := suite.getAppConfigTemplate(tenantId)

	for i := 0; i < len(appConfigs); i++ {
		if appConfigs[i].GetName() == appConfigTemplateName {
			for key := range dataServiceNameIdMap {
				if dataServiceNameIdMap[key] == appConfigs[i].GetDataServiceId() {
					dataServiceNameDefaultAppConfigMap[key] = appConfigs[i].GetId()
				}
			}
		}

	}

	log.Info("Create kafka deployment dataservices")
	namespace := pdsNamespaces[0]
	namespaceId := namespaceNameIdMap[namespace]
	var applicationConfigurationOverrides = make(map[string]string)
	applicationConfigurationOverrides["ZOOKEEPER_CONNECTION_STRING"] = os.Getenv("ZOOKEEPER_CONNECTION_STRING")
	applicationConfigurationOverrides["ZOOKEEPER_PASSWORD"] = os.Getenv("ZOOKEEPER_PASSWORD")

	log.Infof("Appconfig ID %v", dataServiceNameDefaultAppConfigMap[suite.env.DATA_SERVICE])
	log.Infof("Image ID %v", dataServiceNameImagesMap[suite.env.DATA_SERVICE])
	log.Infof("Resource template ID %v", dataServiceDefaultResourceTemplateIdMap[suite.env.DATA_SERVICE])

	deployment, _ :=
		suite.components.DataServiceDeployment.CreateKafkaDeployment(
			applicationConfigurationOverrides,
			projectId,
			deploymentTargetId,
			dnsZone,
			deploymentName,
			namespaceId,
			dataServiceNameDefaultAppConfigMap[suite.env.DATA_SERVICE],
			dataServiceNameImagesMap[suite.env.DATA_SERVICE],
			3,
			serviceType,
			dataServiceDefaultResourceTemplateIdMap[suite.env.DATA_SERVICE],
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
