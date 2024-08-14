package kocra

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	core "k8s.io/client-go/kubernetes/typed/core/v1"
)

func ListNamespace(n core.NamespaceInterface) *v1.NamespaceList {
	namespace, err := n.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return namespace
}

func GetNamespaceDetails(n core.NamespaceInterface, namespace string) *v1.Namespace {
	namespaceDetails, err := n.Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return namespaceDetails
}
