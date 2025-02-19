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

type PrivateDNSResolverOutboundEndpointInitParameters struct {

	// Specifies the Azure Region where the Private DNS Resolver Outbound Endpoint should exist. Changing this forces a new Private DNS Resolver Outbound Endpoint to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The ID of the Subnet that is linked to the Private DNS Resolver Outbound Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Private DNS Resolver Outbound Endpoint.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSResolverOutboundEndpointObservation struct {

	// The ID of the Private DNS Resolver Outbound Endpoint.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the Azure Region where the Private DNS Resolver Outbound Endpoint should exist. Changing this forces a new Private DNS Resolver Outbound Endpoint to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the ID of the Private DNS Resolver Outbound Endpoint. Changing this forces a new Private DNS Resolver Outbound Endpoint to be created.
	PrivateDNSResolverID *string `json:"privateDnsResolverId,omitempty" tf:"private_dns_resolver_id,omitempty"`

	// The ID of the Subnet that is linked to the Private DNS Resolver Outbound Endpoint. Changing this forces a new resource to be created.
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// A mapping of tags which should be assigned to the Private DNS Resolver Outbound Endpoint.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSResolverOutboundEndpointParameters struct {

	// Specifies the Azure Region where the Private DNS Resolver Outbound Endpoint should exist. Changing this forces a new Private DNS Resolver Outbound Endpoint to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// Specifies the ID of the Private DNS Resolver Outbound Endpoint. Changing this forces a new Private DNS Resolver Outbound Endpoint to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.PrivateDNSResolver
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	PrivateDNSResolverID *string `json:"privateDnsResolverId,omitempty" tf:"private_dns_resolver_id,omitempty"`

	// Reference to a PrivateDNSResolver in network to populate privateDnsResolverId.
	// +kubebuilder:validation:Optional
	PrivateDNSResolverIDRef *v1.Reference `json:"privateDnsResolverIdRef,omitempty" tf:"-"`

	// Selector for a PrivateDNSResolver in network to populate privateDnsResolverId.
	// +kubebuilder:validation:Optional
	PrivateDNSResolverIDSelector *v1.Selector `json:"privateDnsResolverIdSelector,omitempty" tf:"-"`

	// The ID of the Subnet that is linked to the Private DNS Resolver Outbound Endpoint. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta2.Subnet
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	SubnetID *string `json:"subnetId,omitempty" tf:"subnet_id,omitempty"`

	// Reference to a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDRef *v1.Reference `json:"subnetIdRef,omitempty" tf:"-"`

	// Selector for a Subnet in network to populate subnetId.
	// +kubebuilder:validation:Optional
	SubnetIDSelector *v1.Selector `json:"subnetIdSelector,omitempty" tf:"-"`

	// A mapping of tags which should be assigned to the Private DNS Resolver Outbound Endpoint.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// PrivateDNSResolverOutboundEndpointSpec defines the desired state of PrivateDNSResolverOutboundEndpoint
type PrivateDNSResolverOutboundEndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PrivateDNSResolverOutboundEndpointParameters `json:"forProvider"`
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
	InitProvider PrivateDNSResolverOutboundEndpointInitParameters `json:"initProvider,omitempty"`
}

// PrivateDNSResolverOutboundEndpointStatus defines the observed state of PrivateDNSResolverOutboundEndpoint.
type PrivateDNSResolverOutboundEndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PrivateDNSResolverOutboundEndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PrivateDNSResolverOutboundEndpoint is the Schema for the PrivateDNSResolverOutboundEndpoints API. Manages a Private DNS Resolver Outbound Endpoint.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type PrivateDNSResolverOutboundEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
	Spec   PrivateDNSResolverOutboundEndpointSpec   `json:"spec"`
	Status PrivateDNSResolverOutboundEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PrivateDNSResolverOutboundEndpointList contains a list of PrivateDNSResolverOutboundEndpoints
type PrivateDNSResolverOutboundEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateDNSResolverOutboundEndpoint `json:"items"`
}

// Repository type metadata.
var (
	PrivateDNSResolverOutboundEndpoint_Kind             = "PrivateDNSResolverOutboundEndpoint"
	PrivateDNSResolverOutboundEndpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: PrivateDNSResolverOutboundEndpoint_Kind}.String()
	PrivateDNSResolverOutboundEndpoint_KindAPIVersion   = PrivateDNSResolverOutboundEndpoint_Kind + "." + CRDGroupVersion.String()
	PrivateDNSResolverOutboundEndpoint_GroupVersionKind = CRDGroupVersion.WithKind(PrivateDNSResolverOutboundEndpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&PrivateDNSResolverOutboundEndpoint{}, &PrivateDNSResolverOutboundEndpointList{})
}
