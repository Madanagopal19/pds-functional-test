package common

import (
	"fmt"
	"os"

	. "github.com/hashicorp/vault/api"
)

type Vault struct {
	host  string
	token string
}

var client *Client

func (v *Vault) Init() error {
	var err error
	log.Info("Initializing vault")
	v.host = os.Getenv("VAULT_HOST")
	v.token = os.Getenv("VAULT_TOKEN")
	if v.host == "" || v.token == "" {
		return fmt.Errorf("either VAULT_HOST or VAULT_TOKEN environment variables are empty")
	}
	config := DefaultConfig()
	client, _ = NewClient(config)
	// Setting the address and token for the client
	err = client.SetAddress(v.host)
	client.SetToken(v.token)
	if err != nil {
		return err
	}
	log.Debug("Looking up token")
	_, err = client.Auth().Token().LookupSelf()
	if err != nil {
		return err
	}
	log.Debug("Token lookup success")
	return nil
}

func (v *Vault) GetSecret(secretName string) (*Secret, error) {
	log.Infof("Getting secret name: %s", secretName)
	var secret *Secret
	secret, err := client.Logical().Read(fmt.Sprintf("secret/%s", secretName))
	if err != nil {
		return secret, fmt.Errorf("error getting secret %s", secretName)
	}
	return secret, nil
}
