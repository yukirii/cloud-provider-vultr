module github.com/yukirii/cloud-provider-vultr

go 1.13

replace k8s.io/api => k8s.io/api v0.0.0-20191114100352-16d7abae0d2a

replace k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20191114105449-027877536833

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20191028221656-72ed19daf4bb

replace k8s.io/apiserver => k8s.io/apiserver v0.0.0-20191114103151-9ca1dc586682

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20191114110141-0a35778df828

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20191114101535-6c5935290e33

replace k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20191114112024-4bbba8331835

replace k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20191114111741-81bb9acf592d

replace k8s.io/code-generator => k8s.io/code-generator v0.0.0-20191004115455-8e001e5d1894

replace k8s.io/component-base => k8s.io/component-base v0.0.0-20191114102325-35a9586014f7

replace k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190828162817-608eb1dad4ac

replace k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20191114112310-0da609c4ca2d

replace k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20191114103820-f023614fb9ea

replace k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20191114111510-6d1ed697a64b

replace k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20191114110717-50a77e50d7d9

replace k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20191114111229-2e90afcb56c7

replace k8s.io/kubectl => k8s.io/kubectl v0.0.0-20191114113550-6123e1c827f7

replace k8s.io/kubelet => k8s.io/kubelet v0.0.0-20191114110954-d67a8e7e2200

replace k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20191114112655-db9be3e678bb

replace k8s.io/metrics => k8s.io/metrics v0.0.0-20191114105837-a4a2842dc51b

replace k8s.io/node-api => k8s.io/node-api v0.0.0-20191114112948-fde05759caf8

replace k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20191114104439-68caf20693ac

replace k8s.io/sample-cli-plugin => k8s.io/sample-cli-plugin v0.0.0-20191114110435-31b16e91580f

replace k8s.io/sample-controller => k8s.io/sample-controller v0.0.0-20191114104921-b2770fad52e3

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2

require (
	github.com/JamesClonk/vultr v2.0.1+incompatible
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190717042225-c3de453c63f4 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.0 // indirect
	github.com/go-kit/kit v0.9.0 // indirect
	github.com/go-logfmt/logfmt v0.4.0 // indirect
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4 // indirect
	github.com/prometheus/procfs v0.0.5 // indirect
	golang.org/x/sys v0.0.0-20191010194322-b09406accb47 // indirect
	k8s.io/api v0.0.0
	k8s.io/apimachinery v0.0.0
	k8s.io/cloud-provider v0.0.0
	k8s.io/kubernetes v1.16.3 // indirect
	sigs.k8s.io/cluster-api v0.2.7
)
