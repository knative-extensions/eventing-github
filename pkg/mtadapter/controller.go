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

	"knative.dev/eventing/pkg/adapter/v2"
	"knative.dev/pkg/controller"
	"knative.dev/pkg/logging"

	"knative.dev/eventing-github/pkg/apis/sources/v1alpha1"
	githubsourceinformer "knative.dev/eventing-github/pkg/client/injection/informers/sources/v1alpha1/githubsource"
	githubsourcereconciler "knative.dev/eventing-github/pkg/client/injection/reconciler/sources/v1alpha1/githubsource"
)

// MTAdapter allows the multi-tenant adapter to expose methods the reconciler
// can call while reconciling a source object.
type MTAdapter interface {
	// Registers a HTTP handler for the given source.
	RegisterHandlerFor(context.Context, *v1alpha1.GitHubSource) error
}

// NewController returns a constructor for the event source's Reconciler.
// This constructor initializes the controller and registers event handlers to
// enqueue events.
func NewController(component string) adapter.ControllerConstructor {
	return func(ctx context.Context, a adapter.Adapter) *controller.Impl {
		r := &Reconciler{
			mtAdapter: a.(MTAdapter),
		}

		impl := githubsourcereconciler.NewImpl(ctx, r, controllerOpts(component))

		logging.FromContext(ctx).Info("Setting up event handlers")

		// Watch for githubsource objects
		githubsourceInformer := githubsourceinformer.Get(ctx)
		githubsourceInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

		return impl
	}
}

// controllerOpts returns a callback function that sets the controller's agent
// name and configures the reconciler to skip status updates.
func controllerOpts(component string) controller.OptionsFn {
	return func(impl *controller.Impl) controller.Options {
		return controller.Options{
			AgentName:         component,
			SkipStatusUpdates: true,
		}
	}
}
