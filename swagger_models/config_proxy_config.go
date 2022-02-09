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

// ConfigProxyConfig config proxy config
//
// swagger:model configProxyConfig
type ConfigProxyConfig struct {

	// exceptions separated by commas
	Exceptions string `json:"exceptions,omitempty"`

	// enable network level proxy in the form of WPAD
	NetworkProxyEnable bool `json:"networkProxyEnable,omitempty"`

	// Direct URL for wpad.dat download
	NetworkProxyURL string `json:"networkProxyURL,omitempty"`

	// or pacfile can be in place of others
	// base64 encoded
	Pacfile string `json:"pacfile,omitempty"`

	// dedicated per protocol information
	Proxies []*ConfigProxyServer `json:"proxies"`

	// Uploaded proxy certificate or certificate chain for MITM
	// this may be needed either in explicit (has ProxyServer items), automatic
	// (networkProxyEnable) or transparent (network layer not aware of proxy)
	ProxyCertPEM []strfmt.Base64 `json:"proxyCertPEM"`
}

// Validate validates this config proxy config
func (m *ConfigProxyConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProxies(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigProxyConfig) validateProxies(formats strfmt.Registry) error {
	if swag.IsZero(m.Proxies) { // not required
		return nil
	}

	for i := 0; i < len(m.Proxies); i++ {
		if swag.IsZero(m.Proxies[i]) { // not required
			continue
		}

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this config proxy config based on the context it is used
func (m *ConfigProxyConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProxies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigProxyConfig) contextValidateProxies(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Proxies); i++ {

		if m.Proxies[i] != nil {
			if err := m.Proxies[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("proxies" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("proxies" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigProxyConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigProxyConfig) UnmarshalBinary(b []byte) error {
	var res ConfigProxyConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
