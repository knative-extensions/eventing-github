/*
Copyright 2021 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package mtadapter

import (
	"context"
	"fmt"
	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"go.uber.org/zap"

	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/logging"

	githubsourceinformer "knative.dev/eventing-github/pkg/client/injection/informers/sources/v1alpha1/githubsource"
	"knative.dev/eventing-github/pkg/common"
	"knative.dev/eventing-github/pkg/mtadapter/router"
)

// envConfig contains a set of configuration values injected via environment variables.
type envConfig struct {
	adapter.EnvConfig

	// Environment variable containing the number of the HTTP port the
	// event handler listens on.
	EnvPort uint16 `envconfig:"PORT" default:"8080"`
}

// NewEnvConfig returns an accessor for the adapter's envConfig.
func NewEnvConfig() adapter.EnvConfigAccessor {
	return &envConfig{}
}

// gitHubAdapter converts incoming GitHub webhook events to CloudEvents.
type gitHubAdapter struct {
	logger *zap.SugaredLogger

	ceClient cloudevents.Client
	port     uint16
	router   *router.Router
}

// NewAdapter is a constructor for a GitHubSource receive adapter.
// It satisfies adapter.AdapterConstructor.
func NewAdapter(ctx context.Context, processed adapter.EnvConfigAccessor, ceClient cloudevents.Client) adapter.Adapter {
	logger := logging.FromContext(ctx)
	env := processed.(*envConfig)

	lister := githubsourceinformer.Get(ctx).Lister()
	router := router.New(logger, lister)

	return &gitHubAdapter{
		logger:   logger,
		ceClient: ceClient,
		port:     env.EnvPort,

		router: router,
	}
}

// Start implements adapter.Adapter.
func (a *gitHubAdapter) Start(ctx context.Context) error {
	// Start our multi-tenant server receiving GitHub events
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.port),
		Handler: a.router,
	}

	done := make(chan bool, 1)
	go common.GracefulShutdown(server, a.logger, ctx.Done(), done)

	a.logger.Infof("Server is ready to handle requests at %s", server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("could not listen on %s: %v", server.Addr, err)
	}

	<-done
	a.logger.Infof("Server stopped")
	return nil
}
