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

// AdapterUsage Adapter Usage
//
// - ADAPTER_USAGE_UNSPECIFIED: Adapter unspecified
//  - ADAPTER_USAGE_MANAGEMENT: Adapter can be used by EVE as well as other Edge applications
//  - ADAPTER_USAGE_APP_DIRECT: Adapter is directly used by one edge application
//  - ADAPTER_USAGE_APP_SHARED: Adapter can be shared by different network instances
//  - ADAPTER_USAGE_DISABLED: Adapter disabled, for future use
//
// swagger:model AdapterUsage
type AdapterUsage string

func NewAdapterUsage(value AdapterUsage) *AdapterUsage {
	v := value
	return &v
}

const (

	// AdapterUsageADAPTERUSAGEUNSPECIFIED captures enum value "ADAPTER_USAGE_UNSPECIFIED"
	AdapterUsageADAPTERUSAGEUNSPECIFIED AdapterUsage = "ADAPTER_USAGE_UNSPECIFIED"

	// AdapterUsageADAPTERUSAGEMANAGEMENT captures enum value "ADAPTER_USAGE_MANAGEMENT"
	AdapterUsageADAPTERUSAGEMANAGEMENT AdapterUsage = "ADAPTER_USAGE_MANAGEMENT"

	// AdapterUsageADAPTERUSAGEAPPDIRECT captures enum value "ADAPTER_USAGE_APP_DIRECT"
	AdapterUsageADAPTERUSAGEAPPDIRECT AdapterUsage = "ADAPTER_USAGE_APP_DIRECT"

	// AdapterUsageADAPTERUSAGEAPPSHARED captures enum value "ADAPTER_USAGE_APP_SHARED"
	AdapterUsageADAPTERUSAGEAPPSHARED AdapterUsage = "ADAPTER_USAGE_APP_SHARED"

	// AdapterUsageADAPTERUSAGEDISABLED captures enum value "ADAPTER_USAGE_DISABLED"
	AdapterUsageADAPTERUSAGEDISABLED AdapterUsage = "ADAPTER_USAGE_DISABLED"
)

// for schema
var adapterUsageEnum []interface{}

func init() {
	var res []AdapterUsage
	if err := json.Unmarshal([]byte(`["ADAPTER_USAGE_UNSPECIFIED","ADAPTER_USAGE_MANAGEMENT","ADAPTER_USAGE_APP_DIRECT","ADAPTER_USAGE_APP_SHARED","ADAPTER_USAGE_DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		adapterUsageEnum = append(adapterUsageEnum, v)
	}
}

func (m AdapterUsage) validateAdapterUsageEnum(path, location string, value AdapterUsage) error {
	if err := validate.EnumCase(path, location, value, adapterUsageEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this adapter usage
func (m AdapterUsage) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAdapterUsageEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this adapter usage based on context it is used
func (m AdapterUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}