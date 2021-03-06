/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
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
	businessv1 "tkestack.io/tke/api/business/v1"
)

// FakeChartGroups implements ChartGroupInterface
type FakeChartGroups struct {
	Fake *FakeBusinessV1
	ns   string
}

var chartgroupsResource = schema.GroupVersionResource{Group: "business.tkestack.io", Version: "v1", Resource: "chartgroups"}

var chartgroupsKind = schema.GroupVersionKind{Group: "business.tkestack.io", Version: "v1", Kind: "ChartGroup"}

// Get takes name of the chartGroup, and returns the corresponding chartGroup object, and an error if there is any.
func (c *FakeChartGroups) Get(name string, options v1.GetOptions) (result *businessv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(chartgroupsResource, c.ns, name), &businessv1.ChartGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.ChartGroup), err
}

// List takes label and field selectors, and returns the list of ChartGroups that match those selectors.
func (c *FakeChartGroups) List(opts v1.ListOptions) (result *businessv1.ChartGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(chartgroupsResource, chartgroupsKind, c.ns, opts), &businessv1.ChartGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &businessv1.ChartGroupList{ListMeta: obj.(*businessv1.ChartGroupList).ListMeta}
	for _, item := range obj.(*businessv1.ChartGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chartGroups.
func (c *FakeChartGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(chartgroupsResource, c.ns, opts))

}

// Create takes the representation of a chartGroup and creates it.  Returns the server's representation of the chartGroup, and an error, if there is any.
func (c *FakeChartGroups) Create(chartGroup *businessv1.ChartGroup) (result *businessv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(chartgroupsResource, c.ns, chartGroup), &businessv1.ChartGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.ChartGroup), err
}

// Update takes the representation of a chartGroup and updates it. Returns the server's representation of the chartGroup, and an error, if there is any.
func (c *FakeChartGroups) Update(chartGroup *businessv1.ChartGroup) (result *businessv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(chartgroupsResource, c.ns, chartGroup), &businessv1.ChartGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.ChartGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChartGroups) UpdateStatus(chartGroup *businessv1.ChartGroup) (*businessv1.ChartGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(chartgroupsResource, "status", c.ns, chartGroup), &businessv1.ChartGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.ChartGroup), err
}

// Delete takes name of the chartGroup and deletes it. Returns an error if one occurs.
func (c *FakeChartGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(chartgroupsResource, c.ns, name), &businessv1.ChartGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChartGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(chartgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &businessv1.ChartGroupList{})
	return err
}

// Patch applies the patch and returns the patched chartGroup.
func (c *FakeChartGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *businessv1.ChartGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(chartgroupsResource, c.ns, name, pt, data, subresources...), &businessv1.ChartGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.ChartGroup), err
}
