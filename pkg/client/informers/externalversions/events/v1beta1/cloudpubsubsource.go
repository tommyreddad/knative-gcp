/*
Copyright 2020 Google LLC

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

package v1beta1

import (
	"context"
	time "time"

	eventsv1beta1 "github.com/google/knative-gcp/pkg/apis/events/v1beta1"
	versioned "github.com/google/knative-gcp/pkg/client/clientset/versioned"
	internalinterfaces "github.com/google/knative-gcp/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/google/knative-gcp/pkg/client/listers/events/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CloudPubSubSourceInformer provides access to a shared informer and lister for
// CloudPubSubSources.
type CloudPubSubSourceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.CloudPubSubSourceLister
}

type cloudPubSubSourceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCloudPubSubSourceInformer constructs a new informer for CloudPubSubSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCloudPubSubSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCloudPubSubSourceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCloudPubSubSourceInformer constructs a new informer for CloudPubSubSource type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCloudPubSubSourceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1beta1().CloudPubSubSources(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EventsV1beta1().CloudPubSubSources(namespace).Watch(context.TODO(), options)
			},
		},
		&eventsv1beta1.CloudPubSubSource{},
		resyncPeriod,
		indexers,
	)
}

func (f *cloudPubSubSourceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCloudPubSubSourceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cloudPubSubSourceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eventsv1beta1.CloudPubSubSource{}, f.defaultInformer)
}

func (f *cloudPubSubSourceInformer) Lister() v1beta1.CloudPubSubSourceLister {
	return v1beta1.NewCloudPubSubSourceLister(f.Informer().GetIndexer())
}
