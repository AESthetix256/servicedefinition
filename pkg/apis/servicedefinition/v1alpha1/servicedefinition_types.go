package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ServicedefinitionSpec defines the desired state of Servicedefinition
// +k8s:openapi-gen=true
type ServicedefinitionSpec struct {
	Webserver	WebserverSpec	`json:"webserver,omitempty"`
	Database	DatabaseSpec	`json:"database,omitempty"`
	CMS			CMSSpec	`json:"cms,omitempty"`
	Service 	ServiceSpec	`json:"service,omitempty"`
	Domain 		string	`json:"domain,omitempty"`
	Owner 		string 	`json:"owner"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

type DatabaseSpec struct {
	Name	string	`json:"name,omitempty"`
	Version 	string 	`json:"version,omitempty"`
}

type WebserverSpec struct {
	Name	string	`json:"name,omitempty"`
	Version 	string 	`json:"version,omitempty"`
}

type CMSSpec struct {
	Name	string	`json:"name,omitempty"`
	Version 	string 	`json:"version,omitempty"`
}

type ServiceSpec struct {
	Name	string	`json:"name,omitempty"`
	Version 	string 	`json:"version,omitempty"`
}

// ServicedefinitionStatus defines the observed state of Servicedefinition
// +k8s:openapi-gen=true
type ServicedefinitionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Servicedefinition is the Schema for the servicedefinitions API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=servicedefinitions,scope=Namespaced
type Servicedefinition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServicedefinitionSpec   `json:"spec,omitempty"`
	Status ServicedefinitionStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ServicedefinitionList contains a list of Servicedefinition
type ServicedefinitionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Servicedefinition `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Servicedefinition{}, &ServicedefinitionList{})
}
