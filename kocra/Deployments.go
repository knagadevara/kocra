package kocra

import (
	"context"
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func ListDeployments(clientset *kubernetes.Clientset, namespace string) *v1.DeploymentList {
	deployments, err := clientset.AppsV1().
		Deployments(namespace).
		List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return deployments
}

func GetDeploymentDetails(clientset *kubernetes.Clientset, namespace string, depName string) *v1.Deployment {
	deploymentDetails, err := clientset.AppsV1().
		Deployments(namespace).Get(context.TODO(), depName, metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return deploymentDetails
}
