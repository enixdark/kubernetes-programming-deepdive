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

	clientset "enixdark/kubernetes-programming-deepdive/Sample-controller/pkg/client/clientset/versioned"
	informers "enixdark/kubernetes-programming-deepdive/Sample-controller/pkg/client/informers/externalversions"
	"k8s.io/sample-controller/pkg/signals"
)