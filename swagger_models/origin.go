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

// Origin OriginType : enum specifies the Object orgigin type
//
// - ORIGIN_UNSPECIFIED: default options, which says no Operation/Invalid Operation
//  - ORIGIN_IMPORTED: Object imported from global enterprise.
//  - ORIGIN_LOCAL: Objectl created locally.
//  - ORIGIN_GLOBAL: Object created in global store,
// to use this type user should have root previlage.
//
// swagger:model Origin
type Origin string

func NewOrigin(value Origin) *Origin {
	v := value
	return &v
}

const (

	// OriginORIGINUNSPECIFIED captures enum value "ORIGIN_UNSPECIFIED"
	OriginORIGINUNSPECIFIED Origin = "ORIGIN_UNSPECIFIED"

	// OriginORIGINIMPORTED captures enum value "ORIGIN_IMPORTED"
	OriginORIGINIMPORTED Origin = "ORIGIN_IMPORTED"

	// OriginORIGINLOCAL captures enum value "ORIGIN_LOCAL"
	OriginORIGINLOCAL Origin = "ORIGIN_LOCAL"

	// OriginORIGINGLOBAL captures enum value "ORIGIN_GLOBAL"
	OriginORIGINGLOBAL Origin = "ORIGIN_GLOBAL"
)

// for schema
var originEnum []interface{}

func init() {
	var res []Origin
	if err := json.Unmarshal([]byte(`["ORIGIN_UNSPECIFIED","ORIGIN_IMPORTED","ORIGIN_LOCAL","ORIGIN_GLOBAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		originEnum = append(originEnum, v)
	}
}

func (m Origin) validateOriginEnum(path, location string, value Origin) error {
	if err := validate.EnumCase(path, location, value, originEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this origin
func (m Origin) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOriginEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this origin based on context it is used
func (m Origin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}