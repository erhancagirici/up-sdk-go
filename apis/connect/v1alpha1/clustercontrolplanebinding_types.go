// Copyright 2024 Upbound Inc.
// All rights reserved

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterControlPlaneBinding binds an API service represented by an APIServiceExport
// in an Upbound Space into a consumer cluster.
// This object lives in the consumer cluster.
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=`.status.message`
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={connect},shortName={clustercpbinding,clustercpbindings}
type ClusterControlPlaneBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +required
	Spec   ControlPlaneBindingSpec   `json:"spec,omitempty"`
	Status ControlPlaneBindingStatus `json:"status,omitempty"`
}

// ClusterControlPlaneBindingList contains a list of ControlPlaneBindings.
// +kubebuilder:object:root=true
type ClusterControlPlaneBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterControlPlaneBinding `json:"items"`
}

var (
	// ClusterControlPlaneBindingKind is kind of ClusterControlPlaneBinding
	ClusterControlPlaneBindingKind = reflect.TypeOf(ClusterControlPlaneBinding{}).Name()
)

func init() {
	SchemeBuilder.Register(&ClusterControlPlaneBinding{}, &ClusterControlPlaneBindingList{})
}
