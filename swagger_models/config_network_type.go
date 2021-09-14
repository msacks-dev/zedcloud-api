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

// ConfigNetworkType config network type
//
// swagger:model configNetworkType
type ConfigNetworkType string

func NewConfigNetworkType(value ConfigNetworkType) *ConfigNetworkType {
	v := value
	return &v
}

const (

	// ConfigNetworkTypeNETWORKTYPENOOP captures enum value "NETWORKTYPENOOP"
	ConfigNetworkTypeNETWORKTYPENOOP ConfigNetworkType = "NETWORKTYPENOOP"

	// ConfigNetworkTypeV4 captures enum value "V4"
	ConfigNetworkTypeV4 ConfigNetworkType = "V4"

	// ConfigNetworkTypeV6 captures enum value "V6"
	ConfigNetworkTypeV6 ConfigNetworkType = "V6"

	// ConfigNetworkTypeCryptoV4 captures enum value "CryptoV4"
	ConfigNetworkTypeCryptoV4 ConfigNetworkType = "CryptoV4"

	// ConfigNetworkTypeCryptoV6 captures enum value "CryptoV6"
	ConfigNetworkTypeCryptoV6 ConfigNetworkType = "CryptoV6"

	// ConfigNetworkTypeCryptoEID captures enum value "CryptoEID"
	ConfigNetworkTypeCryptoEID ConfigNetworkType = "CryptoEID"
)

// for schema
var configNetworkTypeEnum []interface{}

func init() {
	var res []ConfigNetworkType
	if err := json.Unmarshal([]byte(`["NETWORKTYPENOOP","V4","V6","CryptoV4","CryptoV6","CryptoEID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		configNetworkTypeEnum = append(configNetworkTypeEnum, v)
	}
}

func (m ConfigNetworkType) validateConfigNetworkTypeEnum(path, location string, value ConfigNetworkType) error {
	if err := validate.EnumCase(path, location, value, configNetworkTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this config network type
func (m ConfigNetworkType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConfigNetworkTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this config network type based on context it is used
func (m ConfigNetworkType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}