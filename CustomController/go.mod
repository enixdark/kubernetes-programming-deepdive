module github.com/enixdark/kubernetes-programming-deepdive/CustomController

go 1.12

require (
	k8s.io/api v0.0.0-20190425012535-181e1f9c52c1
	k8s.io/apimachinery v0.0.0-20190425132440-17f84483f500
	k8s.io/client-go v0.0.0-20190425172711-65184652c889
	k8s.io/code-generator v0.0.0-20190419212335-ff26e7842f9d
	k8s.io/klog v0.3.0
)

replace (
	golang.org/x/sync => golang.org/x/sync v0.0.0-20181108010431-42b317875d0f
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190209173611-3b5209105503
	golang.org/x/tools => golang.org/x/tools v0.0.0-20190313210603-aa82965741a9
	k8s.io/api => k8s.io/api v0.0.0-20190425012535-181e1f9c52c1
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190425132440-17f84483f500
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190425172711-65184652c889
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190419212335-ff26e7842f9d
)
