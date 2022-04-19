package common

import (
	"fmt"
	"net/http"
	"time"
)

// TargetCluster khash
type TargetCluster struct {
	kubeconfig string
}

const (
	pdsSystemNamespace = "pds-system"
)

func (targetCluster *TargetCluster) RegisterToControlPlane(controlPlaneUrl string, helmChartversion string, bearerToken string, tenantId string) error {
	log.Info("Test control plane url connectivity.")
	_, err := isReachbale(controlPlaneUrl)
	if err != nil {
		return fmt.Errorf("unable to reach the control plane with following error - %v", err)
	}
	apiEndpoint := fmt.Sprintf(controlPlaneUrl + "/api")
	cmd := fmt.Sprintf("helm install --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%s --set tenantId=%s "+
		"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)

	_, _, err = ExecShell(cmd)
	if err != nil {
		log.Error(err)
	}
	time.Sleep(30 * time.Second)
	log.Info("Verify pods status.")
	VerifyPodsStatus(pdsSystemNamespace, targetCluster.kubeconfig)
	return err
}

func isReachbale(url string) (bool, error) {
	timeout := time.Duration(15 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(url)
	if err != nil {
		log.Error(err.Error())
		return false, err
	}
	return true, nil
}

// NewTargetCluster lsajajsklj
func NewTargetCluster(context string) *TargetCluster {
	return &TargetCluster{
		kubeconfig: context,
	}
}
