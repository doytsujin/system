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

	v1alpha1 "github.com/projectriff/system/pkg/apis/bindings/v1alpha1"
)

// FakeImageBindings implements ImageBindingInterface
type FakeImageBindings struct {
	Fake *FakeBindingsV1alpha1
	ns   string
}

var imagebindingsResource = schema.GroupVersionResource{Group: "bindings.projectriff.io", Version: "v1alpha1", Resource: "imagebindings"}

var imagebindingsKind = schema.GroupVersionKind{Group: "bindings.projectriff.io", Version: "v1alpha1", Kind: "ImageBinding"}

// Get takes name of the imageBinding, and returns the corresponding imageBinding object, and an error if there is any.
func (c *FakeImageBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.ImageBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(imagebindingsResource, c.ns, name), &v1alpha1.ImageBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageBinding), err
}

// List takes label and field selectors, and returns the list of ImageBindings that match those selectors.
func (c *FakeImageBindings) List(opts v1.ListOptions) (result *v1alpha1.ImageBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(imagebindingsResource, imagebindingsKind, c.ns, opts), &v1alpha1.ImageBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ImageBindingList{ListMeta: obj.(*v1alpha1.ImageBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.ImageBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested imageBindings.
func (c *FakeImageBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(imagebindingsResource, c.ns, opts))

}

// Create takes the representation of a imageBinding and creates it.  Returns the server's representation of the imageBinding, and an error, if there is any.
func (c *FakeImageBindings) Create(imageBinding *v1alpha1.ImageBinding) (result *v1alpha1.ImageBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(imagebindingsResource, c.ns, imageBinding), &v1alpha1.ImageBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageBinding), err
}

// Update takes the representation of a imageBinding and updates it. Returns the server's representation of the imageBinding, and an error, if there is any.
func (c *FakeImageBindings) Update(imageBinding *v1alpha1.ImageBinding) (result *v1alpha1.ImageBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(imagebindingsResource, c.ns, imageBinding), &v1alpha1.ImageBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeImageBindings) UpdateStatus(imageBinding *v1alpha1.ImageBinding) (*v1alpha1.ImageBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(imagebindingsResource, "status", c.ns, imageBinding), &v1alpha1.ImageBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageBinding), err
}

// Delete takes name of the imageBinding and deletes it. Returns an error if one occurs.
func (c *FakeImageBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(imagebindingsResource, c.ns, name), &v1alpha1.ImageBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeImageBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(imagebindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ImageBindingList{})
	return err
}

// Patch applies the patch and returns the patched imageBinding.
func (c *FakeImageBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ImageBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(imagebindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ImageBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ImageBinding), err
}
