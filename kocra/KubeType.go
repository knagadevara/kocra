package kocra

type KubeUser struct {
	Name string `yaml:"name"`
	User struct {
		Token string `yaml:"token"`
	} `yaml:"user"`
}

type KubeContext struct {
	ContextName  string `yaml:"name"`
	ClustDetails struct {
		ClusterName string `yaml:"cluster"`
		NameSpace   string `yaml:"namespace"`
		UserName    string `yaml:"user"`
	} `yaml:"context"`
}

type KubeCluster struct {
	ClusterName  string `yaml:"name"`
	ClustDetails struct {
		CA     string `yaml:"certificate-authority-data"`
		Server string `yaml:"server"`
	} `yaml:"cluster"`
}

type KubeConfig struct {
	APIVersion     string        `yaml:"apiVersion"`
	Kind           string        `yaml:"kind"`
	CurrentContext string        `yaml:"current-context"`
	Preferences    struct{}      `yaml:"preferences"`
	Clusters       []KubeCluster `yaml:"clusters"`
	Contexts       []KubeContext `yaml:"contexts"`
	Users          []KubeUser    `yaml:"users"`
}

type KubeConfigI interface {
	GetKubeConfig() *KubeConfig
}

type AllClusterContext []KubeConfig
