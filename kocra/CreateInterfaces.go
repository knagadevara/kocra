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
	dsc "k8s.io/client-go/kubernetes/typed/discovery/v1"
	evnt "k8s.io/client-go/kubernetes/typed/events/v1"
	flc "k8s.io/client-go/kubernetes/typed/flowcontrol/v1"
	nwg "k8s.io/client-go/kubernetes/typed/networking/v1"
	nd "k8s.io/client-go/kubernetes/typed/node/v1"
	ply "k8s.io/client-go/kubernetes/typed/policy/v1"
	rbc "k8s.io/client-go/kubernetes/typed/rbac/v1"
	sch "k8s.io/client-go/kubernetes/typed/scheduling/v1"
	stg "k8s.io/client-go/kubernetes/typed/storage/v1"
	"k8s.io/client-go/rest"
)

func IfcCoreV1(clientset *kubernetes.Clientset) *core.CoreV1Interface {
	obj := clientset.CoreV1()
	return &obj
}

func IfcAppsV1(clientset *kubernetes.Clientset) *apps.AppsV1Interface {
	obj := clientset.AppsV1()
	return &obj
}

func IfcAuthenticationV1(clientset *kubernetes.Clientset) *athn.AuthenticationV1Interface {
	obj := clientset.AuthenticationV1()
	return &obj
}

func IfcAuthorizationV1(clientset *kubernetes.Clientset) *athr.AuthorizationV1Interface {
	obj := clientset.AuthorizationV1()
	return &obj
}

func IfcAutoscalingV1(clientset *kubernetes.Clientset) *auto.AutoscalingV1Interface {
	obj := clientset.AutoscalingV1()
	return &obj
}

func IfcBatchV1(clientset *kubernetes.Clientset) *bth.BatchV1Interface {
	obj := clientset.BatchV1()
	return &obj
}

func IfcCertificatesV1(clientset *kubernetes.Clientset) *cft.CertificatesV1Interface {
	obj := clientset.CertificatesV1()
	return &obj
}

func IfcCoordinationV1(clientset *kubernetes.Clientset) *crd.CoordinationV1Interface {
	obj := clientset.CoordinationV1()
	return &obj
}

func IfcDiscoveryV1(clientset *kubernetes.Clientset) *dsc.DiscoveryV1Interface {
	obj := clientset.DiscoveryV1()
	return &obj
}

func IfcDiscovery(clientset *kubernetes.Clientset) *discovery.DiscoveryInterface {
	obj := clientset.Discovery()
	return &obj
}

func IfcEventsV1(clientset *kubernetes.Clientset) *evnt.EventsV1Interface {
	obj := clientset.EventsV1()
	return &obj
}

func IfcFlowcontrolV1(clientset *kubernetes.Clientset) *flc.FlowcontrolV1Interface {
	obj := clientset.FlowcontrolV1()
	return &obj
}

func IfcNetworkingV1(clientset *kubernetes.Clientset) *nwg.NetworkingV1Interface {
	obj := clientset.NetworkingV1()
	return &obj
}

func IfcNodeV1(clientset *kubernetes.Clientset) *nd.NodeV1Interface {
	obj := clientset.NodeV1()
	return &obj
}

func IfcPolicyV1(clientset *kubernetes.Clientset) *ply.PolicyV1Interface {
	obj := clientset.PolicyV1()
	return &obj
}

func IfcRbacV1(clientset *kubernetes.Clientset) *rbc.RbacV1Interface {
	obj := clientset.RbacV1()
	return &obj
}

func IfcStorageV1(clientset *kubernetes.Clientset) *stg.StorageV1Interface {
	obj := clientset.StorageV1()
	return &obj
}

func IfcSchedulingV1(clientset *kubernetes.Clientset) *sch.SchedulingV1Interface {
	obj := clientset.SchedulingV1()
	return &obj
}

func IfcComponentStatus(cv1 core.CoreV1Interface) *core.ComponentStatusInterface {
	obj := cv1.ComponentStatuses()
	return &obj
}

func IfcNamespace(cv1 core.CoreV1Interface) *core.NamespaceInterface {
	obj := cv1.Namespaces()
	return &obj
}

func IfcNode(cv1 core.CoreV1Interface) *core.NodeInterface {
	obj := cv1.Nodes()
	return &obj
}

func IfcPersistentVolume(cv1 core.CoreV1Interface) *core.PersistentVolumeInterface {
	obj := cv1.PersistentVolumes()
	return &obj
}

func IfcEndpoints(cv1 core.CoreV1Interface, namespace string) *core.EndpointsInterface {
	obj := cv1.Endpoints(namespace)
	return &obj
}

