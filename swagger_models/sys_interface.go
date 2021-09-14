// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SysInterface sysInterface payload detail
//
// system interfaces that needs to be used by dom0
//
// swagger:model sysInterface
type SysInterface struct {

	// cost of using this interface. Default is 0.
	// Maximum: 255
	Cost int64 `json:"cost,omitempty"`

	// Adapter Udage
	IntfUsage *AdapterUsage `json:"intfUsage,omitempty"`

	// name of interface in the manifest to which this network or adapter maps to
	Intfname string `json:"intfname,omitempty"`

	// IP address: we will be needing this in cae of static network
	Ipaddr string `json:"ipaddr,omitempty"`

	// mac address needs to be over-written in some cases
	Macaddr string `json:"macaddr,omitempty"`

	// network name: if attaching a network use netname
	Netname string `json:"netname,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`
}

// Validate validates this sys interface
func (m *SysInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntfUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysInterface) validateCost(formats strfmt.Registry) error {
	if swag.IsZero(m.Cost) { // not required
		return nil
	}

	if err := validate.MaximumInt("cost", "body", m.Cost, 255, false); err != nil {
		return err
	}

	return nil
}

func (m *SysInterface) validateIntfUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.IntfUsage) { // not required
		return nil
	}

	if m.IntfUsage != nil {
		if err := m.IntfUsage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intfUsage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sys interface based on the context it is used
func (m *SysInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntfUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysInterface) contextValidateIntfUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.IntfUsage != nil {
		if err := m.IntfUsage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("intfUsage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SysInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SysInterface) UnmarshalBinary(b []byte) error {
	var res SysInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}