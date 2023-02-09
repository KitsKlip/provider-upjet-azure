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

type VirtualHubRouteTableObservation struct {

	// The ID of the Virtual Hub Route Table.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type VirtualHubRouteTableParameters struct {

	// List of labels associated with this route table.
	// +kubebuilder:validation:Optional
	Labels []*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// A route block as defined below.
	// +kubebuilder:validation:Optional
	Route []VirtualHubRouteTableRouteParameters `json:"route,omitempty" tf:"route,omitempty"`

	// The ID of the Virtual Hub within which this route table should be created. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHub
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	VirtualHubID *string `json:"virtualHubId,omitempty" tf:"virtual_hub_id,omitempty"`

	// Reference to a VirtualHub in network to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDRef *v1.Reference `json:"virtualHubIdRef,omitempty" tf:"-"`

	// Selector for a VirtualHub in network to populate virtualHubId.
	// +kubebuilder:validation:Optional
	VirtualHubIDSelector *v1.Selector `json:"virtualHubIdSelector,omitempty" tf:"-"`
}

type VirtualHubRouteTableRouteObservation struct {
}

type VirtualHubRouteTableRouteParameters struct {

	// A list of destination addresses for this route.
	// +kubebuilder:validation:Required
	Destinations []*string `json:"destinations" tf:"destinations,omitempty"`

	// The type of destinations. Possible values are CIDR, ResourceId and Service.
	// +kubebuilder:validation:Required
	DestinationsType *string `json:"destinationsType" tf:"destinations_type,omitempty"`

	// The name which should be used for this route.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The next hop's resource ID.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/network/v1beta1.VirtualHubConnection
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	NextHop *string `json:"nextHop,omitempty" tf:"next_hop,omitempty"`

	// Reference to a VirtualHubConnection in network to populate nextHop.
	// +kubebuilder:validation:Optional
	NextHopRef *v1.Reference `json:"nextHopRef,omitempty" tf:"-"`

	// Selector for a VirtualHubConnection in network to populate nextHop.
	// +kubebuilder:validation:Optional
	NextHopSelector *v1.Selector `json:"nextHopSelector,omitempty" tf:"-"`

	// The type of next hop. Currently the only possible value is ResourceId. Defaults to ResourceId.
	// +kubebuilder:validation:Optional
	NextHopType *string `json:"nextHopType,omitempty" tf:"next_hop_type,omitempty"`
}

// VirtualHubRouteTableSpec defines the desired state of VirtualHubRouteTable
type VirtualHubRouteTableSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VirtualHubRouteTableParameters `json:"forProvider"`
}

// VirtualHubRouteTableStatus defines the observed state of VirtualHubRouteTable.
type VirtualHubRouteTableStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VirtualHubRouteTableObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubRouteTable is the Schema for the VirtualHubRouteTables API. Manages a Virtual Hub Route Table.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type VirtualHubRouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VirtualHubRouteTableSpec   `json:"spec"`
	Status            VirtualHubRouteTableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualHubRouteTableList contains a list of VirtualHubRouteTables
type VirtualHubRouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualHubRouteTable `json:"items"`
}

// Repository type metadata.
var (
	VirtualHubRouteTable_Kind             = "VirtualHubRouteTable"
	VirtualHubRouteTable_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VirtualHubRouteTable_Kind}.String()
	VirtualHubRouteTable_KindAPIVersion   = VirtualHubRouteTable_Kind + "." + CRDGroupVersion.String()
	VirtualHubRouteTable_GroupVersionKind = CRDGroupVersion.WithKind(VirtualHubRouteTable_Kind)
)

func init() {
	SchemeBuilder.Register(&VirtualHubRouteTable{}, &VirtualHubRouteTableList{})
}
