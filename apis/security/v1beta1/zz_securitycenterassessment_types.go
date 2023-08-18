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

type SecurityCenterAssessmentInitParameters struct {

	// A map of additional data to associate with the assessment.
	AdditionalData map[string]*string `json:"additionalData,omitempty" tf:"additional_data,omitempty"`

	// A status block as defined below.
	Status []StatusInitParameters `json:"status,omitempty" tf:"status,omitempty"`
}

type SecurityCenterAssessmentObservation struct {

	// A map of additional data to associate with the assessment.
	AdditionalData map[string]*string `json:"additionalData,omitempty" tf:"additional_data,omitempty"`

	// The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
	AssessmentPolicyID *string `json:"assessmentPolicyId,omitempty" tf:"assessment_policy_id,omitempty"`

	// The ID of the Security Center Assessment.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A status block as defined below.
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`

	// The ID of the target resource. Changing this forces a new security Assessment to be created.
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`
}

type SecurityCenterAssessmentParameters struct {

	// A map of additional data to associate with the assessment.
	// +kubebuilder:validation:Optional
	AdditionalData map[string]*string `json:"additionalData,omitempty" tf:"additional_data,omitempty"`

	// The ID of the security Assessment policy to apply to this resource. Changing this forces a new security Assessment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/security/v1beta1.SecurityCenterAssessmentPolicy
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	AssessmentPolicyID *string `json:"assessmentPolicyId,omitempty" tf:"assessment_policy_id,omitempty"`

	// Reference to a SecurityCenterAssessmentPolicy in security to populate assessmentPolicyId.
	// +kubebuilder:validation:Optional
	AssessmentPolicyIDRef *v1.Reference `json:"assessmentPolicyIdRef,omitempty" tf:"-"`

	// Selector for a SecurityCenterAssessmentPolicy in security to populate assessmentPolicyId.
	// +kubebuilder:validation:Optional
	AssessmentPolicyIDSelector *v1.Selector `json:"assessmentPolicyIdSelector,omitempty" tf:"-"`

	// A status block as defined below.
	// +kubebuilder:validation:Optional
	Status []StatusParameters `json:"status,omitempty" tf:"status,omitempty"`

	// The ID of the target resource. Changing this forces a new security Assessment to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/compute/v1beta1.LinuxVirtualMachineScaleSet
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	TargetResourceID *string `json:"targetResourceId,omitempty" tf:"target_resource_id,omitempty"`

	// Reference to a LinuxVirtualMachineScaleSet in compute to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDRef *v1.Reference `json:"targetResourceIdRef,omitempty" tf:"-"`

	// Selector for a LinuxVirtualMachineScaleSet in compute to populate targetResourceId.
	// +kubebuilder:validation:Optional
	TargetResourceIDSelector *v1.Selector `json:"targetResourceIdSelector,omitempty" tf:"-"`
}

type StatusInitParameters struct {

	// Specifies the cause of the assessment status.
	Cause *string `json:"cause,omitempty" tf:"cause,omitempty"`

	// Specifies the programmatic code of the assessment status. Possible values are Healthy, Unhealthy and NotApplicable.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Specifies the human readable description of the assessment status.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type StatusObservation struct {

	// Specifies the cause of the assessment status.
	Cause *string `json:"cause,omitempty" tf:"cause,omitempty"`

	// Specifies the programmatic code of the assessment status. Possible values are Healthy, Unhealthy and NotApplicable.
	Code *string `json:"code,omitempty" tf:"code,omitempty"`

	// Specifies the human readable description of the assessment status.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

type StatusParameters struct {

	// Specifies the cause of the assessment status.
	// +kubebuilder:validation:Optional
	Cause *string `json:"cause,omitempty" tf:"cause,omitempty"`

	// Specifies the programmatic code of the assessment status. Possible values are Healthy, Unhealthy and NotApplicable.
	// +kubebuilder:validation:Optional
	Code *string `json:"code" tf:"code,omitempty"`

	// Specifies the human readable description of the assessment status.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`
}

// SecurityCenterAssessmentSpec defines the desired state of SecurityCenterAssessment
type SecurityCenterAssessmentSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SecurityCenterAssessmentParameters `json:"forProvider"`
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
	InitProvider SecurityCenterAssessmentInitParameters `json:"initProvider,omitempty"`
}

// SecurityCenterAssessmentStatus defines the observed state of SecurityCenterAssessment.
type SecurityCenterAssessmentStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SecurityCenterAssessmentObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessment is the Schema for the SecurityCenterAssessments API. Manages the Security Center Assessment for Azure Security Center.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type SecurityCenterAssessment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.status) || has(self.initProvider.status)",message="status is a required parameter"
	Spec   SecurityCenterAssessmentSpec   `json:"spec"`
	Status SecurityCenterAssessmentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecurityCenterAssessmentList contains a list of SecurityCenterAssessments
type SecurityCenterAssessmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecurityCenterAssessment `json:"items"`
}

// Repository type metadata.
var (
	SecurityCenterAssessment_Kind             = "SecurityCenterAssessment"
	SecurityCenterAssessment_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SecurityCenterAssessment_Kind}.String()
	SecurityCenterAssessment_KindAPIVersion   = SecurityCenterAssessment_Kind + "." + CRDGroupVersion.String()
	SecurityCenterAssessment_GroupVersionKind = CRDGroupVersion.WithKind(SecurityCenterAssessment_Kind)
)

func init() {
	SchemeBuilder.Register(&SecurityCenterAssessment{}, &SecurityCenterAssessmentList{})
}
