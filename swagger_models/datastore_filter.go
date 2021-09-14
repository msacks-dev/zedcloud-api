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

// DatastoreFilter datastore filter
//
// swagger:model DatastoreFilter
type DatastoreFilter struct {

	// Datastore type to be matched
	DsType *DatastoreType `json:"dsType,omitempty"`

	// Datastore name pattern to be matched.
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9_.-]{3,256}
	NamePattern string `json:"namePattern,omitempty"`
}

// Validate validates this datastore filter
func (m *DatastoreFilter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDsType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamePattern(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreFilter) validateDsType(formats strfmt.Registry) error {
	if swag.IsZero(m.DsType) { // not required
		return nil
	}

	if m.DsType != nil {
		if err := m.DsType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dsType")
			}
			return err
		}
	}

	return nil
}

func (m *DatastoreFilter) validateNamePattern(formats strfmt.Registry) error {
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

// ContextValidate validate this datastore filter based on the context it is used
func (m *DatastoreFilter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDsType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatastoreFilter) contextValidateDsType(ctx context.Context, formats strfmt.Registry) error {

	if m.DsType != nil {
		if err := m.DsType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dsType")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatastoreFilter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatastoreFilter) UnmarshalBinary(b []byte) error {
	var res DatastoreFilter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}