package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/golang/glog"
	"github.com/enixdark/kubernetes-programming-deepdive/CustomResourceDefinition/pkg/apis/app/v1alpha1"
	apiextension "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	// clientcmdapi "k8s.io/client-go/tools/clientcmd/api"
)

var (
	// Set during build
	version string
	kubeconfig string
	proxyURL = flag.String("proxy", "",
		`If specified, it is assumed that a kubctl proxy server is running on the
		given url and creates a proxy client. In case it is not given InCluster
		kubernetes setup will be used`)
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "path to Kubernetes config file")
	flag.Parse()
}

func main() {

	var err error

	var config *rest.Config

	if kubeconfig == "" {
		fmt.Printf("using in-cluster configuration")
		config, err = rest.InClusterConfig()
	} else {
		fmt.Printf("using configuration from '%s'", kubeconfig)
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	 
	kubeclient, err := apiextension.NewForConfig(config)
	if err != nil {
		glog.Fatalf("Failed to create client: %v", err)
	}

	err = v1alpha1.CreateCRD(kubeclient)
	if err != nil {
		glog.Fatalf("Failed to create crd: %v", err)
	}

	time.Sleep(5 * time.Second)

	crdclient, err := v1alpha1.NewClient(config)

	if err != nil {
		panic(err)
	}

	SslConfig := &v1alpha1.SslConfig{
		ObjectMeta: meta_v1.ObjectMeta{
			Name:   "sslconfigobj",
			Labels: map[string]string{"mylabel": "crd"},
			GenerateName: "sslconfigobj",
		},
		Spec: v1alpha1.SslConfigSpec{
			Cert: "my-cert",
			Key: "my-key",
			Domain: "*.velotio.com",
		},
		Status: v1alpha1.SslConfigStatus{
			State: "created",
			Message: "Created, not processed yet",
		},
	}

	resp, err := crdclient.SslConfigs("default").Create(SslConfig)

	if err != nil {
		fmt.Printf("error while creating object: %v\n", err)
	} else {
		fmt.Printf("object created: %v\n", resp)
	}

	obj, err := crdclient.SslConfigs("default").Get(SslConfig.ObjectMeta.Name)
	if err != nil {
		glog.Infof("error while getting the object %v\n", err)
	}
	fmt.Printf("SslConfig Objects Found: \n%v\n", obj)
	select {}
}