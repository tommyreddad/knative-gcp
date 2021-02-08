/*
Copyright 2021 Google LLC

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/google/knative-gcp/pkg/apis/intevents/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PullSubscriptionLister helps list PullSubscriptions.
// All objects returned here must be treated as read-only.
type PullSubscriptionLister interface {
	// List lists all PullSubscriptions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PullSubscription, err error)
	// PullSubscriptions returns an object that can list and get PullSubscriptions.
	PullSubscriptions(namespace string) PullSubscriptionNamespaceLister
	PullSubscriptionListerExpansion
}

// pullSubscriptionLister implements the PullSubscriptionLister interface.
type pullSubscriptionLister struct {
	indexer cache.Indexer
}

// NewPullSubscriptionLister returns a new PullSubscriptionLister.
func NewPullSubscriptionLister(indexer cache.Indexer) PullSubscriptionLister {
	return &pullSubscriptionLister{indexer: indexer}
}

// List lists all PullSubscriptions in the indexer.
func (s *pullSubscriptionLister) List(selector labels.Selector) (ret []*v1.PullSubscription, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PullSubscription))
	})
	return ret, err
}

// PullSubscriptions returns an object that can list and get PullSubscriptions.
func (s *pullSubscriptionLister) PullSubscriptions(namespace string) PullSubscriptionNamespaceLister {
	return pullSubscriptionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PullSubscriptionNamespaceLister helps list and get PullSubscriptions.
// All objects returned here must be treated as read-only.
type PullSubscriptionNamespaceLister interface {
	// List lists all PullSubscriptions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PullSubscription, err error)
	// Get retrieves the PullSubscription from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PullSubscription, error)
	PullSubscriptionNamespaceListerExpansion
}

// pullSubscriptionNamespaceLister implements the PullSubscriptionNamespaceLister
// interface.
type pullSubscriptionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PullSubscriptions in the indexer for a given namespace.
func (s pullSubscriptionNamespaceLister) List(selector labels.Selector) (ret []*v1.PullSubscription, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PullSubscription))
	})
	return ret, err
}

// Get retrieves the PullSubscription from the indexer for a given namespace and name.
func (s pullSubscriptionNamespaceLister) Get(name string) (*v1.PullSubscription, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("pullsubscription"), name)
	}
	return obj.(*v1.PullSubscription), nil
}
