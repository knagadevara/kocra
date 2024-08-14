package kocra

import (
	"fmt"

	v1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

func GetEtcdState(podIface *v1.PodInterface) {

	mv1LOpts := MakeMetaListLabels("component=etcd", "tier=control-plane")
	for _, pod := range ListPods(*podIface, *mv1LOpts).Items {
		podDetails := GetPodDetails(*podIface, pod.Name)
		fmt.Printf("Name: %v\n", podDetails.Name)
		fmt.Printf("Status: %v\n", podDetails.Status)
		fmt.Println()
	}
}
