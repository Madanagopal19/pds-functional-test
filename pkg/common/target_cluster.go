package common

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// TargetCluster struct
type TargetCluster struct {
	kubeconfig string
}

func (targetCluster *TargetCluster) RegisterToControlPlane(controlPlaneUrl string, helmChartversion string, bearerToken string, tenantId string) error {
	log.Info("Test control plane url connectivity.")
	_, err := isReachbale(controlPlaneUrl)
	if err != nil {
		return fmt.Errorf("unable to reach the control plane with following error - %v", err)
	}
	var cmd, currentHelmVerison string
	apiEndpoint := fmt.Sprintf(controlPlaneUrl + "/api")
	log.Infof("Verify if the namespace %s already exits.", pdsSystemNamespace)
	isExist := IsNamespaceExist(pdsSystemNamespace, targetCluster.kubeconfig)
	isRegistered := false
	if isExist {
		log.Infof("%s namespace already exists.", pdsSystemNamespace)
		pods := ListPods(pdsSystemNamespace, targetCluster.kubeconfig)
		if len(pods) > 0 {
			log.Warnf("Target cluster is already registered to control plane.")
			cmd = fmt.Sprintf("helm list -A --kubeconfig %s", targetCluster.kubeconfig)
			currentHelmVerison, _ = GetCurrentHelmVersion(targetCluster.kubeconfig)
			if !strings.EqualFold(currentHelmVerison, helmChartversion) {
				log.Infof("Upgrading PDS helm chart from %v to %v", currentHelmVerison, helmChartversion)
				cmd = fmt.Sprintf("helm upgrade --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%s --set tenantId=%s "+
					"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)
			}
			isRegistered = true
		}
	}

	if !isRegistered {
		log.Infof("Installing PDS ( helm version -  %v)", helmChartversion)
		cmd = fmt.Sprintf("helm install --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%s --set tenantId=%s "+
			"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)
	}
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Warn("Kindly remove the PDS chart properly and retry. CMD>> helm uninstall  pds --namespace pds-system --kubeconfig $KUBECONFIG")
		log.Error(err)
		return err
	}
	log.Infof("Terminal output -> %v", output)
	time.Sleep(30 * time.Second)
	log.Infof("Watch states of pods in %s namespace", pdsSystemNamespace)
	WatchPodsStatus(pdsSystemNamespace, targetCluster.kubeconfig)

	time.Sleep(30 * time.Second)
	log.Infof("Verify the health of all the pods in %s namespace", pdsSystemNamespace)
	CheckPodsHealth(pdsSystemNamespace, targetCluster.kubeconfig)
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

func (targetCluster *TargetCluster) CreatePDSNamespace(name string) error {
	log.Infof("Create namespace.")
	if IsNamespaceExist(name, targetCluster.kubeconfig) {
		log.Infof("Namespace %s already exists.", name)
	} else {
		cmd := fmt.Sprintf("kubectl create namespace %v --kubeconfig %s", name, targetCluster.kubeconfig)
		output, _, err := ExecShell(cmd)
		if err != nil {
			log.Error(err)
			return err
		}
		log.Infof("Terminal output ----> %v", output)
	}

	log.Infof("Add PDS label.")
	cmd := fmt.Sprintf("kubectl label namespaces %v pds.portworx.com/available=true --overwrite=true --kubeconfig %s", name, targetCluster.kubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Error(err)
		return err
	}
	log.Infof("Terminal output ----> %v", output)

	return nil
}

// NewTargetCluster lsajajsklj
func NewTargetCluster(context string) *TargetCluster {
	return &TargetCluster{
		kubeconfig: context,
	}
}
