/*
Copyright The Kubernetes Authors.

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
	stablev1beta1 "redis-controller/pkg/apis/stable/v1beta1"
	versioned "redis-controller/pkg/generated/clientset/versioned"
	internalinterfaces "redis-controller/pkg/generated/informers/externalversions/internalinterfaces"
	v1beta1 "redis-controller/pkg/generated/listers/stable/v1beta1"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// RedisInformer provides access to a shared informer and lister for
// Redises.
type RedisInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.RedisLister
}

type redisInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewRedisInformer constructs a new informer for Redis type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewRedisInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredRedisInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredRedisInformer constructs a new informer for Redis type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredRedisInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StableV1beta1().Redises(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.StableV1beta1().Redises(namespace).Watch(context.TODO(), options)
			},
		},
		&stablev1beta1.Redis{},
		resyncPeriod,
		indexers,
	)
}

func (f *redisInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredRedisInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *redisInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&stablev1beta1.Redis{}, f.defaultInformer)
}

func (f *redisInformer) Lister() v1beta1.RedisLister {
	return v1beta1.NewRedisLister(f.Informer().GetIndexer())
}