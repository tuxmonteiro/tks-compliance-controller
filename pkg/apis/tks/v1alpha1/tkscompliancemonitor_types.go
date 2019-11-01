package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TksComplianceMonitorSpec defines the desired state of TksComplianceMonitor
// +k8s:openapi-gen=true
type TksComplianceMonitorSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// TksComplianceMonitorStatus defines the observed state of TksComplianceMonitor
// +k8s:openapi-gen=true
type TksComplianceMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TksComplianceMonitor is the Schema for the tkscompliancemonitors API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=tkscompliancemonitors,scope=Namespaced
type TksComplianceMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TksComplianceMonitorSpec   `json:"spec,omitempty"`
	Status TksComplianceMonitorStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TksComplianceMonitorList contains a list of TksComplianceMonitor
type TksComplianceMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TksComplianceMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TksComplianceMonitor{}, &TksComplianceMonitorList{})
}
