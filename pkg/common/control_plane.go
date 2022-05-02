package common

import (
	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	api "github.com/portworx/pds-functional-test/pkg/api"
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

func (cp *ControlPlane) CreateStorageOptionTemplate(tenantId string, fg bool, fs string, name string, repl int32, secure bool) (*pds.ModelsStorageOptionsTemplate, error) {
	log.Info("Creating storage option template.")
	st := cp.components.StorageSettingsTemplate
	template, err := st.CreateTemplate(tenantId, fg, fs, name, repl, secure)
	if err != nil {
		log.Errorf("Storage template creation failed with error - %v", err)
	}
	return template, nil

}
func (cp *ControlPlane) CreateResourceSettingTemplate(tenantId string, name string, dataServiceName string) error {
	log.Info("Creating Resource setting template %s for the data service %s.", name, dataServiceName)
	dsComp := cp.components.DataService
	dataServices, _ := dsComp.ListDataServices()
	var dataServiceId string
	for _, ds := range dataServices {
		if ds.GetName() == dataServiceName {
			dataServiceId = ds.GetId()
		}

	}
	rtComp := cp.components.ResourceSettingsTemplate
	templates, _ := rtComp.ListTemplates(tenantId)
	isExists := false
	for _, template := range templates {
		if template.GetName() == name {
			isExists = true
		}
	}
	if !isExists {
		_, err := rtComp.CreateTemplate(tenantId, "2", "1", dataServiceId, "4G", "2G", name, "10G")
		if err != nil {
			log.Errorf("Storage template creation failed with error - %v", err)
			return err
		}
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
