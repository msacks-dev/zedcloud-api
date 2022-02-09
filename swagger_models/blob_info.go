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

// BlobInfo blob info
//
// swagger:model BlobInfo
type BlobInfo struct {

	// state
	State *SWState `json:"State,omitempty"`

	// err info
	ErrInfo *DeviceError `json:"errInfo,omitempty"`

	// progress percentage
	ProgressPercentage int64 `json:"progressPercentage,omitempty"`

	// resource
	Resource *VolInstResource `json:"resource,omitempty"`

	// sha256
	Sha256 string `json:"sha256,omitempty"`

	// usage
	Usage *DeviceObjectUsageInfo `json:"usage,omitempty"`
}

// Validate validates this blob info
func (m *BlobInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlobInfo) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("State")
			}
			return err
		}
	}

	return nil
}

func (m *BlobInfo) validateErrInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrInfo) { // not required
		return nil
	}

	if m.ErrInfo != nil {
		if err := m.ErrInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errInfo")
			}
			return err
		}
	}

	return nil
}

func (m *BlobInfo) validateResource(formats strfmt.Registry) error {
	if swag.IsZero(m.Resource) { // not required
		return nil
	}

	if m.Resource != nil {
		if err := m.Resource.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *BlobInfo) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	if m.Usage != nil {
		if err := m.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this blob info based on the context it is used
func (m *BlobInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlobInfo) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("State")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("State")
			}
			return err
		}
	}

	return nil
}

func (m *BlobInfo) contextValidateErrInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrInfo != nil {
		if err := m.ErrInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errInfo")
			}
			return err
		}
	}

	return nil
}

func (m *BlobInfo) contextValidateResource(ctx context.Context, formats strfmt.Registry) error {

	if m.Resource != nil {
		if err := m.Resource.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resource")
			}
			return err
		}
	}

	return nil
}

func (m *BlobInfo) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if m.Usage != nil {
		if err := m.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlobInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlobInfo) UnmarshalBinary(b []byte) error {
	var res BlobInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
