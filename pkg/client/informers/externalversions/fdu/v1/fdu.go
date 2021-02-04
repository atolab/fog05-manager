/*
* Copyright (c) 2018,2021 ADLINK Technology Inc.
*
* This program and the accompanying materials are made available under the
* terms of the Eclipse Public License 2.0 which is available at
* http://www.eclipse.org/legal/epl-2.0, or the Apache Software License 2.0
* which is available at https://www.apache.org/licenses/LICENSE-2.0.
*
* SPDX-License-Identifier: EPL-2.0 OR Apache-2.0
* Contributors:
*   ADLINK fog05 team, <fog05@adlink-labs.tech>
 */

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	fduv1 "github.com/atolab/fog05-manager/pkg/apis/fdu/v1"
	versioned "github.com/atolab/fog05-manager/pkg/client/clientset/versioned"
	internalinterfaces "github.com/atolab/fog05-manager/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/atolab/fog05-manager/pkg/client/listers/fdu/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FDUInformer provides access to a shared informer and lister for
// FDUs.
type FDUInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.FDULister
}

type fDUInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFDUInformer constructs a new informer for FDU type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFDUInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFDUInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFDUInformer constructs a new informer for FDU type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFDUInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Fog05V1().FDUs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Fog05V1().FDUs(namespace).Watch(context.TODO(), options)
			},
		},
		&fduv1.FDU{},
		resyncPeriod,
		indexers,
	)
}

func (f *fDUInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFDUInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *fDUInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&fduv1.FDU{}, f.defaultInformer)
}

func (f *fDUInformer) Lister() v1.FDULister {
	return v1.NewFDULister(f.Informer().GetIndexer())
}
