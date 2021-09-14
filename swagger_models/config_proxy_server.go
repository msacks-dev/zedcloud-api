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
)

// ConfigProxyServer config proxy server
//
// swagger:model configProxyServer
type ConfigProxyServer struct {

	// port
	Port int64 `json:"port,omitempty"`

	// proto
	Proto *ConfigproxyProto `json:"proto,omitempty"`

	// server
	Server string `json:"server,omitempty"`
}

// Validate validates this config proxy server
func (m *ConfigProxyServer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProto(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigProxyServer) validateProto(formats strfmt.Registry) error {
	if swag.IsZero(m.Proto) { // not required
		return nil
	}

	if m.Proto != nil {
		if err := m.Proto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proto")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config proxy server based on the context it is used
func (m *ConfigProxyServer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProto(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigProxyServer) contextValidateProto(ctx context.Context, formats strfmt.Registry) error {

	if m.Proto != nil {
		if err := m.Proto.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proto")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigProxyServer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigProxyServer) UnmarshalBinary(b []byte) error {
	var res ConfigProxyServer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}