// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConfigAppInstanceConfig The complete configuration for an Application Instance
// When changing key fields such as the drives/volumeRefs or the number
// of interfaces, the controller is required to issue a purge command i.e.,
// increase the purge counter. Otherwise there will be an error (The controller
// can also issue a purge command to re-construct the content of the first
// drive/volumeRef without any changes.)
// Some changes such as ACL changes in the interfaces do not require a restart,
// but all other changes (such as fixedresources and adapters) require a
// restart command i.e., an increase to the restart counter. The restart counter
// can also be increased to cause an application instance restart without
// any other change to the application instance.
//
// swagger:model configAppInstanceConfig
type ConfigAppInstanceConfig struct {

	// Set activate to start the application instance; clear it to stop it.
	Activate bool `json:"activate,omitempty"`

	// Physical adapters such as eth1 or USB controllers and GPUs assigned
	// to the application instance.
	// The Name in Adapter should be set to PhysicalIO.assigngrp
	Adapters []*ConfigAdapter `json:"adapters"`

	// contains the encrypted userdata
	CipherData *ConfigCipherBlock `json:"cipherData,omitempty"`

	// The static IP address assigned on the NetworkAdapter which App Container
	// stats collection uses. If the 'collectStatsIPAddr' is not empty and valid,
	// it enables the container stats collection for this App.
	// During App instance creation, after user enables the collection of stats
	// from App, cloud needs to make sure at least one 'Local' type of Network-Instance
	// is assigned to the App interface, and based on the subnet of the NI, statically
	// assign an IP address on the same subnet, e.g. 10.1.0.100
	CollectStatsIPAddr string `json:"collectStatsIPAddr,omitempty"`

	// displayname
	Displayname string `json:"displayname,omitempty"`

	// VolumeRefs, if supported by EVE, will supersede drives. Drives still
	// exist for backward compatibility.
	// Drives will be deprecated in the future.
	// The order here is critical because they are presented to the VM or
	// container in the order they are listed, e.g., the first VM image
	// will be the root disk.
	Drives []*ConfigDrive `json:"drives"`

	// fixedresources
	Fixedresources *ConfigVMConfig `json:"fixedresources,omitempty"`

	// NetworkAdapter are virtual adapters assigned to the application
	// The order here is critical because they are presented to the VM or
	// container in the order they are listed, e.g., the first NetworkAdapter
	// will appear in a Linux VM as eth0. Also, the MAC address is determined
	// based on the order in the list.
	Interfaces []*ConfigNetworkAdapter `json:"interfaces"`

	// metadata type to use for app if provided inside userData
	MetaDataType *ConfigMetaDataType `json:"metaDataType,omitempty"`

	// profile_list is a set of strings which can be used to control which sets
	// of applications are run. Combined with the activate flag above.
	// If the profile list is empty it means wildcard; application will
	// be started independent of the global or local profile specified for the
	// device.
	ProfileList []string `json:"profile_list"`

	// The EVE behavior for a purge command is to restart the application instance
	// with the first drive/volumeRef recreated from its origin.
	Purge *ConfigInstanceOpsCmd `json:"purge,omitempty"`

	// Config flag if the app-instance should be made accessible
	// through a remote console session established by the device.
	RemoteConsole bool `json:"remoteConsole,omitempty"`

	// The device behavior for a restart command (if counter increased)
	// is to restart the application instance
	// Increasing this multiple times does not imply the application instance
	// will restart more than once.
	// EVE can assume that the adapters did not change.
	Restart *ConfigInstanceOpsCmd `json:"restart,omitempty"`

	// App Instance initialization configuration data provided by user
	// This will be used as "user-data" in cloud-init
	// Empty string will indicate that cloud-init is not required
	// It is also used to carry environment variables for containers.
	// XXX will be deprecated and replaced by the cipherData below.
	UserData string `json:"userData,omitempty"`

	// uuidandversion
	Uuidandversion *ConfigUUIDandVersion `json:"uuidandversion,omitempty"`

	// The volumes to be attached to the app-instance.
	// The order here is critical because they are presented to the VM or
	// container in the order they are listed, e.g., the first VM image
	// will be the root disk.
	// Note that since the name volumeRef was used before and deprecated
	// python protobuf seems to require that we use a different name.
	VolumeRefList []*ConfigVolumeRef `json:"volumeRefList"`
}

