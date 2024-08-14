package main

import "github.com/knagadevara/kocra/kocra"

func main() {
	kubeConfig := kocra.GetKubeConfig()
	clientset := kocra.ForgeClient(kubeConfig)
	cv1 := kocra.CreateCoreV1Iface(clientset)
	podIface := kocra.CreatePodIface(*cv1, "kube-system")
	kocra.GetEtcdState(podIface)
}
