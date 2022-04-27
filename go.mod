module knative.dev/eventing-github

go 1.16

require (
	cloud.google.com/go/iam v0.2.0 // indirect
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/google/go-cmp v0.5.7
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
	golang.org/x/oauth2 v0.0.0-20220223155221-ee480838109b
	gonum.org/v1/gonum v0.0.0-20190331200053-3d26580ed485 // indirect
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.23.5
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.23.5
	knative.dev/eventing v0.31.1-0.20220427080352-3abb43601181
	knative.dev/hack v0.0.0-20220427014036-5f473869d377
	knative.dev/pkg v0.0.0-20220427013826-1f681e126af6
	knative.dev/serving v0.31.1-0.20220427014148-6b3905bfa89f
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
