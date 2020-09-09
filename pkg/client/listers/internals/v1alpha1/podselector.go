/*
Copyright 2019 The Knative Authors

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

package v1alpha1

import (
	v1alpha1 "github.com/rccrdpccl/bindings/pkg/apis/internals/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PodSelectorLister helps list PodSelectors.
type PodSelectorLister interface {
	// List lists all PodSelectors in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PodSelector, err error)
	// PodSelectors returns an object that can list and get PodSelectors.
	PodSelectors(namespace string) PodSelectorNamespaceLister
	PodSelectorListerExpansion
}

// podSelectorLister implements the PodSelectorLister interface.
type podSelectorLister struct {
	indexer cache.Indexer
}

// NewPodSelectorLister returns a new PodSelectorLister.
func NewPodSelectorLister(indexer cache.Indexer) PodSelectorLister {
	return &podSelectorLister{indexer: indexer}
}

// List lists all PodSelectors in the indexer.
func (s *podSelectorLister) List(selector labels.Selector) (ret []*v1alpha1.PodSelector, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodSelector))
	})
	return ret, err
}

// PodSelectors returns an object that can list and get PodSelectors.
func (s *podSelectorLister) PodSelectors(namespace string) PodSelectorNamespaceLister {
	return podSelectorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PodSelectorNamespaceLister helps list and get PodSelectors.
type PodSelectorNamespaceLister interface {
	// List lists all PodSelectors in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PodSelector, err error)
	// Get retrieves the PodSelector from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PodSelector, error)
	PodSelectorNamespaceListerExpansion
}

// podSelectorNamespaceLister implements the PodSelectorNamespaceLister
// interface.
type podSelectorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PodSelectors in the indexer for a given namespace.
func (s podSelectorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PodSelector, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PodSelector))
	})
	return ret, err
}

// Get retrieves the PodSelector from the indexer for a given namespace and name.
func (s podSelectorNamespaceLister) Get(name string) (*v1alpha1.PodSelector, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("podselector"), name)
	}
	return obj.(*v1alpha1.PodSelector), nil
}
