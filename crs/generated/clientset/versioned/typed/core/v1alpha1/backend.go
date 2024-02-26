//
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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/haproxytech/kubernetes-ingress/crs/api/core/v1alpha1"
	scheme "github.com/haproxytech/kubernetes-ingress/crs/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackendsGetter has a method to return a BackendInterface.
// A group's client should implement this interface.
type BackendsGetter interface {
	Backends(namespace string) BackendInterface
}

// BackendInterface has methods to work with Backend resources.
type BackendInterface interface {
	Create(ctx context.Context, backend *v1alpha1.Backend, opts v1.CreateOptions) (*v1alpha1.Backend, error)
	Update(ctx context.Context, backend *v1alpha1.Backend, opts v1.UpdateOptions) (*v1alpha1.Backend, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Backend, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BackendList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Backend, err error)
	BackendExpansion
}

// backends implements BackendInterface
type backends struct {
	client rest.Interface
	ns     string
}

// newBackends returns a Backends
func newBackends(c *CoreV1alpha1Client, namespace string) *backends {
	return &backends{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backend, and returns the corresponding backend object, and an error if there is any.
func (c *backends) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Backend, err error) {
	result = &v1alpha1.Backend{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backends").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Backends that match those selectors.
func (c *backends) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BackendList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BackendList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backends.
func (c *backends) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a backend and creates it.  Returns the server's representation of the backend, and an error, if there is any.
func (c *backends) Create(ctx context.Context, backend *v1alpha1.Backend, opts v1.CreateOptions) (result *v1alpha1.Backend, err error) {
	result = &v1alpha1.Backend{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backend).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a backend and updates it. Returns the server's representation of the backend, and an error, if there is any.
func (c *backends) Update(ctx context.Context, backend *v1alpha1.Backend, opts v1.UpdateOptions) (result *v1alpha1.Backend, err error) {
	result = &v1alpha1.Backend{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backends").
		Name(backend.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(backend).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the backend and deletes it. Returns an error if one occurs.
func (c *backends) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backends").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backends) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backends").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched backend.
func (c *backends) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Backend, err error) {
	result = &v1alpha1.Backend{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backends").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
