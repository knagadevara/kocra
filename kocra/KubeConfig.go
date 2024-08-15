package kocra

import (
	"fmt"
	"log"
	"os"

	utl "github.com/knagadevara/GoUtility"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetKCpath() string {
	pathKubeConfig := os.Getenv("KUBECONFIG")
	if pathKubeConfig == "" {
		fmt.Println("Please enter Absolute Path to config file.")
		rdr := utl.GetNewStdInRdr()
		kcfgPath := utl.GetString()(rdr)
		return kcfgPath
	}
	return pathKubeConfig
}

func GetKubeConfig() *rest.Config {
	kcPath := GetKCpath()
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", kcPath)
	if err != nil {
		if clientcmd.IsConfigurationInvalid(err) {
			log.Fatalln(err)
			return nil
		}
	}
	return kubeConfig
}

func ForgeClient(kc *rest.Config) *kubernetes.Clientset {
	clientset, err := kubernetes.NewForConfig(kc)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	return clientset
}
