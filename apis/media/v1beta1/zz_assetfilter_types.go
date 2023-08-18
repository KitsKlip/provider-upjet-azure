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

type AssetFilterInitParameters struct {

	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate *float64 `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate,omitempty"`

	// A presentation_time_range block as defined below.
	PresentationTimeRange []PresentationTimeRangeInitParameters `json:"presentationTimeRange,omitempty" tf:"presentation_time_range,omitempty"`

	// One or more track_selection blocks as defined below.
	TrackSelection []TrackSelectionInitParameters `json:"trackSelection,omitempty" tf:"track_selection,omitempty"`
}

type AssetFilterObservation struct {

	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	AssetID *string `json:"assetId,omitempty" tf:"asset_id,omitempty"`

	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	FirstQualityBitrate *float64 `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate,omitempty"`

	// The ID of the Asset Filter.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A presentation_time_range block as defined below.
	PresentationTimeRange []PresentationTimeRangeObservation `json:"presentationTimeRange,omitempty" tf:"presentation_time_range,omitempty"`

	// One or more track_selection blocks as defined below.
	TrackSelection []TrackSelectionObservation `json:"trackSelection,omitempty" tf:"track_selection,omitempty"`
}

type AssetFilterParameters struct {

	// The Asset ID for which the Asset Filter should be created. Changing this forces a new Asset Filter to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/media/v1beta1.Asset
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AssetID *string `json:"assetId,omitempty" tf:"asset_id,omitempty"`

	// Reference to a Asset in media to populate assetId.
	// +kubebuilder:validation:Optional
	AssetIDRef *v1.Reference `json:"assetIdRef,omitempty" tf:"-"`

	// Selector for a Asset in media to populate assetId.
	// +kubebuilder:validation:Optional
	AssetIDSelector *v1.Selector `json:"assetIdSelector,omitempty" tf:"-"`

	// The first quality bitrate. Sets the first video track to appear in the Live Streaming playlist to allow HLS native players to start downloading from this quality level at the beginning.
	// +kubebuilder:validation:Optional
	FirstQualityBitrate *float64 `json:"firstQualityBitrate,omitempty" tf:"first_quality_bitrate,omitempty"`

	// A presentation_time_range block as defined below.
	// +kubebuilder:validation:Optional
	PresentationTimeRange []PresentationTimeRangeParameters `json:"presentationTimeRange,omitempty" tf:"presentation_time_range,omitempty"`

	// One or more track_selection blocks as defined below.
	// +kubebuilder:validation:Optional
	TrackSelection []TrackSelectionParameters `json:"trackSelection,omitempty" tf:"track_selection,omitempty"`
}

type ConditionInitParameters struct {

	// The condition operation to test a track property against. Supported values are Equal and NotEqual.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The track property to compare. Supported values are Bitrate, FourCC, Language, Name and Type. Check documentation for more details.
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// The track property value to match or not match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionObservation struct {

	// The condition operation to test a track property against. Supported values are Equal and NotEqual.
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The track property to compare. Supported values are Bitrate, FourCC, Language, Name and Type. Check documentation for more details.
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// The track property value to match or not match.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ConditionParameters struct {

	// The condition operation to test a track property against. Supported values are Equal and NotEqual.
	// +kubebuilder:validation:Optional
	Operation *string `json:"operation,omitempty" tf:"operation,omitempty"`

	// The track property to compare. Supported values are Bitrate, FourCC, Language, Name and Type. Check documentation for more details.
	// +kubebuilder:validation:Optional
	Property *string `json:"property,omitempty" tf:"property,omitempty"`

	// The track property value to match or not match.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type PresentationTimeRangeInitParameters struct {

	// The absolute end time boundary. Applies to Video on Demand (VoD).
	// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by unit_timescale_in_miliseconds, so an end_in_units of 180 would be for 3 minutes. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_miliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	EndInUnits *float64 `json:"endInUnits,omitempty" tf:"end_in_units,omitempty"`

	// Indicates whether the end_in_units property must be present. If true, end_in_units must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end,omitempty"`

	// The relative to end right edge. Applies to Live Streaming only.
	// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by unit_timescale_in_miliseconds. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
	LiveBackoffInUnits *float64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units,omitempty"`

	// The relative to end sliding window. Applies to Live Streaming only. Use presentation_window_in_units to apply a sliding window of fragments to include in a playlist. The unit is defined by unit_timescale_in_miliseconds. For example, set presentation_window_in_units to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
	PresentationWindowInUnits *float64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units,omitempty"`

	// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by unit_timescale_in_miliseconds, so a start_in_units of 15 would be for 15 seconds. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_miliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	StartInUnits *float64 `json:"startInUnits,omitempty" tf:"start_in_units,omitempty"`

	// Specified as the number of miliseconds in one unit timescale. For example, if you want to set a start_in_units at 30 seconds, you would use a value of 30 when using the unit_timescale_in_miliseconds in 1000. Or if you want to set start_in_units in 30 miliseconds, you would use a value of 30 when using the unit_timescale_in_miliseconds in 1. Applies timescale to start_in_units, start_timescale and presentation_window_in_timescale and live_backoff_in_timescale.
	UnitTimescaleInMiliseconds *float64 `json:"unitTimescaleInMiliseconds,omitempty" tf:"unit_timescale_in_miliseconds,omitempty"`
}

