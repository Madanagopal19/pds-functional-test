package common

import (
	"encoding/json"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	api "github.com/portworx/pds-functional-test/pkg/api"
)

// ControlPlane PDS
type ControlPlane struct {
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

func (cp *ControlPlane) GetRegistrationToken(tenantId string) string {
	log.Info("Fetch the registration token.")

	saClient := cp.components.ServiceAccount
	serviceAccounts, _ := saClient.ListServiceAccounts(tenantId)
	var agentWriterId string
	for _, sa := range serviceAccounts {
		if sa.GetName() == "Default-AgentWriter" {
			agentWriterId = sa.GetId()
		}
	}
	token, _ := saClient.GetServiceAccountToken(agentWriterId)
	return token.GetToken()
}

func (cp *ControlPlane) CreateDeafaultStorageTemplate(tenantId string, name string) (*pds.ModelsStorageOptionsTemplate, error) {
	log.Info("Creating storage option template.")
	st := cp.components.StorageSettingsTemplate
	templates, _ := st.ListTemplates(tenantId)
	for _, template := range templates {
		if template.GetName() == name {
			templateId := template.GetId()
			log.Infof("Template Name %s,  Template ID %s", name, templateId)
			return st.GetTemplate(templateId)
		}
	}
	template, err := st.CreateTemplate(tenantId, false, "ext4", name, 2, false)
	if err != nil {
		log.Errorf("Storage template creation failed with error - %v", err)

	}
	return template, err
}

func (cp *ControlPlane) GetDataserviceId(dataServiceName string) string {
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
func (cp *ControlPlane) CreateResourceSettingTemplate(tenantId string, cpuLimit string, cpuRequest string, dataServiceId string, memoryLimit string, memoryRequest string, templateName string, storageRequest string) (*pds.ModelsResourceSettingsTemplate, error) {
	rt := cp.components.ResourceSettingsTemplate
	templates, _ := rt.ListTemplates(tenantId)
	for _, template := range templates {
		if (template.GetName() == templateName) && (template.GetDataServiceId() == dataServiceId) {
			templateId := template.GetId()
			log.Infof("Template Name %s,  Template ID %s", templateName, templateId)
			return rt.GetTemplate(templateId)
		}
	}
	rsTemplate, err := rt.CreateTemplate(tenantId, cpuLimit, cpuRequest, dataServiceId, memoryLimit, memoryRequest, templateName, storageRequest)
	if err != nil {
		log.Errorf("resource template creation failed with error - %v", err)
	}
	return rsTemplate, err
}

//CreateAppconfigTemplates unmarshal the json data to the struct and creates app config templates
func (cp *ControlPlane) CreateAppconfigTemplates(tenantId string, dataServiceName string, templateName string, data string) (*pds.ModelsApplicationConfigurationTemplate, error) {
	ap := cp.components.AppConfigTemplate
	dataServiceID := cp.GetDataserviceId(dataServiceName)
	log.Infof("DataserviceID %s of %s ", dataServiceID, dataServiceName)
	templates, _ := ap.ListTemplates(tenantId)
	for _, template := range templates {
		if (template.GetName() == templateName) && (template.GetDataServiceId() == dataServiceID) {
			templateId := template.GetId()
			log.Infof("Template Name %s,  Template ID %s", templateName, templateId)
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

	apTemplate, err := ap.CreateTemplate(tenantId, dataServiceID, templateName, pdsData)
	if err != nil {
		log.Errorf("App config template creation failed with error - %v", err)
	}
	return apTemplate, err
}

//CreateDefaultResourceSettingTemplate func Creates Resource setting templates with Default values for available dataservices
func (cp *ControlPlane) CreateDefaultResourceSettingTemplate(tenantId string, templateName string) error {
	dataService := []string{"PostgreSQL", "ZooKeeper", "Kafka", "RabbitMQ", "Cassandra"}
	for _, services := range dataService {
		dataServiceID := cp.GetDataserviceId(services)
		log.Infof("Creating Resource setting template %s for the data service %s.", templateName, services)
		_, err := cp.CreateResourceSettingTemplate(tenantId, "2", "1", dataServiceID, "4G", "2G", templateName, "50G")
		if err != nil {
			log.Errorf("Storage template creation failed with error - %v", err)
			return err
		}
	}
	return nil
}

//CreateDefaultAppconfigTemplate func Creates AppConfig templates with Default values for available dataservices
func (cp *ControlPlane) CreateDefaultAppconfigTemplate(tenantId string, templateName string) error {
	var appTemplate *pds.ModelsApplicationConfigurationTemplate
	var err error
	cassConfdata := `{
		"config_items": [
		  {
			"key": "CASSANDRA_AUTHORIZER",
			"value": "AllowAllAuthorizer",
			"deploy_time": false
		  },
		  {
			"key": "CASSANDRA_AUTHENTICATOR",
			"value": "AllowAllAuthenticator",
			"deploy_time": false
		  },
		  {
			"key": "HEAP_NEWSIZE",
			"value": "400M",
			"deploy_time": false
		  },
		  {
			"key": "MAX_HEAP_SIZE",
			"value": "1G",
			"deploy_time": false
		  },
		  {
			"key": "CASSANDRA_RACK",
			"value": "rack1",
			"deploy_time": false
		  },
		  {
			"key": "CASSANDRA_DC",
			"value": "dc1",
			"deploy_time": false
		  }
		]
	  }`
	kafkaConfData := `{
		"config_items": [
		  {
			"key": "heapSize",
			"value": "400M",
			"deploy_time": false
		  }
		]
	  }`
	zkConfData := `{
		"config_items": [
		  {
			"key": "ZOO_4LW_COMMANDS_WHITELIST",
			"value": "*",
			"deploy_time": false
		  }
		]
	  }`
	psqlConfData := `{
		"config_items": [
		  {
			"key": "PG_DATABASE",
			"value": "pds",
			"deploy_time": false
		  }
		]
	  }`
	rmqConfData := `{
		"config_items": [
		  {
			"key": "RABBITMQ_FORCE_BOOT",
			"value": "yes",
			"deploy_time": false
		  }
		]
	  }`

	appConfigData := map[string]string{"PostgreSQL": psqlConfData, "ZooKeeper": zkConfData, "Kafka": kafkaConfData, "RabbitMQ": rmqConfData, "Cassandra": cassConfdata}

	for dataService, appConfData := range appConfigData {
		log.Infof("Creating Default Template for Dataservice %s", dataService)
		appTemplate, err = cp.CreateAppconfigTemplates(tenantId, dataService, templateName, appConfData)
		if err != nil {
			log.Errorf("App Config template creation failed with error - %v", err)
			return err
		}
		log.Infof("App Config Template ID %s", appTemplate.GetId())
	}
	return nil
}

// NewTargetCluster lsajajsklj
func NewControlPlane(url string, components *api.Components) *ControlPlane {
	return &ControlPlane{
		controlPlaneUrl: url,
		components:      components,
	}
}
