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

// AAASuccessResponseType a a a success response type
//
// swagger:model AAASuccessResponseType
type AAASuccessResponseType string

func NewAAASuccessResponseType(value AAASuccessResponseType) *AAASuccessResponseType {
	v := value
	return &v
}

const (

	// AAASuccessResponseTypeINVALID captures enum value "__INVALID__"
	AAASuccessResponseTypeINVALID AAASuccessResponseType = "__INVALID__"

	// AAASuccessResponseTypeAAASuccessResponseTypeLogin captures enum value "AAASuccessResponseTypeLogin"
	AAASuccessResponseTypeAAASuccessResponseTypeLogin AAASuccessResponseType = "AAASuccessResponseTypeLogin"

	// AAASuccessResponseTypeAAASuccessResponseTypeRefresh captures enum value "AAASuccessResponseTypeRefresh"
	AAASuccessResponseTypeAAASuccessResponseTypeRefresh AAASuccessResponseType = "AAASuccessResponseTypeRefresh"

	// AAASuccessResponseTypeAAASuccessResponseTypePermisson captures enum value "AAASuccessResponseTypePermisson"
	AAASuccessResponseTypeAAASuccessResponseTypePermisson AAASuccessResponseType = "AAASuccessResponseTypePermisson"

	// AAASuccessResponseTypeAAASuccessResponseTypeLogout captures enum value "AAASuccessResponseTypeLogout"
	AAASuccessResponseTypeAAASuccessResponseTypeLogout AAASuccessResponseType = "AAASuccessResponseTypeLogout"

	// AAASuccessResponseTypeAAASuccessResponseTypeSessionDetails captures enum value "AAASuccessResponseTypeSessionDetails"
	AAASuccessResponseTypeAAASuccessResponseTypeSessionDetails AAASuccessResponseType = "AAASuccessResponseTypeSessionDetails"

	// AAASuccessResponseTypeAAASuccessResponseTypeCredentialChange captures enum value "AAASuccessResponseTypeCredentialChange"
	AAASuccessResponseTypeAAASuccessResponseTypeCredentialChange AAASuccessResponseType = "AAASuccessResponseTypeCredentialChange"

	// AAASuccessResponseTypeAAASuccessResponseTypeQueryAllSessionDetails captures enum value "AAASuccessResponseTypeQueryAllSessionDetails"
	AAASuccessResponseTypeAAASuccessResponseTypeQueryAllSessionDetails AAASuccessResponseType = "AAASuccessResponseTypeQueryAllSessionDetails"

	// AAASuccessResponseTypeAAASuccessResponseTypeGenerateToken captures enum value "AAASuccessResponseTypeGenerateToken"
	AAASuccessResponseTypeAAASuccessResponseTypeGenerateToken AAASuccessResponseType = "AAASuccessResponseTypeGenerateToken"
)

// for schema
var aAASuccessResponseTypeEnum []interface{}

func init() {
	var res []AAASuccessResponseType
	if err := json.Unmarshal([]byte(`["__INVALID__","AAASuccessResponseTypeLogin","AAASuccessResponseTypeRefresh","AAASuccessResponseTypePermisson","AAASuccessResponseTypeLogout","AAASuccessResponseTypeSessionDetails","AAASuccessResponseTypeCredentialChange","AAASuccessResponseTypeQueryAllSessionDetails","AAASuccessResponseTypeGenerateToken"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAASuccessResponseTypeEnum = append(aAASuccessResponseTypeEnum, v)
	}
}

func (m AAASuccessResponseType) validateAAASuccessResponseTypeEnum(path, location string, value AAASuccessResponseType) error {
	if err := validate.EnumCase(path, location, value, aAASuccessResponseTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a success response type
func (m AAASuccessResponseType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAASuccessResponseTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a success response type based on context it is used
func (m AAASuccessResponseType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}