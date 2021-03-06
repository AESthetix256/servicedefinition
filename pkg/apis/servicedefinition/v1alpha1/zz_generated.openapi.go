// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/servicedefinition/v1alpha1.CMSSpec":                 schema_pkg_apis_servicedefinition_v1alpha1_CMSSpec(ref),
		"./pkg/apis/servicedefinition/v1alpha1.DatabaseSpec":            schema_pkg_apis_servicedefinition_v1alpha1_DatabaseSpec(ref),
		"./pkg/apis/servicedefinition/v1alpha1.ServiceSpec":             schema_pkg_apis_servicedefinition_v1alpha1_ServiceSpec(ref),
		"./pkg/apis/servicedefinition/v1alpha1.Servicedefinition":       schema_pkg_apis_servicedefinition_v1alpha1_Servicedefinition(ref),
		"./pkg/apis/servicedefinition/v1alpha1.ServicedefinitionSpec":   schema_pkg_apis_servicedefinition_v1alpha1_ServicedefinitionSpec(ref),
		"./pkg/apis/servicedefinition/v1alpha1.ServicedefinitionStatus": schema_pkg_apis_servicedefinition_v1alpha1_ServicedefinitionStatus(ref),
		"./pkg/apis/servicedefinition/v1alpha1.WebserverSpec":           schema_pkg_apis_servicedefinition_v1alpha1_WebserverSpec(ref),
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_CMSSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_DatabaseSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_ServiceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_Servicedefinition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Servicedefinition is the Schema for the servicedefinitions API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/servicedefinition/v1alpha1.ServicedefinitionSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/servicedefinition/v1alpha1.ServicedefinitionStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/servicedefinition/v1alpha1.ServicedefinitionSpec", "./pkg/apis/servicedefinition/v1alpha1.ServicedefinitionStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_ServicedefinitionSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"domain": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifications",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"owner": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"purpose": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"refClass": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"webserver": {
						SchemaProps: spec.SchemaProps{
							Description: "Services",
							Ref:         ref("./pkg/apis/servicedefinition/v1alpha1.WebserverSpec"),
						},
					},
					"database": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/servicedefinition/v1alpha1.DatabaseSpec"),
						},
					},
					"cms": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/servicedefinition/v1alpha1.CMSSpec"),
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/servicedefinition/v1alpha1.ServiceSpec"),
						},
					},
				},
				Required: []string{"owner", "purpose"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/servicedefinition/v1alpha1.CMSSpec", "./pkg/apis/servicedefinition/v1alpha1.DatabaseSpec", "./pkg/apis/servicedefinition/v1alpha1.ServiceSpec", "./pkg/apis/servicedefinition/v1alpha1.WebserverSpec"},
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_ServicedefinitionStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ServicedefinitionStatus defines the observed state of Servicedefinition",
				Type:        []string{"object"},
			},
		},
	}
}

func schema_pkg_apis_servicedefinition_v1alpha1_WebserverSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}
