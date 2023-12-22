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

package v2alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v2alpha1 "github.com/kubestellar/kubestellar/pkg/apis/edge/v2alpha1"
	scheme "github.com/kubestellar/kubestellar/pkg/client/clientset/versioned/scheme"
)

// EdgePlacementDecisionsGetter has a method to return a EdgePlacementDecisionInterface.
// A group's client should implement this interface.
type EdgePlacementDecisionsGetter interface {
	EdgePlacementDecisions() EdgePlacementDecisionInterface
}

// EdgePlacementDecisionInterface has methods to work with EdgePlacementDecision resources.
type EdgePlacementDecisionInterface interface {
	Create(ctx context.Context, edgePlacementDecision *v2alpha1.EdgePlacementDecision, opts v1.CreateOptions) (*v2alpha1.EdgePlacementDecision, error)
	Update(ctx context.Context, edgePlacementDecision *v2alpha1.EdgePlacementDecision, opts v1.UpdateOptions) (*v2alpha1.EdgePlacementDecision, error)
	UpdateStatus(ctx context.Context, edgePlacementDecision *v2alpha1.EdgePlacementDecision, opts v1.UpdateOptions) (*v2alpha1.EdgePlacementDecision, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2alpha1.EdgePlacementDecision, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2alpha1.EdgePlacementDecisionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.EdgePlacementDecision, err error)
	EdgePlacementDecisionExpansion
}

// edgePlacementDecisions implements EdgePlacementDecisionInterface
type edgePlacementDecisions struct {
	client rest.Interface
}

// newEdgePlacementDecisions returns a EdgePlacementDecisions
func newEdgePlacementDecisions(c *EdgeV2alpha1Client) *edgePlacementDecisions {
	return &edgePlacementDecisions{
		client: c.RESTClient(),
	}
}

// Get takes name of the edgePlacementDecision, and returns the corresponding edgePlacementDecision object, and an error if there is any.
func (c *edgePlacementDecisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.EdgePlacementDecision, err error) {
	result = &v2alpha1.EdgePlacementDecision{}
	err = c.client.Get().
		Resource("edgeplacementdecisions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EdgePlacementDecisions that match those selectors.
func (c *edgePlacementDecisions) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.EdgePlacementDecisionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2alpha1.EdgePlacementDecisionList{}
	err = c.client.Get().
		Resource("edgeplacementdecisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested edgePlacementDecisions.
func (c *edgePlacementDecisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("edgeplacementdecisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a edgePlacementDecision and creates it.  Returns the server's representation of the edgePlacementDecision, and an error, if there is any.
func (c *edgePlacementDecisions) Create(ctx context.Context, edgePlacementDecision *v2alpha1.EdgePlacementDecision, opts v1.CreateOptions) (result *v2alpha1.EdgePlacementDecision, err error) {
	result = &v2alpha1.EdgePlacementDecision{}
	err = c.client.Post().
		Resource("edgeplacementdecisions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(edgePlacementDecision).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a edgePlacementDecision and updates it. Returns the server's representation of the edgePlacementDecision, and an error, if there is any.
func (c *edgePlacementDecisions) Update(ctx context.Context, edgePlacementDecision *v2alpha1.EdgePlacementDecision, opts v1.UpdateOptions) (result *v2alpha1.EdgePlacementDecision, err error) {
	result = &v2alpha1.EdgePlacementDecision{}
	err = c.client.Put().
		Resource("edgeplacementdecisions").
		Name(edgePlacementDecision.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(edgePlacementDecision).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *edgePlacementDecisions) UpdateStatus(ctx context.Context, edgePlacementDecision *v2alpha1.EdgePlacementDecision, opts v1.UpdateOptions) (result *v2alpha1.EdgePlacementDecision, err error) {
	result = &v2alpha1.EdgePlacementDecision{}
	err = c.client.Put().
		Resource("edgeplacementdecisions").
		Name(edgePlacementDecision.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(edgePlacementDecision).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the edgePlacementDecision and deletes it. Returns an error if one occurs.
func (c *edgePlacementDecisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("edgeplacementdecisions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *edgePlacementDecisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("edgeplacementdecisions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched edgePlacementDecision.
func (c *edgePlacementDecisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.EdgePlacementDecision, err error) {
	result = &v2alpha1.EdgePlacementDecision{}
	err = c.client.Patch(pt).
		Resource("edgeplacementdecisions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}