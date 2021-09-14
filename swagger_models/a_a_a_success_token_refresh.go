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

// AAASuccessTokenRefresh Token64 refresh was successful
//
// swagger:model AAASuccessTokenRefresh
type AAASuccessTokenRefresh struct {

	// token
	Token *Token64 `json:"token,omitempty"`
}

// Validate validates this a a a success token refresh
func (m *AAASuccessTokenRefresh) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessTokenRefresh) validateToken(formats strfmt.Registry) error {
	if swag.IsZero(m.Token) { // not required
		return nil
	}

	if m.Token != nil {
		if err := m.Token.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a success token refresh based on the context it is used
func (m *AAASuccessTokenRefresh) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateToken(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessTokenRefresh) contextValidateToken(ctx context.Context, formats strfmt.Registry) error {

	if m.Token != nil {
		if err := m.Token.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("token")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAASuccessTokenRefresh) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAASuccessTokenRefresh) UnmarshalBinary(b []byte) error {
	var res AAASuccessTokenRefresh
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}