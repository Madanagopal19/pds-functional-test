// Package api comprises of all the components and associated CRUD functionality
package api

import (
	"context"
	status "net/http"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	log "github.com/sirupsen/logrus"
)

type DataServiceDeployment struct {
	context   context.Context
	apiClient *pds.APIClient
}

func (ds *DataServiceDeployment) ListDeployments(projectId string) ([]pds.ModelsDeployment, error) {
	dsClient := ds.apiClient.DeploymentsApi

	dsModels, res, err := dsClient.ApiProjectsIdDeploymentsGet(ds.context, projectId).Execute()

	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return dsModels.GetData(), err
}

func (ds *DataServiceDeployment) CreateDeployment(projectId string, deploymentTargetId string, name string, namespaceId string, appConfigId string, imageId string, nodeCount int32, serviceType string, resourceTemplateId string, storageTemplateId string) (*pds.ModelsDeployment, error) {
	dsClient := ds.apiClient.DeploymentsApi
	dns_zone := "portworx.pds-dns-dev.io"
	createRequest := pds.ControllersCreateProjectDeployment{
		// ApplicationConfigurationOverrides:  &appConfigOverride,
		ApplicationConfigurationTemplateId: &appConfigId,
		DeploymentTargetId:                 &deploymentTargetId,
		DnsZone:                            &dns_zone,
		ImageId:                            &imageId,
		// LoadBalancerSourceRanges: lbSourceRange,
		Name:        &name,
		NamespaceId: &namespaceId,
		NodeCount:   &nodeCount,
		// ScheduledBackup:                    &scheduledBackup,
		ResourceSettingsTemplateId: &resourceTemplateId,
		ServiceType:                &serviceType,
		StorageOptionsTemplateId:   &storageTemplateId,
	}
	dsModel, res, err := dsClient.ApiProjectsIdDeploymentsPost(ds.context, projectId).Body(createRequest).Execute()

	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiProjectsIdDeploymentsPost``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return dsModel, err
}
func (ds *DataServiceDeployment) CreateDeploymentWithScehduleBackup(projectId string, deploymentTargetId string, name string, namespaceId string, appConfigId string, imageId string, nodeCount int32, serviceType string, resourceTemplateId string, storageTemplateId string, backupPolicyId string, backupTargetId string) (*pds.ModelsDeployment, error) {
	dsClient := ds.apiClient.DeploymentsApi
	dns_zone := "portworx.pds-dns-dev.io"
	scheduledBackup := pds.ControllersCreateDeploymentScheduledBackup{
		BackupPolicyId: &backupPolicyId,
		BackupTargetId: &backupTargetId,
	}
	createRequest := pds.ControllersCreateProjectDeployment{
		ApplicationConfigurationTemplateId: &appConfigId,
		DeploymentTargetId:                 &deploymentTargetId,
		DnsZone:                            &dns_zone,
		ImageId:                            &imageId,
		Name:                               &name,
		NamespaceId:                        &namespaceId,
		NodeCount:                          &nodeCount,
		ScheduledBackup:                    &scheduledBackup,
		ResourceSettingsTemplateId:         &resourceTemplateId,
		ServiceType:                        &serviceType,
		StorageOptionsTemplateId:           &storageTemplateId,
	}
	dsModel, res, err := dsClient.ApiProjectsIdDeploymentsPost(ds.context, projectId).Body(createRequest).Execute()

	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiProjectsIdDeploymentsPost``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return dsModel, err
}

func (ds *DataServiceDeployment) GetDeployment(deploymentId string) (*pds.ModelsDeployment, error) {
	dsClient := ds.apiClient.DeploymentsApi
	dsModel, res, err := dsClient.ApiDeploymentsIdGet(ds.context, deploymentId).Execute()
	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return dsModel, err
}

func (ds *DataServiceDeployment) GetDeploymentSatus(deploymentId string) (*pds.DeploymentsStatus, error) {
	dsClient := ds.apiClient.DeploymentsApi
	dsModel, res, err := dsClient.ApiDeploymentsIdStatusGet(ds.context, deploymentId).Execute()

	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return dsModel, err
}

func (ds *DataServiceDeployment) GetConnectionDetails(deploymentId string) (pds.DeploymentsConnectionDetails, error) {
	dsClient := ds.apiClient.DeploymentsApi
	dsModel, res, err := dsClient.ApiDeploymentsIdConnectionInfoGet(ds.context, deploymentId).Execute()

	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return dsModel.GetConnectionDetails(), err
}

func (ds *DataServiceDeployment) DeleteDeployment(deploymentId string) (*status.Response, error) {
	dsClient := ds.apiClient.DeploymentsApi
	res, err := dsClient.ApiDeploymentsIdDelete(ds.context, deploymentId).Execute()
	if res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
	}
	return res, err
}
