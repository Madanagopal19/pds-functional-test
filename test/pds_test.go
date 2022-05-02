package test

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	api "github.com/portworx/pds-functional-test/pkg/api"
	. "github.com/portworx/pds-functional-test/pkg/common"
	logger "github.com/portworx/pds-functional-test/pkg/logger"
	"github.com/stretchr/testify/suite"
)

var (
	accountId, tenantId, projectId string

	log                      = logger.Log
	accountName              = "Portworx"
	S3BackupTarget           = "pds-qa-s3-target"
	S3CompatibleBackupTarget = "pds-qa-s3-compatible-target"
	BLOBBackuptarget         = "pds-qa-blob-target"
	serviceType              = "LoadBalancer"
	pdsNamespaces            []string

	supportedDataServices = map[string]string{"cas": "Cassandra", "zk": "ZooKeeper", "kf": "Kafka", "rmq": "RabbitMQ", "pg": "PostgreSQL"}

	backupSupportedDataService = map[string]string{"cas": "Cassandra", "pg": "PostgreSQL"}
	//futureSupportedDataService = map[string]string{"mdb": "Mongodb", "red": "Redis", "con": "Consul", "cbs": "Couchbase", "dse": "DatastaxEnterprise", "ess": "Elasticsearch"}
)

type PDSTestSuite struct {
	suite.Suite
	ControlPlane    *ControlPlane
	TargetCluster   *TargetCluster
	TestEnvironment *Environment
	ctx             context.Context
	apiClient       *pds.APIClient
	components      *api.Components
	env             Environment
}

const (
	duration  = 900
	sleepTime = 10

	defaultNumPods = 3
	dnsZone        = "portworx.pds-dns.io"

	// FIX-ME
	// Create the template manually for all the data serices with below name (i.e QaDefault)
	storageTemplateName   = "QaDefault"
	resourceTemplateName  = "QaDefault"
	appConfigTemplateName = "QaDefault"
	deploymentName        = "automation"
	templateName          = "QaDefault"
)

var (
	dataServiceDefaultResourceTemplateIdMap = make(map[string]string)
	dataServiceNameIdMap                    = make(map[string]string)
	dataServiceNameVersionMap               = make(map[string][]string)
	dataServiceIdImagesMap                  = make(map[string]string)
	dataServiceNameDefaultAppConfigMap      = make(map[string]string)
	deployementIdNameMap                    = make(map[string]string)
	namespaceNameIdMap                      = make(map[string]string)
	deployementIdnameWithSchBkpMap          = make(map[string]string)
	deployementIdnameWithAdhocBkpMap        = make(map[string]string)
	storageTemplateID                       string
)

