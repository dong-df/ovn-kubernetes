/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	routeadvertisementsv1 "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/routeadvertisements/v1"
	types "github.com/ovn-org/ovn-kubernetes/go-controller/pkg/crd/types"
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// RouteAdvertisementsSpecApplyConfiguration represents a declarative configuration of the RouteAdvertisementsSpec type for use
// with apply.
type RouteAdvertisementsSpecApplyConfiguration struct {
	TargetVRF                *string                                   `json:"targetVRF,omitempty"`
	NetworkSelectors         *types.NetworkSelectors                   `json:"networkSelectors,omitempty"`
	NodeSelector             *metav1.LabelSelectorApplyConfiguration   `json:"nodeSelector,omitempty"`
	FRRConfigurationSelector *metav1.LabelSelectorApplyConfiguration   `json:"frrConfigurationSelector,omitempty"`
	Advertisements           []routeadvertisementsv1.AdvertisementType `json:"advertisements,omitempty"`
}

// RouteAdvertisementsSpecApplyConfiguration constructs a declarative configuration of the RouteAdvertisementsSpec type for use with
// apply.
func RouteAdvertisementsSpec() *RouteAdvertisementsSpecApplyConfiguration {
	return &RouteAdvertisementsSpecApplyConfiguration{}
}

// WithTargetVRF sets the TargetVRF field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TargetVRF field is set to the value of the last call.
func (b *RouteAdvertisementsSpecApplyConfiguration) WithTargetVRF(value string) *RouteAdvertisementsSpecApplyConfiguration {
	b.TargetVRF = &value
	return b
}

// WithNetworkSelectors sets the NetworkSelectors field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkSelectors field is set to the value of the last call.
func (b *RouteAdvertisementsSpecApplyConfiguration) WithNetworkSelectors(value types.NetworkSelectors) *RouteAdvertisementsSpecApplyConfiguration {
	b.NetworkSelectors = &value
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *RouteAdvertisementsSpecApplyConfiguration) WithNodeSelector(value *metav1.LabelSelectorApplyConfiguration) *RouteAdvertisementsSpecApplyConfiguration {
	b.NodeSelector = value
	return b
}

// WithFRRConfigurationSelector sets the FRRConfigurationSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FRRConfigurationSelector field is set to the value of the last call.
func (b *RouteAdvertisementsSpecApplyConfiguration) WithFRRConfigurationSelector(value *metav1.LabelSelectorApplyConfiguration) *RouteAdvertisementsSpecApplyConfiguration {
	b.FRRConfigurationSelector = value
	return b
}

// WithAdvertisements adds the given value to the Advertisements field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Advertisements field.
func (b *RouteAdvertisementsSpecApplyConfiguration) WithAdvertisements(values ...routeadvertisementsv1.AdvertisementType) *RouteAdvertisementsSpecApplyConfiguration {
	for i := range values {
		b.Advertisements = append(b.Advertisements, values[i])
	}
	return b
}
