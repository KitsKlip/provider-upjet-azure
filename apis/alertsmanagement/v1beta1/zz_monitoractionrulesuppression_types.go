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

type ConditionAlertContextObservation struct {
}

type ConditionAlertContextParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionAlertRuleIDObservation struct {
}

type ConditionAlertRuleIDParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionDescriptionObservation struct {
}

type ConditionDescriptionParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorObservation struct {
}

type ConditionMonitorParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionMonitorServiceObservation struct {
}

type ConditionMonitorServiceParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionSeverityObservation struct {
}

type ConditionSeverityParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type ConditionTargetResourceTypeObservation struct {
}

type ConditionTargetResourceTypeParameters struct {

	// The operator for a given condition. Possible values are Equals and NotEquals.
	// +kubebuilder:validation:Required
	Operator *string `json:"operator" tf:"operator,omitempty"`

	// A list of values to match for a given condition. The values should be valid resource types.
	// +kubebuilder:validation:Required
	Values []*string `json:"values" tf:"values,omitempty"`
}

type MonitorActionRuleSuppressionConditionObservation struct {
}

type MonitorActionRuleSuppressionConditionParameters struct {

	// A alert_context block as defined below.
	// +kubebuilder:validation:Optional
	AlertContext []ConditionAlertContextParameters `json:"alertContext,omitempty" tf:"alert_context,omitempty"`

	// A alert_rule_id block as defined below.
	// +kubebuilder:validation:Optional
	AlertRuleID []ConditionAlertRuleIDParameters `json:"alertRuleId,omitempty" tf:"alert_rule_id,omitempty"`

	// A description block as defined below.
	// +kubebuilder:validation:Optional
	Description []ConditionDescriptionParameters `json:"description,omitempty" tf:"description,omitempty"`

	// A monitor block as defined below.
	// +kubebuilder:validation:Optional
	Monitor []ConditionMonitorParameters `json:"monitor,omitempty" tf:"monitor,omitempty"`

	// A monitor_service as block defined below.
	// +kubebuilder:validation:Optional
	MonitorService []ConditionMonitorServiceParameters `json:"monitorService,omitempty" tf:"monitor_service,omitempty"`

	// A severity block as defined below.
	// +kubebuilder:validation:Optional
	Severity []ConditionSeverityParameters `json:"severity,omitempty" tf:"severity,omitempty"`

	// A target_resource_type block as defined below.
	// +kubebuilder:validation:Optional
	TargetResourceType []ConditionTargetResourceTypeParameters `json:"targetResourceType,omitempty" tf:"target_resource_type,omitempty"`
}

type MonitorActionRuleSuppressionObservation struct {

	// The ID of the Monitor Action Rule.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorActionRuleSuppressionParameters struct {

	// A condition block as defined below.
	// +kubebuilder:validation:Optional
	Condition []MonitorActionRuleSuppressionConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// Specifies a description for the Action Rule.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Is the Action Rule enabled? Defaults to true.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// Specifies the name of the resource group in which the Monitor Action Rule should exist. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// A scope block as defined below.
	// +kubebuilder:validation:Optional
	Scope []MonitorActionRuleSuppressionScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`

	// A suppression block as defined below.
	// +kubebuilder:validation:Required
	Suppression []SuppressionParameters `json:"suppression" tf:"suppression,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MonitorActionRuleSuppressionScopeObservation struct {
}

type MonitorActionRuleSuppressionScopeParameters struct {

	// A list of resource IDs of the given scope type which will be the target of action rule.
	// +kubebuilder:validation:Required
	ResourceIds []*string `json:"resourceIds" tf:"resource_ids,omitempty"`

	// Specifies the type of target scope. Possible values are ResourceGroup and Resource.
	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type ScheduleObservation struct {
}

type ScheduleParameters struct {

	// specifies the recurrence UTC end datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Required
	EndDateUtc *string `json:"endDateUtc" tf:"end_date_utc,omitempty"`

	// specifies the list of dayOfMonth to recurrence. Possible values are between 1 - 31. Required if recurrence_type is Monthly.
	// +kubebuilder:validation:Optional
	RecurrenceMonthly []*float64 `json:"recurrenceMonthly,omitempty" tf:"recurrence_monthly,omitempty"`

	// specifies the list of dayOfWeek to recurrence. Possible values are Sunday, Monday, Tuesday, Wednesday, Thursday, Friday and  Saturday.
	// +kubebuilder:validation:Optional
	RecurrenceWeekly []*string `json:"recurrenceWeekly,omitempty" tf:"recurrence_weekly,omitempty"`

	// specifies the recurrence UTC start datetime (Y-m-d'T'H:M:S'Z').
	// +kubebuilder:validation:Required
	StartDateUtc *string `json:"startDateUtc" tf:"start_date_utc,omitempty"`
}

type SuppressionObservation struct {
}

type SuppressionParameters struct {

	// Specifies the type of suppression. Possible values are Always, Daily, Monthly, Once, and Weekly.
	// +kubebuilder:validation:Required
	RecurrenceType *string `json:"recurrenceType" tf:"recurrence_type,omitempty"`

	// A schedule block as defined below. Required if recurrence_type is Daily, Monthly, Once or Weekly.
	// +kubebuilder:validation:Optional
	Schedule []ScheduleParameters `json:"schedule,omitempty" tf:"schedule,omitempty"`
}

// MonitorActionRuleSuppressionSpec defines the desired state of MonitorActionRuleSuppression
type MonitorActionRuleSuppressionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorActionRuleSuppressionParameters `json:"forProvider"`
}

// MonitorActionRuleSuppressionStatus defines the observed state of MonitorActionRuleSuppression.
type MonitorActionRuleSuppressionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorActionRuleSuppressionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionRuleSuppression is the Schema for the MonitorActionRuleSuppressions API. Manages an Monitor Action Rule which type is suppression.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MonitorActionRuleSuppression struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MonitorActionRuleSuppressionSpec   `json:"spec"`
	Status            MonitorActionRuleSuppressionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorActionRuleSuppressionList contains a list of MonitorActionRuleSuppressions
type MonitorActionRuleSuppressionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MonitorActionRuleSuppression `json:"items"`
}

// Repository type metadata.
var (
	MonitorActionRuleSuppression_Kind             = "MonitorActionRuleSuppression"
	MonitorActionRuleSuppression_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MonitorActionRuleSuppression_Kind}.String()
	MonitorActionRuleSuppression_KindAPIVersion   = MonitorActionRuleSuppression_Kind + "." + CRDGroupVersion.String()
	MonitorActionRuleSuppression_GroupVersionKind = CRDGroupVersion.WithKind(MonitorActionRuleSuppression_Kind)
)

func init() {
	SchemeBuilder.Register(&MonitorActionRuleSuppression{}, &MonitorActionRuleSuppressionList{})
}
