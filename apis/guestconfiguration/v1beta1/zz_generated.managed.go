/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PolicyVirtualMachineConfigurationAssignment.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PolicyVirtualMachineConfigurationAssignment) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PolicyVirtualMachineConfigurationAssignment.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PolicyVirtualMachineConfigurationAssignment) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PolicyVirtualMachineConfigurationAssignment.
func (mg *PolicyVirtualMachineConfigurationAssignment) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
