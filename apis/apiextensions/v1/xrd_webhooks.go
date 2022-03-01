/*
Copyright 2022 The Crossplane Authors.

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

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/errors"
)

// +kubebuilder:webhook:verbs=create;update;delete,path=/validate-apiextensions-crossplane-io-v1-compositeresourcedefinition,mutating=false,failurePolicy=fail,groups=apiextensions.crossplane.io,resources=compositeresourcedefinitions,versions=v1,name=compositeresourcedefinitions.apiextensions.crossplane.io,sideEffects=None,admissionReviewVersions=v1

// ValidateCreate is run for creation actions.
func (in *CompositeResourceDefinition) ValidateCreate() error {
	return nil
}

// ValidateUpdate is run for update actions.
func (in *CompositeResourceDefinition) ValidateUpdate(old runtime.Object) error {
	oldObj, ok := old.(*CompositeResourceDefinition)
	if !ok {
		return errors.Errorf("unexpected type")
	}
	if in.Spec.Group != oldObj.Spec.Group {
		return errors.Errorf("spec.group is immutable")
	}
	return nil
}

// ValidateDelete is run for delete actions.
func (in *CompositeResourceDefinition) ValidateDelete() error {
	return nil
}

// SetupWebhookWithManager sets up webhook with manager.
func (in *CompositeResourceDefinition) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(in).
		Complete()
}
