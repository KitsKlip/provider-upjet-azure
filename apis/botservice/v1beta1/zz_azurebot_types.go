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




type AzureBotInitParameters struct {


// The CMK Key Vault Key URL that will be used to encrypt the Bot with the Customer Managed Encryption Key.
CmkKeyVaultKeyURL *string `json:"cmkKeyVaultKeyUrl,omitempty" tf:"cmk_key_vault_key_url,omitempty"`

// The Application Insights API Key to associate with this Azure Bot Service.
DeveloperAppInsightsAPIKeySecretRef *v1.SecretKeySelector `json:"developerAppInsightsApiKeySecretRef,omitempty" tf:"-"`

// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("app_id",true)
DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

// Reference to a ApplicationInsights in insights to populate developerAppInsightsApplicationId.
// +kubebuilder:validation:Optional
DeveloperAppInsightsApplicationIDRef *v1.Reference `json:"developerAppInsightsApplicationIdRef,omitempty" tf:"-"`

// Selector for a ApplicationInsights in insights to populate developerAppInsightsApplicationId.
// +kubebuilder:validation:Optional
DeveloperAppInsightsApplicationIDSelector *v1.Selector `json:"developerAppInsightsApplicationIdSelector,omitempty" tf:"-"`

// The Application Insight Key to associate with this Azure Bot Service.
DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

// The name that the Azure Bot Service will be displayed as. This defaults to the value set for name if not specified.
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// The Azure Bot Service endpoint.
Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

// The Icon Url of the Azure Bot Service. Defaults to https://docs.botframework.com/static/devportal/client/images/bot-framework-default.png.
IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

// Is local authentication enabled? Defaults to true.
LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
Location *string `json:"location,omitempty" tf:"location,omitempty"`

// A list of LUIS App IDs to associate with this Azure Bot Service.
LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

// The LUIS key to associate with this Azure Bot Service.
LuisKeySecretRef *v1.SecretKeySelector `json:"luisKeySecretRef,omitempty" tf:"-"`

// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

// The ID of the Microsoft App Managed Identity for this Azure Bot Service. Changing this forces a new resource to be created.
MicrosoftAppMsiID *string `json:"microsoftAppMsiId,omitempty" tf:"microsoft_app_msi_id,omitempty"`

// The Tenant ID of the Microsoft App for this Azure Bot Service. Changing this forces a new resource to be created.
MicrosoftAppTenantID *string `json:"microsoftAppTenantId,omitempty" tf:"microsoft_app_tenant_id,omitempty"`

// The Microsoft App Type for this Azure Bot Service. Possible values are MultiTenant, SingleTenant and UserAssignedMSI. Changing this forces a new resource to be created.
MicrosoftAppType *string `json:"microsoftAppType,omitempty" tf:"microsoft_app_type,omitempty"`

// Whether public network access is enabled. Defaults to true.
PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

// The SKU of the Azure Bot Service. Accepted values are F0 or S1. Changing this forces a new resource to be created.
Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

// Is the streaming endpoint enabled for this Azure Bot Service. Defaults to false.
StreamingEndpointEnabled *bool `json:"streamingEndpointEnabled,omitempty" tf:"streaming_endpoint_enabled,omitempty"`

// A mapping of tags which should be assigned to this Azure Bot Service.
// +mapType=granular
Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}