type PresentationTimeRangeObservation struct {

	// The absolute end time boundary. Applies to Video on Demand (VoD).
	// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by unit_timescale_in_miliseconds, so an end_in_units of 180 would be for 3 minutes. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_miliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	EndInUnits *float64 `json:"endInUnits,omitempty" tf:"end_in_units,omitempty"`

	// Indicates whether the end_in_units property must be present. If true, end_in_units must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end,omitempty"`

	// The relative to end right edge. Applies to Live Streaming only.
	// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by unit_timescale_in_miliseconds. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
	LiveBackoffInUnits *float64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units,omitempty"`

	// The relative to end sliding window. Applies to Live Streaming only. Use presentation_window_in_units to apply a sliding window of fragments to include in a playlist. The unit is defined by unit_timescale_in_miliseconds. For example, set presentation_window_in_units to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
	PresentationWindowInUnits *float64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units,omitempty"`

	// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by unit_timescale_in_miliseconds, so a start_in_units of 15 would be for 15 seconds. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_miliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	StartInUnits *float64 `json:"startInUnits,omitempty" tf:"start_in_units,omitempty"`

	// Specified as the number of miliseconds in one unit timescale. For example, if you want to set a start_in_units at 30 seconds, you would use a value of 30 when using the unit_timescale_in_miliseconds in 1000. Or if you want to set start_in_units in 30 miliseconds, you would use a value of 30 when using the unit_timescale_in_miliseconds in 1. Applies timescale to start_in_units, start_timescale and presentation_window_in_timescale and live_backoff_in_timescale.
	UnitTimescaleInMiliseconds *float64 `json:"unitTimescaleInMiliseconds,omitempty" tf:"unit_timescale_in_miliseconds,omitempty"`
}

