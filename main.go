package main

import (
	"os"

	utl "github.com/knagadevara/GoUtility"
	"github.com/knagadevara/kocra/kocra"
)

func main() {
	kcfgPath := os.Getenv("HOME") + os.Getenv("KUBECONFIG")
	flbuf := utl.LoadFile(kcfgPath)
	kocra.Initiate(flbuf)
}
