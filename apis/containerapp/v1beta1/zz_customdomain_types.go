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

type CustomDomainInitParameters_2 struct {

	// The Certificate Binding type. Possible values include Disabled and SniEnabled.  Required with container_app_environment_certificate_id. Changing this forces a new resource to be created.
	// The Binding type. Possible values include `Disabled` and `SniEnabled`.
	CertificateBindingType *string `json:"certificateBindingType,omitempty" tf:"certificate_binding_type,omitempty"`

	// The ID of the Container App Environment Certificate to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerapp/v1beta1.EnvironmentCertificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	ContainerAppEnvironmentCertificateID *string `json:"containerAppEnvironmentCertificateId,omitempty" tf:"container_app_environment_certificate_id,omitempty"`

	// Reference to a EnvironmentCertificate in containerapp to populate containerAppEnvironmentCertificateId.
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentCertificateIDRef *v1.Reference `json:"containerAppEnvironmentCertificateIdRef,omitempty" tf:"-"`

	// Selector for a EnvironmentCertificate in containerapp to populate containerAppEnvironmentCertificateId.
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentCertificateIDSelector *v1.Selector `json:"containerAppEnvironmentCertificateIdSelector,omitempty" tf:"-"`
}

type CustomDomainObservation_2 struct {

	// The Certificate Binding type. Possible values include Disabled and SniEnabled.  Required with container_app_environment_certificate_id. Changing this forces a new resource to be created.
	// The Binding type. Possible values include `Disabled` and `SniEnabled`.
	CertificateBindingType *string `json:"certificateBindingType,omitempty" tf:"certificate_binding_type,omitempty"`

	// The ID of the Container App Environment Certificate to use. Changing this forces a new resource to be created.
	ContainerAppEnvironmentCertificateID *string `json:"containerAppEnvironmentCertificateId,omitempty" tf:"container_app_environment_certificate_id,omitempty"`

	// The ID of the Container App to which this Custom Domain should be bound. Changing this forces a new resource to be created.
	ContainerAppID *string `json:"containerAppId,omitempty" tf:"container_app_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type CustomDomainParameters_2 struct {

	// The Certificate Binding type. Possible values include Disabled and SniEnabled.  Required with container_app_environment_certificate_id. Changing this forces a new resource to be created.
	// The Binding type. Possible values include `Disabled` and `SniEnabled`.
	// +kubebuilder:validation:Optional
	CertificateBindingType *string `json:"certificateBindingType,omitempty" tf:"certificate_binding_type,omitempty"`

	// The ID of the Container App Environment Certificate to use. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerapp/v1beta1.EnvironmentCertificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentCertificateID *string `json:"containerAppEnvironmentCertificateId,omitempty" tf:"container_app_environment_certificate_id,omitempty"`

	// Reference to a EnvironmentCertificate in containerapp to populate containerAppEnvironmentCertificateId.
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentCertificateIDRef *v1.Reference `json:"containerAppEnvironmentCertificateIdRef,omitempty" tf:"-"`

	// Selector for a EnvironmentCertificate in containerapp to populate containerAppEnvironmentCertificateId.
	// +kubebuilder:validation:Optional
	ContainerAppEnvironmentCertificateIDSelector *v1.Selector `json:"containerAppEnvironmentCertificateIdSelector,omitempty" tf:"-"`

	// The ID of the Container App to which this Custom Domain should be bound. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/containerapp/v1beta2.ContainerApp
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ContainerAppID *string `json:"containerAppId,omitempty" tf:"container_app_id,omitempty"`

	// Reference to a ContainerApp in containerapp to populate containerAppId.
	// +kubebuilder:validation:Optional
	ContainerAppIDRef *v1.Reference `json:"containerAppIdRef,omitempty" tf:"-"`

	// Selector for a ContainerApp in containerapp to populate containerAppId.
	// +kubebuilder:validation:Optional
	ContainerAppIDSelector *v1.Selector `json:"containerAppIdSelector,omitempty" tf:"-"`
}

// CustomDomainSpec defines the desired state of CustomDomain
type CustomDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CustomDomainParameters_2 `json:"forProvider"`
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
	InitProvider CustomDomainInitParameters_2 `json:"initProvider,omitempty"`
}

// CustomDomainStatus defines the observed state of CustomDomain.
type CustomDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CustomDomainObservation_2 `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// CustomDomain is the Schema for the CustomDomains API. Manages a Container App Custom Domain.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type CustomDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomDomainSpec   `json:"spec"`
	Status            CustomDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CustomDomainList contains a list of CustomDomains
type CustomDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomDomain `json:"items"`
}

// Repository type metadata.
var (
	CustomDomain_Kind             = "CustomDomain"
	CustomDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CustomDomain_Kind}.String()
	CustomDomain_KindAPIVersion   = CustomDomain_Kind + "." + CRDGroupVersion.String()
	CustomDomain_GroupVersionKind = CRDGroupVersion.WithKind(CustomDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&CustomDomain{}, &CustomDomainList{})
}
