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

func (targetCluster *TargetCluster) RegisterToControlPlane(controlPlaneUrl string, helmChartversion string, bearerToken string, tenantId string, clusterType string) error {
	log.Info("Test control plane url connectivity.")
	_, err := isReachbale(controlPlaneUrl)
	if err != nil {
		return fmt.Errorf("unable to reach the control plane with following error - %v", err)
	}
	var cmd string
	apiEndpoint := fmt.Sprintf(controlPlaneUrl + "/api")
	log.Infof("Verify if the namespace %s already exits.", pdsSystemNamespace)
	isExist := IsNamespaceExist(pdsSystemNamespace, targetCluster.kubeconfig)
	isRegistered := false
	if isExist {
		log.Infof("%s namespace already exists.", pdsSystemNamespace)
		pods := ListPods(pdsSystemNamespace, targetCluster.kubeconfig)
		if len(pods) > 0 {
			log.Warn("Target cluster is already registered to control plane.")
			if !targetCluster.isLatestHelm() {
				log.Infof("Upgrading PDS helm chart to %v", helmChartversion)
				if strings.EqualFold(clusterType, "ocp") {
					cmd = fmt.Sprintf("helm upgrade --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%v --set platform=ocp --set tenantId=%s "+
						"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)
				} else {
					cmd = fmt.Sprintf("helm upgrade --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%v --set tenantId=%s "+
						"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)
				}

			}
			isRegistered = true
		} else {
			log.Infof("Just the %s namespace exists, but no pods are available.", pdsSystemNamespace)
		}
	}

	if !isRegistered {
		log.Infof("Installing PDS ( helm version -  %v)", helmChartversion)
		if strings.EqualFold(clusterType, "ocp") {
			cmd = fmt.Sprintf("helm install --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%s --set platform=ocp --set tenantId=%s "+
				"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)
		} else {
			cmd = fmt.Sprintf("helm install --create-namespace --namespace=%s pds pds-target --repo=https://portworx.github.io/pds-charts --version=%s --set tenantId=%s "+
				"--set bearerToken=%s --set apiEndpoint=%s --kubeconfig %s", pdsSystemNamespace, helmChartversion, tenantId, bearerToken, apiEndpoint, targetCluster.kubeconfig)
		}

	}
	log.Warn(cmd)
	time.Sleep(10 * time.Second)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Warn("Kindly remove the PDS chart properly and retry if that helps(or slack us for more details). CMD>> helm uninstall  pds --namespace pds-system --kubeconfig $KUBECONFIG")
		log.Panic(err)
		return err
	}
	log.Infof("Terminal output -> %v", output)
	time.Sleep(20 * time.Second)
	log.Infof("Watch states of pods in %s namespace", pdsSystemNamespace)
	WatchPodsStatus(pdsSystemNamespace, targetCluster.kubeconfig)

	time.Sleep(20 * time.Second)
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
	time.Sleep(2 * time.Minute)
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

func (targetCluster *TargetCluster) isLatestHelm() bool {
	cmd := fmt.Sprintf(" helm ls --all -n pds-system --kubeconfig %s ", targetCluster.kubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Panic(err)
	}
	log.Infof(output)
	cmd = fmt.Sprintf(" helm ls --all -n pds-system --kubeconfig %s | tail -n+2 | awk '{print $8}' ", targetCluster.kubeconfig)
	output, _, err = ExecShell(cmd)
	if err != nil {
		log.Panic(err)
	}
	output = strings.TrimSpace(output)
	log.Infof("Helm chart Status- %v", strings.EqualFold(output, "pending-upgrade"))
	return !strings.EqualFold(output, "pending-upgrade")

}

func (targetCluster *TargetCluster) DeRegisterFromControlPlane() {
	log.Info("Derisgtering the target cluster from control plane.")
	cmd := fmt.Sprintf("helm uninstall --namespace=pds-system pds --kubeconfig %s ", targetCluster.kubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Panic(err)
	}
	log.Info(output)
	log.Info("Deleting the pds-system namespace.")
	DeleteNamespace("pds-system", targetCluster.kubeconfig)

}

// NewTargetCluster lsajajsklj
func NewTargetCluster(context string) *TargetCluster {
	return &TargetCluster{
		kubeconfig: context,
	}
}
