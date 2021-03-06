/*
Copyright 2016 The Kubernetes Authors.

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

package informers

import (
	"reflect"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/kubernetes/pkg/api/v1"
	extensions "k8s.io/kubernetes/pkg/apis/extensions/v1beta1"
	"k8s.io/kubernetes/pkg/client/cache"
	"k8s.io/kubernetes/pkg/client/legacylisters"
)

// DaemonSetInformer is type of SharedIndexInformer which watches and lists all pods.
// Interface provides constructor for informer and lister for pods
type DaemonSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() *listers.StoreToDaemonSetLister
}

type daemonSetInformer struct {
	*sharedInformerFactory
}

func (f *daemonSetInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(&extensions.DaemonSet{})
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return f.client.Extensions().DaemonSets(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return f.client.Extensions().DaemonSets(v1.NamespaceAll).Watch(options)
			},
		},
		&extensions.DaemonSet{},
		f.defaultResync,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	f.informers[informerType] = informer

	return informer
}

func (f *daemonSetInformer) Lister() *listers.StoreToDaemonSetLister {
	informer := f.Informer()
	return &listers.StoreToDaemonSetLister{Store: informer.GetIndexer()}
}

// DeploymentInformer is a type of SharedIndexInformer which watches and lists all deployments.
type DeploymentInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() *listers.StoreToDeploymentLister
}

type deploymentInformer struct {
	*sharedInformerFactory
}

func (f *deploymentInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(&extensions.Deployment{})
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return f.client.Extensions().Deployments(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return f.client.Extensions().Deployments(v1.NamespaceAll).Watch(options)
			},
		},
		&extensions.Deployment{},
		f.defaultResync,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	f.informers[informerType] = informer

	return informer
}

func (f *deploymentInformer) Lister() *listers.StoreToDeploymentLister {
	informer := f.Informer()
	return &listers.StoreToDeploymentLister{Indexer: informer.GetIndexer()}
}

// ReplicaSetInformer is a type of SharedIndexInformer which watches and lists all replicasets.
type ReplicaSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() *listers.StoreToReplicaSetLister
}

type replicaSetInformer struct {
	*sharedInformerFactory
}

func (f *replicaSetInformer) Informer() cache.SharedIndexInformer {
	f.lock.Lock()
	defer f.lock.Unlock()

	informerType := reflect.TypeOf(&extensions.ReplicaSet{})
	informer, exists := f.informers[informerType]
	if exists {
		return informer
	}
	informer = cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return f.client.Extensions().ReplicaSets(v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return f.client.Extensions().ReplicaSets(v1.NamespaceAll).Watch(options)
			},
		},
		&extensions.ReplicaSet{},
		f.defaultResync,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)
	f.informers[informerType] = informer

	return informer
}

func (f *replicaSetInformer) Lister() *listers.StoreToReplicaSetLister {
	informer := f.Informer()
	return &listers.StoreToReplicaSetLister{Indexer: informer.GetIndexer()}
}
