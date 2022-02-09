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

// ConfigKeyExchangeScheme Security Key Exchange Method
//
// swagger:model configKeyExchangeScheme
type ConfigKeyExchangeScheme string

func NewConfigKeyExchangeScheme(value ConfigKeyExchangeScheme) *ConfigKeyExchangeScheme {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConfigKeyExchangeScheme.
func (m ConfigKeyExchangeScheme) Pointer() *ConfigKeyExchangeScheme {
	return &m
}

const (

	// ConfigKeyExchangeSchemeKEANONE captures enum value "KEA_NONE"
	ConfigKeyExchangeSchemeKEANONE ConfigKeyExchangeScheme = "KEA_NONE"

	// ConfigKeyExchangeSchemeKEAECDH captures enum value "KEA_ECDH"
	ConfigKeyExchangeSchemeKEAECDH ConfigKeyExchangeScheme = "KEA_ECDH"
)

// for schema
var configKeyExchangeSchemeEnum []interface{}

func init() {
	var res []ConfigKeyExchangeScheme
	if err := json.Unmarshal([]byte(`["KEA_NONE","KEA_ECDH"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configKeyExchangeSchemeEnum = append(configKeyExchangeSchemeEnum, v)
	}
}

func (m ConfigKeyExchangeScheme) validateConfigKeyExchangeSchemeEnum(path, location string, value ConfigKeyExchangeScheme) error {
	if err := validate.EnumCase(path, location, value, configKeyExchangeSchemeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config key exchange scheme
func (m ConfigKeyExchangeScheme) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigKeyExchangeSchemeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config key exchange scheme based on context it is used
func (m ConfigKeyExchangeScheme) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
