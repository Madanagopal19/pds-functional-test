package api

import (
	"context"
	status "net/http"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	log "github.com/sirupsen/logrus"
)

type StorageSettingsTemplate struct {
	context   context.Context
	apiClient *pds.APIClient
}

func (st *StorageSettingsTemplate) ListTemplates(tenantId string) ([]pds.ModelsStorageOptionsTemplate, error) {
	stClient := st.apiClient.StorageOptionsTemplatesApi
	log.Infof("Get list of storage templates for tenant ID - %v", tenantId)
	pdsStorageTemplates, res, err := stClient.ApiTenantsIdStorageOptionsTemplatesGet(st.context, tenantId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
		return nil, err
	}
	return pdsStorageTemplates.GetData(), nil
}

func (st *StorageSettingsTemplate) GetTemplate(templateId string) (*pds.ModelsStorageOptionsTemplate, error) {
	stClient := st.apiClient.StorageOptionsTemplatesApi
	log.Infof("Get storage template details for UUID - %v", templateId)
	stModel, res, err := stClient.ApiStorageOptionsTemplatesIdGet(st.context, templateId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
		return nil, err
	}
	return stModel, nil
}

func (st *StorageSettingsTemplate) CreateTemplate(tenantId string, fg bool, fs string, name string, repl int32, secure bool) (*pds.ModelsStorageOptionsTemplate, error) {
	stClient := st.apiClient.StorageOptionsTemplatesApi
	log.Info("Create new storage template.")
	createRequest := pds.ControllersCreateStorageOptionsTemplatesRequest{&fg, &fs, &name, &repl, &secure}
	stModel, res, err := stClient.ApiTenantsIdStorageOptionsTemplatesPost(st.context, tenantId).Body(createRequest).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
		return nil, err
	}
	return stModel, nil
}

func (st *StorageSettingsTemplate) UpdateTemplate(templateId string, fg bool, fs string, name string, repl int32, secure bool) (*pds.ModelsStorageOptionsTemplate, error) {
	stClient := st.apiClient.StorageOptionsTemplatesApi
	log.Info("Create new storage template.")
	updateRequest := pds.ControllersUpdateStorageOptionsTemplateRequest{&fg, &fs, &name, &repl, &secure}
	stModel, res, err := stClient.ApiStorageOptionsTemplatesIdPut(st.context, templateId).Body(updateRequest).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
		return nil, err
	}
	return stModel, nil
}

func (st *StorageSettingsTemplate) DeleteTemplate(templateId string) (*status.Response, error) {
	stClient := st.apiClient.StorageOptionsTemplatesApi
	log.Infof("Delete strogae template: %v", templateId)
	res, err := stClient.ApiStorageOptionsTemplatesIdDelete(st.context, templateId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Error("Full HTTP response: %v\n", res)
		return nil, err
	}
	return res, nil
}
