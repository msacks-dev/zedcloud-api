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

// HealthDesc health desc
//
// swagger:model HealthDesc
type HealthDesc struct {

	// brief health
	BriefHealth *BriefHealth `json:"briefHealth,omitempty"`

	// env name
	EnvName string `json:"envName,omitempty"`

	// host name
	HostName string `json:"hostName,omitempty"`

	// service instance
	ServiceInstance string `json:"serviceInstance,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`
}

// Validate validates this health desc
func (m *HealthDesc) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBriefHealth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthDesc) validateBriefHealth(formats strfmt.Registry) error {
	if swag.IsZero(m.BriefHealth) { // not required
		return nil
	}

	if m.BriefHealth != nil {
		if err := m.BriefHealth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("briefHealth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this health desc based on the context it is used
func (m *HealthDesc) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBriefHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthDesc) contextValidateBriefHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.BriefHealth != nil {
		if err := m.BriefHealth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("briefHealth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthDesc) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthDesc) UnmarshalBinary(b []byte) error {
	var res HealthDesc
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}