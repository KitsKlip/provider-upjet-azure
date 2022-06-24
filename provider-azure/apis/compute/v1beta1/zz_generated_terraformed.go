/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	"github.com/pkg/errors"

	"github.com/upbound/upjet/pkg/resource"
	"github.com/upbound/upjet/pkg/resource/json"
)

// GetTerraformResourceType returns Terraform resource type for this AvailabilitySet
func (mg *AvailabilitySet) GetTerraformResourceType() string {
	return "azurerm_availability_set"
}

// GetConnectionDetailsMapping for this AvailabilitySet
func (tr *AvailabilitySet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this AvailabilitySet
func (tr *AvailabilitySet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this AvailabilitySet
func (tr *AvailabilitySet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this AvailabilitySet
func (tr *AvailabilitySet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this AvailabilitySet
func (tr *AvailabilitySet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this AvailabilitySet
func (tr *AvailabilitySet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this AvailabilitySet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *AvailabilitySet) LateInitialize(attrs []byte) (bool, error) {
	params := &AvailabilitySetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *AvailabilitySet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this DedicatedHost
func (mg *DedicatedHost) GetTerraformResourceType() string {
	return "azurerm_dedicated_host"
}

// GetConnectionDetailsMapping for this DedicatedHost
func (tr *DedicatedHost) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this DedicatedHost
func (tr *DedicatedHost) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DedicatedHost
func (tr *DedicatedHost) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DedicatedHost
func (tr *DedicatedHost) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DedicatedHost
func (tr *DedicatedHost) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DedicatedHost
func (tr *DedicatedHost) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DedicatedHost using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DedicatedHost) LateInitialize(attrs []byte) (bool, error) {
	params := &DedicatedHostParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DedicatedHost) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this DiskAccess
func (mg *DiskAccess) GetTerraformResourceType() string {
	return "azurerm_disk_access"
}

// GetConnectionDetailsMapping for this DiskAccess
func (tr *DiskAccess) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this DiskAccess
func (tr *DiskAccess) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DiskAccess
func (tr *DiskAccess) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DiskAccess
func (tr *DiskAccess) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DiskAccess
func (tr *DiskAccess) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DiskAccess
func (tr *DiskAccess) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DiskAccess using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DiskAccess) LateInitialize(attrs []byte) (bool, error) {
	params := &DiskAccessParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DiskAccess) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this DiskEncryptionSet
func (mg *DiskEncryptionSet) GetTerraformResourceType() string {
	return "azurerm_disk_encryption_set"
}

// GetConnectionDetailsMapping for this DiskEncryptionSet
func (tr *DiskEncryptionSet) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this DiskEncryptionSet
func (tr *DiskEncryptionSet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this DiskEncryptionSet
func (tr *DiskEncryptionSet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this DiskEncryptionSet
func (tr *DiskEncryptionSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this DiskEncryptionSet
func (tr *DiskEncryptionSet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this DiskEncryptionSet
func (tr *DiskEncryptionSet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this DiskEncryptionSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *DiskEncryptionSet) LateInitialize(attrs []byte) (bool, error) {
	params := &DiskEncryptionSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *DiskEncryptionSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Image
func (mg *Image) GetTerraformResourceType() string {
	return "azurerm_image"
}

// GetConnectionDetailsMapping for this Image
func (tr *Image) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Image
func (tr *Image) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Image
func (tr *Image) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Image
func (tr *Image) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Image
func (tr *Image) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Image
func (tr *Image) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Image using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Image) LateInitialize(attrs []byte) (bool, error) {
	params := &ImageParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Image) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LinuxVirtualMachine
func (mg *LinuxVirtualMachine) GetTerraformResourceType() string {
	return "azurerm_linux_virtual_machine"
}

// GetConnectionDetailsMapping for this LinuxVirtualMachine
func (tr *LinuxVirtualMachine) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"admin_password": "spec.forProvider.adminPasswordSecretRef", "custom_data": "spec.forProvider.customDataSecretRef"}
}

// GetObservation of this LinuxVirtualMachine
func (tr *LinuxVirtualMachine) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LinuxVirtualMachine
func (tr *LinuxVirtualMachine) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LinuxVirtualMachine
func (tr *LinuxVirtualMachine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LinuxVirtualMachine
func (tr *LinuxVirtualMachine) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LinuxVirtualMachine
func (tr *LinuxVirtualMachine) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LinuxVirtualMachine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LinuxVirtualMachine) LateInitialize(attrs []byte) (bool, error) {
	params := &LinuxVirtualMachineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LinuxVirtualMachine) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this LinuxVirtualMachineScaleSet
func (mg *LinuxVirtualMachineScaleSet) GetTerraformResourceType() string {
	return "azurerm_linux_virtual_machine_scale_set"
}

// GetConnectionDetailsMapping for this LinuxVirtualMachineScaleSet
func (tr *LinuxVirtualMachineScaleSet) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"admin_password": "spec.forProvider.adminPasswordSecretRef", "custom_data": "spec.forProvider.customDataSecretRef", "extension[*].protected_settings": "spec.forProvider.extension[*].protectedSettingsSecretRef"}
}

// GetObservation of this LinuxVirtualMachineScaleSet
func (tr *LinuxVirtualMachineScaleSet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this LinuxVirtualMachineScaleSet
func (tr *LinuxVirtualMachineScaleSet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this LinuxVirtualMachineScaleSet
func (tr *LinuxVirtualMachineScaleSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this LinuxVirtualMachineScaleSet
func (tr *LinuxVirtualMachineScaleSet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this LinuxVirtualMachineScaleSet
func (tr *LinuxVirtualMachineScaleSet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this LinuxVirtualMachineScaleSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *LinuxVirtualMachineScaleSet) LateInitialize(attrs []byte) (bool, error) {
	params := &LinuxVirtualMachineScaleSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *LinuxVirtualMachineScaleSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ManagedDisk
func (mg *ManagedDisk) GetTerraformResourceType() string {
	return "azurerm_managed_disk"
}

// GetConnectionDetailsMapping for this ManagedDisk
func (tr *ManagedDisk) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ManagedDisk
func (tr *ManagedDisk) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ManagedDisk
func (tr *ManagedDisk) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ManagedDisk
func (tr *ManagedDisk) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ManagedDisk
func (tr *ManagedDisk) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ManagedDisk
func (tr *ManagedDisk) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ManagedDisk using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ManagedDisk) LateInitialize(attrs []byte) (bool, error) {
	params := &ManagedDiskParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ManagedDisk) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this OrchestratedVirtualMachineScaleSet
func (mg *OrchestratedVirtualMachineScaleSet) GetTerraformResourceType() string {
	return "azurerm_orchestrated_virtual_machine_scale_set"
}

// GetConnectionDetailsMapping for this OrchestratedVirtualMachineScaleSet
func (tr *OrchestratedVirtualMachineScaleSet) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"extension[*].protected_settings": "spec.forProvider.extension[*].protectedSettingsSecretRef", "os_profile[*].custom_data": "spec.forProvider.osProfile[*].customDataSecretRef", "os_profile[*].linux_configuration[*].admin_password": "spec.forProvider.osProfile[*].linuxConfiguration[*].adminPasswordSecretRef", "os_profile[*].windows_configuration[*].admin_password": "spec.forProvider.osProfile[*].windowsConfiguration[*].adminPasswordSecretRef"}
}

// GetObservation of this OrchestratedVirtualMachineScaleSet
func (tr *OrchestratedVirtualMachineScaleSet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this OrchestratedVirtualMachineScaleSet
func (tr *OrchestratedVirtualMachineScaleSet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this OrchestratedVirtualMachineScaleSet
func (tr *OrchestratedVirtualMachineScaleSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this OrchestratedVirtualMachineScaleSet
func (tr *OrchestratedVirtualMachineScaleSet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this OrchestratedVirtualMachineScaleSet
func (tr *OrchestratedVirtualMachineScaleSet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this OrchestratedVirtualMachineScaleSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *OrchestratedVirtualMachineScaleSet) LateInitialize(attrs []byte) (bool, error) {
	params := &OrchestratedVirtualMachineScaleSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *OrchestratedVirtualMachineScaleSet) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this ProximityPlacementGroup
func (mg *ProximityPlacementGroup) GetTerraformResourceType() string {
	return "azurerm_proximity_placement_group"
}

// GetConnectionDetailsMapping for this ProximityPlacementGroup
func (tr *ProximityPlacementGroup) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this ProximityPlacementGroup
func (tr *ProximityPlacementGroup) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this ProximityPlacementGroup
func (tr *ProximityPlacementGroup) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this ProximityPlacementGroup
func (tr *ProximityPlacementGroup) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this ProximityPlacementGroup
func (tr *ProximityPlacementGroup) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this ProximityPlacementGroup
func (tr *ProximityPlacementGroup) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this ProximityPlacementGroup using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *ProximityPlacementGroup) LateInitialize(attrs []byte) (bool, error) {
	params := &ProximityPlacementGroupParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *ProximityPlacementGroup) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this SharedImageGallery
func (mg *SharedImageGallery) GetTerraformResourceType() string {
	return "azurerm_shared_image_gallery"
}

// GetConnectionDetailsMapping for this SharedImageGallery
func (tr *SharedImageGallery) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this SharedImageGallery
func (tr *SharedImageGallery) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this SharedImageGallery
func (tr *SharedImageGallery) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this SharedImageGallery
func (tr *SharedImageGallery) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this SharedImageGallery
func (tr *SharedImageGallery) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this SharedImageGallery
func (tr *SharedImageGallery) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this SharedImageGallery using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *SharedImageGallery) LateInitialize(attrs []byte) (bool, error) {
	params := &SharedImageGalleryParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *SharedImageGallery) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this Snapshot
func (mg *Snapshot) GetTerraformResourceType() string {
	return "azurerm_snapshot"
}

// GetConnectionDetailsMapping for this Snapshot
func (tr *Snapshot) GetConnectionDetailsMapping() map[string]string {
	return nil
}

// GetObservation of this Snapshot
func (tr *Snapshot) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this Snapshot
func (tr *Snapshot) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this Snapshot
func (tr *Snapshot) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this Snapshot
func (tr *Snapshot) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this Snapshot
func (tr *Snapshot) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this Snapshot using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *Snapshot) LateInitialize(attrs []byte) (bool, error) {
	params := &SnapshotParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *Snapshot) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this WindowsVirtualMachine
func (mg *WindowsVirtualMachine) GetTerraformResourceType() string {
	return "azurerm_windows_virtual_machine"
}

// GetConnectionDetailsMapping for this WindowsVirtualMachine
func (tr *WindowsVirtualMachine) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"additional_unattend_content[*].content": "spec.forProvider.additionalUnattendContent[*].contentSecretRef", "admin_password": "spec.forProvider.adminPasswordSecretRef", "custom_data": "spec.forProvider.customDataSecretRef"}
}

// GetObservation of this WindowsVirtualMachine
func (tr *WindowsVirtualMachine) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this WindowsVirtualMachine
func (tr *WindowsVirtualMachine) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this WindowsVirtualMachine
func (tr *WindowsVirtualMachine) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this WindowsVirtualMachine
func (tr *WindowsVirtualMachine) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this WindowsVirtualMachine
func (tr *WindowsVirtualMachine) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this WindowsVirtualMachine using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *WindowsVirtualMachine) LateInitialize(attrs []byte) (bool, error) {
	params := &WindowsVirtualMachineParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *WindowsVirtualMachine) GetTerraformSchemaVersion() int {
	return 0
}

// GetTerraformResourceType returns Terraform resource type for this WindowsVirtualMachineScaleSet
func (mg *WindowsVirtualMachineScaleSet) GetTerraformResourceType() string {
	return "azurerm_windows_virtual_machine_scale_set"
}

// GetConnectionDetailsMapping for this WindowsVirtualMachineScaleSet
func (tr *WindowsVirtualMachineScaleSet) GetConnectionDetailsMapping() map[string]string {
	return map[string]string{"additional_unattend_content[*].content": "spec.forProvider.additionalUnattendContent[*].contentSecretRef", "admin_password": "spec.forProvider.adminPasswordSecretRef", "custom_data": "spec.forProvider.customDataSecretRef", "extension[*].protected_settings": "spec.forProvider.extension[*].protectedSettingsSecretRef"}
}

// GetObservation of this WindowsVirtualMachineScaleSet
func (tr *WindowsVirtualMachineScaleSet) GetObservation() (map[string]interface{}, error) {
	o, err := json.TFParser.Marshal(tr.Status.AtProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(o, &base)
}

// SetObservation for this WindowsVirtualMachineScaleSet
func (tr *WindowsVirtualMachineScaleSet) SetObservation(obs map[string]interface{}) error {
	p, err := json.TFParser.Marshal(obs)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Status.AtProvider)
}

// GetID returns ID of underlying Terraform resource of this WindowsVirtualMachineScaleSet
func (tr *WindowsVirtualMachineScaleSet) GetID() string {
	if tr.Status.AtProvider.ID == nil {
		return ""
	}
	return *tr.Status.AtProvider.ID
}

// GetParameters of this WindowsVirtualMachineScaleSet
func (tr *WindowsVirtualMachineScaleSet) GetParameters() (map[string]interface{}, error) {
	p, err := json.TFParser.Marshal(tr.Spec.ForProvider)
	if err != nil {
		return nil, err
	}
	base := map[string]interface{}{}
	return base, json.TFParser.Unmarshal(p, &base)
}

// SetParameters for this WindowsVirtualMachineScaleSet
func (tr *WindowsVirtualMachineScaleSet) SetParameters(params map[string]interface{}) error {
	p, err := json.TFParser.Marshal(params)
	if err != nil {
		return err
	}
	return json.TFParser.Unmarshal(p, &tr.Spec.ForProvider)
}

// LateInitialize this WindowsVirtualMachineScaleSet using its observed tfState.
// returns True if there are any spec changes for the resource.
func (tr *WindowsVirtualMachineScaleSet) LateInitialize(attrs []byte) (bool, error) {
	params := &WindowsVirtualMachineScaleSetParameters{}
	if err := json.TFParser.Unmarshal(attrs, params); err != nil {
		return false, errors.Wrap(err, "failed to unmarshal Terraform state parameters for late-initialization")
	}
	opts := []resource.GenericLateInitializerOption{resource.WithZeroValueJSONOmitEmptyFilter(resource.CNameWildcard)}

	li := resource.NewGenericLateInitializer(opts...)
	return li.LateInitialize(&tr.Spec.ForProvider, params)
}

// GetTerraformSchemaVersion returns the associated Terraform schema version
func (tr *WindowsVirtualMachineScaleSet) GetTerraformSchemaVersion() int {
	return 0
}
