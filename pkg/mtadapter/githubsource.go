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

	corev1 "k8s.io/api/core/v1"

	"knative.dev/pkg/controller"
	"knative.dev/pkg/reconciler"

	"knative.dev/eventing-github/pkg/apis/sources/v1alpha1"
	githubsourcereconciler "knative.dev/eventing-github/pkg/client/injection/reconciler/sources/v1alpha1/githubsource"
)

// Reconciler manages HTTP routes based on the GitHubSource objects in the
// controller's cache.
type Reconciler struct {
	mtAdapter MTAdapter
}

// Check that Reconciler implements reconciler.Interface.
var _ githubsourcereconciler.Interface = (*Reconciler)(nil)

// ReconcileKind implements reconciler.Interface.
func (r *Reconciler) ReconcileKind(ctx context.Context, src *v1alpha1.GitHubSource) reconciler.Event {
	if !src.Status.IsReady() {
		// Mark that error as permanent so we don't retry until the
		// source's status has been updated, which automatically
		// triggers a new reconciliation.
		return controller.NewPermanentError(reconciler.NewEvent(corev1.EventTypeWarning, "NotReady",
			"GitHubSource is not ready yet. Skipping adapter configuration"))
	}

	return r.mtAdapter.RegisterHandlerFor(ctx, src)
}