type AzureBotObservation struct {


// The CMK Key Vault Key URL that will be used to encrypt the Bot with the Customer Managed Encryption Key.
CmkKeyVaultKeyURL *string `json:"cmkKeyVaultKeyUrl,omitempty" tf:"cmk_key_vault_key_url,omitempty"`

// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

// The Application Insight Key to associate with this Azure Bot Service.
DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

// The name that the Azure Bot Service will be displayed as. This defaults to the value set for name if not specified.
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// The Azure Bot Service endpoint.
Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

// The ID of the Azure Bot Service.
ID *string `json:"id,omitempty" tf:"id,omitempty"`

// The Icon Url of the Azure Bot Service. Defaults to https://docs.botframework.com/static/devportal/client/images/bot-framework-default.png.
IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

// Is local authentication enabled? Defaults to true.
LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
Location *string `json:"location,omitempty" tf:"location,omitempty"`

// A list of LUIS App IDs to associate with this Azure Bot Service.
LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

// The ID of the Microsoft App Managed Identity for this Azure Bot Service. Changing this forces a new resource to be created.
MicrosoftAppMsiID *string `json:"microsoftAppMsiId,omitempty" tf:"microsoft_app_msi_id,omitempty"`

// The Tenant ID of the Microsoft App for this Azure Bot Service. Changing this forces a new resource to be created.
MicrosoftAppTenantID *string `json:"microsoftAppTenantId,omitempty" tf:"microsoft_app_tenant_id,omitempty"`

// The Microsoft App Type for this Azure Bot Service. Possible values are MultiTenant, SingleTenant and UserAssignedMSI. Changing this forces a new resource to be created.
MicrosoftAppType *string `json:"microsoftAppType,omitempty" tf:"microsoft_app_type,omitempty"`

// Whether public network access is enabled. Defaults to true.
PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

// The SKU of the Azure Bot Service. Accepted values are F0 or S1. Changing this forces a new resource to be created.
Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

// Is the streaming endpoint enabled for this Azure Bot Service. Defaults to false.
StreamingEndpointEnabled *bool `json:"streamingEndpointEnabled,omitempty" tf:"streaming_endpoint_enabled,omitempty"`

// A mapping of tags which should be assigned to this Azure Bot Service.
// +mapType=granular
Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}


