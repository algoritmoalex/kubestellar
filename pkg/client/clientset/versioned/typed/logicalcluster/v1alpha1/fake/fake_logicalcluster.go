/*
Copyright The KubeStellar Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kubestellar/kubestellar/pkg/apis/logicalcluster/v1alpha1"
)

// FakeLogicalClusters implements LogicalClusterInterface
type FakeLogicalClusters struct {
	Fake *FakeLogicalclusterV1alpha1
	ns   string
}

var logicalclustersResource = schema.GroupVersionResource{Group: "logicalcluster.kubestellar.io", Version: "v1alpha1", Resource: "logicalclusters"}

var logicalclustersKind = schema.GroupVersionKind{Group: "logicalcluster.kubestellar.io", Version: "v1alpha1", Kind: "LogicalCluster"}

// Get takes name of the logicalCluster, and returns the corresponding logicalCluster object, and an error if there is any.
func (c *FakeLogicalClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.LogicalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logicalclustersResource, c.ns, name), &v1alpha1.LogicalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicalCluster), err
}

// List takes label and field selectors, and returns the list of LogicalClusters that match those selectors.
func (c *FakeLogicalClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogicalClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logicalclustersResource, logicalclustersKind, c.ns, opts), &v1alpha1.LogicalClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogicalClusterList{ListMeta: obj.(*v1alpha1.LogicalClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogicalClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logicalClusters.
func (c *FakeLogicalClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logicalclustersResource, c.ns, opts))

}

// Create takes the representation of a logicalCluster and creates it.  Returns the server's representation of the logicalCluster, and an error, if there is any.
func (c *FakeLogicalClusters) Create(ctx context.Context, logicalCluster *v1alpha1.LogicalCluster, opts v1.CreateOptions) (result *v1alpha1.LogicalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logicalclustersResource, c.ns, logicalCluster), &v1alpha1.LogicalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicalCluster), err
}

// Update takes the representation of a logicalCluster and updates it. Returns the server's representation of the logicalCluster, and an error, if there is any.
func (c *FakeLogicalClusters) Update(ctx context.Context, logicalCluster *v1alpha1.LogicalCluster, opts v1.UpdateOptions) (result *v1alpha1.LogicalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logicalclustersResource, c.ns, logicalCluster), &v1alpha1.LogicalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicalCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogicalClusters) UpdateStatus(ctx context.Context, logicalCluster *v1alpha1.LogicalCluster, opts v1.UpdateOptions) (*v1alpha1.LogicalCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logicalclustersResource, "status", c.ns, logicalCluster), &v1alpha1.LogicalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicalCluster), err
}

// Delete takes name of the logicalCluster and deletes it. Returns an error if one occurs.
func (c *FakeLogicalClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(logicalclustersResource, c.ns, name, opts), &v1alpha1.LogicalCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogicalClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logicalclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogicalClusterList{})
	return err
}

// Patch applies the patch and returns the patched logicalCluster.
func (c *FakeLogicalClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.LogicalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logicalclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.LogicalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LogicalCluster), err
}