package kocra

import (
	apps "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/rest"
)

func CreateControllerRevisionIface(av1 apps.AppsV1Client, namespace string) *apps.ControllerRevisionInterface {
	obj := av1.ControllerRevisions(namespace)
	return &obj
}

func CreateDaemonSetIface(av1 apps.AppsV1Client, namespace string) *apps.DaemonSetInterface {
	obj := av1.DaemonSets(namespace)
	return &obj
}

func CreateDeploymentIface(av1 apps.AppsV1Client, namespace string) *apps.DeploymentInterface {
	obj := av1.Deployments(namespace)
	return &obj
}

func CreateReplicaSetIface(av1 apps.AppsV1Client, namespace string) *apps.ReplicaSetInterface {
	obj := av1.ReplicaSets(namespace)
	return &obj
}

func CreateStatefulSetIface(av1 apps.AppsV1Client, namespace string) *apps.StatefulSetInterface {
	obj := av1.StatefulSets(namespace)
	return &obj
}

func CreateAppRESTClientIface(av1 apps.AppsV1Client) *rest.Interface {
	obj := av1.RESTClient()
	return &obj
}
