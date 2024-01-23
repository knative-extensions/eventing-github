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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
	clientset "knative.dev/eventing/pkg/client/clientset/versioned"
	eventingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1"
	fakeeventingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1/fake"
	eventingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta1"
	fakeeventingv1beta1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta1/fake"
	eventingv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta2"
	fakeeventingv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta2/fake"
	eventingv1beta3 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta3"
	fakeeventingv1beta3 "knative.dev/eventing/pkg/client/clientset/versioned/typed/eventing/v1beta3/fake"
	flowsv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1"
	fakeflowsv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/flows/v1/fake"
	messagingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1"
	fakemessagingv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/messaging/v1/fake"
	sourcesv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1"
	fakesourcesv1 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1/fake"
	sourcesv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1beta2"
	fakesourcesv1beta2 "knative.dev/eventing/pkg/client/clientset/versioned/typed/sources/v1beta2/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// EventingV1beta1 retrieves the EventingV1beta1Client
func (c *Clientset) EventingV1beta1() eventingv1beta1.EventingV1beta1Interface {
	return &fakeeventingv1beta1.FakeEventingV1beta1{Fake: &c.Fake}
}

// EventingV1beta2 retrieves the EventingV1beta2Client
func (c *Clientset) EventingV1beta2() eventingv1beta2.EventingV1beta2Interface {
	return &fakeeventingv1beta2.FakeEventingV1beta2{Fake: &c.Fake}
}

// EventingV1beta3 retrieves the EventingV1beta3Client
func (c *Clientset) EventingV1beta3() eventingv1beta3.EventingV1beta3Interface {
	return &fakeeventingv1beta3.FakeEventingV1beta3{Fake: &c.Fake}
}

// EventingV1 retrieves the EventingV1Client
func (c *Clientset) EventingV1() eventingv1.EventingV1Interface {
	return &fakeeventingv1.FakeEventingV1{Fake: &c.Fake}
}

// FlowsV1 retrieves the FlowsV1Client
func (c *Clientset) FlowsV1() flowsv1.FlowsV1Interface {
	return &fakeflowsv1.FakeFlowsV1{Fake: &c.Fake}
}

// MessagingV1 retrieves the MessagingV1Client
func (c *Clientset) MessagingV1() messagingv1.MessagingV1Interface {
	return &fakemessagingv1.FakeMessagingV1{Fake: &c.Fake}
}

// SourcesV1beta2 retrieves the SourcesV1beta2Client
func (c *Clientset) SourcesV1beta2() sourcesv1beta2.SourcesV1beta2Interface {
	return &fakesourcesv1beta2.FakeSourcesV1beta2{Fake: &c.Fake}
}

// SourcesV1 retrieves the SourcesV1Client
func (c *Clientset) SourcesV1() sourcesv1.SourcesV1Interface {
	return &fakesourcesv1.FakeSourcesV1{Fake: &c.Fake}
}
