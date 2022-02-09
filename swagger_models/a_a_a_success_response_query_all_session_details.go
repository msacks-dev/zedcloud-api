// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AAASuccessResponseQueryAllSessionDetails a a a success response query all session details
//
// swagger:model AAASuccessResponseQueryAllSessionDetails
type AAASuccessResponseQueryAllSessionDetails struct {

	// cause
	Cause *AAASuccessResponseQueryAllSessionDetailsCause `json:"cause,omitempty"`

	// session details
	SessionDetails []*SessionDetails `json:"sessionDetails"`
}

// Validate validates this a a a success response query all session details
func (m *AAASuccessResponseQueryAllSessionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCause(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessResponseQueryAllSessionDetails) validateCause(formats strfmt.Registry) error {
	if swag.IsZero(m.Cause) { // not required
		return nil
	}

	if m.Cause != nil {
		if err := m.Cause.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *AAASuccessResponseQueryAllSessionDetails) validateSessionDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.SessionDetails); i++ {
		if swag.IsZero(m.SessionDetails[i]) { // not required
			continue
		}

		if m.SessionDetails[i] != nil {
			if err := m.SessionDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessionDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sessionDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this a a a success response query all session details based on the context it is used
func (m *AAASuccessResponseQueryAllSessionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCause(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AAASuccessResponseQueryAllSessionDetails) contextValidateCause(ctx context.Context, formats strfmt.Registry) error {

	if m.Cause != nil {
		if err := m.Cause.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cause")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cause")
			}
			return err
		}
	}

	return nil
}

func (m *AAASuccessResponseQueryAllSessionDetails) contextValidateSessionDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SessionDetails); i++ {

		if m.SessionDetails[i] != nil {
			if err := m.SessionDetails[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sessionDetails" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sessionDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AAASuccessResponseQueryAllSessionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AAASuccessResponseQueryAllSessionDetails) UnmarshalBinary(b []byte) error {
	var res AAASuccessResponseQueryAllSessionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
