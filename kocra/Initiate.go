package kocra

import (
	"fmt"
	"os"

	utl "github.com/knagadevara/GoUtility"
)

func GetKCpath() string {
	pathKubeConfig := os.Getenv("KUBECONFIG")
	if pathKubeConfig == "" {
		fmt.Println("Please enter the path to config file.")
		rdr := utl.GetNewStdInRdr()
		kcfgPath := utl.GetString()(rdr)
		return kcfgPath
	}
	return pathKubeConfig
}

func GetKubeConfig() *KubeConfig {
	kcPath := GetKCpath()
	ymlBuf := utl.OperateFile(kcPath, os.O_RDONLY, 0644)
	kubeConfig := utl.RdYamlFileToStruct[KubeConfig](ymlBuf)
	return &kubeConfig
}
