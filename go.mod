module knative.dev/eventing-github

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/google/go-cmp v0.5.5
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20210402161424-2e8d93401602
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/eventing v0.22.1-0.20210415222903-e74573361fe6
	knative.dev/hack v0.0.0-20210325223819-b6ab329907d3
	knative.dev/pkg v0.0.0-20210416161310-b80a1926251c
	knative.dev/serving v0.22.1-0.20210417191805-44f60331bbfd
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
