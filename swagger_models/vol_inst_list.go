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

// VolInstList vol inst list
//
// swagger:model VolInstList
type VolInstList struct {

	// cfg list
	CfgList []*VolInstConfig `json:"cfgList"`

	// list
	List []*VolInstShortConfig `json:"list"`

	// next
	Next *Cursor `json:"next,omitempty"`
}

// Validate validates this vol inst list
func (m *VolInstList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfgList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolInstList) validateCfgList(formats strfmt.Registry) error {
	if swag.IsZero(m.CfgList) { // not required
		return nil
	}

	for i := 0; i < len(m.CfgList); i++ {
		if swag.IsZero(m.CfgList[i]) { // not required
			continue
		}

		if m.CfgList[i] != nil {
			if err := m.CfgList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfgList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VolInstList) validateList(formats strfmt.Registry) error {
	if swag.IsZero(m.List) { // not required
		return nil
	}

	for i := 0; i < len(m.List); i++ {
		if swag.IsZero(m.List[i]) { // not required
			continue
		}

		if m.List[i] != nil {
			if err := m.List[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VolInstList) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this vol inst list based on the context it is used
func (m *VolInstList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCfgList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VolInstList) contextValidateCfgList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CfgList); i++ {

		if m.CfgList[i] != nil {
			if err := m.CfgList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cfgList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VolInstList) contextValidateList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.List); i++ {

		if m.List[i] != nil {
			if err := m.List[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("list" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VolInstList) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("next")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VolInstList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VolInstList) UnmarshalBinary(b []byte) error {
	var res VolInstList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}