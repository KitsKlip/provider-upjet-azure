/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type IntegrationRuntimeInitParameters struct {

	// The integration runtime reference to associate with the Data Factory Linked Service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the integration runtime.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type IntegrationRuntimeObservation struct {

	// The integration runtime reference to associate with the Data Factory Linked Service.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// A map of parameters to associate with the integration runtime.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type IntegrationRuntimeParameters struct {

	// The integration runtime reference to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// A map of parameters to associate with the integration runtime.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`
}

type LinkedCustomServiceInitParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An integration_runtime block as defined below.
	IntegrationRuntime []IntegrationRuntimeInitParameters `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The type of data stores that will be connected to Data Factory. For full list of supported data stores, please refer to Azure Data Factory connector. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A JSON object that contains the properties of the Data Factory Linked Service.
	TypePropertiesJSON *string `json:"typePropertiesJson,omitempty" tf:"type_properties_json,omitempty"`
}

type LinkedCustomServiceObservation struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// The description for the Data Factory Linked Service.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The ID of the Data Factory Linked Service.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// An integration_runtime block as defined below.
	IntegrationRuntime []IntegrationRuntimeObservation `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The type of data stores that will be connected to Data Factory. For full list of supported data stores, please refer to Azure Data Factory connector. Changing this forces a new resource to be created.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A JSON object that contains the properties of the Data Factory Linked Service.
	TypePropertiesJSON *string `json:"typePropertiesJson,omitempty" tf:"type_properties_json,omitempty"`
}

type LinkedCustomServiceParameters struct {

	// A map of additional properties to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	AdditionalProperties map[string]*string `json:"additionalProperties,omitempty" tf:"additional_properties,omitempty"`

	// List of tags that can be used for describing the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Annotations []*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// The Data Factory ID in which to associate the Linked Service with. Changing this forces a new resource.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/datafactory/v1beta1.Factory
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DataFactoryID *string `json:"dataFactoryId,omitempty" tf:"data_factory_id,omitempty"`

	// Reference to a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDRef *v1.Reference `json:"dataFactoryIdRef,omitempty" tf:"-"`

	// Selector for a Factory in datafactory to populate dataFactoryId.
	// +kubebuilder:validation:Optional
	DataFactoryIDSelector *v1.Selector `json:"dataFactoryIdSelector,omitempty" tf:"-"`

	// The description for the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// An integration_runtime block as defined below.
	// +kubebuilder:validation:Optional
	IntegrationRuntime []IntegrationRuntimeParameters `json:"integrationRuntime,omitempty" tf:"integration_runtime,omitempty"`

	// A map of parameters to associate with the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	Parameters map[string]*string `json:"parameters,omitempty" tf:"parameters,omitempty"`

	// The type of data stores that will be connected to Data Factory. For full list of supported data stores, please refer to Azure Data Factory connector. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// A JSON object that contains the properties of the Data Factory Linked Service.
	// +kubebuilder:validation:Optional
	TypePropertiesJSON *string `json:"typePropertiesJson,omitempty" tf:"type_properties_json,omitempty"`
}

// LinkedCustomServiceSpec defines the desired state of LinkedCustomService
type LinkedCustomServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkedCustomServiceParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider LinkedCustomServiceInitParameters `json:"initProvider,omitempty"`
}

// LinkedCustomServiceStatus defines the observed state of LinkedCustomService.
type LinkedCustomServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkedCustomServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedCustomService is the Schema for the LinkedCustomServices API. Manages a Linked Service (connection) between a resource and Azure Data Factory. This is a generic resource that supports all different Linked Service Types.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type LinkedCustomService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || has(self.initProvider.type)",message="type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.typePropertiesJson) || has(self.initProvider.typePropertiesJson)",message="typePropertiesJson is a required parameter"
	Spec   LinkedCustomServiceSpec   `json:"spec"`
	Status LinkedCustomServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkedCustomServiceList contains a list of LinkedCustomServices
type LinkedCustomServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkedCustomService `json:"items"`
}

// Repository type metadata.
var (
	LinkedCustomService_Kind             = "LinkedCustomService"
	LinkedCustomService_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkedCustomService_Kind}.String()
	LinkedCustomService_KindAPIVersion   = LinkedCustomService_Kind + "." + CRDGroupVersion.String()
	LinkedCustomService_GroupVersionKind = CRDGroupVersion.WithKind(LinkedCustomService_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkedCustomService{}, &LinkedCustomServiceList{})
}
