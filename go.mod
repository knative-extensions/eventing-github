module knative.dev/eventing-github

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/google/go-cmp v0.5.4
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20201208152858-08078c50e5b5
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/eventing v0.20.1-0.20210126110930-81ce89802f86
	knative.dev/hack v0.0.0-20210120165453-8d623a0af457
	knative.dev/pkg v0.0.0-20210125222030-6040b3af4803
	knative.dev/serving v0.20.1-0.20210122011153-d0a0e3827242
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
