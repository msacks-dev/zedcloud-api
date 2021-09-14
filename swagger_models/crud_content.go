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

// CrudContent Content, with a type encoded either as protobuf or json
//
// swagger:model CrudContent
type CrudContent struct {

	// clazz
	Clazz *ModelClazz `json:"clazz,omitempty"`

	// encoding
	Encoding *CrudContentEncoding `json:"encoding,omitempty"`

	// json
	JSON *CrudContentJSON `json:"json,omitempty"`

	// oid
	Oid *Identifier64 `json:"oid,omitempty"`

	// protobuf
	Protobuf *CrudContentProtobuf `json:"protobuf,omitempty"`
}

// Validate validates this crud content
func (m *CrudContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClazz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncoding(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJSON(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtobuf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrudContent) validateClazz(formats strfmt.Registry) error {
	if swag.IsZero(m.Clazz) { // not required
		return nil
	}

	if m.Clazz != nil {
		if err := m.Clazz.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clazz")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) validateEncoding(formats strfmt.Registry) error {
	if swag.IsZero(m.Encoding) { // not required
		return nil
	}

	if m.Encoding != nil {
		if err := m.Encoding.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encoding")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) validateJSON(formats strfmt.Registry) error {
	if swag.IsZero(m.JSON) { // not required
		return nil
	}

	if m.JSON != nil {
		if err := m.JSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("json")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) validateOid(formats strfmt.Registry) error {
	if swag.IsZero(m.Oid) { // not required
		return nil
	}

	if m.Oid != nil {
		if err := m.Oid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oid")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) validateProtobuf(formats strfmt.Registry) error {
	if swag.IsZero(m.Protobuf) { // not required
		return nil
	}

	if m.Protobuf != nil {
		if err := m.Protobuf.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protobuf")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this crud content based on the context it is used
func (m *CrudContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClazz(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEncoding(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateJSON(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOid(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtobuf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrudContent) contextValidateClazz(ctx context.Context, formats strfmt.Registry) error {

	if m.Clazz != nil {
		if err := m.Clazz.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clazz")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) contextValidateEncoding(ctx context.Context, formats strfmt.Registry) error {

	if m.Encoding != nil {
		if err := m.Encoding.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encoding")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) contextValidateJSON(ctx context.Context, formats strfmt.Registry) error {

	if m.JSON != nil {
		if err := m.JSON.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("json")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) contextValidateOid(ctx context.Context, formats strfmt.Registry) error {

	if m.Oid != nil {
		if err := m.Oid.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oid")
			}
			return err
		}
	}

	return nil
}

func (m *CrudContent) contextValidateProtobuf(ctx context.Context, formats strfmt.Registry) error {

	if m.Protobuf != nil {
		if err := m.Protobuf.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("protobuf")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrudContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrudContent) UnmarshalBinary(b []byte) error {
	var res CrudContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}