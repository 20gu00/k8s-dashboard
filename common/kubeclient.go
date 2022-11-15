package common

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"path/filepath"
)

var (
	kubeclient *kubernetes.Clientset
)

func kubeClient() error {
	var err error
	var config *rest.Config
	//no flag
	kubeConfig := filepath.Join(homedir.HomeDir(), ".kube", "config")

	config, err = rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
		if err != nil {
			return err
		}
	}

	//实例化clientset(kube)
	kubeclient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return err
	}

	return nil
}

//同事屏蔽了细节
func NewKubeClient() *kubernetes.Clientset {
	return kubeclient
}
