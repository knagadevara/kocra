package kocra

import (
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/kubernetes"
	apps "k8s.io/client-go/kubernetes/typed/apps/v1"
	athn "k8s.io/client-go/kubernetes/typed/authentication/v1"
	athr "k8s.io/client-go/kubernetes/typed/authorization/v1"
	auto "k8s.io/client-go/kubernetes/typed/autoscaling/v1"
	bth "k8s.io/client-go/kubernetes/typed/batch/v1"
	cft "k8s.io/client-go/kubernetes/typed/certificates/v1"
	crd "k8s.io/client-go/kubernetes/typed/coordination/v1"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
	dsc1 "k8s.io/client-go/kubernetes/typed/discovery/v1"
	evnt "k8s.io/client-go/kubernetes/typed/events/v1"
	flc "k8s.io/client-go/kubernetes/typed/flowcontrol/v1"
	nwg "k8s.io/client-go/kubernetes/typed/networking/v1"
	nd "k8s.io/client-go/kubernetes/typed/node/v1"
	ply "k8s.io/client-go/kubernetes/typed/policy/v1"
	rbc "k8s.io/client-go/kubernetes/typed/rbac/v1"
	slg "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	stg "k8s.io/client-go/kubernetes/typed/storage/v1"
)

func CreateCoreV1Iface(clientset *kubernetes.Clientset) *core.CoreV1Interface {
	obj := clientset.CoreV1()
	return &obj
}

func CreateAppsV1Iface(clientset *kubernetes.Clientset) *apps.AppsV1Interface {
	obj := clientset.AppsV1()
	return &obj
}

func CreateAuthenticationV1Iface(clientset *kubernetes.Clientset) *athn.AuthenticationV1Interface {
	obj := clientset.AuthenticationV1()
	return &obj
}

func CreateAuthorizationV1Iface(clientset *kubernetes.Clientset) *athr.AuthorizationV1Interface {
	obj := clientset.AuthorizationV1()
	return &obj
}

func CreateAutoscalingV1Iface(clientset *kubernetes.Clientset) *auto.AutoscalingV1Interface {
	obj := clientset.AutoscalingV1()
	return &obj
}

func CreateBatchV1Iface(clientset *kubernetes.Clientset) *bth.BatchV1Interface {
	obj := clientset.BatchV1()
	return &obj
}

func CreateCertificatesV1Iface(clientset *kubernetes.Clientset) *cft.CertificatesV1Interface {
	obj := clientset.CertificatesV1()
	return &obj
}

func CreateCoordinationV1Iface(clientset *kubernetes.Clientset) *crd.CoordinationV1Interface {
	obj := clientset.CoordinationV1()
	return &obj
}

func CreateDiscoveryV1Iface(clientset *kubernetes.Clientset) *dsc1.DiscoveryV1Interface {
	obj := clientset.DiscoveryV1()
	return &obj
}

func CreateDiscoveryIface(clientset *kubernetes.Clientset) *discovery.DiscoveryInterface {
	obj := clientset.Discovery()
	return &obj
}

func CreateEventsV1Iface(clientset *kubernetes.Clientset) *evnt.EventsV1Interface {
	obj := clientset.EventsV1()
	return &obj
}

func CreateFlowcontrolV1Iface(clientset *kubernetes.Clientset) *flc.FlowcontrolV1Interface {
	obj := clientset.FlowcontrolV1()
	return &obj
}

func CreateNetworkingV1Iface(clientset *kubernetes.Clientset) *nwg.NetworkingV1Interface {
	obj := clientset.NetworkingV1()
	return &obj
}

func CreateNodeV1Iface(clientset *kubernetes.Clientset) *nd.NodeV1Interface {
	obj := clientset.NodeV1()
	return &obj
}

func CreatePolicyV1Iface(clientset *kubernetes.Clientset) *ply.PolicyV1Interface {
	obj := clientset.PolicyV1()
	return &obj
}

func CreateRbacV1Iface(clientset *kubernetes.Clientset) *rbc.RbacV1Interface {
	obj := clientset.RbacV1()
	return &obj
}

func CreateStorageV1Iface(clientset *kubernetes.Clientset) *stg.StorageV1Interface {
	obj := clientset.StorageV1()
	return &obj
}

func CreateSchedulingV1Iface(clientset *kubernetes.Clientset) *slg.SchedulingV1Interface {
	obj := clientset.SchedulingV1()
	return &obj
}
