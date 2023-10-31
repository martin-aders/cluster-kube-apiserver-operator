// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

// EtcdBackupSpecApplyConfiguration represents an declarative configuration of the EtcdBackupSpec type for use
// with apply.
type EtcdBackupSpecApplyConfiguration struct {
	PVCName *string `json:"pvcName,omitempty"`
}

// EtcdBackupSpecApplyConfiguration constructs an declarative configuration of the EtcdBackupSpec type for use with
// apply.
func EtcdBackupSpec() *EtcdBackupSpecApplyConfiguration {
	return &EtcdBackupSpecApplyConfiguration{}
}

// WithPVCName sets the PVCName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PVCName field is set to the value of the last call.
func (b *EtcdBackupSpecApplyConfiguration) WithPVCName(value string) *EtcdBackupSpecApplyConfiguration {
	b.PVCName = &value
	return b
}