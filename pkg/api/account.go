// Package api comprises of all the components and associated CRUD functionality
package api

import (
	"context"
	status "net/http"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	logger "github.com/portworx/pds-functional-test/pkg/logger"
)

type Account struct {
	Context   context.Context
	apiClient *pds.APIClient
}

var log = logger.Log

func (account *Account) GetAccountsList() ([]pds.ModelsAccount, error) {
	client := account.apiClient.AccountsApi
	log.Info("Get list of Accounts.")
	accountsModel, res, err := client.ApiAccountsGet(account.Context).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Error when calling `ApiAccountsGet``: %v\n", err)
		log.Errorf("Full HTTP response: %v\n", res)
		return nil, err
	}
	return accountsModel.GetData(), nil
}

func (account *Account) GetAccount(accountId string) (*pds.ModelsAccount, error) {
	client := account.apiClient.AccountsApi
	log.Info("Get the account detail having UUID: %v", accountId)
	accountModel, res, err := client.ApiAccountsIdGet(account.Context, accountId).Execute()

	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Full HTTP response: %v\n", res)
		log.Errorf("Error when calling `ApiAccountsIdGet``: %v\n", err)
		return nil, err
	}
	return accountModel, nil
}

func (account *Account) GetAccountUsers(accountId string) ([]pds.ModelsUser, error) {
	client := account.apiClient.AccountsApi
	accountInfo, _ := account.GetAccount(accountId)
	log.Info("Get the users belong to the account having name: %v", accountInfo.GetName())
	usersModel, res, err := client.ApiAccountsIdUsersGet(account.Context, accountId).Execute()
	if err != nil && res.StatusCode != status.StatusOK {
		log.Errorf("Full HTTP response: %v\n", res)
		log.Errorf("Error when calling `ApiAccountsIdUsersGet``: %v\n", err)
		return nil, err
	}
	return usersModel.GetData(), nil
}
