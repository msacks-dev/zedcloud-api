// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IDState Id state
//
// swagger:model IdState
type IDState string

func NewIDState(value IDState) *IDState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated IDState.
func (m IDState) Pointer() *IDState {
	return &m
}

const (

	// IDStateIDSTATEUNSPECIFIED captures enum value "ID_STATE_UNSPECIFIED"
	IDStateIDSTATEUNSPECIFIED IDState = "ID_STATE_UNSPECIFIED"

	// IDStateIDSTATENOTVERIFIED captures enum value "ID_STATE_NOT_VERIFIED"
	IDStateIDSTATENOTVERIFIED IDState = "ID_STATE_NOT_VERIFIED"

	// IDStateIDSTATEVERIFIED captures enum value "ID_STATE_VERIFIED"
	IDStateIDSTATEVERIFIED IDState = "ID_STATE_VERIFIED"
)

// for schema
var idStateEnum []interface{}

func init() {
	var res []IDState
	if err := json.Unmarshal([]byte(`["ID_STATE_UNSPECIFIED","ID_STATE_NOT_VERIFIED","ID_STATE_VERIFIED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		idStateEnum = append(idStateEnum, v)
	}
}

func (m IDState) validateIDStateEnum(path, location string, value IDState) error {
	if err := validate.EnumCase(path, location, value, idStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this Id state
func (m IDState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIDStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this Id state based on context it is used
func (m IDState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