func IfcEvents(cv1 core.CoreV1Interface, namespace string) *core.EventInterface {
	obj := cv1.Events(namespace)
	return &obj
}

func IfcLimitRange(cv1 core.CoreV1Interface, namespace string) *core.LimitRangeInterface {
	obj := cv1.LimitRanges(namespace)
	return &obj
}

func IfcPersistentVolumeClaims(cv1 core.CoreV1Interface, namespace string) *core.PersistentVolumeClaimInterface {
	obj := cv1.PersistentVolumeClaims(namespace)
	return &obj
}

func IfcPodTemplate(cv1 core.CoreV1Interface, namespace string) *core.PodTemplateInterface {
	obj := cv1.PodTemplates(namespace)
	return &obj
}

func IfcPod(cv1 core.CoreV1Interface, namespace string) *core.PodInterface {
	obj := cv1.Pods(namespace)
	return &obj
}

func IfcReplicationController(cv1 core.CoreV1Interface, namespace string) *core.ReplicationControllerInterface {
	obj := cv1.ReplicationControllers(namespace)
	return &obj
}

func IfcResourceQuota(cv1 core.CoreV1Interface, namespace string) *core.ResourceQuotaInterface {
	obj := cv1.ResourceQuotas(namespace)
	return &obj
}

func IfcSecret(cv1 core.CoreV1Interface, namespace string) *core.SecretInterface {
	obj := cv1.Secrets(namespace)
	return &obj
}

func IfcServiceAccount(cv1 core.CoreV1Interface, namespace string) *core.ServiceAccountInterface {
	obj := cv1.ServiceAccounts(namespace)
	return &obj
}

func IfcServices(cv1 core.CoreV1Interface, namespace string) *core.ServiceInterface {
	obj := cv1.Services(namespace)
	return &obj
}

func IfcCoreRESTClient(cv1 core.CoreV1Interface) *rest.Interface {
	obj := cv1.RESTClient()
	return &obj
}

func IfcControllerRevision(av1 apps.AppsV1Client, namespace string) *apps.ControllerRevisionInterface {
	obj := av1.ControllerRevisions(namespace)
	return &obj
}

func IfcDaemonSet(av1 apps.AppsV1Client, namespace string) *apps.DaemonSetInterface {
	obj := av1.DaemonSets(namespace)
	return &obj
}

func IfcDeployment(av1 apps.AppsV1Client, namespace string) *apps.DeploymentInterface {
	obj := av1.Deployments(namespace)
	return &obj
}

func IfcReplicaSet(av1 apps.AppsV1Client, namespace string) *apps.ReplicaSetInterface {
	obj := av1.ReplicaSets(namespace)
	return &obj
}

func IfcStatefulSet(av1 apps.AppsV1Client, namespace string) *apps.StatefulSetInterface {
	obj := av1.StatefulSets(namespace)
	return &obj
}

func IfcAppRESTClient(av1 apps.AppsV1Client) *rest.Interface {
	obj := av1.RESTClient()
	return &obj
}

func IfcAutoHorizontalPodAutoscalers(autoV1 auto.AutoscalingV1Interface, namespace string) *auto.HorizontalPodAutoscalerInterface {
	obj := autoV1.HorizontalPodAutoscalers(namespace)
	return &obj
}

func IfcBatchCronJobs(bthV1 bth.BatchV1Interface, namespace string) *bth.CronJobInterface {
	obj := bthV1.CronJobs(namespace)
	return &obj
}

func IfcBatchJobs(bthV1 bth.BatchV1Interface, namespace string) *bth.JobInterface {
	obj := bthV1.Jobs(namespace)
	return &obj
}

func IfcAthnSelfSubjectReviews(athnV1 athn.AuthenticationV1Client) *athn.SelfSubjectReviewInterface {
	obj := athnV1.SelfSubjectReviews()
	return &obj
}

func IfcAthnTokenReviews(athnV1 athn.AuthenticationV1Client) *athn.TokenReviewInterface {
	obj := athnV1.TokenReviews()
	return &obj
}

func IfcAthrSelfSubjectReviews(athrV1 athr.AuthorizationV1Interface) *athr.SelfSubjectAccessReviewInterface {
	obj := athrV1.SelfSubjectAccessReviews()
	return &obj
}

func IfcAthrSelfSubjectRulesReviews(athrV1 athr.AuthorizationV1Interface) *athr.SelfSubjectRulesReviewInterface {
	obj := athrV1.SelfSubjectRulesReviews()
	return &obj
}

