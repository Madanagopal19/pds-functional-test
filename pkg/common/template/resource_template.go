package template

import (
	"encoding/json"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	"github.com/portworx/pds-functional-test/pkg/api"
	"github.com/portworx/pds-functional-test/pkg/logger"
)

// ControlPlane PDS
type NewTemplates struct {
	controlPlaneUrl string
	components      *api.Components
}

type ConfigData struct {
	ConfigItems []struct {
		Key        string `json:"key"`
		Value      string `json:"value"`
		DeployTime bool   `json:"deploy_time"`
	} `json:"config_items"`
}

var log = logger.Log

func (cp *NewTemplates) GetDataserviceId(dataServiceName string) string {
	log.Infof("Getting the data service ID for %s ", dataServiceName)
	dsComp := cp.components.DataService
	dataServices, err := dsComp.ListDataServices()
	if err != nil {
		log.Panicf("Unable to list data services: %v", err)
	}
	var dataServiceId string
	for _, ds := range dataServices {
		if ds.GetName() == dataServiceName {
			dataServiceId = ds.GetId()
		}
	}
	return dataServiceId
}

//CreateResourceSettingTemplate creates resource setting templates based on passed dataservice ID
func (cp *NewTemplates) CreateResourceSettingTemplate(tenantId string, cpuLimit string, cpuRequest string, dataServiceName string, memoryLimit string, memoryRequest string, name string, storageRequest string) (*pds.ModelsResourceSettingsTemplate, error) {
	rt := cp.components.ResourceSettingsTemplate
	dataServiceId := cp.GetDataserviceId(dataServiceName)
	log.Infof("DataserviceID %s of %s ", dataServiceId, dataServiceName)
	templates, _ := rt.ListTemplates(tenantId)
	for _, template := range templates {
		if (template.GetName() == name) && (template.GetDataServiceId() == dataServiceId) {
			templateId := template.GetId()
			log.Infof("Template Name %s,  Template ID %s", name, templateId)
			return rt.GetTemplate(templateId)
		}
	}
	rsTemplate, err := rt.CreateTemplate(tenantId, cpuLimit, cpuRequest, dataServiceId, memoryLimit, memoryRequest, name, storageRequest)
	if err != nil {
		log.Errorf("resource template creation failed with error - %v", err)
	}
	return rsTemplate, err
}

//CreateAppconfigTemplates unmarshal the json data to the struct and creates app config templates
func (cp *NewTemplates) CreateAppconfigTemplates(tenantId string, dataServiceName string, name string, data string) (*pds.ModelsApplicationConfigurationTemplate, error) {
	ap := cp.components.AppConfigTemplate
	dataServiceID := cp.GetDataserviceId(dataServiceName)
	log.Infof("DataserviceID %s of %s ", dataServiceID, dataServiceName)
	templates, _ := ap.ListTemplates(tenantId)
	for _, template := range templates {
		if (template.GetName() == name) && (template.GetDataServiceId() == dataServiceID) {
			templateId := template.GetId()
			log.Infof("Template Name %s,  Template ID %s", name, templateId)
			return ap.GetTemplate(templateId)
		}
	}

	var myConfigdata ConfigData
	err := json.Unmarshal([]byte(data), &myConfigdata)
	if err != nil {
		log.Panicf("Unable to unmarshal json data: %v", err)
	}

	var pdsData []pds.ModelsConfigItem
	var testDefaultData pds.ModelsConfigItem

	for index, _ := range myConfigdata.ConfigItems {
		testDefaultData.Key = &myConfigdata.ConfigItems[index].Key
		testDefaultData.Value = &myConfigdata.ConfigItems[index].Value
		testDefaultData.DeployTime = &myConfigdata.ConfigItems[index].DeployTime

		pdsData = append(pdsData, testDefaultData)
	}

	apTemplate, err := ap.CreateTemplate(tenantId, dataServiceID, name, pdsData)
	if err != nil {
		log.Errorf("App config template creation failed with error - %v", err)
	}
	return apTemplate, err

}

// NewTargetCluster lsajajsklj
func CreateTemplates(url string, components *api.Components) *NewTemplates {
	return &NewTemplates{
		controlPlaneUrl: url,
		components:      components,
	}
}
