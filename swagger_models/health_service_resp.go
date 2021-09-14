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

// HealthServiceResp health service resp
//
// swagger:model HealthServiceResp
type HealthServiceResp struct {

	// health desc
	HealthDesc []*HealthDesc `json:"healthDesc"`

	// health service
	HealthService *HealthServiceSubType `json:"healthService,omitempty"`

	// hresult
	Hresult *ZsrvResponse `json:"hresult,omitempty"`
}

// Validate validates this health service resp
func (m *HealthServiceResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHealthDesc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealthService(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHresult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthServiceResp) validateHealthDesc(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthDesc) { // not required
		return nil
	}

	for i := 0; i < len(m.HealthDesc); i++ {
		if swag.IsZero(m.HealthDesc[i]) { // not required
			continue
		}

		if m.HealthDesc[i] != nil {
			if err := m.HealthDesc[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("healthDesc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HealthServiceResp) validateHealthService(formats strfmt.Registry) error {
	if swag.IsZero(m.HealthService) { // not required
		return nil
	}

	if m.HealthService != nil {
		if err := m.HealthService.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthService")
			}
			return err
		}
	}

	return nil
}

func (m *HealthServiceResp) validateHresult(formats strfmt.Registry) error {
	if swag.IsZero(m.Hresult) { // not required
		return nil
	}

	if m.Hresult != nil {
		if err := m.Hresult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hresult")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this health service resp based on the context it is used
func (m *HealthServiceResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHealthDesc(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealthService(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHresult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HealthServiceResp) contextValidateHealthDesc(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.HealthDesc); i++ {

		if m.HealthDesc[i] != nil {
			if err := m.HealthDesc[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("healthDesc" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *HealthServiceResp) contextValidateHealthService(ctx context.Context, formats strfmt.Registry) error {

	if m.HealthService != nil {
		if err := m.HealthService.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("healthService")
			}
			return err
		}
	}

	return nil
}

func (m *HealthServiceResp) contextValidateHresult(ctx context.Context, formats strfmt.Registry) error {

	if m.Hresult != nil {
		if err := m.Hresult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hresult")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HealthServiceResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HealthServiceResp) UnmarshalBinary(b []byte) error {
	var res HealthServiceResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}