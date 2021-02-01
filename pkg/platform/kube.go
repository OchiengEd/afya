package platform

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// Client function returns the clientset with with to
// interact with the container platform in which the
// pod would be running
func Client() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = getKubeConfig()
		if err != nil {
			return nil, err
		}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return clientset, nil
}

func getKubeConfig() (*rest.Config, error) {
	homePath := homedir.HomeDir()
	if homePath == "" {
		return nil, fmt.Errorf("HOME environment is empty")
	}

	configPath := filepath.Join(homePath, ".kube", "config")

	return clientcmd.BuildConfigFromFlags("", configPath)
}
