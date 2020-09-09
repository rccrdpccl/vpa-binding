/*
Copyright 2019 The Knative Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	duckv1 "knative.dev/pkg/apis/duck/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodSelector is a Knative-style Binding for binding VPA objects to kservice
type PodSelector struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +optional
	Spec PodSelectorSpec `json:"spec,omitempty"`

	// +optional
	Status PodSelectorStatus `json:"status,omitempty"`
}

type ReplicaCount int

type SelectorString string

// VPABindingSpec holds the desired state of the VPABinding (from the client).
type PodSelectorSpec struct {
	Replicas ReplicaCount `json:"replicas""`
	Selector SelectorString `json:"selector""`
}

// VPABindingStatus communicates the observed state of the VPABinding (from the controller).
type PodSelectorStatus struct {
	duckv1.Status `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VPABindingList is a list of VPABinding resources
type PodSelectorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PodSelector `json:"items"`
}