type AzureBotParameters struct {


// The CMK Key Vault Key URL that will be used to encrypt the Bot with the Customer Managed Encryption Key.
// +kubebuilder:validation:Optional
CmkKeyVaultKeyURL *string `json:"cmkKeyVaultKeyUrl,omitempty" tf:"cmk_key_vault_key_url,omitempty"`

// The Application Insights API Key to associate with this Azure Bot Service.
// +kubebuilder:validation:Optional
DeveloperAppInsightsAPIKeySecretRef *v1.SecretKeySelector `json:"developerAppInsightsApiKeySecretRef,omitempty" tf:"-"`

// The resource ID of the Application Insights instance to associate with this Azure Bot Service.
// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/insights/v1beta1.ApplicationInsights
// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("app_id",true)
// +kubebuilder:validation:Optional
DeveloperAppInsightsApplicationID *string `json:"developerAppInsightsApplicationId,omitempty" tf:"developer_app_insights_application_id,omitempty"`

// Reference to a ApplicationInsights in insights to populate developerAppInsightsApplicationId.
// +kubebuilder:validation:Optional
DeveloperAppInsightsApplicationIDRef *v1.Reference `json:"developerAppInsightsApplicationIdRef,omitempty" tf:"-"`

// Selector for a ApplicationInsights in insights to populate developerAppInsightsApplicationId.
// +kubebuilder:validation:Optional
DeveloperAppInsightsApplicationIDSelector *v1.Selector `json:"developerAppInsightsApplicationIdSelector,omitempty" tf:"-"`

// The Application Insight Key to associate with this Azure Bot Service.
// +kubebuilder:validation:Optional
DeveloperAppInsightsKey *string `json:"developerAppInsightsKey,omitempty" tf:"developer_app_insights_key,omitempty"`

// The name that the Azure Bot Service will be displayed as. This defaults to the value set for name if not specified.
// +kubebuilder:validation:Optional
DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

// The Azure Bot Service endpoint.
// +kubebuilder:validation:Optional
Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

// The Icon Url of the Azure Bot Service. Defaults to https://docs.botframework.com/static/devportal/client/images/bot-framework-default.png.
// +kubebuilder:validation:Optional
IconURL *string `json:"iconUrl,omitempty" tf:"icon_url,omitempty"`

// Is local authentication enabled? Defaults to true.
// +kubebuilder:validation:Optional
LocalAuthenticationEnabled *bool `json:"localAuthenticationEnabled,omitempty" tf:"local_authentication_enabled,omitempty"`

// The supported Azure location where the Azure Bot Service should exist. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
Location *string `json:"location,omitempty" tf:"location,omitempty"`

// A list of LUIS App IDs to associate with this Azure Bot Service.
// +kubebuilder:validation:Optional
LuisAppIds []*string `json:"luisAppIds,omitempty" tf:"luis_app_ids,omitempty"`

// The LUIS key to associate with this Azure Bot Service.
// +kubebuilder:validation:Optional
LuisKeySecretRef *v1.SecretKeySelector `json:"luisKeySecretRef,omitempty" tf:"-"`

// The Microsoft Application ID for the Azure Bot Service. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
MicrosoftAppID *string `json:"microsoftAppId,omitempty" tf:"microsoft_app_id,omitempty"`

// The ID of the Microsoft App Managed Identity for this Azure Bot Service. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
MicrosoftAppMsiID *string `json:"microsoftAppMsiId,omitempty" tf:"microsoft_app_msi_id,omitempty"`

// The Tenant ID of the Microsoft App for this Azure Bot Service. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
MicrosoftAppTenantID *string `json:"microsoftAppTenantId,omitempty" tf:"microsoft_app_tenant_id,omitempty"`

// The Microsoft App Type for this Azure Bot Service. Possible values are MultiTenant, SingleTenant and UserAssignedMSI. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
MicrosoftAppType *string `json:"microsoftAppType,omitempty" tf:"microsoft_app_type,omitempty"`

// Whether public network access is enabled. Defaults to true.
// +kubebuilder:validation:Optional
PublicNetworkAccessEnabled *bool `json:"publicNetworkAccessEnabled,omitempty" tf:"public_network_access_enabled,omitempty"`

// The name of the Resource Group where the Azure Bot Service should exist. Changing this forces a new resource to be created.
// +crossplane:generate:reference:type=github.com/upbound/provider-azure/apis/azure/v1beta1.ResourceGroup
// +kubebuilder:validation:Optional
ResourceGroupName *string `json:"resourceGroupName,omitempty" tf:"resource_group_name,omitempty"`

// Reference to a ResourceGroup in azure to populate resourceGroupName.
// +kubebuilder:validation:Optional
ResourceGroupNameRef *v1.Reference `json:"resourceGroupNameRef,omitempty" tf:"-"`

// Selector for a ResourceGroup in azure to populate resourceGroupName.
// +kubebuilder:validation:Optional
ResourceGroupNameSelector *v1.Selector `json:"resourceGroupNameSelector,omitempty" tf:"-"`

// The SKU of the Azure Bot Service. Accepted values are F0 or S1. Changing this forces a new resource to be created.
// +kubebuilder:validation:Optional
Sku *string `json:"sku,omitempty" tf:"sku,omitempty"`

// Is the streaming endpoint enabled for this Azure Bot Service. Defaults to false.
// +kubebuilder:validation:Optional
StreamingEndpointEnabled *bool `json:"streamingEndpointEnabled,omitempty" tf:"streaming_endpoint_enabled,omitempty"`

// A mapping of tags which should be assigned to this Azure Bot Service.
// +kubebuilder:validation:Optional
// +mapType=granular
Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AzureBotSpec defines the desired state of AzureBot
type AzureBotSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider       AzureBotParameters `json:"forProvider"`
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
	InitProvider       AzureBotInitParameters `json:"initProvider,omitempty"`
}

// AzureBotStatus defines the observed state of AzureBot.
type AzureBotStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider          AzureBotObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion


// AzureBot is the Schema for the AzureBots API. Manages an Azure Bot Service.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azure}
type AzureBot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.location) || (has(self.initProvider) && has(self.initProvider.location))",message="spec.forProvider.location is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.microsoftAppId) || (has(self.initProvider) && has(self.initProvider.microsoftAppId))",message="spec.forProvider.microsoftAppId is a required parameter"
// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.sku) || (has(self.initProvider) && has(self.initProvider.sku))",message="spec.forProvider.sku is a required parameter"
	Spec              AzureBotSpec   `json:"spec"`
	Status            AzureBotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureBotList contains a list of AzureBots
type AzureBotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureBot `json:"items"`
}

// Repository type metadata.
var (
	AzureBot_Kind             = "AzureBot"
	AzureBot_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AzureBot_Kind}.String()
	AzureBot_KindAPIVersion   = AzureBot_Kind + "." + CRDGroupVersion.String()
	AzureBot_GroupVersionKind = CRDGroupVersion.WithKind(AzureBot_Kind)
)

func init() {
	SchemeBuilder.Register(&AzureBot{}, &AzureBotList{})
}
