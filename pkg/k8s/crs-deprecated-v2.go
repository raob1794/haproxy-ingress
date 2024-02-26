// Copyright 2019 HAProxy Technologies LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package k8s

import (
	"k8s.io/client-go/tools/cache"

	corev1alpha2 "github.com/haproxytech/kubernetes-ingress/crs/api/core/v1alpha2"
	v1 "github.com/haproxytech/kubernetes-ingress/crs/api/ingress/v1"
	"github.com/haproxytech/kubernetes-ingress/crs/converters"
	informers "github.com/haproxytech/kubernetes-ingress/crs/generated/informers/externalversions"
	"github.com/haproxytech/kubernetes-ingress/pkg/store"
)

type GlobalCRV1Alpha2 struct{}

type DefaultsCRV1Alpha2 struct{}

type BackendCRV1Alpha2 struct{}

func NewGlobalCRV1Alpha2() GlobalCRV1Alpha2 {
	return GlobalCRV1Alpha2{}
}

func NewDefaultsCRV1Alpha2() DefaultsCRV1Alpha2 {
	return DefaultsCRV1Alpha2{}
}

func NewBackendCRV1Alpha2() BackendCRV1Alpha2 {
	return BackendCRV1Alpha2{}
}

func (c GlobalCRV1Alpha2) GetKind() string {
	return "Global"
}

func (c GlobalCRV1Alpha2) GetInformer(eventChan chan SyncDataEvent, factory informers.SharedInformerFactory) cache.SharedIndexInformer { //nolint:ireturn
	informer := factory.Core().V1alpha2().Globals().Informer()

	sendToChannel := func(eventChan chan SyncDataEvent, object interface{}, status store.Status) {
		dataV1Alpha2, ok := object.(*corev1alpha2.Global)
		if !ok {
			logger.Warning(CRSGroupVersionV1alpha2 + ": type mismatch with Global kind")
			return
		}
		logger.Warningf("Global CR defined in API %s is DEPRECATED, please upgrade", CRSGroupVersionV1alpha2)
		data := &v1.Global{}
		data.TypeMeta = dataV1Alpha2.TypeMeta
		data.ObjectMeta = dataV1Alpha2.ObjectMeta
		data.Spec = converters.DeepConvertGlobalSpecA2toV1(dataV1Alpha2.Spec)

		logger.Debugf("%s %s: %s", data.GetNamespace(), status, data.GetName())
		if status == store.DELETED {
			eventChan <- SyncDataEvent{SyncType: SyncType(c.GetKind()), Namespace: data.GetNamespace(), Name: data.GetName(), Data: nil}
			return
		}
		eventChan <- SyncDataEvent{SyncType: SyncType(c.GetKind()), Namespace: data.GetNamespace(), Name: data.GetName(), Data: data}
	}

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.ADDED)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			sendToChannel(eventChan, newObj, store.MODIFIED)
		},
		DeleteFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.DELETED)
		},
	})
	logger.Error(err)
	return informer
}

func (c DefaultsCRV1Alpha2) GetKind() string {
	return "Defaults"
}

func (c DefaultsCRV1Alpha2) GetInformer(eventChan chan SyncDataEvent, factory informers.SharedInformerFactory) cache.SharedIndexInformer { //nolint:ireturn
	informer := factory.Core().V1alpha2().Defaults().Informer()

	sendToChannel := func(eventChan chan SyncDataEvent, object interface{}, status store.Status) {
		dataV1Alpha2, ok := object.(*corev1alpha2.Defaults)
		if !ok {
			logger.Warning(CRSGroupVersionV1alpha2 + ": type mismatch with Defaults kind")
			return
		}
		logger.Warningf("Defaults CR defined in API %s is DEPRECATED, please upgrade", CRSGroupVersionV1alpha2)
		data := &v1.Defaults{}
		data.TypeMeta = dataV1Alpha2.TypeMeta
		data.ObjectMeta = dataV1Alpha2.ObjectMeta
		data.Spec = converters.DeepConvertDefaultsSpecA2toV1(dataV1Alpha2.Spec)
		logger.Debugf("%s %s: %s", data.GetNamespace(), status, data.GetName())
		if status == store.DELETED {
			eventChan <- SyncDataEvent{SyncType: SyncType(c.GetKind()), Namespace: data.GetNamespace(), Name: data.GetName(), Data: nil}
			return
		}
		eventChan <- SyncDataEvent{SyncType: SyncType(c.GetKind()), Namespace: data.GetNamespace(), Name: data.GetName(), Data: data}
	}

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.ADDED)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			sendToChannel(eventChan, newObj, store.MODIFIED)
		},
		DeleteFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.DELETED)
		},
	})
	logger.Error(err)
	return informer
}

func (c BackendCRV1Alpha2) GetKind() string {
	return "Backend"
}

func (c BackendCRV1Alpha2) GetInformer(eventChan chan SyncDataEvent, factory informers.SharedInformerFactory) cache.SharedIndexInformer { //nolint:ireturn
	informer := factory.Core().V1alpha2().Backends().Informer()

	sendToChannel := func(eventChan chan SyncDataEvent, object interface{}, status store.Status) {
		dataV1Alpha2, ok := object.(*corev1alpha2.Backend)
		if !ok {
			logger.Warning(CRSGroupVersionV1alpha2 + ": type mismatch with Backend kind")
			return
		}
		logger.Warningf("Backend CR defined in API %s is DEPRECATED, please upgrade", CRSGroupVersionV1alpha2)
		data := &v1.Backend{}
		data.TypeMeta = dataV1Alpha2.TypeMeta
		data.ObjectMeta = dataV1Alpha2.ObjectMeta
		data.Spec = converters.DeepConvertBackendSpecA2toV1(dataV1Alpha2.Spec)

		logger.Debugf("%s %s: %s", data.GetNamespace(), status, data.GetName())
		if status == store.DELETED {
			eventChan <- SyncDataEvent{SyncType: SyncType(c.GetKind()), Namespace: data.GetNamespace(), Name: data.GetName(), Data: nil}
			return
		}
		eventChan <- SyncDataEvent{SyncType: SyncType(c.GetKind()), Namespace: data.GetNamespace(), Name: data.GetName(), Data: data}
	}

	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.ADDED)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			sendToChannel(eventChan, newObj, store.MODIFIED)
		},
		DeleteFunc: func(obj interface{}) {
			sendToChannel(eventChan, obj, store.DELETED)
		},
	})
	logger.Error(err)
	return informer
}
