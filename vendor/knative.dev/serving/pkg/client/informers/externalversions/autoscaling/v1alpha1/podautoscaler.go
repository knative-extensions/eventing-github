/*
Copyright 2022 The Knative Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	apisautoscalingv1alpha1 "knative.dev/serving/pkg/apis/autoscaling/v1alpha1"
	versioned "knative.dev/serving/pkg/client/clientset/versioned"
	internalinterfaces "knative.dev/serving/pkg/client/informers/externalversions/internalinterfaces"
	autoscalingv1alpha1 "knative.dev/serving/pkg/client/listers/autoscaling/v1alpha1"
)

// PodAutoscalerInformer provides access to a shared informer and lister for
// PodAutoscalers.
type PodAutoscalerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() autoscalingv1alpha1.PodAutoscalerLister
}

type podAutoscalerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodAutoscalerInformer constructs a new informer for PodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodAutoscalerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodAutoscalerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodAutoscalerInformer constructs a new informer for PodAutoscaler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodAutoscalerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha1().PodAutoscalers(namespace).List(context.Background(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha1().PodAutoscalers(namespace).Watch(context.Background(), options)
			},
			ListWithContextFunc: func(ctx context.Context, options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha1().PodAutoscalers(namespace).List(ctx, options)
			},
			WatchFuncWithContext: func(ctx context.Context, options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AutoscalingV1alpha1().PodAutoscalers(namespace).Watch(ctx, options)
			},
		},
		&apisautoscalingv1alpha1.PodAutoscaler{},
		resyncPeriod,
		indexers,
	)
}

func (f *podAutoscalerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodAutoscalerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podAutoscalerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisautoscalingv1alpha1.PodAutoscaler{}, f.defaultInformer)
}

func (f *podAutoscalerInformer) Lister() autoscalingv1alpha1.PodAutoscalerLister {
	return autoscalingv1alpha1.NewPodAutoscalerLister(f.Informer().GetIndexer())
}
