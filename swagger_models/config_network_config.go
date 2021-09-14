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

// ConfigNetworkConfig config network config
//
// swagger:model configNetworkConfig
type ConfigNetworkConfig struct {

	// dns
	DNS []*ConfigZnetStaticDNSEntry `json:"dns"`

	// enterprise proxy
	EntProxy *ConfigProxyConfig `json:"entProxy,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// network ip specification
	IP *Configipspec `json:"ip,omitempty"`

	// type
	Type *ConfigNetworkType `json:"type,omitempty"`

	// wireless specification
	Wireless *ConfigWirelessConfig `json:"wireless,omitempty"`
}

// Validate validates this config network config
func (m *ConfigNetworkConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntProxy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWireless(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigNetworkConfig) validateDNS(formats strfmt.Registry) error {
	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	for i := 0; i < len(m.DNS); i++ {
		if swag.IsZero(m.DNS[i]) { // not required
			continue
		}

		if m.DNS[i] != nil {
			if err := m.DNS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigNetworkConfig) validateEntProxy(formats strfmt.Registry) error {
	if swag.IsZero(m.EntProxy) { // not required
		return nil
	}

	if m.EntProxy != nil {
		if err := m.EntProxy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entProxy")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkConfig) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkConfig) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkConfig) validateWireless(formats strfmt.Registry) error {
	if swag.IsZero(m.Wireless) { // not required
		return nil
	}

	if m.Wireless != nil {
		if err := m.Wireless.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireless")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config network config based on the context it is used
func (m *ConfigNetworkConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntProxy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWireless(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigNetworkConfig) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNS); i++ {

		if m.DNS[i] != nil {
			if err := m.DNS[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigNetworkConfig) contextValidateEntProxy(ctx context.Context, formats strfmt.Registry) error {

	if m.EntProxy != nil {
		if err := m.EntProxy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entProxy")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkConfig) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkConfig) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkConfig) contextValidateWireless(ctx context.Context, formats strfmt.Registry) error {

	if m.Wireless != nil {
		if err := m.Wireless.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wireless")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigNetworkConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigNetworkConfig) UnmarshalBinary(b []byte) error {
	var res ConfigNetworkConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}