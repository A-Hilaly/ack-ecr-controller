// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PullThroughCacheRuleSpec defines the desired state of PullThroughCacheRule.
//
// The details of a pull through cache rule.
type PullThroughCacheRuleSpec struct {
	// The repository name prefix to use when caching images from the source registry.
	// +kubebuilder:validation:Required
	ECRRepositoryPrefix *string `json:"ecrRepositoryPrefix"`
	// The Amazon Web Services account ID associated with the registry to create
	// the pull through cache rule for. If you do not specify a registry, the default
	// registry is assumed.
	RegistryID *string `json:"registryID,omitempty"`
	// The registry URL of the upstream public registry to use as the source for
	// the pull through cache rule.
	// +kubebuilder:validation:Required
	UpstreamRegistryURL *string `json:"upstreamRegistryURL"`
}

// PullThroughCacheRuleStatus defines the observed state of PullThroughCacheRule
type PullThroughCacheRuleStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// The date and time, in JavaScript date format, when the pull through cache
	// rule was created.
	// +kubebuilder:validation:Optional
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
}

// PullThroughCacheRule is the Schema for the PullThroughCacheRules API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type PullThroughCacheRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PullThroughCacheRuleSpec   `json:"spec,omitempty"`
	Status            PullThroughCacheRuleStatus `json:"status,omitempty"`
}

// PullThroughCacheRuleList contains a list of PullThroughCacheRule
// +kubebuilder:object:root=true
type PullThroughCacheRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PullThroughCacheRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PullThroughCacheRule{}, &PullThroughCacheRuleList{})
}
