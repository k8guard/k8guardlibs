package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)


func LoadClientset() (*kubernetes.Clientset, error) {

	var config *rest.Config
	var err error

	if _, err := os.Stat("/root/.kube/config"); os.IsNotExist(err) {
		config, err = rest.InClusterConfig()
	} else {
		config, err = clientcmd.BuildConfigFromFlags("", "/root/.kube/config")

	}
	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	return clientset, nil
}

