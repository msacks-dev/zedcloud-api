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

// SysModelFilter sys model filter
//
// swagger:model SysModelFilter
type SysModelFilter struct {

	// System defined universally unique Id of the brand.
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12
	BrandID string `json:"brandId,omitempty"`

	// Model name pattern to be matched.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9_.-]{3,256}
	NamePattern string `json:"namePattern,omitempty"`

	// origin of object
	OriginType *Origin `json:"originType,omitempty"`
}

// Validate validates this sys model filter
func (m *SysModelFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBrandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamePattern(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysModelFilter) validateBrandID(formats strfmt.Registry) error {
	if swag.IsZero(m.BrandID) { // not required
		return nil
	}

	if err := validate.Pattern("brandId", "body", m.BrandID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12`); err != nil {
		return err
	}

	return nil
}

func (m *SysModelFilter) validateNamePattern(formats strfmt.Registry) error {
	if swag.IsZero(m.NamePattern) { // not required
		return nil
	}

	if err := validate.MinLength("namePattern", "body", m.NamePattern, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("namePattern", "body", m.NamePattern, 256); err != nil {
		return err
	}

	if err := validate.Pattern("namePattern", "body", m.NamePattern, `[a-zA-Z0-9_.-]{3,256}`); err != nil {
		return err
	}

	return nil
}

func (m *SysModelFilter) validateOriginType(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginType) { // not required
		return nil
	}

	if m.OriginType != nil {
		if err := m.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sys model filter based on the context it is used
func (m *SysModelFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysModelFilter) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginType != nil {
		if err := m.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SysModelFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SysModelFilter) UnmarshalBinary(b []byte) error {
	var res SysModelFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
