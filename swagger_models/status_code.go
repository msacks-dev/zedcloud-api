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

// StatusCode status codes for internal use (Gilas and Niles)
//
// swagger:model StatusCode
type StatusCode string

func NewStatusCode(value StatusCode) *StatusCode {
	v := value
	return &v
}

const (

	// StatusCodeSTATUSCODEUNSPECIFIED captures enum value "STATUS_CODE_UNSPECIFIED"
	StatusCodeSTATUSCODEUNSPECIFIED StatusCode = "STATUS_CODE_UNSPECIFIED"

	// StatusCodeSTATUSCODESUCCESS captures enum value "STATUS_CODE_SUCCESS"
	StatusCodeSTATUSCODESUCCESS StatusCode = "STATUS_CODE_SUCCESS"

	// StatusCodeSTATUSCODEFAILURE captures enum value "STATUS_CODE_FAILURE"
	StatusCodeSTATUSCODEFAILURE StatusCode = "STATUS_CODE_FAILURE"

	// StatusCodeSTATUSCODEPARTIALSUCCESS captures enum value "STATUS_CODE_PARTIAL_SUCCESS"
	StatusCodeSTATUSCODEPARTIALSUCCESS StatusCode = "STATUS_CODE_PARTIAL_SUCCESS"
)

// for schema
var statusCodeEnum []interface{}

func init() {
	var res []StatusCode
	if err := json.Unmarshal([]byte(`["STATUS_CODE_UNSPECIFIED","STATUS_CODE_SUCCESS","STATUS_CODE_FAILURE","STATUS_CODE_PARTIAL_SUCCESS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusCodeEnum = append(statusCodeEnum, v)
	}
}

func (m StatusCode) validateStatusCodeEnum(path, location string, value StatusCode) error {
	if err := validate.EnumCase(path, location, value, statusCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this status code
func (m StatusCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this status code based on context it is used
func (m StatusCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}