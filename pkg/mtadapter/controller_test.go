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
	"testing"

	"github.com/stretchr/testify/assert"
	rt "knative.dev/pkg/reconciler/testing"

	// Fake injection clients and informers
	_ "knative.dev/eventing-github/pkg/client/injection/informers/sources/v1alpha1/githubsource/fake"
	_ "knative.dev/eventing/pkg/client/injection/client/fake"
	_ "knative.dev/pkg/client/injection/kube/client/fake"
)

func TestNew(t *testing.T) {
	const testComponent = "test_component"

	ctx, _ := rt.SetupFakeContext(t)

	c := NewController(testComponent)(ctx, &gitHubAdapter{})
	assert.NotNil(t, c)
}
