// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	context "context"
	time "time"

	apiconsolev1 "github.com/openshift/api/console/v1"
	versioned "github.com/openshift/client-go/console/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/console/informers/externalversions/internalinterfaces"
	consolev1 "github.com/openshift/client-go/console/listers/console/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ConsoleNotificationInformer provides access to a shared informer and lister for
// ConsoleNotifications.
type ConsoleNotificationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() consolev1.ConsoleNotificationLister
}

type consoleNotificationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewConsoleNotificationInformer constructs a new informer for ConsoleNotification type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewConsoleNotificationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredConsoleNotificationInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredConsoleNotificationInformer constructs a new informer for ConsoleNotification type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredConsoleNotificationInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleNotifications().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConsoleV1().ConsoleNotifications().Watch(context.TODO(), options)
			},
		},
		&apiconsolev1.ConsoleNotification{},
		resyncPeriod,
		indexers,
	)
}

func (f *consoleNotificationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredConsoleNotificationInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *consoleNotificationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apiconsolev1.ConsoleNotification{}, f.defaultInformer)
}

func (f *consoleNotificationInformer) Lister() consolev1.ConsoleNotificationLister {
	return consolev1.NewConsoleNotificationLister(f.Informer().GetIndexer())
}
