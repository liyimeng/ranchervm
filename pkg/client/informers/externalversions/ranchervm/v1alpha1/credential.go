/*
Copyright 2018 Rancher Labs, Inc.

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
	time "time"

	ranchervm_v1alpha1 "github.com/rancher/vm/pkg/apis/ranchervm/v1alpha1"
	versioned "github.com/rancher/vm/pkg/client/clientset/versioned"
	internalinterfaces "github.com/rancher/vm/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/rancher/vm/pkg/client/listers/ranchervm/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CredentialInformer provides access to a shared informer and lister for
// Credentials.
type CredentialInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CredentialLister
}

type credentialInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCredentialInformer constructs a new informer for Credential type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCredentialInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCredentialInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCredentialInformer constructs a new informer for Credential type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCredentialInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VirtualmachineV1alpha1().Credentials().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.VirtualmachineV1alpha1().Credentials().Watch(options)
			},
		},
		&ranchervm_v1alpha1.Credential{},
		resyncPeriod,
		indexers,
	)
}

func (f *credentialInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCredentialInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *credentialInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ranchervm_v1alpha1.Credential{}, f.defaultInformer)
}

func (f *credentialInformer) Lister() v1alpha1.CredentialLister {
	return v1alpha1.NewCredentialLister(f.Informer().GetIndexer())
}
