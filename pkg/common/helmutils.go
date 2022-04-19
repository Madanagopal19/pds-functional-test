package common

// import (
// 	"flag"
// 	"fmt"

// 	"k8s.io/client-go/kubernetes"
// 	"k8s.io/client-go/tools/clientcmd"
// 	"k8s.io/helm/pkg/helm"
// 	"k8s.io/helm/pkg/helm/portforwarder"
// )

// func GetHelmList(namespace string, pathToKubeconfig string) {
// 	kubeconfig := flag.String("kubeconfig", "", pathToKubeconfig)
// 	flag.Parse()
// 	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	client, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	tillerTunnel, _ := portforwarder.New(namespace, client, config)

// 	host := fmt.Sprintf("127.0.0.1:%d", tillerTunnel.Local)

// 	helmClient := helm.NewClient(helm.Host(host))

// 	// list/print releases
// 	resp, _ := helmClient.ListReleases()
// 	for _, release := range resp.Releases {
// 		log.InfoFn(release.GetName())
// 		log.InfoFn(release)
// 	}

// }
