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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "github.com/kubeflow/kfserving/pkg/apis/serving/v1alpha2"
	scheme "github.com/kubeflow/kfserving/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// InferenceServicesGetter has a method to return a InferenceServiceInterface.
// A group's client should implement this interface.
type InferenceServicesGetter interface {
	InferenceServices(namespace string) InferenceServiceInterface
}

// InferenceServiceInterface has methods to work with InferenceService resources.
type InferenceServiceInterface interface {
	Create(*v1alpha2.InferenceService) (*v1alpha2.InferenceService, error)
	Update(*v1alpha2.InferenceService) (*v1alpha2.InferenceService, error)
	UpdateStatus(*v1alpha2.InferenceService) (*v1alpha2.InferenceService, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.InferenceService, error)
	List(opts v1.ListOptions) (*v1alpha2.InferenceServiceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.InferenceService, err error)
	InferenceServiceExpansion
}

// inferenceServices implements InferenceServiceInterface
type inferenceServices struct {
	client rest.Interface
	ns     string
}

// newInferenceServices returns a InferenceServices
func newInferenceServices(c *ServingV1alpha2Client, namespace string) *inferenceServices {
	return &inferenceServices{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the inferenceService, and returns the corresponding inferenceService object, and an error if there is any.
func (c *inferenceServices) Get(name string, options v1.GetOptions) (result *v1alpha2.InferenceService, err error) {
	result = &v1alpha2.InferenceService{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of InferenceServices that match those selectors.
func (c *inferenceServices) List(opts v1.ListOptions) (result *v1alpha2.InferenceServiceList, err error) {
	result = &v1alpha2.InferenceServiceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested inferenceServices.
func (c *inferenceServices) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a inferenceService and creates it.  Returns the server's representation of the inferenceService, and an error, if there is any.
func (c *inferenceServices) Create(inferenceService *v1alpha2.InferenceService) (result *v1alpha2.InferenceService, err error) {
	result = &v1alpha2.InferenceService{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("inferenceservices").
		Body(inferenceService).
		Do().
		Into(result)
	return
}

// Update takes the representation of a inferenceService and updates it. Returns the server's representation of the inferenceService, and an error, if there is any.
func (c *inferenceServices) Update(inferenceService *v1alpha2.InferenceService) (result *v1alpha2.InferenceService, err error) {
	result = &v1alpha2.InferenceService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(inferenceService.Name).
		Body(inferenceService).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *inferenceServices) UpdateStatus(inferenceService *v1alpha2.InferenceService) (result *v1alpha2.InferenceService, err error) {
	result = &v1alpha2.InferenceService{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(inferenceService.Name).
		SubResource("status").
		Body(inferenceService).
		Do().
		Into(result)
	return
}

// Delete takes name of the inferenceService and deletes it. Returns an error if one occurs.
func (c *inferenceServices) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inferenceservices").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *inferenceServices) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("inferenceservices").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched inferenceService.
func (c *inferenceServices) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.InferenceService, err error) {
	result = &v1alpha2.InferenceService{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("inferenceservices").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}