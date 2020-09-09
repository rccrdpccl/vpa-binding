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
	"context"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
	"knative.dev/pkg/apis/duck"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/tracker"
)

const (
	// VPABindingConditionReady is set when the binding has been applied to the subjects.
	VPABindingConditionReady = apis.ConditionReady
	VPABindingLabelKey = "vpa-binding"
)

var vpaCondSet = apis.NewLivingConditionSet()

// GetGroupVersionKind implements kmeta.OwnerRefable
func (fb *VPABinding) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("SQLBinding")
}

// GetSubject implements Bindable
func (fb *VPABinding) GetSubject() tracker.Reference {
	return fb.Spec.Subject
}

// GetBindingStatus implements Bindable
func (fb *VPABinding) GetBindingStatus() duck.BindableStatus {
	return &fb.Status
}

// SetObservedGeneration implements BindableStatus
func (fbs *VPABindingStatus) SetObservedGeneration(gen int64) {
	fbs.ObservedGeneration = gen
}

func (fbs *VPABindingStatus) InitializeConditions() {
	vpaCondSet.Manage(fbs).InitializeConditions()
}

func (fbs *VPABindingStatus) MarkBindingUnavailable(reason, message string) {
	vpaCondSet.Manage(fbs).MarkFalse(VPABindingConditionReady, reason, message)
}

func (fbs *VPABindingStatus) MarkBindingAvailable() {
	vpaCondSet.Manage(fbs).MarkTrue(VPABindingConditionReady)
}

func (fb *VPABinding) Do(ctx context.Context, ps *duckv1.WithPod) {
	ps.Spec.Template.Labels[VPABindingLabelKey] = fb.ObjectMeta.Name
}

func (fb *VPABinding) Undo(ctx context.Context, ps *duckv1.WithPod) {
	delete(ps.Spec.Template.Labels, VPABindingLabelKey)
}
