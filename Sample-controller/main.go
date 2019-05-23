package main

import (
	"flag"
	"time"

	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/enixdark/kubernetes-programming-deepdive/Sample-controller/pkg/client/clientset/versioned"
	informers "github.com/enixdark/kubernetes-programming-deepdive/Sample-controller/pkg/client/informers/externalversions"
	"github.com/enixdark/kubernetes-programming-deepdive/Sample-controller/pkg/signals"
)

var (
	masterURL string
	kubeconfig string
)

func main() {
	flag.Parse()

	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)

	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	exampleClient, err := clientset.NewForConfig(cfg)

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second * 30)
	exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)

	controller := NewController(
		kubeClient,
		exampleClient,
		kubeInformerFactory.Apps().V1().Deployments(),
		exampleInformerFactory.SampleController().V1alpha1().Foos()
	)

	kubeInformerFactory.Start(stopCh)
	exampleInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}