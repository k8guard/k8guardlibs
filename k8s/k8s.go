package k8s

import (
	"errors"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


func LoadClientset() (kubernetes.Clientset, error) {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	configOverrides := &clientcmd.ConfigOverrides{}
	kubeConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(loadingRules, configOverrides)
	config, err := kubeConfig.ClientConfig()
	if err != nil {
		return kubernetes.Clientset{}, errors.New("Failed loading client config")
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return kubernetes.Clientset{}, errors.New("Failed getting clientset")
	}
	return *clientset, nil
}

