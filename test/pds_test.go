package test

import (
	"context"
	"fmt"
	"net/url"
	"testing"

	pds "github.com/portworx/pds-api-go-client/pds/v1alpha1"
	api "github.com/portworx/pds-functional-test/pkg/api"
	. "github.com/portworx/pds-functional-test/pkg/common"
	logger "github.com/portworx/pds-functional-test/pkg/logger"
	"github.com/stretchr/testify/suite"
)

// Logger alias
var log = logger.Log

type PDSTestSuite struct {
	suite.Suite
	TargetCluster   *TargetCluster
	TestEnvironment *Environment
	ctx             context.Context
	apiClient       *pds.APIClient
	components      *api.Components
	env             Environment
}

func (s *PDSTestSuite) SetupSuite() {

	// Perform basic setup with sanity checks.
	log.Info("Check for environmental variable.")
	s.env = MustHaveEnvVariables()

	log.Info("Get target cluster.")
	s.TargetCluster = NewTargetCluster(s.env.TARGET_KUBECONFIG)

	endpointUrl, err := url.Parse(s.env.CONTROL_PLANE_URL)
	if err != nil {
		s.T().Errorf("Unable to access the URL: %s", s.env.CONTROL_PLANE_URL)
	}
	apiConf := pds.NewConfiguration()
	apiConf.Host = endpointUrl.Host
	apiConf.Scheme = endpointUrl.Scheme

	// // Use Configuration or context with WithValue (above)
	s.ctx = context.WithValue(context.Background(), pds.ContextAPIKeys, map[string]pds.APIKey{"ApiKeyAuth": {Key: GetBearerToken(true), Prefix: "Bearer"}})
	s.apiClient = pds.NewAPIClient(apiConf)
	s.components = api.NewComponents(s.ctx, s.apiClient)

}

func (s *PDSTestSuite) AfterTest(suiteName, testName string) {
	if s.T().Failed() {
		log.Errorf(fmt.Sprintf("Failed test %s:", testName))
	}
}

func TestPDSSuite(t *testing.T) {
	suite.Run(t, new(PDSTestSuite))
}
