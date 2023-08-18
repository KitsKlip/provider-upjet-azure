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

type ExportDataOptionsInitParameters struct {

	// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: WeekToDate, MonthToDate, BillingMonthToDate, TheLast7Days, TheLastMonth, TheLastBillingMonth, Custom.
	TimeFrame *string `json:"timeFrame,omitempty" tf:"time_frame,omitempty"`

	// The type of the query. Possible values are ActualCost, AmortizedCost and Usage.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ExportDataOptionsObservation struct {

	// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: WeekToDate, MonthToDate, BillingMonthToDate, TheLast7Days, TheLastMonth, TheLastBillingMonth, Custom.
	TimeFrame *string `json:"timeFrame,omitempty" tf:"time_frame,omitempty"`

	// The type of the query. Possible values are ActualCost, AmortizedCost and Usage.
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ExportDataOptionsParameters struct {

	// The time frame for pulling data for the query. If custom, then a specific time period must be provided. Possible values include: WeekToDate, MonthToDate, BillingMonthToDate, TheLast7Days, TheLastMonth, TheLastBillingMonth, Custom.
	// +kubebuilder:validation:Optional
	TimeFrame *string `json:"timeFrame" tf:"time_frame,omitempty"`

	// The type of the query. Possible values are ActualCost, AmortizedCost and Usage.
	// +kubebuilder:validation:Optional
	Type *string `json:"type" tf:"type,omitempty"`
}

type ExportDataStorageLocationInitParameters struct {

	// The path of the directory where exports will be uploaded. Changing this forces a new resource to be created.
	RootFolderPath *string `json:"rootFolderPath,omitempty" tf:"root_folder_path,omitempty"`
}

type ExportDataStorageLocationObservation struct {

	// The Resource Manager ID of the container where exports will be uploaded. Changing this forces a new resource to be created.
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// The path of the directory where exports will be uploaded. Changing this forces a new resource to be created.
	RootFolderPath *string `json:"rootFolderPath,omitempty" tf:"root_folder_path,omitempty"`
}

type ExportDataStorageLocationParameters struct {

	// The Resource Manager ID of the container where exports will be uploaded. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/storage/v1beta1.Container
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("resource_manager_id",true)
	// +kubebuilder:validation:Optional
	ContainerID *string `json:"containerId,omitempty" tf:"container_id,omitempty"`

	// Reference to a Container in storage to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDRef *v1.Reference `json:"containerIdRef,omitempty" tf:"-"`

	// Selector for a Container in storage to populate containerId.
	// +kubebuilder:validation:Optional
	ContainerIDSelector *v1.Selector `json:"containerIdSelector,omitempty" tf:"-"`

	// The path of the directory where exports will be uploaded. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	RootFolderPath *string `json:"rootFolderPath" tf:"root_folder_path,omitempty"`
}

type ResourceGroupCostManagementExportInitParameters struct {

	// Is the cost management export active? Default is true.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// A export_data_options block as defined below.
	ExportDataOptions []ExportDataOptionsInitParameters `json:"exportDataOptions,omitempty" tf:"export_data_options,omitempty"`

	// A export_data_storage_location block as defined below.
	ExportDataStorageLocation []ExportDataStorageLocationInitParameters `json:"exportDataStorageLocation,omitempty" tf:"export_data_storage_location,omitempty"`

	// The date the export will stop capturing information.
	RecurrencePeriodEndDate *string `json:"recurrencePeriodEndDate,omitempty" tf:"recurrence_period_end_date,omitempty"`

	// The date the export will start capturing information.
	RecurrencePeriodStartDate *string `json:"recurrencePeriodStartDate,omitempty" tf:"recurrence_period_start_date,omitempty"`

	// How often the requested information will be exported. Valid values include Annually, Daily, Monthly, Weekly.
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`
}

type ResourceGroupCostManagementExportObservation struct {

	// Is the cost management export active? Default is true.
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// A export_data_options block as defined below.
	ExportDataOptions []ExportDataOptionsObservation `json:"exportDataOptions,omitempty" tf:"export_data_options,omitempty"`

	// A export_data_storage_location block as defined below.
	ExportDataStorageLocation []ExportDataStorageLocationObservation `json:"exportDataStorageLocation,omitempty" tf:"export_data_storage_location,omitempty"`

	// The ID of the Cost Management Export for this Resource Group.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The date the export will stop capturing information.
	RecurrencePeriodEndDate *string `json:"recurrencePeriodEndDate,omitempty" tf:"recurrence_period_end_date,omitempty"`

	// The date the export will start capturing information.
	RecurrencePeriodStartDate *string `json:"recurrencePeriodStartDate,omitempty" tf:"recurrence_period_start_date,omitempty"`

	// How often the requested information will be exported. Valid values include Annually, Daily, Monthly, Weekly.
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`

	// The id of the resource group on which to create an export. Changing this forces a new resource to be created.
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`
}

type ResourceGroupCostManagementExportParameters struct {

	// Is the cost management export active? Default is true.
	// +kubebuilder:validation:Optional
	Active *bool `json:"active,omitempty" tf:"active,omitempty"`

	// A export_data_options block as defined below.
	// +kubebuilder:validation:Optional
	ExportDataOptions []ExportDataOptionsParameters `json:"exportDataOptions,omitempty" tf:"export_data_options,omitempty"`

	// A export_data_storage_location block as defined below.
	// +kubebuilder:validation:Optional
	ExportDataStorageLocation []ExportDataStorageLocationParameters `json:"exportDataStorageLocation,omitempty" tf:"export_data_storage_location,omitempty"`

	// The date the export will stop capturing information.
	// +kubebuilder:validation:Optional
	RecurrencePeriodEndDate *string `json:"recurrencePeriodEndDate,omitempty" tf:"recurrence_period_end_date,omitempty"`

	// The date the export will start capturing information.
	// +kubebuilder:validation:Optional
	RecurrencePeriodStartDate *string `json:"recurrencePeriodStartDate,omitempty" tf:"recurrence_period_start_date,omitempty"`

	// How often the requested information will be exported. Valid values include Annually, Daily, Monthly, Weekly.
	// +kubebuilder:validation:Optional
	RecurrenceType *string `json:"recurrenceType,omitempty" tf:"recurrence_type,omitempty"`

	// The id of the resource group on which to create an export. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	ResourceGroupID *string `json:"resourceGroupId,omitempty" tf:"resource_group_id,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDRef *v1.Reference `json:"resourceGroupIdRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupId.
	// +kubebuilder:validation:Optional
	ResourceGroupIDSelector *v1.Selector `json:"resourceGroupIdSelector,omitempty" tf:"-"`
}

// ResourceGroupCostManagementExportSpec defines the desired state of ResourceGroupCostManagementExport
type ResourceGroupCostManagementExportSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ResourceGroupCostManagementExportParameters `json:"forProvider"`
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
	InitProvider ResourceGroupCostManagementExportInitParameters `json:"initProvider,omitempty"`
}

