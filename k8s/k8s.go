package k8s

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
)

func LoadClientset() (*kubernetes.Clientset, error) {

	configFilePath := "/root/.kube/config"

	if _, err := os.Stat("/root/.kube/config"); os.IsNotExist(err) {
		configFilePath = ""
	}

	config, err := clientcmd.BuildConfigFromFlags("", configFilePath)

	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return &kubernetes.Clientset{}, err
	}
	return clientset, nil
}
