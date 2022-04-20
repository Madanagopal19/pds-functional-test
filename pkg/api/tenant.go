package api

import (
	"context"
	status "net/http"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
)

type Tenant struct {
	context   context.Context
	apiClient *pds.APIClient
}

func (tenant *Tenant) GetTenantsList(accountId string) ([]pds.ModelsTenant, error) {
	tenantClient := tenant.apiClient.TenantsApi
	log.Info("Get list of Accounts.")
	tenantsModel, res, err := tenantClient.ApiAccountsIdTenantsGet(tenant.context, accountId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsIdTenantsGet``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return tenantsModel.GetData(), nil
}

func (tenant *Tenant) GetTenant(tenantId string) (*pds.ModelsTenant, error) {
	tenantClient := tenant.apiClient.TenantsApi
	log.Info("Get list of Accounts.")
	tenantModel, res, err := tenantClient.ApiTenantsIdGet(tenant.context, tenantId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiTenantsIdGet``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return tenantModel, nil
}
