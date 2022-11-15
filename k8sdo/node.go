package k8sdo

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func GetNode(kubeclient *kubernetes.Clientset, label string) ([]corev1.Node, error) {
	//过滤
	option := metav1.ListOptions{}
	if label != "" {
		option.LabelSelector = label
	}

	nodeList, err := kubeclient.CoreV1().Nodes().List(context.Background(), option)
	if err != nil {
		return nil, err
	}

	return nodeList.Items, nil
}
