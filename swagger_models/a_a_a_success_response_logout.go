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

// AAASuccessResponseLogout a a a success response logout
//
// swagger:model AAASuccessResponseLogout
type AAASuccessResponseLogout struct {

	// original
	Original *OpaqueToken64 `json:"original,omitempty"`
}

// Validate validates this a a a success response logout
func (m *AAASuccessResponseLogout) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOriginal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessResponseLogout) validateOriginal(formats strfmt.Registry) error {
	if swag.IsZero(m.Original) { // not required
		return nil
	}

	if m.Original != nil {
		if err := m.Original.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this a a a success response logout based on the context it is used
func (m *AAASuccessResponseLogout) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOriginal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessResponseLogout) contextValidateOriginal(ctx context.Context, formats strfmt.Registry) error {

	if m.Original != nil {
		if err := m.Original.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("original")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAASuccessResponseLogout) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAASuccessResponseLogout) UnmarshalBinary(b []byte) error {
	var res AAASuccessResponseLogout
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}