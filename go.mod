module github.com/rccrdpccl/bindings

go 1.14

require (
	k8s.io/api v0.18.8
	k8s.io/apimachinery v0.18.8
	k8s.io/autoscaler/vertical-pod-autoscaler v0.0.0-20200910092546-63259fb5dd89
	k8s.io/client-go v11.0.1-0.20190805182717-6502b5e7b1b5+incompatible
	k8s.io/code-generator v0.18.8
	k8s.io/kube-openapi v0.0.0-20200410145947-bcb3869e6f29
	knative.dev/pkg v0.0.0-20200911235400-de640e81d149
	knative.dev/test-infra v0.0.0-20200911201000-3f90e7c8f2fa
)

replace (
	github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2

	k8s.io/api => k8s.io/api v0.18.3
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.18.3
	k8s.io/apimachinery => k8s.io/apimachinery v0.18.3

	k8s.io/client-go => k8s.io/client-go v0.18.3
	k8s.io/code-generator => k8s.io/code-generator v0.18.3
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190816220812-743ec37842bf
)
