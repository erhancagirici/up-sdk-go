// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ControlPlaneBinding binds an API service represented by an APIServiceExport
// in an Upbound Space into a consumer cluster.
// This object lives in the consumer cluster.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Namespaced,categories={connect},shortName={cpbinding,cpbindings}
type ControlPlaneBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   ControlPlaneBindingSpec   `json:"spec,omitempty"`
	Status ControlPlaneBindingStatus `json:"status,omitempty"`
}

// ControlPlaneBindingList contains a list of ControlPlaneBindings.
// +kubebuilder:object:root=true
type ControlPlaneBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ControlPlaneBinding `json:"items"`
}

// ControlPlaneBindingSpec defines the configuration for an ControlPlaneBinding.
type ControlPlaneBindingSpec struct {
	// +required
	DefaultControlPlaneTarget ControlPlaneTarget `json:"defaultControlPlaneTarget"`
	// +optional
	Targets []LooseControlPlaneTarget `json:"targets,omitempty"`
}

type ControlPlaneTarget struct {
	// +required
	IdentitySecretRef xpv1.SecretKeySelector `json:"identitySecretRef"`
	// +required
	SpaceConfigSecretRef xpv1.SecretKeySelector `json:"spaceConfigSecretRef"`
	// +required
	ControlPlaneRef ControlPlaneRef `json:"controlPlaneRef"`
	// +required
	TargetNamespace string `json:"targetNamespace"`
}

// LooseControlPlaneTarget is the same as ControlPlaneTarget
type LooseControlPlaneTarget struct {
	// +required
	SourceSelector SourceSelector `json:"sourceSelector"`
	// +optional
	IdentitySecretRef *xpv1.SecretKeySelector `json:"identitySecretRef,omitempty"`
	// +optional
	SpaceConfigSecretRef *xpv1.SecretKeySelector `json:"spaceConfigSecretRef,omitempty"`
	// +optional
	ControlPlaneRef *ControlPlaneRef `json:"controlPlaneRef,omitempty"`
	// +optional
	TargetNamespace string `json:"targetNamespace"`
}

type SourceSelector struct {
	// +optional
	// +kubebuilder:validation:MinProperties=1
	MatchLabels map[string]string `json:"matchLabels,omitempty"`
	// +optional
	// +kubebuilder:validation:MinLength=1
	GVK string `json:"gvk,omitempty"`
	// TODO(erhan): validate GVK format?
}

// ControlPlaneRef identifies an ControlPlaneBindingRequest in the
// service provider Upbound Space.
type ControlPlaneRef struct {
	// name is the name of the ControlPlaneBindingRequest object
	// +required
	Name string `json:"name,omitempty"`
	// group is the Space group of the ControlPlaneBindingRequest object
	// +optional
	Group string `json:"group,omitempty"`
}

// ControlPlaneBindingStatus defines the status of a ControlPlaneBinding.
// Reports the ControlPlaneExport that was bound to, and the success status.
type ControlPlaneBindingStatus struct {
	xpv1.ResourceStatus `json:",inline"`

	// Message provides human-readable information about the current status of
	// the ControlPlaneBinding.
	// +optional
	Message string `json:"message,omitempty"`
}

var (
	// ControlPlaneBindingKind is kind of ControlPlaneBinding
	ControlPlaneBindingKind = reflect.TypeOf(ControlPlaneBinding{}).Name()
)

func init() {
	SchemeBuilder.Register(&ControlPlaneBinding{}, &ControlPlaneBindingList{})
}
