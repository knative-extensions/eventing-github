module knative.dev/eventing-github

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.4.1
	github.com/google/go-cmp v0.5.6
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.17.0
	golang.org/x/oauth2 v0.0.0-20210514164344-f6687ab2804c
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.20.7
	k8s.io/apimachinery v0.20.7
	k8s.io/client-go v0.20.7
	knative.dev/eventing v0.24.2
	knative.dev/hack v0.0.0-20210622141627-e28525d8d260
	knative.dev/pkg v0.0.0-20210902173607-953af0138c75
	knative.dev/serving v0.24.0
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
