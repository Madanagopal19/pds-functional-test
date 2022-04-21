package common

import (
	api "github.com/portworx/pds-functional-test/pkg/api"
)

// ControlPlabe khash
type ControlPlane struct {
	controlPlaneUrl string
	components      *api.Components
}

func (cp *ControlPlane) CreateStorageOptionTemplate(tenantId string, name string) error {
	log.Info("Creating storage option template.")
	st := cp.components.StorageSettingsTemplate
	templates, _ := st.ListTemplates(tenantId)
	isExists := false
	for _, template := range templates {
		if template.GetName() == name {
			isExists = true
		}
	}
	if !isExists {
		_, err := st.CreateTemplate(tenantId, false, "ext4", name, 2, false)
		if err != nil {
			log.Errorf("Storage template creation failed with error - %v", err)
			return err
		}
	}
	return nil

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
func NewControlPlane(context string) *ControlPlane {
	return &ControlPlane{
		controlPlaneUrl: context,
	}
}
