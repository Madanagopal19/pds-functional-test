package common

import (
	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	api "github.com/portworx/pds-functional-test/pkg/api"
	"github.com/portworx/pds-functional-test/pkg/common/template"
)

// ControlPlane PDS
type ControlPlane struct {
	controlPlaneUrl string
	components      *api.Components
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

func (cp *ControlPlane) CreateDefaultResourceSettingTemplate(tenantId string, name string, dataServiceName string) error {
	log.Info("Creating Resource setting template %s for the data service %s.", name, dataServiceName)
	temp := template.CreateTemplates(cp.controlPlaneUrl, cp.components)

	_, err := temp.CreateResourceSettingTemplate(tenantId, "2", "1", dataServiceName, "4G", "2G", name, "10G")
	if err != nil {
		log.Errorf("Storage template creation failed with error - %v", err)
		return err
	}
	return nil
}

func (cp *ControlPlane) CreateDefaultAppconfigTemplate(tenantId string, name string, dataServiceName string) error {
	var data string
	log.Infof("Creating Default %s App config Template ", dataServiceName)
	switch dataServiceName {
	case "Cassandra":
		data = `{
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

	case "Kafka":
		data = `{
			"config_items": [
			  {
				"key": "heapSize",
				"value": "400M",
				"deploy_time": false
			  }
			]
		  }`
	case "ZooKeeper":
		data = `{
			"config_items": [
			  {
				"key": "ZOO_4LW_COMMANDS_WHITELIST",
				"value": "*",
				"deploy_time": false
			  }
			]
		  }`
	case "PostgreSQL":
		data = `{
			"config_items": [
			  {
				"key": "PG_DATABASE",
				"value": "pds",
				"deploy_time": false
			  }
			]
		  }`
	case "RabbitMQ":
		data = `{
			"config_items": [
			  {
				"key": "RABBITMQ_FORCE_BOOT",
				"value": "yes",
				"deploy_time": false
			  }
			]
		  }`
	default:
		log.Panicf("Data Service Name is required to Create App Config Templates")
	}

	temp := template.CreateTemplates(cp.controlPlaneUrl, cp.components)
	appTemplate, err := temp.CreateAppconfigTemplates(tenantId, dataServiceName, name, data)
	if err != nil {
		log.Errorf("Storage template creation failed with error - %v", err)
		return err
	}
	log.Infof("App Config Template ID %s", appTemplate.GetId())
	return nil

}

// NewTargetCluster lsajajsklj
func NewControlPlane(url string, components *api.Components) *ControlPlane {
	return &ControlPlane{
		controlPlaneUrl: url,
		components:      components,
	}
}
