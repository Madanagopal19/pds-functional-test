package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	logger "github.com/portworx/pds-functional-test/pkg/logger"
)

const (
	envControlPlaneURL   = "CONTROL_PLANE_URL"
	envTargetKubeconfig  = "TARGET_KUBECONFIG"
	envVaultHost         = "VAULT_HOST"
	envVaultToken        = "VAULT_TOKEN"
	envPDSUserCredential = "PDS_USER_CREDENTIAL"
	envPDSSecretKey      = "PDS_SECRET_KEY"
)

// Environment lhasha
type Environment struct {
	CONTROL_PLANE_URL   string
	TARGET_KUBECONFIG   string
	VAULT_HOST          string
	VAULT_TOKEN         string
	PDS_USER_CREDENTIAL string
	PDS_SECRET_KEY      string
}

type BearerToken struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Expires_in    uint64 `json:"expires_in"`
	Refresh_token string `json:"refresh_token"`
}

var log = logger.Log

// MustHaveEnvVariables ljsas
func MustHaveEnvVariables() Environment {
	return Environment{
		CONTROL_PLANE_URL:   mustGetEnvVariable(envControlPlaneURL),
		TARGET_KUBECONFIG:   mustGetEnvVariable(envTargetKubeconfig),
		VAULT_HOST:          mustGetEnvVariable(envVaultHost),
		VAULT_TOKEN:         mustGetEnvVariable(envVaultToken),
		PDS_USER_CREDENTIAL: mustGetEnvVariable(envPDSUserCredential),
		PDS_SECRET_KEY:      mustGetEnvVariable(envPDSSecretKey),
	}
}

// mustGetEnvVariable jasljla
func mustGetEnvVariable(key string) string {
	value, isExist := os.LookupEnv(key)
	if !isExist {
		log.Errorf("Key: %v doesn't exist", key)

	}
	return value
}

func GetBearerToken(isAdmin bool) string {
	vault := Vault{}
	err := vault.Init()
	if err != nil {
		// return "Error in getting bearer token", err
		log.Fatalln(err)
	}
	pdsSecret := os.Getenv("PDS_SECRET_KEY")
	pdsUserCreds := os.Getenv("PDS_USER_CREDENTIAL")
	log.Info("Getting login credentials")
	log.Info(pdsSecret)
	log.Info(pdsUserCreds)
	secret, err := vault.GetSecret(pdsSecret)
	if err != nil {
		// return "Error in getting bearer token", err
		log.Fatalln(err)
	}
	userCreds, err := vault.GetSecret(pdsUserCreds)
	if err != nil {
		// return "Error in getting bearer token", err
		log.Fatalln(err)
	}
	clientId := fmt.Sprintf("%s", secret.Data["client-id"])
	clientSecret := fmt.Sprintf("%s", secret.Data["client-secret"])
	url := "http://master-staging-api.portworx.dev/api/protocol/openid-connect/token"
	grantType := "password"
	var username, password string
	if isAdmin {
		username = fmt.Sprintf("%s", userCreds.Data["pds-account-admin-email"])
		password = fmt.Sprintf("%s", userCreds.Data["pds-account-admin-password"])
	} else {
		username = fmt.Sprintf("%s", userCreds.Data["pds-account-user-email"])
		password = fmt.Sprintf("%s", userCreds.Data["pds-account-user-password"])
	}

	postBody, _ := json.Marshal(map[string]string{
		"grant_type":    grantType,
		"client_id":     clientId,
		"client_secret": clientSecret,
		"username":      username,
		"password":      password,
	})
	requestBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()
	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	var bearerToken = new(BearerToken)
	err = json.Unmarshal(body, &bearerToken)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}

	return bearerToken.Access_token

}
