package kocra

import (
	core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
)

func CreateComponentStatusIface(cv1 core.CoreV1Interface) *core.ComponentStatusInterface {
	obj := cv1.ComponentStatuses()
	return &obj
}

func CreateNamespaceIface(cv1 core.CoreV1Interface) *core.NamespaceInterface {
	obj := cv1.Namespaces()
	return &obj
}

func CreateNodeIface(cv1 core.CoreV1Interface) *core.NodeInterface {
	obj := cv1.Nodes()
	return &obj
}

func CreatePersistentVolumeIface(cv1 core.CoreV1Interface) *core.PersistentVolumeInterface {
	obj := cv1.PersistentVolumes()
	return &obj
}

func CreateEndpointsIface(cv1 core.CoreV1Interface, namespace string) *core.EndpointsInterface {
	obj := cv1.Endpoints(namespace)
	return &obj
}

func CreateEventsIface(cv1 core.CoreV1Interface, namespace string) *core.EventInterface {
	obj := cv1.Events(namespace)
	return &obj
}

func CreateLimitRangeIface(cv1 core.CoreV1Interface, namespace string) *core.LimitRangeInterface {
	obj := cv1.LimitRanges(namespace)
	return &obj
}

func CreatePersistentVolumeClaimsIface(cv1 core.CoreV1Interface, namespace string) *core.PersistentVolumeClaimInterface {
	obj := cv1.PersistentVolumeClaims(namespace)
	return &obj
}

func CreatePodTemplateIface(cv1 core.CoreV1Interface, namespace string) *core.PodTemplateInterface {
	obj := cv1.PodTemplates(namespace)
	return &obj
}

func CreatePodIface(cv1 core.CoreV1Interface, namespace string) *core.PodInterface {
	obj := cv1.Pods(namespace)
	return &obj
}

func CreateReplicationControllerIface(cv1 core.CoreV1Interface, namespace string) *core.ReplicationControllerInterface {
	obj := cv1.ReplicationControllers(namespace)
	return &obj
}

func CreateResourceQuotaIface(cv1 core.CoreV1Interface, namespace string) *core.ResourceQuotaInterface {
	obj := cv1.ResourceQuotas(namespace)
	return &obj
}

func CreateSecretIface(cv1 core.CoreV1Interface, namespace string) *core.SecretInterface {
	obj := cv1.Secrets(namespace)
	return &obj
}

func CreateServiceAccountIface(cv1 core.CoreV1Interface, namespace string) *core.ServiceAccountInterface {
	obj := cv1.ServiceAccounts(namespace)
	return &obj
}

func CreateServicesIface(cv1 core.CoreV1Interface, namespace string) *core.ServiceInterface {
	obj := cv1.Services(namespace)
	return &obj
}

func CreateCoreRESTClientIface(cv1 core.CoreV1Interface) *rest.Interface {
	obj := cv1.RESTClient()
	return &obj
}
