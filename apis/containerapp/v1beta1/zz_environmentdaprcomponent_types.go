// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type EnvironmentDaprComponentInitParameters struct {

	// The Dapr Component Type. For example state.azure.blobstorage. Changing this forces a new resource to be created.
	// The Dapr Component Type. For example `state.azure.blobstorage`.
	ComponentType *string `json:"componentType,omitempty" tf:"component_type,omitempty"`

	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to false
	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to `false`
	IgnoreErrors *bool `json:"ignoreErrors,omitempty" tf:"ignore_errors,omitempty"`

	// The timeout for component initialisation as a ISO8601 formatted string. e.g. 5s, 2h, 1m. Defaults to 5s.
	// The component initialisation timeout in ISO8601 format. e.g. `5s`, `2h`, `1m`. Defaults to `5s`.
	InitTimeout *string `json:"initTimeout,omitempty" tf:"init_timeout,omitempty"`

	// One or more metadata blocks as detailed below.
	Metadata []MetadataInitParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A list of scopes to which this component applies.
	// A list of scopes to which this component applies. e.g. a Container App's `dapr.app_id` value.
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// A secret block as detailed below.
	Secret []EnvironmentDaprComponentSecretInitParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// The version of the component.
	// The version of the component.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type EnvironmentDaprComponentObservation struct {

	// The Dapr Component Type. For example state.azure.blobstorage. Changing this forces a new resource to be created.
	// The Dapr Component Type. For example `state.azure.blobstorage`.
	ComponentType *string `json:"componentType,omitempty" tf:"component_type,omitempty"`

	// The ID of the Container App Managed Environment for this Dapr Component. Changing this forces a new resource to be created.
	// The Container App Managed Environment ID to configure this Dapr component on.
	ContainerAppEnvironmentID *string `json:"containerAppEnvironmentId,omitempty" tf:"container_app_environment_id,omitempty"`

	// The ID of the Container App Environment Dapr Component
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to false
	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to `false`
	IgnoreErrors *bool `json:"ignoreErrors,omitempty" tf:"ignore_errors,omitempty"`

	// The timeout for component initialisation as a ISO8601 formatted string. e.g. 5s, 2h, 1m. Defaults to 5s.
	// The component initialisation timeout in ISO8601 format. e.g. `5s`, `2h`, `1m`. Defaults to `5s`.
	InitTimeout *string `json:"initTimeout,omitempty" tf:"init_timeout,omitempty"`

	// One or more metadata blocks as detailed below.
	Metadata []MetadataObservation `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A list of scopes to which this component applies.
	// A list of scopes to which this component applies. e.g. a Container App's `dapr.app_id` value.
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// A secret block as detailed below.
	Secret []EnvironmentDaprComponentSecretObservation `json:"secret,omitempty" tf:"secret,omitempty"`

	// The version of the component.
	// The version of the component.
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type EnvironmentDaprComponentParameters struct {

	// The Dapr Component Type. For example state.azure.blobstorage. Changing this forces a new resource to be created.
	// The Dapr Component Type. For example `state.azure.blobstorage`.
	// +kubebuilder:validation:Optional
	ComponentType *string `json:"componentType,omitempty" tf:"component_type,omitempty"`

	// The ID of the Container App Managed Environment for this Dapr Component. Changing this forces a new resource to be created.
	// The Container App Managed Environment ID to configure this Dapr component on.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerapp/v1beta1.Environment
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentID *string `json:"containerAppEnvironmentId,omitempty" tf:"container_app_environment_id,omitempty"`

	// Reference to a Environment in containerapp to populate containerAppEnvironmentId.
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentIDRef *v1.Reference `json:"containerAppEnvironmentIdRef,omitempty" tf:"-"`

	// Selector for a Environment in containerapp to populate containerAppEnvironmentId.
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentIDSelector *v1.Selector `json:"containerAppEnvironmentIdSelector,omitempty" tf:"-"`

	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to false
	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to `false`
	// +kubebuilder:validation:Optional
	IgnoreErrors *bool `json:"ignoreErrors,omitempty" tf:"ignore_errors,omitempty"`

	// The timeout for component initialisation as a ISO8601 formatted string. e.g. 5s, 2h, 1m. Defaults to 5s.
	// The component initialisation timeout in ISO8601 format. e.g. `5s`, `2h`, `1m`. Defaults to `5s`.
	// +kubebuilder:validation:Optional
	InitTimeout *string `json:"initTimeout,omitempty" tf:"init_timeout,omitempty"`

	// One or more metadata blocks as detailed below.
	// +kubebuilder:validation:Optional
	Metadata []MetadataParameters `json:"metadata,omitempty" tf:"metadata,omitempty"`

	// A list of scopes to which this component applies.
	// A list of scopes to which this component applies. e.g. a Container App's `dapr.app_id` value.
	// +kubebuilder:validation:Optional
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// A secret block as detailed below.
	// +kubebuilder:validation:Optional
	Secret []EnvironmentDaprComponentSecretParameters `json:"secret,omitempty" tf:"secret,omitempty"`

	// The version of the component.
	// The version of the component.
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type EnvironmentDaprComponentSecretInitParameters struct {

	// The identity to use for accessing key vault reference.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Container App Environment Dapr Component
	// The Key Vault Secret ID. Could be either one of `id` or `versionless_id`.
	KeyVaultSecretID *string `json:"keyVaultSecretId,omitempty" tf:"key_vault_secret_id,omitempty"`

	// The Secret name.
	// The secret name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value for this secret.
	// The value for this secret.
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type EnvironmentDaprComponentSecretObservation struct {

	// The identity to use for accessing key vault reference.
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Container App Environment Dapr Component
	// The Key Vault Secret ID. Could be either one of `id` or `versionless_id`.
	KeyVaultSecretID *string `json:"keyVaultSecretId,omitempty" tf:"key_vault_secret_id,omitempty"`

	// The Secret name.
	// The secret name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type EnvironmentDaprComponentSecretParameters struct {

	// The identity to use for accessing key vault reference.
	// +kubebuilder:validation:Optional
	Identity *string `json:"identity,omitempty" tf:"identity,omitempty"`

	// The ID of the Container App Environment Dapr Component
	// The Key Vault Secret ID. Could be either one of `id` or `versionless_id`.
	// +kubebuilder:validation:Optional
	KeyVaultSecretID *string `json:"keyVaultSecretId,omitempty" tf:"key_vault_secret_id,omitempty"`

	// The Secret name.
	// The secret name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The value for this secret.
	// The value for this secret.
	// +kubebuilder:validation:Optional
	ValueSecretRef *v1.SecretKeySelector `json:"valueSecretRef,omitempty" tf:"-"`
}

type MetadataInitParameters struct {

	// The name of the Metadata configuration item.
	// The name of the Metadata configuration item.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of a secret specified in the secrets block that contains the value for this metadata configuration item.
	// The name of a secret specified in the `secrets` block that contains the value for this metadata configuration item.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// The value for this metadata configuration item.
	// The value for this metadata configuration item.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MetadataObservation struct {

	// The name of the Metadata configuration item.
	// The name of the Metadata configuration item.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The name of a secret specified in the secrets block that contains the value for this metadata configuration item.
	// The name of a secret specified in the `secrets` block that contains the value for this metadata configuration item.
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// The value for this metadata configuration item.
	// The value for this metadata configuration item.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type MetadataParameters struct {

	// The name of the Metadata configuration item.
	// The name of the Metadata configuration item.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The name of a secret specified in the secrets block that contains the value for this metadata configuration item.
	// The name of a secret specified in the `secrets` block that contains the value for this metadata configuration item.
	// +kubebuilder:validation:Optional
	SecretName *string `json:"secretName,omitempty" tf:"secret_name,omitempty"`

	// The value for this metadata configuration item.
	// The value for this metadata configuration item.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

// EnvironmentDaprComponentSpec defines the desired state of EnvironmentDaprComponent
type EnvironmentDaprComponentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EnvironmentDaprComponentParameters `json:"forProvider"`
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
	InitProvider EnvironmentDaprComponentInitParameters `json:"initProvider,omitempty"`
}

// EnvironmentDaprComponentStatus defines the observed state of EnvironmentDaprComponent.
type EnvironmentDaprComponentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EnvironmentDaprComponentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// EnvironmentDaprComponent is the Schema for the EnvironmentDaprComponents API. Manages a Dapr Component for a Container App Environment.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type EnvironmentDaprComponent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.componentType) || (has(self.initProvider) && has(self.initProvider.componentType))",message="spec.forProvider.componentType is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.version) || (has(self.initProvider) && has(self.initProvider.version))",message="spec.forProvider.version is a required parameter"
	Spec   EnvironmentDaprComponentSpec   `json:"spec"`
	Status EnvironmentDaprComponentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EnvironmentDaprComponentList contains a list of EnvironmentDaprComponents
type EnvironmentDaprComponentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EnvironmentDaprComponent `json:"items"`
}

// Repository type metadata.
var (
	EnvironmentDaprComponent_Kind             = "EnvironmentDaprComponent"
	EnvironmentDaprComponent_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EnvironmentDaprComponent_Kind}.String()
	EnvironmentDaprComponent_KindAPIVersion   = EnvironmentDaprComponent_Kind + "." + CRDGroupVersion.String()
	EnvironmentDaprComponent_GroupVersionKind = CRDGroupVersion.WithKind(EnvironmentDaprComponent_Kind)
)

func init() {
	SchemeBuilder.Register(&EnvironmentDaprComponent{}, &EnvironmentDaprComponentList{})
}