func (suite *PDSTestSuite) SetupSuite() {

	log.Info(`
		==========================================================================
		@owner: PDS-QA team
		Please go through https://github.com/portworx/pds-functional-test
		Right now we supported only basic sanity tests.
		Resources(Creation/deletion) as part of the runs.
			- PDS Helm chart will be installed to the lastest supported version w.r.t your control plane.
			- Namespaces - pds-automation-*
			- PVC / PV 
		- Prerequsite
			- Please make sure kubectl and helm are installed.
			- "QaDefault" Storage option / Resource / Appconfig template should be present.
			- Create the template manually for all the data serices having name as QaDefault if its not already populated.

		==========================================================================
	`)
	// Perform basic setup with sanity checks.
	log.Info("Check for environmental variable.")
	suite.env = MustHaveEnvVariables()

	endpointUrl, err := url.Parse(suite.env.CONTROL_PLANE_URL)
	if err != nil {
		log.Panicf("Unable to access the URL: %s", suite.env.CONTROL_PLANE_URL)
	}
	apiConf := pds.NewConfiguration()
	apiConf.Host = endpointUrl.Host
	apiConf.Scheme = endpointUrl.Scheme

	// Use Configuration or context with WithValue (above)
	suite.ctx = context.WithValue(context.Background(), pds.ContextAPIKeys, map[string]pds.APIKey{"ApiKeyAuth": {Key: "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiIsImp0aSI6IjgwZTcxMDU5MmQ1NmI3MDA0Mjg2ODEyNTYyYjI2M2VlMWQzMzM1YWNkYTkyNmRkZWE2M2U5MGEyYjE1ZGQ3MjdhYTMzNTU4NzMyZmViMzNiIn0.eyJhdWQiOiIzIiwianRpIjoiODBlNzEwNTkyZDU2YjcwMDQyODY4MTI1NjJiMjYzZWUxZDMzMzVhY2RhOTI2ZGRlYTYzZTkwYTJiMTVkZDcyN2FhMzM1NTg3MzJmZWIzM2IiLCJpYXQiOjE2NTE0NjI4ODMsIm5iZiI6MTY1MTQ2Mjg4MywiZXhwIjoxNjUxNTQ5MjgzLCJzdWIiOiIyMTk5NiIsImlzcyI6Imh0dHBzOlwvXC9hcGljZW50cmFsLnBvcnR3b3J4LmNvbVwvYXBpIiwibmFtZSI6Ik1hZGFuYWdvcGFsIEFydW5hY2hhbGFtIiwiZW1haWwiOiJtYXJ1bmFjaGFsYW1AcHVyZXN0b3JhZ2UuY29tIiwic2NvcGVzIjpbXSwicm9sZXMiOlsicHgtYmFja3VwLWluZnJhLmFkbWluIl19.IcrrTJeuT3JEGcg7h0nh5HTRLiRORyd92K0G48bEZLRbB0c2FHYsUxRHbS4o7Y4YqJFrx1kjpk-eS6FQv5pgZXu-eLbtVmc-vEBjApwULxrnn2-TlnyzzjBqSPyDURSGpVLyeoa5l_w7DPjZqDW6Z8H-EfVRLgobBYO-j5Vt9bw9vPsqIbWrB6zr_yEmx8JwgU3iZasQG0sQvAxw8owivzwUK3l5I-_WaLeB_b8Sl5t-p0NbqDu08grUoTwKAjl6lTyMcpq80NFYTLEJzpaDkkukHB3swwGEYYlUojUJARifESaKxHXZwm3eQhn7HmcIQmCXVlwzsxmr6JTaNJ1nXjVbthBUwWtOAiYD-oZharzaYyRFwqHwXytfVA-onKcuAKL0xo7DZy5MY0hzvieIm35DmXZ3ApiEH3_N0AGfz0tjW7KU1m1AJ1XN0JhAa8BciDINvfvLIpmwCWMWnANKVwltcE9JzHnz4lH5whgTy1XjVa7KqiRhMkQsH35inf_fkm5uxcH5WpbuuBDubCp_5B7wEmsptr5GuYJU9JbghAuaNTHfuJQdg4A1hDYkvNhotbxCHpIOflfiYVJqJuvyeleTVti4utP2maCRDyikFQdiaRsaOd7O3caw2m_7CsXyxZ8qrkgfxKc94E_uq3ijvAccvKPsIb_D8Xu40nZx43k"}})
	suite.apiClient = pds.NewAPIClient(apiConf)
	suite.components = api.NewComponents(suite.ctx, suite.apiClient)

	log.Info("Get control plane cluster.")
	suite.ControlPlane = NewControlPlane(suite.env.CONTROL_PLANE_URL, suite.components)

	log.Info("Get target cluster.")
	suite.TargetCluster = NewTargetCluster(suite.env.TARGET_KUBECONFIG)

	acc := suite.components.Account
	accounts, _ := acc.GetAccountsList()
	if strings.EqualFold(suite.env.CLUSTER_TYPE, "onprem") {
		serviceType = "ClusterIP"
	}

	for i := 0; i < len(accounts); i++ {
		log.Infof("Account Name: %v", accounts[i].GetName())
		if accounts[i].GetName() == accountName {
			accountId = accounts[i].GetId()
		}
	}
	log.Infof("Account Detail- Name: %s, UUID: %s ", accountName, accountId)
	tnts := suite.components.Tenant
	tenants, _ := tnts.GetTenantsList(accountId)
	tenantId = tenants[0].GetId()
	tenantName := tenants[0].GetName()
	log.Infof("Tenant Details- Name: %s, UUID: %s ", tenantName, tenantId)
	projcts := suite.components.Project
	projects, _ := projcts.GetprojectsList(tenantId)
	projectId = projects[0].GetId()
	projectName := projects[0].GetName()
	log.Infof("Project Details- Name: %s, UUID: %s ", projectName, projectId)

	log.Info("Get helm version")
	version, _ := suite.components.ApiVersion.GetHelmChartVersion()
	log.Infof("Helm chart Version : %s ", version)

	clusterId, err := GetClusterId(suite.env.TARGET_KUBECONFIG)
	if err != nil {
		log.Panicf("Unable to fetch the cluster Id")
	}

	log.Infof("Register cluster %s to control plane %v ", clusterId, suite.env.CONTROL_PLANE_URL)
	err = suite.TargetCluster.RegisterToControlPlane(suite.env.CONTROL_PLANE_URL, version, suite.ControlPlane.GetRegistrationToken(tenantId), tenantId)
	if err != nil {
		log.Panicf("Unable to register the target cluster to control plane %v", suite.env.CONTROL_PLANE_URL)
	}
	log.Info("Creating namespaces for data service deployment")
	pdsNamespaces = []string{"pds-automation-1", "pds-automation-2", "pds-automation-3"}
	for _, ns := range pdsNamespaces {
		log.Infof("Namespace name - %s", ns)
		suite.TargetCluster.CreatePDSNamespace(ns)
	}

	log.Infof("Creating Storage Template")
	suite.ControlPlane = NewControlPlane(suite.env.CONTROL_PLANE_URL, suite.components)
	template, err := suite.ControlPlane.CreateStorageOptionTemplate(tenantId, false, "ext4", storageTemplateName, 2, false)
	storageTemplateID = *template.Id
	log.Infof("storage template id %s", storageTemplateID)
	if err != nil {
		log.Panicf("Storage template creation failed with error - %v", err)
	}
}

