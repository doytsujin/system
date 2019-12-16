/*
Copyright 2019 the original author or authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/projectriff/system/pkg/apis/streaming/v1alpha1"
)

// FakeInMemoryProviders implements InMemoryProviderInterface
type FakeInMemoryProviders struct {
	Fake *FakeStreamingV1alpha1
	ns   string
}

var inmemoryprovidersResource = schema.GroupVersionResource{Group: "streaming.projectriff.io", Version: "v1alpha1", Resource: "inmemoryproviders"}

var inmemoryprovidersKind = schema.GroupVersionKind{Group: "streaming.projectriff.io", Version: "v1alpha1", Kind: "InMemoryProvider"}

// Get takes name of the inMemoryProvider, and returns the corresponding inMemoryProvider object, and an error if there is any.
func (c *FakeInMemoryProviders) Get(name string, options v1.GetOptions) (result *v1alpha1.InMemoryProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(inmemoryprovidersResource, c.ns, name), &v1alpha1.InMemoryProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InMemoryProvider), err
}

// List takes label and field selectors, and returns the list of InMemoryProviders that match those selectors.
func (c *FakeInMemoryProviders) List(opts v1.ListOptions) (result *v1alpha1.InMemoryProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(inmemoryprovidersResource, inmemoryprovidersKind, c.ns, opts), &v1alpha1.InMemoryProviderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InMemoryProviderList{ListMeta: obj.(*v1alpha1.InMemoryProviderList).ListMeta}
	for _, item := range obj.(*v1alpha1.InMemoryProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested inMemoryProviders.
func (c *FakeInMemoryProviders) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(inmemoryprovidersResource, c.ns, opts))

}

// Create takes the representation of a inMemoryProvider and creates it.  Returns the server's representation of the inMemoryProvider, and an error, if there is any.
func (c *FakeInMemoryProviders) Create(inMemoryProvider *v1alpha1.InMemoryProvider) (result *v1alpha1.InMemoryProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(inmemoryprovidersResource, c.ns, inMemoryProvider), &v1alpha1.InMemoryProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InMemoryProvider), err
}

// Update takes the representation of a inMemoryProvider and updates it. Returns the server's representation of the inMemoryProvider, and an error, if there is any.
func (c *FakeInMemoryProviders) Update(inMemoryProvider *v1alpha1.InMemoryProvider) (result *v1alpha1.InMemoryProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(inmemoryprovidersResource, c.ns, inMemoryProvider), &v1alpha1.InMemoryProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InMemoryProvider), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeInMemoryProviders) UpdateStatus(inMemoryProvider *v1alpha1.InMemoryProvider) (*v1alpha1.InMemoryProvider, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(inmemoryprovidersResource, "status", c.ns, inMemoryProvider), &v1alpha1.InMemoryProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InMemoryProvider), err
}

// Delete takes name of the inMemoryProvider and deletes it. Returns an error if one occurs.
func (c *FakeInMemoryProviders) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(inmemoryprovidersResource, c.ns, name), &v1alpha1.InMemoryProvider{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInMemoryProviders) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(inmemoryprovidersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InMemoryProviderList{})
	return err
}

// Patch applies the patch and returns the patched inMemoryProvider.
func (c *FakeInMemoryProviders) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InMemoryProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(inmemoryprovidersResource, c.ns, name, pt, data, subresources...), &v1alpha1.InMemoryProvider{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InMemoryProvider), err
}