// ResourceGroupCostManagementExportStatus defines the observed state of ResourceGroupCostManagementExport.
type ResourceGroupCostManagementExportStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ResourceGroupCostManagementExportObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupCostManagementExport is the Schema for the ResourceGroupCostManagementExports API. Manages an Azure Cost Management Export for a Resource Group.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type ResourceGroupCostManagementExport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.exportDataOptions) || has(self.initProvider.exportDataOptions)",message="exportDataOptions is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.exportDataStorageLocation) || has(self.initProvider.exportDataStorageLocation)",message="exportDataStorageLocation is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recurrencePeriodEndDate) || has(self.initProvider.recurrencePeriodEndDate)",message="recurrencePeriodEndDate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recurrencePeriodStartDate) || has(self.initProvider.recurrencePeriodStartDate)",message="recurrencePeriodStartDate is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.recurrenceType) || has(self.initProvider.recurrenceType)",message="recurrenceType is a required parameter"
	Spec   ResourceGroupCostManagementExportSpec   `json:"spec"`
	Status ResourceGroupCostManagementExportStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupCostManagementExportList contains a list of ResourceGroupCostManagementExports
type ResourceGroupCostManagementExportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroupCostManagementExport `json:"items"`
}

// Repository type metadata.
var (
	ResourceGroupCostManagementExport_Kind             = "ResourceGroupCostManagementExport"
	ResourceGroupCostManagementExport_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ResourceGroupCostManagementExport_Kind}.String()
	ResourceGroupCostManagementExport_KindAPIVersion   = ResourceGroupCostManagementExport_Kind + "." + CRDGroupVersion.String()
	ResourceGroupCostManagementExport_GroupVersionKind = CRDGroupVersion.WithKind(ResourceGroupCostManagementExport_Kind)
)

func init() {
	SchemeBuilder.Register(&ResourceGroupCostManagementExport{}, &ResourceGroupCostManagementExportList{})
}
