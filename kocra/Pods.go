package kocra

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func ListPods(p core.PodInterface, metaLstOps metav1.ListOptions) *v1.PodList {
	pods, err := p.List(context.TODO(), metaLstOps)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return pods
}

func GetPodDetails(p core.PodInterface, podName string) *v1.Pod {
	podDetails, err := p.Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return podDetails
}
