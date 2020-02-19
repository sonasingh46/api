/*
Copyright 2020 The OpenEBS Authors

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

package v1

import (
	time "time"

	cstorv1 "github.com/openebs/api/pkg/apis/cstor/v1"
	versioned "github.com/openebs/api/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openebs/api/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openebs/api/pkg/client/listers/cstor/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CStorVolumePolicyInformer provides access to a shared informer and lister for
// CStorVolumePolicies.
type CStorVolumePolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CStorVolumePolicyLister
}

type cStorVolumePolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCStorVolumePolicyInformer constructs a new informer for CStorVolumePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCStorVolumePolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCStorVolumePolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCStorVolumePolicyInformer constructs a new informer for CStorVolumePolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCStorVolumePolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CstorV1().CStorVolumePolicies(namespace).List(options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CstorV1().CStorVolumePolicies(namespace).Watch(options)
			},
		},
		&cstorv1.CStorVolumePolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *cStorVolumePolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCStorVolumePolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cStorVolumePolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cstorv1.CStorVolumePolicy{}, f.defaultInformer)
}

func (f *cStorVolumePolicyInformer) Lister() v1.CStorVolumePolicyLister {
	return v1.NewCStorVolumePolicyLister(f.Informer().GetIndexer())
}
