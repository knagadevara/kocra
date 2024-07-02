package kocra

import (
	"fmt"

	utl "github.com/knagadevara/GoUtility"
)

func GetYmlFile() {
	fmt.Println("Please enter the path to config file.")
	rdr := utl.GetNewStdInRdr()
	kcfgPath := utl.GetString()(rdr)
	ymlBuf := utl.LoadFile(kcfgPath)
	kubeConfig := utl.YmlToStruct[KubeConfig](ymlBuf)
	fmt.Println(kubeConfig)
}
