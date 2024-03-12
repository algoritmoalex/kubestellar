/*
Copyright 2023 The KubeStellar Authors.

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

package ocm

import (
	"context"
	"fmt"
	"os"

	clusterv1 "open-cluster-management.io/api/cluster/v1"
	workv1 "open-cluster-management.io/api/work/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

func GetOCMClient(kubeconfig *rest.Config) client.Client {
	scheme := runtime.NewScheme()
	httpClient, err := rest.HTTPClientFor(kubeconfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating HTTPClient: %v\n", err)
		os.Exit(1)
	}
	mapper, err := apiutil.NewDiscoveryRESTMapper(kubeconfig, httpClient)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating NewDiscoveryRESTMapper: %v\n", err)
		os.Exit(1)
	}
	if err := workv1.AddToScheme(scheme); err != nil {
		fmt.Fprintf(os.Stderr, "Error adding to schema: %v\n", err)
		os.Exit(1)
	}
	if err := clusterv1.AddToScheme(scheme); err != nil {
		fmt.Fprintf(os.Stderr, "Error adding to schema: %v\n", err)
		os.Exit(1)
	}
	c, err := client.New(kubeconfig, client.Options{Scheme: scheme, Mapper: mapper})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating client: %v\n", err)
		os.Exit(1)
	}
	return c
}

func FindClustersBySelectors(ocmClient client.Client, selectors []metav1.LabelSelector) (sets.Set[string], error) {
	clusters := &clusterv1.ManagedClusterList{}
	labelSelectors := []labels.Selector{}
	for _, s := range selectors {
		selector, err := metav1.LabelSelectorAsSelector(&s)
		if err != nil {
			return nil, err
		}
		labelSelectors = append(labelSelectors, selector)
	}
	labelsMap, err := convertLabelSelectorsToMap(labelSelectors)
	if err != nil {
		return nil, err
	}
	if err := ocmClient.List(context.TODO(), clusters, client.MatchingLabels(labelsMap)); err != nil {
		return nil, err
	}
	if len(clusters.Items) == 0 {
		return nil, nil
	}

	clusterNames := sets.New[string]()
	for _, cluster := range clusters.Items {
		clusterNames.Insert(cluster.GetName())
	}

	return clusterNames, nil
}

func convertLabelSelectorsToMap(labelSelectors []labels.Selector) (map[string]string, error) {
	labelsMap := make(map[string]string)
	for _, selector := range labelSelectors {
		strSelector := selector.String()
		labelSet, err := labels.ConvertSelectorToLabelsMap(strSelector)
		if err != nil {
			return nil, err
		}
		for key, value := range labelSet {
			labelsMap[key] = value
		}
	}
	return labelsMap, nil
}
