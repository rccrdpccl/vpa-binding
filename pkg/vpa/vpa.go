/*
Copyright 2020 The Knative Authors

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

package vpa

import (
	"context"
	alpha1 "github.com/rccrdpccl/bindings/pkg/apis/bindings/v1alpha1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/errors"
	"knative.dev/pkg/logging"
	"knative.dev/pkg/webhook/psbinding"

	"github.com/rccrdpccl/bindings/pkg/apis/internals/v1alpha1"
	clientset "github.com/rccrdpccl/bindings/pkg/client/clientset/versioned"
	v1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	autoscalingv1 "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/apis/autoscaling.k8s.io/v1"
	vpa "k8s.io/autoscaler/vertical-pod-autoscaler/pkg/client/clientset/versioned"
	"sync"
)
var once sync.Once

type VPASubresourcesReconciler struct {
	vpaBindingClient *clientset.Clientset
	vpaClient *vpa.Clientset
}

var vpaSubresourcesReconciler VPASubresourcesReconciler

func NewVPASubresourcesReconciler(c *clientset.Clientset, vpaClient *vpa.Clientset) VPASubresourcesReconciler {
	once.Do(func() {
		vpaSubresourcesReconciler = VPASubresourcesReconciler{
			vpaBindingClient: c,
			vpaClient: vpaClient,
		}
	})
	return vpaSubresourcesReconciler
}

func getPodSelectorName(vpaBinding *alpha1.VPABinding) string {
	return vpaBinding.Name + "-pod-selector"
}

func getVPAName(vpaBinding *alpha1.VPABinding) string {
	return vpaBinding.Name + "-vpa"
}

func (vsr *VPASubresourcesReconciler) Reconcile(ctx context.Context, fb psbinding.Bindable) error {
	vpaBinding := fb.(*alpha1.VPABinding)

	podSelector, err := vsr.createPodSelector(ctx, vpaBinding)
	if err != nil {
		return err
	}
	vpaCreated, err := vsr.createVPA(ctx, vpaBinding, podSelector)
	if err != nil {
		return err
	}
	logging.FromContext(ctx).Infof("Created VPA %v", vpaCreated)
	return nil
}

func (vsr *VPASubresourcesReconciler) ReconcileDeletion(ctx context.Context, fb psbinding.Bindable) error {
	vpaBinding := fb.(*alpha1.VPABinding)

	var errorList []error
	podSelectorName := getPodSelectorName(vpaBinding)
	vpaName := getVPAName(vpaBinding)
	podSelectorErr := vsr.vpaBindingClient.InternalsV1alpha1().PodSelectors(vpaBinding.Namespace).Delete(ctx, podSelectorName, metav1.DeleteOptions{})
	if podSelectorErr != nil {
		logging.FromContext(ctx).Infof("Could not delete PodSelector %s", podSelectorName)
		errorList = append(errorList, podSelectorErr)
	}
	vpaErr := vsr.vpaClient.AutoscalingV1().VerticalPodAutoscalers(vpaBinding.Namespace).Delete(ctx, vpaName, metav1.DeleteOptions{})
	if vpaErr != nil {
		logging.FromContext(ctx).Infof("Could not delete VerticalPodAutoscaler %s", vpaName)
		errorList = append(errorList, vpaErr)
	}
	if len(errorList) > 0 {
		return errors.NewAggregate(errorList)
	}
	return nil
}

func (vsr *VPASubresourcesReconciler) createPodSelector(ctx context.Context, vpaBinding *alpha1.VPABinding) (*v1alpha1.PodSelector, error){
	podSelectorsClient := vsr.vpaBindingClient.InternalsV1alpha1().PodSelectors(vpaBinding.Namespace)
	name := getPodSelectorName(vpaBinding)
	podSelector, err := podSelectorsClient.Get(ctx, name, metav1.GetOptions{})
	if err == nil {
		logging.FromContext(ctx).Infof("PodSelector %s already exists", name)
		return podSelector, err
	}
	if apierrs.IsNotFound(err) {

		newPodSelector := v1alpha1.PodSelector{
			TypeMeta: metav1.TypeMeta{
				Kind:       "PodSelector",
				APIVersion: v1alpha1.SchemeGroupVersion.String(),
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      name,
				Namespace: vpaBinding.Namespace,
			},
			Spec:   v1alpha1.PodSelectorSpec{
				Replicas: 0,
				Selector: v1alpha1.SelectorString(alpha1.VPABindingLabelKey + "=" + vpaBinding.Name),
			},
			Status: v1alpha1.PodSelectorStatus{},
		}
		logging.FromContext(ctx).Infof("PodSelector %s not found, creating new: %v", name, newPodSelector)

		podSelector, err = podSelectorsClient.Create(ctx, &newPodSelector, metav1.CreateOptions{})
		return &newPodSelector, err
	}
	return podSelector, err
}

func (vsr *VPASubresourcesReconciler) createVPA(ctx context.Context, vpaBinding *alpha1.VPABinding, podSelector *v1alpha1.PodSelector) (*autoscalingv1.VerticalPodAutoscaler, error) {
	vs := vsr.vpaClient.AutoscalingV1().VerticalPodAutoscalers(podSelector.Namespace)
	vpaName := getVPAName(vpaBinding)


	updateMode := autoscalingv1.UpdateModeAuto

	vpa, err := vs.Get(ctx, vpaName, metav1.GetOptions{})
	if err == nil {
		logging.FromContext(ctx).Infof("VPA %s already exists", vpaName)
		// check diff, update labels
		return vpa, err
	}

	if apierrs.IsNotFound(err) {
		vpaObj := &autoscalingv1.VerticalPodAutoscaler{
			ObjectMeta: metav1.ObjectMeta{
				Name:                       vpaName,
				Namespace:                  podSelector.ObjectMeta.Namespace,
			},
			Spec:       autoscalingv1.VerticalPodAutoscalerSpec{
				TargetRef:      &v1.CrossVersionObjectReference{
					Kind:       podSelector.TypeMeta.Kind,
					Name:       podSelector.GetName(),
					APIVersion: podSelector.TypeMeta.APIVersion,
				},
				UpdatePolicy:   &autoscalingv1.PodUpdatePolicy{
					UpdateMode: &updateMode,
				},
				ResourcePolicy: nil,
			},
		}
		logging.FromContext(ctx).Infof("VPA %s not found, creating new", vpaName)
		vpa, err = vs.Create(ctx, vpaObj, metav1.CreateOptions{})
	}

	return vpa, err
}