type PresentationTimeRangeParameters struct {

	// The absolute end time boundary. Applies to Video on Demand (VoD).
	// For the Live Streaming presentation, it is silently ignored and applied when the presentation ends and the stream becomes VoD. This is a long value that represents an absolute end point of the presentation, rounded to the closest next GOP start. The unit is defined by unit_timescale_in_miliseconds, so an end_in_units of 180 would be for 3 minutes. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_miliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	// +kubebuilder:validation:Optional
	EndInUnits *float64 `json:"endInUnits,omitempty" tf:"end_in_units,omitempty"`

	// Indicates whether the end_in_units property must be present. If true, end_in_units must be specified or a bad request code is returned. Applies to Live Streaming only. Allowed values: false, true.
	// +kubebuilder:validation:Optional
	ForceEnd *bool `json:"forceEnd,omitempty" tf:"force_end,omitempty"`

	// The relative to end right edge. Applies to Live Streaming only.
	// This value defines the latest live position that a client can seek to. Using this property, you can delay live playback position and create a server-side buffer for players. The unit is defined by unit_timescale_in_miliseconds. The maximum live back off duration is 300 seconds. For example, a value of 20 means that the latest available content is 20 seconds delayed from the real live edge.
	// +kubebuilder:validation:Optional
	LiveBackoffInUnits *float64 `json:"liveBackoffInUnits,omitempty" tf:"live_backoff_in_units,omitempty"`

	// The relative to end sliding window. Applies to Live Streaming only. Use presentation_window_in_units to apply a sliding window of fragments to include in a playlist. The unit is defined by unit_timescale_in_miliseconds. For example, set presentation_window_in_units to 120 to apply a two-minute sliding window. Media within 2 minutes of the live edge will be included in the playlist. If a fragment straddles the boundary, the entire fragment will be included in the playlist. The minimum presentation window duration is 60 seconds.
	// +kubebuilder:validation:Optional
	PresentationWindowInUnits *float64 `json:"presentationWindowInUnits,omitempty" tf:"presentation_window_in_units,omitempty"`

	// The absolute start time boundary. Applies to Video on Demand (VoD) or Live Streaming. This is a long value that represents an absolute start point of the stream. The value gets rounded to the closest next GOP start. The unit is defined by unit_timescale_in_miliseconds, so a start_in_units of 15 would be for 15 seconds. Use start_in_units and end_in_units to trim the fragments that will be in the playlist (manifest). For example, start_in_units set to 20 and end_in_units set to 60 using unit_timescale_in_miliseconds in 1000 will generate a playlist that contains fragments from between 20 seconds and 60 seconds of the VoD presentation. If a fragment straddles the boundary, the entire fragment will be included in the manifest.
	// +kubebuilder:validation:Optional
	StartInUnits *float64 `json:"startInUnits,omitempty" tf:"start_in_units,omitempty"`

	// Specified as the number of miliseconds in one unit timescale. For example, if you want to set a start_in_units at 30 seconds, you would use a value of 30 when using the unit_timescale_in_miliseconds in 1000. Or if you want to set start_in_units in 30 miliseconds, you would use a value of 30 when using the unit_timescale_in_miliseconds in 1. Applies timescale to start_in_units, start_timescale and presentation_window_in_timescale and live_backoff_in_timescale.
	// +kubebuilder:validation:Optional
	UnitTimescaleInMiliseconds *float64 `json:"unitTimescaleInMiliseconds,omitempty" tf:"unit_timescale_in_miliseconds,omitempty"`
}

type TrackSelectionInitParameters struct {

	// One or more condition blocks as defined above.
	Condition []ConditionInitParameters `json:"condition,omitempty" tf:"condition,omitempty"`
}

type TrackSelectionObservation struct {

	// One or more condition blocks as defined above.
	Condition []ConditionObservation `json:"condition,omitempty" tf:"condition,omitempty"`
}

type TrackSelectionParameters struct {

	// One or more condition blocks as defined above.
	// +kubebuilder:validation:Optional
	Condition []ConditionParameters `json:"condition" tf:"condition,omitempty"`
}

// AssetFilterSpec defines the desired state of AssetFilter
type AssetFilterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AssetFilterParameters `json:"forProvider"`
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
	InitProvider AssetFilterInitParameters `json:"initProvider,omitempty"`
}

// AssetFilterStatus defines the observed state of AssetFilter.
type AssetFilterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AssetFilterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AssetFilter is the Schema for the AssetFilters API. Manages an Azure Media Asset Filter.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AssetFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AssetFilterSpec   `json:"spec"`
	Status            AssetFilterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AssetFilterList contains a list of AssetFilters
type AssetFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AssetFilter `json:"items"`
}

// Repository type metadata.
var (
	AssetFilter_Kind             = "AssetFilter"
	AssetFilter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AssetFilter_Kind}.String()
	AssetFilter_KindAPIVersion   = AssetFilter_Kind + "." + CRDGroupVersion.String()
	AssetFilter_GroupVersionKind = CRDGroupVersion.WithKind(AssetFilter_Kind)
)

func init() {
	SchemeBuilder.Register(&AssetFilter{}, &AssetFilterList{})
}
