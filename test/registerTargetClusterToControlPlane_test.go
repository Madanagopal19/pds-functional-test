package test

import (
	"math/rand"
	"reflect"

	. "github.com/portworx/pds-functional-test/pkg/common"
)

const accountName = "Portworx"

func (s *PDSTestSuite) TestRegisterTargetToControlPlane() {
	var accountId string
	acc := s.components.Account
	accounts, _ := acc.GetAccountsList()

	for i := 0; i < len(accounts); i++ {
		log.Infof("Account Name: %v", accounts[i].GetName())
		if accounts[i].GetName() == accountName {
			accountId = accounts[i].GetId()
		}
	}
	log.Infof("Account Detail- Name: %s, UUID: %s ", accountName, accountId)
	tnts := s.components.Tenant
	tenants, _ := tnts.GetTenantsList(accountId)
	tenantId := tenants[0].GetId()
	tenantName := tenants[0].GetName()
	log.Infof("Tenant Details- Name: %s, UUID: %s ", tenantName, tenantId)
	projcts := s.components.Project
	projects, _ := projcts.GetprojectsList(tenantId)
	projectId := projects[0].GetId()
	projectName := projects[0].GetName()
	log.Infof("Project Details- Name: %s, UUID: %s ", projectName, projectId)

	log.Info("Get helm version")
	version, _ := s.components.ApiVersion.GetHelmChartVersion()
	log.Infof("Helm chart Version : %s ", version)

	log.Infof("Register Target cluster to control plane %v ", s.env.CONTROL_PLANE_URL)
	err := s.TargetCluster.RegisterToControlPlane(s.env.CONTROL_PLANE_URL, version, GetBearerToken(true), tenantId)
	if err != nil {
		log.Panicf("Unable to register the target cluster to control plane %v", s.env.CONTROL_PLANE_URL)
	}

}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()

	return keys[rand.Intn(len(keys))].Interface()
}
