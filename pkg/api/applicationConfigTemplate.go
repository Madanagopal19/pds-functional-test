// Package api comprises of all the components and associated CRUD functionality
package api

import (
	"context"
	status "net/http"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
)

type AppConfigTemplate struct {
	context   context.Context
	apiClient *pds.APIClient
}

func (at *AppConfigTemplate) ListTemplates(tenantId string) ([]pds.ModelsApplicationConfigurationTemplate, error) {
	atClient := at.apiClient.ApplicationConfigurationTemplatesApi
	atModel, res, err := atClient.ApiTenantsIdApplicationConfigurationTemplatesGet(at.context, tenantId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiTenantsIdApplicationConfigurationTemplatesGet``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return atModel.GetData(), nil
}

func (at *AppConfigTemplate) GetTemplate(templateId string) (*pds.ModelsApplicationConfigurationTemplate, error) {
	atClient := at.apiClient.ApplicationConfigurationTemplatesApi
	log.Infof("Get list of storage templates for tenant ID - %v", templateId)
	atModel, res, err := atClient.ApiApplicationConfigurationTemplatesIdGet(at.context, templateId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiApplicationConfigurationTemplatesIdGet``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return atModel, nil
}

func (at *AppConfigTemplate) CreateTemplate(tenantId string, deployTime bool, key string, value string, dataServiceId string, name string) (*pds.ModelsApplicationConfigurationTemplate, error) {
	atClient := at.apiClient.ApplicationConfigurationTemplatesApi
	log.Info("Create new resource template.")
	data := []pds.ModelsConfigItem{pds.ModelsConfigItem{&deployTime, &key, &value}}
	createRequest := pds.ControllersCreateApplicationConfigurationTemplatesRequest{data, &dataServiceId, &name}
	atModel, res, err := atClient.ApiTenantsIdApplicationConfigurationTemplatesPost(at.context, tenantId).Body(createRequest).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiTenantsIdApplicationConfigurationTemplatesPost``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return atModel, nil
}

func (at *AppConfigTemplate) UpdateTemplate(templateId string, deployTime bool, key string, value string, name string) (*pds.ModelsApplicationConfigurationTemplate, error) {
	atClient := at.apiClient.ApplicationConfigurationTemplatesApi
	log.Info("Create new resource template.")
	data := []pds.ModelsConfigItem{pds.ModelsConfigItem{&deployTime, &key, &value}}
	updateRequest := pds.ControllersUpdateApplicationConfigurationTemplateRequest{data, &name}
	atModel, res, err := atClient.ApiApplicationConfigurationTemplatesIdPut(at.context, templateId).Body(updateRequest).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiApplicationConfigurationTemplatesIdPut``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return atModel, nil
}

func (at *AppConfigTemplate) DeleteTemplate(templateId string) (*status.Response, error) {
	atClient := at.apiClient.ApplicationConfigurationTemplatesApi
	log.Infof("Get list of storage templates for tenant ID - %v", templateId)
	res, err := atClient.ApiApplicationConfigurationTemplatesIdDelete(at.context, templateId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiApplicationConfigurationTemplatesIdDelete``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return res, nil
}
