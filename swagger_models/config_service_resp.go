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
	"github.com/go-openapi/validate"
)

// ConfigServiceResp Edge Device Configuration Response to CLI/UI routed via Kafka
//
// swagger:model ConfigServiceResp
type ConfigServiceResp struct {

	// config
	Config *ConfigEdgeDevConfig `json:"config,omitempty"`

	// created at
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// last known status
	LastKnownStatus *LastKnownStatus `json:"lastKnownStatus,omitempty"`

	// protobuf stringified
	Pconfig string `json:"pconfig,omitempty"`

	// read at
	// Format: date-time
	ReadAt strfmt.DateTime `json:"readAt,omitempty"`

	// result
	Result *ZsrvResponse `json:"result,omitempty"`

	// updated at
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this config service resp
func (m *ConfigServiceResp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastKnownStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReadAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigServiceResp) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigServiceResp) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigServiceResp) validateLastKnownStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.LastKnownStatus) { // not required
		return nil
	}

	if m.LastKnownStatus != nil {
		if err := m.LastKnownStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastKnownStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigServiceResp) validateReadAt(formats strfmt.Registry) error {
	if swag.IsZero(m.ReadAt) { // not required
		return nil
	}

	if err := validate.FormatOf("readAt", "body", "date-time", m.ReadAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConfigServiceResp) validateResult(formats strfmt.Registry) error {
	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigServiceResp) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this config service resp based on the context it is used
func (m *ConfigServiceResp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastKnownStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigServiceResp) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigServiceResp) contextValidateLastKnownStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.LastKnownStatus != nil {
		if err := m.LastKnownStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastKnownStatus")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigServiceResp) contextValidateResult(ctx context.Context, formats strfmt.Registry) error {

	if m.Result != nil {
		if err := m.Result.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigServiceResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigServiceResp) UnmarshalBinary(b []byte) error {
	var res ConfigServiceResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}