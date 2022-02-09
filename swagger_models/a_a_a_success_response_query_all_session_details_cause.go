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

// AAASuccessResponseQueryAllSessionDetailsCause a a a success response query all session details cause
//
// swagger:model AAASuccessResponseQueryAllSessionDetailsCause
type AAASuccessResponseQueryAllSessionDetailsCause string

func NewAAASuccessResponseQueryAllSessionDetailsCause(value AAASuccessResponseQueryAllSessionDetailsCause) *AAASuccessResponseQueryAllSessionDetailsCause {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AAASuccessResponseQueryAllSessionDetailsCause.
func (m AAASuccessResponseQueryAllSessionDetailsCause) Pointer() *AAASuccessResponseQueryAllSessionDetailsCause {
	return &m
}

const (

	// AAASuccessResponseQueryAllSessionDetailsCauseINVALID captures enum value "__INVALID__"
	AAASuccessResponseQueryAllSessionDetailsCauseINVALID AAASuccessResponseQueryAllSessionDetailsCause = "__INVALID__"

	// AAASuccessResponseQueryAllSessionDetailsCauseOK captures enum value "OK"
	AAASuccessResponseQueryAllSessionDetailsCauseOK AAASuccessResponseQueryAllSessionDetailsCause = "OK"

	// AAASuccessResponseQueryAllSessionDetailsCauseFAILED captures enum value "FAILED"
	AAASuccessResponseQueryAllSessionDetailsCauseFAILED AAASuccessResponseQueryAllSessionDetailsCause = "FAILED"
)

// for schema
var aAASuccessResponseQueryAllSessionDetailsCauseEnum []interface{}

func init() {
	var res []AAASuccessResponseQueryAllSessionDetailsCause
	if err := json.Unmarshal([]byte(`["__INVALID__","OK","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAASuccessResponseQueryAllSessionDetailsCauseEnum = append(aAASuccessResponseQueryAllSessionDetailsCauseEnum, v)
	}
}

func (m AAASuccessResponseQueryAllSessionDetailsCause) validateAAASuccessResponseQueryAllSessionDetailsCauseEnum(path, location string, value AAASuccessResponseQueryAllSessionDetailsCause) error {
	if err := validate.EnumCase(path, location, value, aAASuccessResponseQueryAllSessionDetailsCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a success response query all session details cause
func (m AAASuccessResponseQueryAllSessionDetailsCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAASuccessResponseQueryAllSessionDetailsCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a success response query all session details cause based on context it is used
func (m AAASuccessResponseQueryAllSessionDetailsCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
