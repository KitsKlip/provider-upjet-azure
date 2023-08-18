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

type MSSQLElasticPoolInitParameters struct {

	// Specifies the license type applied to this database. Possible values are LicenseIncluded and BasePrice.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the elastic pool. Valid values include SQL_Default, SQL_EastUS_DB_1, SQL_EastUS2_DB_1, SQL_SoutheastAsia_DB_1, SQL_AustraliaEast_DB_1, SQL_NorthEurope_DB_1, SQL_SouthCentralUS_DB_1, SQL_WestUS2_DB_1, SQL_UKSouth_DB_1, SQL_WestEurope_DB_1, SQL_EastUS_DB_2, SQL_EastUS2_DB_2, SQL_WestUS2_DB_2, SQL_SoutheastAsia_DB_2, SQL_AustraliaEast_DB_2, SQL_NorthEurope_DB_2, SQL_SouthCentralUS_DB_2, SQL_UKSouth_DB_2, SQL_WestEurope_DB_2, SQL_AustraliaSoutheast_DB_1, SQL_BrazilSouth_DB_1, SQL_CanadaCentral_DB_1, SQL_CanadaEast_DB_1, SQL_CentralUS_DB_1, SQL_EastAsia_DB_1, SQL_FranceCentral_DB_1, SQL_GermanyWestCentral_DB_1, SQL_CentralIndia_DB_1, SQL_SouthIndia_DB_1, SQL_JapanEast_DB_1, SQL_JapanWest_DB_1, SQL_NorthCentralUS_DB_1, SQL_UKWest_DB_1, SQL_WestUS_DB_1, SQL_AustraliaSoutheast_DB_2, SQL_BrazilSouth_DB_2, SQL_CanadaCentral_DB_2, SQL_CanadaEast_DB_2, SQL_CentralUS_DB_2, SQL_EastAsia_DB_2, SQL_FranceCentral_DB_2, SQL_GermanyWestCentral_DB_2, SQL_CentralIndia_DB_2, SQL_SouthIndia_DB_2, SQL_JapanEast_DB_2, SQL_JapanWest_DB_2, SQL_NorthCentralUS_DB_2, SQL_UKWest_DB_2, SQL_WestUS_DB_2, SQL_WestCentralUS_DB_1, SQL_FranceSouth_DB_1, SQL_WestCentralUS_DB_2, SQL_FranceSouth_DB_2, SQL_SwitzerlandNorth_DB_1, SQL_SwitzerlandNorth_DB_2, SQL_BrazilSoutheast_DB_1, SQL_UAENorth_DB_1, SQL_BrazilSoutheast_DB_2, SQL_UAENorth_DB_2. Defaults to SQL_Default.
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The max data size of the elastic pool in bytes. Conflicts with max_size_gb.
	MaxSizeBytes *float64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`

	// The max data size of the elastic pool in gigabytes. Conflicts with max_size_bytes.
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// A per_database_settings block as defined below.
	PerDatabaseSettings []PerDatabaseSettingsInitParameters `json:"perDatabaseSettings,omitempty" tf:"per_database_settings,omitempty"`

	// A sku block as defined below.
	Sku []SkuInitParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether or not this elastic pool is zone redundant. tier needs to be Premium for DTU based or BusinessCritical for vCore based sku.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MSSQLElasticPoolObservation struct {

	// The ID of the MS SQL Elastic Pool.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Specifies the license type applied to this database. Possible values are LicenseIncluded and BasePrice.
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the elastic pool. Valid values include SQL_Default, SQL_EastUS_DB_1, SQL_EastUS2_DB_1, SQL_SoutheastAsia_DB_1, SQL_AustraliaEast_DB_1, SQL_NorthEurope_DB_1, SQL_SouthCentralUS_DB_1, SQL_WestUS2_DB_1, SQL_UKSouth_DB_1, SQL_WestEurope_DB_1, SQL_EastUS_DB_2, SQL_EastUS2_DB_2, SQL_WestUS2_DB_2, SQL_SoutheastAsia_DB_2, SQL_AustraliaEast_DB_2, SQL_NorthEurope_DB_2, SQL_SouthCentralUS_DB_2, SQL_UKSouth_DB_2, SQL_WestEurope_DB_2, SQL_AustraliaSoutheast_DB_1, SQL_BrazilSouth_DB_1, SQL_CanadaCentral_DB_1, SQL_CanadaEast_DB_1, SQL_CentralUS_DB_1, SQL_EastAsia_DB_1, SQL_FranceCentral_DB_1, SQL_GermanyWestCentral_DB_1, SQL_CentralIndia_DB_1, SQL_SouthIndia_DB_1, SQL_JapanEast_DB_1, SQL_JapanWest_DB_1, SQL_NorthCentralUS_DB_1, SQL_UKWest_DB_1, SQL_WestUS_DB_1, SQL_AustraliaSoutheast_DB_2, SQL_BrazilSouth_DB_2, SQL_CanadaCentral_DB_2, SQL_CanadaEast_DB_2, SQL_CentralUS_DB_2, SQL_EastAsia_DB_2, SQL_FranceCentral_DB_2, SQL_GermanyWestCentral_DB_2, SQL_CentralIndia_DB_2, SQL_SouthIndia_DB_2, SQL_JapanEast_DB_2, SQL_JapanWest_DB_2, SQL_NorthCentralUS_DB_2, SQL_UKWest_DB_2, SQL_WestUS_DB_2, SQL_WestCentralUS_DB_1, SQL_FranceSouth_DB_1, SQL_WestCentralUS_DB_2, SQL_FranceSouth_DB_2, SQL_SwitzerlandNorth_DB_1, SQL_SwitzerlandNorth_DB_2, SQL_BrazilSoutheast_DB_1, SQL_UAENorth_DB_1, SQL_BrazilSoutheast_DB_2, SQL_UAENorth_DB_2. Defaults to SQL_Default.
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The max data size of the elastic pool in bytes. Conflicts with max_size_gb.
	MaxSizeBytes *float64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`

	// The max data size of the elastic pool in gigabytes. Conflicts with max_size_bytes.
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// A per_database_settings block as defined below.
	PerDatabaseSettings []PerDatabaseSettingsObservation `json:"perDatabaseSettings,omitempty" tf:"per_database_settings,omitempty"`

	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server. Changing this forces a new resource to be created.
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// A sku block as defined below.
	Sku []SkuObservation `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether or not this elastic pool is zone redundant. tier needs to be Premium for DTU based or BusinessCritical for vCore based sku.
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type MSSQLElasticPoolParameters struct {

	// Specifies the license type applied to this database. Possible values are LicenseIncluded and BasePrice.
	// +kubebuilder:validation:Optional
	LicenseType *string `json:"licenseType,omitempty" tf:"license_type,omitempty"`

	// Specifies the supported Azure location where the resource exists. Changing this forces a new resource to be created.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the Public Maintenance Configuration window to apply to the elastic pool. Valid values include SQL_Default, SQL_EastUS_DB_1, SQL_EastUS2_DB_1, SQL_SoutheastAsia_DB_1, SQL_AustraliaEast_DB_1, SQL_NorthEurope_DB_1, SQL_SouthCentralUS_DB_1, SQL_WestUS2_DB_1, SQL_UKSouth_DB_1, SQL_WestEurope_DB_1, SQL_EastUS_DB_2, SQL_EastUS2_DB_2, SQL_WestUS2_DB_2, SQL_SoutheastAsia_DB_2, SQL_AustraliaEast_DB_2, SQL_NorthEurope_DB_2, SQL_SouthCentralUS_DB_2, SQL_UKSouth_DB_2, SQL_WestEurope_DB_2, SQL_AustraliaSoutheast_DB_1, SQL_BrazilSouth_DB_1, SQL_CanadaCentral_DB_1, SQL_CanadaEast_DB_1, SQL_CentralUS_DB_1, SQL_EastAsia_DB_1, SQL_FranceCentral_DB_1, SQL_GermanyWestCentral_DB_1, SQL_CentralIndia_DB_1, SQL_SouthIndia_DB_1, SQL_JapanEast_DB_1, SQL_JapanWest_DB_1, SQL_NorthCentralUS_DB_1, SQL_UKWest_DB_1, SQL_WestUS_DB_1, SQL_AustraliaSoutheast_DB_2, SQL_BrazilSouth_DB_2, SQL_CanadaCentral_DB_2, SQL_CanadaEast_DB_2, SQL_CentralUS_DB_2, SQL_EastAsia_DB_2, SQL_FranceCentral_DB_2, SQL_GermanyWestCentral_DB_2, SQL_CentralIndia_DB_2, SQL_SouthIndia_DB_2, SQL_JapanEast_DB_2, SQL_JapanWest_DB_2, SQL_NorthCentralUS_DB_2, SQL_UKWest_DB_2, SQL_WestUS_DB_2, SQL_WestCentralUS_DB_1, SQL_FranceSouth_DB_1, SQL_WestCentralUS_DB_2, SQL_FranceSouth_DB_2, SQL_SwitzerlandNorth_DB_1, SQL_SwitzerlandNorth_DB_2, SQL_BrazilSoutheast_DB_1, SQL_UAENorth_DB_1, SQL_BrazilSoutheast_DB_2, SQL_UAENorth_DB_2. Defaults to SQL_Default.
	// +kubebuilder:validation:Optional
	MaintenanceConfigurationName *string `json:"maintenanceConfigurationName,omitempty" tf:"maintenance_configuration_name,omitempty"`

	// The max data size of the elastic pool in bytes. Conflicts with max_size_gb.
	// +kubebuilder:validation:Optional
	MaxSizeBytes *float64 `json:"maxSizeBytes,omitempty" tf:"max_size_bytes,omitempty"`

	// The max data size of the elastic pool in gigabytes. Conflicts with max_size_bytes.
	// +kubebuilder:validation:Optional
	MaxSizeGb *float64 `json:"maxSizeGb,omitempty" tf:"max_size_gb,omitempty"`

	// A per_database_settings block as defined below.
	// +kubebuilder:validation:Optional
	PerDatabaseSettings []PerDatabaseSettingsParameters `json:"perDatabaseSettings,omitempty" tf:"per_database_settings,omitempty"`

	// The name of the resource group in which to create the elastic pool. This must be the same as the resource group of the underlying SQL server. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
	// +kubebuilder:validation:Optional
	ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

	// Reference to a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

	// Selector for a ResourceGroup in azure to populate resourceGroupName.
	// +kubebuilder:validation:Optional
	ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

	// The name of the SQL Server on which to create the elastic pool. Changing this forces a new resource to be created.
	// +crossplane:generate:reference:type=MSSQLServer
	// +kubebuilder:validation:Optional
	ServerName *string `json:"serverName,omitempty" tf:"server_name,omitempty"`

	// Reference to a MSSQLServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameRef *v1.Reference `json:"serverNameRef,omitempty" tf:"-"`

	// Selector for a MSSQLServer to populate serverName.
	// +kubebuilder:validation:Optional
	ServerNameSelector *v1.Selector `json:"serverNameSelector,omitempty" tf:"-"`

	// A sku block as defined below.
	// +kubebuilder:validation:Optional
	Sku []SkuParameters `json:"sku,omitempty" tf:"sku,omitempty"`

	// A mapping of tags to assign to the resource.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Whether or not this elastic pool is zone redundant. tier needs to be Premium for DTU based or BusinessCritical for vCore based sku.
	// +kubebuilder:validation:Optional
	ZoneRedundant *bool `json:"zoneRedundant,omitempty" tf:"zone_redundant,omitempty"`
}

type PerDatabaseSettingsInitParameters struct {

	// The maximum capacity any one database can consume.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// The minimum capacity all databases are guaranteed.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type PerDatabaseSettingsObservation struct {

	// The maximum capacity any one database can consume.
	MaxCapacity *float64 `json:"maxCapacity,omitempty" tf:"max_capacity,omitempty"`

	// The minimum capacity all databases are guaranteed.
	MinCapacity *float64 `json:"minCapacity,omitempty" tf:"min_capacity,omitempty"`
}

type PerDatabaseSettingsParameters struct {

	// The maximum capacity any one database can consume.
	// +kubebuilder:validation:Optional
	MaxCapacity *float64 `json:"maxCapacity" tf:"max_capacity,omitempty"`

	// The minimum capacity all databases are guaranteed.
	// +kubebuilder:validation:Optional
	MinCapacity *float64 `json:"minCapacity" tf:"min_capacity,omitempty"`
}

type SkuInitParameters struct {

	// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: vCore-based or DTU-based.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The family of hardware Gen4, Gen5, Fsv2 or DC.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either vCore based tier + family pattern (e.g. GP_Gen4, BC_Gen5) or the DTU based BasicPool, StandardPool, or PremiumPool pattern. Possible values are BasicPool, StandardPool, PremiumPool, GP_Gen4, GP_Gen5, GP_Fsv2, GP_DC, BC_Gen4, BC_Gen5 and BC_DC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The tier of the particular SKU. Possible values are GeneralPurpose, BusinessCritical, Basic, Standard, or Premium. For more information see the documentation for your Elasticpool configuration: vCore-based or DTU-based.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SkuObservation struct {

	// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: vCore-based or DTU-based.
	Capacity *float64 `json:"capacity,omitempty" tf:"capacity,omitempty"`

	// The family of hardware Gen4, Gen5, Fsv2 or DC.
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either vCore based tier + family pattern (e.g. GP_Gen4, BC_Gen5) or the DTU based BasicPool, StandardPool, or PremiumPool pattern. Possible values are BasicPool, StandardPool, PremiumPool, GP_Gen4, GP_Gen5, GP_Fsv2, GP_DC, BC_Gen4, BC_Gen5 and BC_DC.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The tier of the particular SKU. Possible values are GeneralPurpose, BusinessCritical, Basic, Standard, or Premium. For more information see the documentation for your Elasticpool configuration: vCore-based or DTU-based.
	Tier *string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type SkuParameters struct {

	// The scale up/out capacity, representing server's compute units. For more information see the documentation for your Elasticpool configuration: vCore-based or DTU-based.
	// +kubebuilder:validation:Optional
	Capacity *float64 `json:"capacity" tf:"capacity,omitempty"`

	// The family of hardware Gen4, Gen5, Fsv2 or DC.
	// +kubebuilder:validation:Optional
	Family *string `json:"family,omitempty" tf:"family,omitempty"`

	// Specifies the SKU Name for this Elasticpool. The name of the SKU, will be either vCore based tier + family pattern (e.g. GP_Gen4, BC_Gen5) or the DTU based BasicPool, StandardPool, or PremiumPool pattern. Possible values are BasicPool, StandardPool, PremiumPool, GP_Gen4, GP_Gen5, GP_Fsv2, GP_DC, BC_Gen4, BC_Gen5 and BC_DC.
	// +kubebuilder:validation:Optional
	Name *string `json:"name" tf:"name,omitempty"`

	// The tier of the particular SKU. Possible values are GeneralPurpose, BusinessCritical, Basic, Standard, or Premium. For more information see the documentation for your Elasticpool configuration: vCore-based or DTU-based.
	// +kubebuilder:validation:Optional
	Tier *string `json:"tier" tf:"tier,omitempty"`
}

// MSSQLElasticPoolSpec defines the desired state of MSSQLElasticPool
type MSSQLElasticPoolSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MSSQLElasticPoolParameters `json:"forProvider"`
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
	InitProvider MSSQLElasticPoolInitParameters `json:"initProvider,omitempty"`
}

// MSSQLElasticPoolStatus defines the observed state of MSSQLElasticPool.
type MSSQLElasticPoolStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MSSQLElasticPoolObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLElasticPool is the Schema for the MSSQLElasticPools API. Manages a SQL Elastic Pool.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type MSSQLElasticPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || has(self.initProvider.location)",message="location is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.perDatabaseSettings) || has(self.initProvider.perDatabaseSettings)",message="perDatabaseSettings is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || has(self.initProvider.sku)",message="sku is a required parameter"
	Spec   MSSQLElasticPoolSpec   `json:"spec"`
	Status MSSQLElasticPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MSSQLElasticPoolList contains a list of MSSQLElasticPools
type MSSQLElasticPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MSSQLElasticPool `json:"items"`
}

// Repository type metadata.
var (
	MSSQLElasticPool_Kind             = "MSSQLElasticPool"
	MSSQLElasticPool_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: MSSQLElasticPool_Kind}.String()
	MSSQLElasticPool_KindAPIVersion   = MSSQLElasticPool_Kind + "." + CRDGroupVersion.String()
	MSSQLElasticPool_GroupVersionKind = CRDGroupVersion.WithKind(MSSQLElasticPool_Kind)
)

func init() {
	SchemeBuilder.Register(&MSSQLElasticPool{}, &MSSQLElasticPoolList{})
}
