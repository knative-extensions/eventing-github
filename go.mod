module knative.dev/eventing-github

go 1.16

require (
	github.com/cloudevents/sdk-go/v2 v2.8.0
	github.com/emicklei/go-restful v2.15.0+incompatible // indirect
	github.com/google/go-cmp v0.5.6
	github.com/google/go-github/v31 v31.0.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.uber.org/zap v1.19.1
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	gonum.org/v1/gonum v0.0.0-20190331200053-3d26580ed485 // indirect
	gopkg.in/go-playground/webhooks.v5 v5.13.0
	k8s.io/api v0.22.5
	k8s.io/apimachinery v0.22.5
	k8s.io/client-go v0.22.5
	knative.dev/eventing v0.29.1-0.20220203104220-17487954d41a
	knative.dev/hack v0.0.0-20220201013531-82bfca153560
	knative.dev/pkg v0.0.0-20220203020920-51be315ed160
	knative.dev/serving v0.29.1-0.20220203055220-1d9529495e17
)

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v0.9.2
