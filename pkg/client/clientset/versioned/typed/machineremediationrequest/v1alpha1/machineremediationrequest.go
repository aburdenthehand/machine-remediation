/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/metal3-io/machine-remediation-request-operator/pkg/apis/machineremediationrequest/v1alpha1"
	scheme "github.com/metal3-io/machine-remediation-request-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineRemediationRequestsGetter has a method to return a MachineRemediationRequestInterface.
// A group's client should implement this interface.
type MachineRemediationRequestsGetter interface {
	MachineRemediationRequests(namespace string) MachineRemediationRequestInterface
}

// MachineRemediationRequestInterface has methods to work with MachineRemediationRequest resources.
type MachineRemediationRequestInterface interface {
	Create(*v1alpha1.MachineRemediationRequest) (*v1alpha1.MachineRemediationRequest, error)
	Update(*v1alpha1.MachineRemediationRequest) (*v1alpha1.MachineRemediationRequest, error)
	UpdateStatus(*v1alpha1.MachineRemediationRequest) (*v1alpha1.MachineRemediationRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MachineRemediationRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.MachineRemediationRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineRemediationRequest, err error)
	MachineRemediationRequestExpansion
}

// machineRemediationRequests implements MachineRemediationRequestInterface
type machineRemediationRequests struct {
	client rest.Interface
	ns     string
}

// newMachineRemediationRequests returns a MachineRemediationRequests
func newMachineRemediationRequests(c *MachineremediationrequestV1alpha1Client, namespace string) *machineRemediationRequests {
	return &machineRemediationRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineRemediationRequest, and returns the corresponding machineRemediationRequest object, and an error if there is any.
func (c *machineRemediationRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineRemediationRequest, err error) {
	result = &v1alpha1.MachineRemediationRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineRemediationRequests that match those selectors.
func (c *machineRemediationRequests) List(opts v1.ListOptions) (result *v1alpha1.MachineRemediationRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.MachineRemediationRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineRemediationRequests.
func (c *machineRemediationRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a machineRemediationRequest and creates it.  Returns the server's representation of the machineRemediationRequest, and an error, if there is any.
func (c *machineRemediationRequests) Create(machineRemediationRequest *v1alpha1.MachineRemediationRequest) (result *v1alpha1.MachineRemediationRequest, err error) {
	result = &v1alpha1.MachineRemediationRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		Body(machineRemediationRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineRemediationRequest and updates it. Returns the server's representation of the machineRemediationRequest, and an error, if there is any.
func (c *machineRemediationRequests) Update(machineRemediationRequest *v1alpha1.MachineRemediationRequest) (result *v1alpha1.MachineRemediationRequest, err error) {
	result = &v1alpha1.MachineRemediationRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		Name(machineRemediationRequest.Name).
		Body(machineRemediationRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *machineRemediationRequests) UpdateStatus(machineRemediationRequest *v1alpha1.MachineRemediationRequest) (result *v1alpha1.MachineRemediationRequest, err error) {
	result = &v1alpha1.MachineRemediationRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		Name(machineRemediationRequest.Name).
		SubResource("status").
		Body(machineRemediationRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineRemediationRequest and deletes it. Returns an error if one occurs.
func (c *machineRemediationRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineRemediationRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machineremediationrequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineRemediationRequest.
func (c *machineRemediationRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineRemediationRequest, err error) {
	result = &v1alpha1.MachineRemediationRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machineremediationrequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
