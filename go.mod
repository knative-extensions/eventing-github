module knative.dev/eventing-github

go 1.16

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/google/go-cmp v0.5.6
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.0
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.21.4
	k8s.io/apimachinery v0.21.4
	k8s.io/client-go v0.21.4
	knative.dev/eventing v0.26.0
	knative.dev/hack v0.0.0-20210806075220-815cd312d65c
	knative.dev/pkg v0.0.0-20210919202233-5ae482141474
	knative.dev/serving v0.26.0
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
