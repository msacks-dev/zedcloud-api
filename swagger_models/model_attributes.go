// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ModelAttributes model attributes
//
// swagger:model ModelAttributes
type ModelAttributes struct {

	// cpus
	Cpus string `json:"cpus,omitempty"`

	// memory
	Memory string `json:"memory,omitempty"`

	// storage
	Storage string `json:"storage,omitempty"`
}

// Validate validates this model attributes
func (m *ModelAttributes) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this model attributes based on context it is used
func (m *ModelAttributes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ModelAttributes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ModelAttributes) UnmarshalBinary(b []byte) error {
	var res ModelAttributes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}