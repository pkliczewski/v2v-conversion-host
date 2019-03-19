package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html

// V2VVmwareSpec defines the desired state of V2VVmware
// +k8s:openapi-gen=true
type V2VVmwareSpec struct {
	Connection string `json:"connection,omitempty"` // name of Secret wit vmware connection details
	TimeToLive string `json:"timeToLive,omitempty"` // for custom garbage collector
}

// V2VVmwareStatus defines the observed state of V2VVmware
// +k8s:openapi-gen=true
type V2VVmwareStatus struct {
	Phase string `json:"phase,omitempty"` // one of the Phase* constants
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// V2VVmware is the Schema for the v2vvmwares API
// +k8s:openapi-gen=true
type V2VVmware struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   V2VVmwareSpec   `json:"spec,omitempty"`
	Status V2VVmwareStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// V2VVmwareList contains a list of V2VVmware
type V2VVmwareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []V2VVmware `json:"items"`
}

func init() {
	SchemeBuilder.Register(&V2VVmware{}, &V2VVmwareList{})
}
