// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DPSServiceDetail d p s service detail
//
// swagger:model DPSServiceDetail
type DPSServiceDetail struct {

	// enrollment
	Enrollment *EnrollmentDetail `json:"enrollment,omitempty"`

	// service detail
	ServiceDetail *AzureResourceAndServiceDetail `json:"serviceDetail,omitempty"`
}

// Validate validates this d p s service detail
func (m *DPSServiceDetail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEnrollment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceDetail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DPSServiceDetail) validateEnrollment(formats strfmt.Registry) error {
	if swag.IsZero(m.Enrollment) { // not required
		return nil
	}

	if m.Enrollment != nil {
		if err := m.Enrollment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enrollment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enrollment")
			}
			return err
		}
	}

	return nil
}

func (m *DPSServiceDetail) validateServiceDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ServiceDetail) { // not required
		return nil
	}

	if m.ServiceDetail != nil {
		if err := m.ServiceDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceDetail")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this d p s service detail based on the context it is used
func (m *DPSServiceDetail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEnrollment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServiceDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DPSServiceDetail) contextValidateEnrollment(ctx context.Context, formats strfmt.Registry) error {

	if m.Enrollment != nil {
		if err := m.Enrollment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("enrollment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("enrollment")
			}
			return err
		}
	}

	return nil
}

func (m *DPSServiceDetail) contextValidateServiceDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.ServiceDetail != nil {
		if err := m.ServiceDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceDetail")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("serviceDetail")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DPSServiceDetail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DPSServiceDetail) UnmarshalBinary(b []byte) error {
	var res DPSServiceDetail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