// Validate validates this config app instance config
func (m *ConfigAppInstanceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdapters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCipherData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDrives(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFixedresources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetaDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUuidandversion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeRefList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigAppInstanceConfig) validateAdapters(formats strfmt.Registry) error {
	if swag.IsZero(m.Adapters) { // not required
		return nil
	}

	for i := 0; i < len(m.Adapters); i++ {
		if swag.IsZero(m.Adapters[i]) { // not required
			continue
		}

		if m.Adapters[i] != nil {
			if err := m.Adapters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adapters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateCipherData(formats strfmt.Registry) error {
	if swag.IsZero(m.CipherData) { // not required
		return nil
	}

	if m.CipherData != nil {
		if err := m.CipherData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cipherData")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateDrives(formats strfmt.Registry) error {
	if swag.IsZero(m.Drives) { // not required
		return nil
	}

	for i := 0; i < len(m.Drives); i++ {
		if swag.IsZero(m.Drives[i]) { // not required
			continue
		}

		if m.Drives[i] != nil {
			if err := m.Drives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateFixedresources(formats strfmt.Registry) error {
	if swag.IsZero(m.Fixedresources) { // not required
		return nil
	}

	if m.Fixedresources != nil {
		if err := m.Fixedresources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fixedresources")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateInterfaces(formats strfmt.Registry) error {
	if swag.IsZero(m.Interfaces) { // not required
		return nil
	}

	for i := 0; i < len(m.Interfaces); i++ {
		if swag.IsZero(m.Interfaces[i]) { // not required
			continue
		}

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateMetaDataType(formats strfmt.Registry) error {
	if swag.IsZero(m.MetaDataType) { // not required
		return nil
	}

	if m.MetaDataType != nil {
		if err := m.MetaDataType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metaDataType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) validatePurge(formats strfmt.Registry) error {
	if swag.IsZero(m.Purge) { // not required
		return nil
	}

	if m.Purge != nil {
		if err := m.Purge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purge")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateRestart(formats strfmt.Registry) error {
	if swag.IsZero(m.Restart) { // not required
		return nil
	}

	if m.Restart != nil {
		if err := m.Restart.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restart")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateUuidandversion(formats strfmt.Registry) error {
	if swag.IsZero(m.Uuidandversion) { // not required
		return nil
	}

	if m.Uuidandversion != nil {
		if err := m.Uuidandversion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidandversion")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) validateVolumeRefList(formats strfmt.Registry) error {
	if swag.IsZero(m.VolumeRefList) { // not required
		return nil
	}

	for i := 0; i < len(m.VolumeRefList); i++ {
		if swag.IsZero(m.VolumeRefList[i]) { // not required
			continue
		}

		if m.VolumeRefList[i] != nil {
			if err := m.VolumeRefList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeRefList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config app instance config based on the context it is used
func (m *ConfigAppInstanceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdapters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCipherData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDrives(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFixedresources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetaDataType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePurge(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRestart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUuidandversion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVolumeRefList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateAdapters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Adapters); i++ {

		if m.Adapters[i] != nil {
			if err := m.Adapters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adapters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateCipherData(ctx context.Context, formats strfmt.Registry) error {

	if m.CipherData != nil {
		if err := m.CipherData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cipherData")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateDrives(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Drives); i++ {

		if m.Drives[i] != nil {
			if err := m.Drives[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("drives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateFixedresources(ctx context.Context, formats strfmt.Registry) error {

	if m.Fixedresources != nil {
		if err := m.Fixedresources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fixedresources")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateInterfaces(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Interfaces); i++ {

		if m.Interfaces[i] != nil {
			if err := m.Interfaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateMetaDataType(ctx context.Context, formats strfmt.Registry) error {

	if m.MetaDataType != nil {
		if err := m.MetaDataType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metaDataType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidatePurge(ctx context.Context, formats strfmt.Registry) error {

	if m.Purge != nil {
		if err := m.Purge.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("purge")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateRestart(ctx context.Context, formats strfmt.Registry) error {

	if m.Restart != nil {
		if err := m.Restart.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restart")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateUuidandversion(ctx context.Context, formats strfmt.Registry) error {

	if m.Uuidandversion != nil {
		if err := m.Uuidandversion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidandversion")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigAppInstanceConfig) contextValidateVolumeRefList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VolumeRefList); i++ {

		if m.VolumeRefList[i] != nil {
			if err := m.VolumeRefList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("volumeRefList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigAppInstanceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigAppInstanceConfig) UnmarshalBinary(b []byte) error {
	var res ConfigAppInstanceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}