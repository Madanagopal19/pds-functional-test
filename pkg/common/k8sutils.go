package common

import (
	"context"
	"fmt"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func ListPods(namespace string, pathToKubeconfig string) []v1.Pod {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}

	podList, err := clientset.CoreV1().Pods(namespace).List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Error(err)
	}
	return podList.Items
}

func CheckPodsHealth(namespace string, pathToKubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}
	pods := ListPods(namespace, pathToKubeconfig)
	log.Infof("There are %d pods present in the namespace %s", len(pods), namespace)
	for _, pod := range pods {
		_, err = clientset.CoreV1().Pods(namespace).Get(context.TODO(), pod.GetName(), metav1.GetOptions{})
		if errors.IsNotFound(err) {
			log.Panicf("Pod %s in namespace %s not found", pod.GetName(), namespace)
		} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
			log.Errorf("Found pod %s (Namespace -  %s) in unhealthy state.", pod.GetName(), namespace)
			log.Panicf("Error getting pod %s in namespace %s: %v",
				pod.GetName(), namespace, statusError.ErrStatus.Message)
		} else if err != nil {
			log.Errorf("Found pod %s (Namespace -  %s) in unhealthy state.", pod.GetName(), namespace)
			log.Panic(err)
		} else {
			log.Infof("Found pod %s (Namespace -  %s) in healthy state.", pod.GetName(), namespace)
		}
		time.Sleep(1 * time.Second)
	}
}

func GetNamespaces(pathToKubeconfig string) ([]v1.Namespace, error) {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}

	namespaces, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Info(err)
	}
	return namespaces.Items, err
}
func GetNamespace(name string, pathToKubeconfig string) (*v1.Namespace, error) {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}

	namespace, err := clientset.CoreV1().Namespaces().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		log.Info(err)
	}
	return namespace, err
}

func IsNamespaceExist(name string, pathToKubeconfig string) bool {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}

	_, err = clientset.CoreV1().Namespaces().Get(context.Background(), name, metav1.GetOptions{})
	return !(errors.IsAlreadyExists(err) || errors.IsNotFound(err))
}

func WatchPodsStatus(namespace string, pathToKubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	watch, err := clientset.CoreV1().Pods(namespace).Watch(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatal(err.Error())
	}
	go func() {
		for event := range watch.ResultChan() {
			log.Infof("Type: %v", event.Type)
			p, ok := event.Object.(*v1.Pod)
			if !ok {
				log.Panic("unexpected type")
			}
			log.Info(p.Status.ContainerStatuses)
			log.Info(p.Status.Phase)
		}
	}()
	time.Sleep(10 * time.Second)
}

func GetClusterId(pathToKubeconfig string) (string, error) {
	log.Infof("Fetch Cluster id ")
	cmd := fmt.Sprintf("kubectl get ns kube-system -o jsonpath={.metadata.uid} --kubeconfig %s", pathToKubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Error(err)
		return "Connection Refused!!", err
	}
	return output, nil
}

func DeleteNamespace(name string, pathToKubeconfig string) error {
	log.Infof("Deleting all the resources in %s namespace.", name)
	cmd := fmt.Sprintf("kubectl delete all --all -n %s --kubeconfig %s", name, pathToKubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Error(err)
	}
	log.Info(output)
	log.Infof("Deleting the namespace %s", name)
	cmd = fmt.Sprintf("kubectl delete ns %s --kubeconfig %s", name, pathToKubeconfig)
	output, _, err = ExecShell(cmd)
	if err != nil {
		log.Error(err)
	}
	log.Info(output)
	return err
}

func DeleteAllPVC(namespace string, pathToKubeconfig string) error {
	log.Infof("Deleting all pvc in namespace %s ", namespace)
	cmd := fmt.Sprintf("kubectl delete pvc --all -n %s --kubeconfig %s", namespace, pathToKubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Error(err)
	}
	log.Info(output)
	return err
}

func DeleteAllReleasedPV(pathToKubeconfig string) error {
	cmd := fmt.Sprintf("kubectl get pv --kubeconfig %s | tail -n+2 | awk '$5 == \"Released\" {print $1}' | xargs kubectl delete pv --kubeconfig %s", pathToKubeconfig, pathToKubeconfig)
	output, _, err := ExecShell(cmd)
	if err != nil {
		log.Error(err)
	}
	log.Info(output)
	return err
}
