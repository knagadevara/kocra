package kocra

import (
	"context"
	"log"

	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	apps "k8s.io/client-go/kubernetes/typed/apps/v1"
)

func ListDeployments(d apps.DeploymentInterface, metaLstOps metav1.ListOptions) *v1.DeploymentList {
	deployments, err := d.List(context.TODO(), metaLstOps)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return deployments
}

func GetDeploymentDetails(d apps.DeploymentInterface, namespace string, deploymentName string) *v1.Deployment {
	deploymentDetails, err := d.Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return deploymentDetails
}
