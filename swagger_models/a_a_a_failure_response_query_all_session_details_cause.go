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

// AAAFailureResponseQueryAllSessionDetailsCause a a a failure response query all session details cause
//
// swagger:model AAAFailureResponseQueryAllSessionDetailsCause
type AAAFailureResponseQueryAllSessionDetailsCause string

func NewAAAFailureResponseQueryAllSessionDetailsCause(value AAAFailureResponseQueryAllSessionDetailsCause) *AAAFailureResponseQueryAllSessionDetailsCause {
	v := value
	return &v
}

const (

	// AAAFailureResponseQueryAllSessionDetailsCauseINVALID captures enum value "__INVALID__"
	AAAFailureResponseQueryAllSessionDetailsCauseINVALID AAAFailureResponseQueryAllSessionDetailsCause = "__INVALID__"

	// AAAFailureResponseQueryAllSessionDetailsCauseAAAFailureResponseQueryAllSessionDetailsCauseUnknown captures enum value "AAAFailureResponseQueryAllSessionDetailsCauseUnknown"
	AAAFailureResponseQueryAllSessionDetailsCauseAAAFailureResponseQueryAllSessionDetailsCauseUnknown AAAFailureResponseQueryAllSessionDetailsCause = "AAAFailureResponseQueryAllSessionDetailsCauseUnknown"

	// AAAFailureResponseQueryAllSessionDetailsCauseAAAFailureResponseQueryAllSessionDetailsCauseUnauthorizedAccess captures enum value "AAAFailureResponseQueryAllSessionDetailsCauseUnauthorizedAccess"
	AAAFailureResponseQueryAllSessionDetailsCauseAAAFailureResponseQueryAllSessionDetailsCauseUnauthorizedAccess AAAFailureResponseQueryAllSessionDetailsCause = "AAAFailureResponseQueryAllSessionDetailsCauseUnauthorizedAccess"

	// AAAFailureResponseQueryAllSessionDetailsCauseAAAFailureResponseQueryAllSessionDetailsCauseUserUnknown captures enum value "AAAFailureResponseQueryAllSessionDetailsCauseUserUnknown"
	AAAFailureResponseQueryAllSessionDetailsCauseAAAFailureResponseQueryAllSessionDetailsCauseUserUnknown AAAFailureResponseQueryAllSessionDetailsCause = "AAAFailureResponseQueryAllSessionDetailsCauseUserUnknown"
)

// for schema
var aAAFailureResponseQueryAllSessionDetailsCauseEnum []interface{}

func init() {
	var res []AAAFailureResponseQueryAllSessionDetailsCause
	if err := json.Unmarshal([]byte(`["__INVALID__","AAAFailureResponseQueryAllSessionDetailsCauseUnknown","AAAFailureResponseQueryAllSessionDetailsCauseUnauthorizedAccess","AAAFailureResponseQueryAllSessionDetailsCauseUserUnknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aAAFailureResponseQueryAllSessionDetailsCauseEnum = append(aAAFailureResponseQueryAllSessionDetailsCauseEnum, v)
	}
}

func (m AAAFailureResponseQueryAllSessionDetailsCause) validateAAAFailureResponseQueryAllSessionDetailsCauseEnum(path, location string, value AAAFailureResponseQueryAllSessionDetailsCause) error {
	if err := validate.EnumCase(path, location, value, aAAFailureResponseQueryAllSessionDetailsCauseEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this a a a failure response query all session details cause
func (m AAAFailureResponseQueryAllSessionDetailsCause) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAAAFailureResponseQueryAllSessionDetailsCauseEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this a a a failure response query all session details cause based on context it is used
func (m AAAFailureResponseQueryAllSessionDetailsCause) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}