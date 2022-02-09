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

// ConfigNetworkAdapter config network adapter
//
// swagger:model configNetworkAdapter
type ConfigNetworkAdapter struct {

	// access port vlan id
	// app interface with access vlan id of zero will be treated as trunk port
	// valid vlan id range: 2 - 4093
	// vlan id 1 is implicitly used by linux bridges
	AccessVlanID int64 `json:"access_vlan_id,omitempty"`

	// firewall
	Acls []*ConfigACE `json:"acls"`

	// addr
	Addr string `json:"addr,omitempty"`

	// more configuration for getting addr/EID
	CryptoEid string `json:"cryptoEid,omitempty"`

	// hostname
	Hostname string `json:"hostname,omitempty"`

	// lispsignature
	Lispsignature string `json:"lispsignature,omitempty"`

	// Used in case of P2V, where we want to specify a macAddress
	// to vif, that is simulated towards app
	MacAddress string `json:"macAddress,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// network Id
	NetworkID string `json:"networkId,omitempty"`

	// pemcert
	// Format: byte
	Pemcert strfmt.Base64 `json:"pemcert,omitempty"`

	// pemprivatekey
	// Format: byte
	Pemprivatekey strfmt.Base64 `json:"pemprivatekey,omitempty"`
}

// Validate validates this config network adapter
func (m *ConfigNetworkAdapter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcls(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigNetworkAdapter) validateAcls(formats strfmt.Registry) error {
	if swag.IsZero(m.Acls) { // not required
		return nil
	}

	for i := 0; i < len(m.Acls); i++ {
		if swag.IsZero(m.Acls[i]) { // not required
			continue
		}

		if m.Acls[i] != nil {
			if err := m.Acls[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config network adapter based on the context it is used
func (m *ConfigNetworkAdapter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAcls(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigNetworkAdapter) contextValidateAcls(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Acls); i++ {

		if m.Acls[i] != nil {
			if err := m.Acls[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("acls" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("acls" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigNetworkAdapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigNetworkAdapter) UnmarshalBinary(b []byte) error {
	var res ConfigNetworkAdapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
