package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretTemplate is a specification for a SecretTemplate resource
type SecretTemplate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretTemplateSpec   `json:"spec"`
}

// SecretTemplateSpec is the spec for a SecretTemplate resource
type SecretTemplateSpec struct {
	Mike          string `json:"mike"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SecretTemplateList is a list of SecretTemplate resources
type SecretTemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SecretTemplate `json:"items"`
}
