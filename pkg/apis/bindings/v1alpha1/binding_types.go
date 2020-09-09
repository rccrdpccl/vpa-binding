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

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
	"knative.dev/pkg/tracker"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VPABinding is a Knative-style Binding for binding VPA objects to kservice
type VPABinding struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the VPABinding (from the client).
	// +optional
	Spec VPABindingSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the VPABinding (from the controller).
	// +optional
	Status VPABindingStatus `json:"status,omitempty"`
}

var (
	// Check that VPABinding can be validated and defaulted.
	_ apis.Validatable   = (*VPABinding)(nil)
	_ apis.Defaultable   = (*VPABinding)(nil)
	_ kmeta.OwnerRefable = (*VPABinding)(nil)
)

// VPABindingSpec holds the desired state of the VPABinding (from the client).
type VPABindingSpec struct {
	// Subject holds a reference to the "addressable" Knative resource which will be bound
	Subject tracker.Reference `json:"subject"`
}

// VPABindingStatus communicates the observed state of the VPABinding (from the controller).
type VPABindingStatus struct {
	duckv1.Status `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VPABindingList is a list of VPABinding resources
type VPABindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []VPABinding `json:"items"`
}