func IfcAthrSubjectAccessReviews(athrV1 athr.AuthorizationV1Interface) *athr.SubjectAccessReviewInterface {
	obj := athrV1.SubjectAccessReviews()
	return &obj
}

func IfcAthrLocalSubjectAccessReviews(athrV1 athr.AuthorizationV1Interface, namespace string) *athr.LocalSubjectAccessReviewInterface {
	obj := athrV1.LocalSubjectAccessReviews(namespace)
	return &obj
}

func IfcCertificateSigningRequest(cftV1 cft.CertificatesV1Interface) *cft.CertificateSigningRequestInterface {
	obj := cftV1.CertificateSigningRequests()
	return &obj
}

func IfcLeases(crdV1 crd.CoordinationV1Interface, namespace string) *crd.LeaseInterface {
	obj := crdV1.Leases(namespace)
	return &obj
}

func IfcEndpointSlices(dscV1 dsc.DiscoveryV1Interface, namespace string) *dsc.EndpointSliceInterface {
	obj := dscV1.EndpointSlices(namespace)
	return &obj
}

func IfcPriorityLevelConfigurations(flcV1 flc.FlowcontrolV1Interface) *flc.PriorityLevelConfigurationInterface {
	obj := flcV1.PriorityLevelConfigurations()
	return &obj
}

func IfcFlowSchemas(flcV1 flc.FlowcontrolV1Interface) *flc.FlowSchemaInterface {
	obj := flcV1.FlowSchemas()
	return &obj
}

func IfcNetIngressClasses(nwgV1 nwg.NetworkingV1Interface) *nwg.IngressClassInterface {
	obj := nwgV1.IngressClasses()
	return &obj
}

func IfcNetIngresses(nwgV1 nwg.NetworkingV1Interface, namespace string) *nwg.IngressInterface {
	obj := nwgV1.Ingresses(namespace)
	return &obj
}

func IfcNetworkPolicies(nwgV1 nwg.NetworkingV1Interface, namespace string) *nwg.NetworkPolicyInterface {
	obj := nwgV1.NetworkPolicies(namespace)
	return &obj
}

func IfcNodeRuntimeClasses(ndV1 nd.NodeV1Interface) *nd.RuntimeClassInterface {
	obj := ndV1.RuntimeClasses()
	return &obj
}

func IfcPolicyEvictions(plyV1 ply.PolicyV1Interface, namespace string) *ply.EvictionInterface {
	obj := plyV1.Evictions(namespace)
	return &obj
}

func IfcPolicyPodDisruptionBudgets(plyV1 ply.PolicyV1Interface, namespace string) *ply.PodDisruptionBudgetInterface {
	obj := plyV1.PodDisruptionBudgets(namespace)
	return &obj
}

func IfcRBCRoleBindings(rbcV1 rbc.RbacV1Interface, namespace string) *rbc.RoleBindingInterface {
	obj := rbcV1.RoleBindings(namespace)
	return &obj
}

func IfcRBCRoles(rbcV1 rbc.RbacV1Interface, namespace string) *rbc.RoleInterface {
	obj := rbcV1.Roles(namespace)
	return &obj
}

func IfcRBCClusterRoleBindings(rbcV1 rbc.RbacV1Interface) *rbc.ClusterRoleBindingInterface {
	obj := rbcV1.ClusterRoleBindings()
	return &obj
}

func IfcRBCClusterRoles(rbcV1 rbc.RbacV1Interface) *rbc.ClusterRoleInterface {
	obj := rbcV1.ClusterRoles()
	return &obj
}

func IfcSchPriorityClasses(schV1 sch.SchedulingV1Interface) *sch.PriorityClassInterface {
	obj := schV1.PriorityClasses()
	return &obj
}

func IfcStgCSIDriver(stgV1 stg.StorageV1Interface) *stg.CSIDriverInterface {
	obj := stgV1.CSIDrivers()
	return &obj
}

func IfcStgCSINode(stgV1 stg.StorageV1Interface) *stg.CSINodeInterface {
	obj := stgV1.CSINodes()
	return &obj
}

func IfcStorageClass(stgV1 stg.StorageV1Interface) *stg.StorageClassInterface {
	obj := stgV1.StorageClasses()
	return &obj
}

func IfcVolumeAttachment(stgV1 stg.StorageV1Interface) *stg.VolumeAttachmentInterface {
	obj := stgV1.VolumeAttachments()
	return &obj
}

func IfcCSIStorageCapacity(stgV1 stg.StorageV1Interface, namespace string) *stg.CSIStorageCapacityInterface {
	obj := stgV1.CSIStorageCapacities(namespace)
	return &obj
}