func (suite *PDSTestSuite) TearDownSuite() {
	log.Info("Sleeping for 5 minutes before teardown.")
	time.Sleep(5 * time.Minute)
	log.Warn("Cleaning all the deployment created as part of this test run")
	log.Info("Sleep for sometime.")
	time.Sleep(1 * time.Minute)
	for id := range deployementIdNameMap {
		log.Infof("Deleting the deployment: %v", id)
		suite.components.DataServiceDeployment.DeleteDeployment(id)
		time.Sleep(sleepTime)
	}

	log.Infof("Deleting the storage template")
	suite.components.StorageSettingsTemplate.DeleteTemplate(storageTemplateID)

	log.Info("Sleep for a minute.")
	time.Sleep(1 * time.Minute)
	for id := range deployementIdnameWithSchBkpMap {
		backups, _ := suite.components.Backup.ListBackup(id)
		for _, backup := range backups {
			backupId := backup.GetId()
			log.Infof("Delete back up having Id - %v", backupId)
			response, err := suite.components.Backup.DeleteBackup(backupId)
			// Success is indicated with 2xx status codes:
			statusOK := response.StatusCode >= 200 && response.StatusCode < 300
			if !statusOK {
				fmt.Println("Non-OK HTTP status:", response.StatusCode)
				// You may read / inspect response body
				log.Error(err)
			}
			time.Sleep(10 * time.Second)
		}
		log.Infof("Deleting the deployment: %v", id)
		suite.components.DataServiceDeployment.DeleteDeployment(id)
		time.Sleep(15 * time.Second)
	}
	log.Info("Sleep for sometime.")
	time.Sleep(1 * time.Minute)
	for id := range deployementIdnameWithAdhocBkpMap {
		backups, _ := suite.components.Backup.ListBackup(id)
		for _, backup := range backups {
			backupId := backup.GetId()
			log.Infof("Delete back up having Id - %v", backupId)
			response, err := suite.components.Backup.DeleteBackup(backupId)
			// Success is indicated with 2xx status codes:
			statusOK := response.StatusCode >= 200 && response.StatusCode < 300
			if !statusOK {
				fmt.Println("Non-OK HTTP status:", response.StatusCode)
				// You may read / inspect response body
				log.Error(err)
			}
			time.Sleep(10 * time.Second)
		}
		log.Infof("Deleting the deployment: %v", id)
		suite.components.DataServiceDeployment.DeleteDeployment(id)
		time.Sleep(25 * time.Second)
	}

	log.Info("Deleting all the Persistent Volume claims created as part of this test run")
	for _, ns := range pdsNamespaces {
		DeleteAllPVC(ns, suite.env.TARGET_KUBECONFIG)
	}

	log.Info("Deleting all the Released Persistent volume")
	DeleteAllReleasedPV(suite.env.TARGET_KUBECONFIG)

	log.Info("Deleting all the namesapce created for deployment.")
	for _, ns := range pdsNamespaces {
		DeleteNamespace(ns, suite.env.TARGET_KUBECONFIG)
	}
	suite.TargetCluster.DeRegisterFromControlPlane()

}

func TestPDSTestSuite(t *testing.T) {
	suite.Run(t, new(PDSTestSuite))
}
