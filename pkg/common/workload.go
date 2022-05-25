package common

import (
	"context"
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//CreatepostgresqlWorkload
func (cp *ControlPlane) CreatepostgresqlWorkload(ctx context.Context, deploymentId string, scalfactor string, iterations string, deploymentName string, namespace string, replicas int32,
	podLabels map[string]string, nodeSelector map[string]string) (*appsv1.Deployment, error) {
	deploymentSpec := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: podLabels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: podLabels,
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:    "pgbench",
							Image:   "madan19/pgbench:pgloadTest1",
							Command: []string{"/pgloadgen.sh"},
							Args:    []string{},
						},
					},
					RestartPolicy: v1.RestartPolicyAlways,
				},
			},
		},
	}

	dataServiceDeployment := cp.components.DataServiceDeployment
	deploymentConnectionDetails, err := dataServiceDeployment.GetConnectionDetails(deploymentId)
	if err != nil {
		return nil, fmt.Errorf("GetConnectionDetails API error: %v", err)
	}
	deploymentNodes := deploymentConnectionDetails.GetNodes()
	log.Infof("Deployment nodes %v", deploymentNodes)
	var dnsEndpoint string
	for _, nodes := range deploymentNodes {
		if !strings.Contains(nodes, "vip") {
			dnsEndpoint = nodes
		}
	}
	log.Infof("Dataservice DNS endpoint %s", dnsEndpoint)
	dataServicePassword, err := dataServiceDeployment.GetDeploymentCredentials(deploymentId)
	if err != nil {
		return nil, fmt.Errorf("GetDeploymentCredentials API error: %v", err)
	}
	pdsPassword := dataServicePassword.GetPassword()

	DataserviceSpec := []string{dnsEndpoint, pdsPassword, scalfactor, iterations}
	for index := range deploymentSpec.Spec.Template.Spec.Containers {
		deploymentSpec.Spec.Template.Spec.Containers[index].Args = DataserviceSpec
	}
	client := Getk8sClient()
	deployment, err := client.AppsV1().Deployments(namespace).Create(ctx, deploymentSpec, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("deployment %q Create API error: %v", deploymentSpec.Name, err)
	}

	err = WaitForDeployment(ctx, client, deployment, namespace, replicas)
	if err != nil {
		return nil, fmt.Errorf("deployment %q failed to come up with the err : %v", deployment.Name, err)
	}
	log.Infof("Workload Started for deployment: %s ", deploymentId)
	return deployment, nil
}

func (cp *ControlPlane) DeleteDeployment(ctx context.Context, deploymentName string, namespace string) error {
	client := Getk8sClient()
	err := client.AppsV1().Deployments(namespace).Delete(ctx, deploymentName, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("deployment delete API error: %v", err)
	}

	log.Infof("Deployment %s Deleted successfully ", deploymentName)
	return nil
}
