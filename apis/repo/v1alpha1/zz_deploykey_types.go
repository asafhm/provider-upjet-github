// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeployKeyInitParameters struct {

	// A boolean qualifying the key to be either read only or read/write.
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Name of the GitHub repository.
	// Name of the GitHub repository.
	// +crossplane:generate:reference:type=Repository
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// A title.
	// A title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type DeployKeyObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A boolean qualifying the key to be either read only or read/write.
	// A boolean qualifying the key to be either read only or read/write.
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Name of the GitHub repository.
	// Name of the GitHub repository.
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// A title.
	// A title.
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

type DeployKeyParameters struct {

	// A SSH key.
	// A SSH key.
	// +kubebuilder:validation:Optional
	KeySecretRef v1.SecretKeySelector `json:"keySecretRef" tf:"-"`

	// A boolean qualifying the key to be either read only or read/write.
	// A boolean qualifying the key to be either read only or read/write.
	// +kubebuilder:validation:Optional
	ReadOnly *bool `json:"readOnly,omitempty" tf:"read_only,omitempty"`

	// Name of the GitHub repository.
	// Name of the GitHub repository.
	// +crossplane:generate:reference:type=Repository
	// +kubebuilder:validation:Optional
	Repository *string `json:"repository,omitempty" tf:"repository,omitempty"`

	// Reference to a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositoryRef *v1.Reference `json:"repositoryRef,omitempty" tf:"-"`

	// Selector for a Repository to populate repository.
	// +kubebuilder:validation:Optional
	RepositorySelector *v1.Selector `json:"repositorySelector,omitempty" tf:"-"`

	// A title.
	// A title.
	// +kubebuilder:validation:Optional
	Title *string `json:"title,omitempty" tf:"title,omitempty"`
}

// DeployKeySpec defines the desired state of DeployKey
type DeployKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeployKeyParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider DeployKeyInitParameters `json:"initProvider,omitempty"`
}

// DeployKeyStatus defines the observed state of DeployKey.
type DeployKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeployKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DeployKey is the Schema for the DeployKeys API. Provides a GitHub repository deploy key resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,github}
type DeployKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.keySecretRef)",message="spec.forProvider.keySecretRef is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.title) || (has(self.initProvider) && has(self.initProvider.title))",message="spec.forProvider.title is a required parameter"
	Spec   DeployKeySpec   `json:"spec"`
	Status DeployKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeployKeyList contains a list of DeployKeys
type DeployKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployKey `json:"items"`
}

// Repository type metadata.
var (
	DeployKey_Kind             = "DeployKey"
	DeployKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeployKey_Kind}.String()
	DeployKey_KindAPIVersion   = DeployKey_Kind + "." + CRDGroupVersion.String()
	DeployKey_GroupVersionKind = CRDGroupVersion.WithKind(DeployKey_Kind)
)

func init() {
	SchemeBuilder.Register(&DeployKey{}, &DeployKeyList{})
}
