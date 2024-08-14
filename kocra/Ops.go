package kocra

import (
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func MakeMetaListLabels(strArray ...string) *metav1.ListOptions {
	var metaLstOps metav1.ListOptions
	metaLstOps.LabelSelector = strings.Join(strArray, ",")
	return &metaLstOps
}
