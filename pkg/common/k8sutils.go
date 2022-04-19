package common

import (
	"context"
	"time"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func GetPods(namespace string, pathToKubeconfig string) []v1.Pod {
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

func VerifyPodsStatus(namespace string, pathToKubeconfig string) {
	config, err := clientcmd.BuildConfigFromFlags("", pathToKubeconfig)
	if err != nil {
		log.Panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Panic(err)
	}
	watchPodsStatus(namespace, pathToKubeconfig)
	pods := GetPods(namespace, pathToKubeconfig)
	log.Infof("There are %d pods in the cluster in the namespace %s", len(pods), namespace)
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
		time.Sleep(10 * time.Second)
	}
}

func GetNamespaces(pathToKubeconfig string) []v1.Namespace {
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
		log.Panic(err)
	}
	return namespaces.Items
}

func watchPodsStatus(namespace string, pathToKubeconfig string) {
	// kubeconfig = flag.String("kubeconfig", "", pathToKubeconfig)
	// flag.Parse()
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
