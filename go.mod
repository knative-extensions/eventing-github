module knative.dev/eventing-github

go 1.15

require (
	github.com/cloudevents/sdk-go/v2 v2.2.0
	github.com/google/go-cmp v0.5.4
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.6.1
	go.uber.org/zap v1.16.0
	golang.org/x/oauth2 v0.0.0-20210126194326-f9ce19ea3013
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.19.7
	k8s.io/apimachinery v0.19.7
	k8s.io/client-go v0.19.7
	knative.dev/eventing v0.21.1-0.20210309092525-37e702765dbc
	knative.dev/hack v0.0.0-20210305150220-f99a25560134
	knative.dev/pkg v0.0.0-20210309024624-0f8d8de5949d
	knative.dev/serving v0.21.1-0.20210309024225-21868c65649c
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
