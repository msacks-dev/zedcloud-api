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
	"github.com/go-openapi/validate"
)

// SysModel SysModel  payload details
//
// SysModel consists of various model attributes like id, name, title, brandId etc
//
// swagger:model SysModel
type SysModel struct {

	// PCR templates keyed by EVE version
	PCRTemplates []*PCRTemplate `json:"PCRTemplates"`

	// attr
	// Required: true
	Attr map[string]string `json:"attr"`

	// System defined universally unique Id of the brand.
	// Required: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12
	BrandID *string `json:"brandId"`

	// Detailed description of the model.
	// Max Length: 256
	Description string `json:"description,omitempty"`

	// System defined universally unique Id of the model.
	// Read Only: true
	// Pattern: [a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}
	ID string `json:"id,omitempty"`

	// List of IoMembers
	IoMemberList []*IoMember `json:"ioMemberList"`

	// logo
	Logo map[string]string `json:"logo,omitempty"`

	// user defined model name
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9][a-zA-Z0-9_.-]+
	Name *string `json:"name"`

	// origin of object
	OriginType *Origin `json:"originType,omitempty"`

	// origin and parent related details
	ParentDetail *ObjectParentDetail `json:"parentDetail,omitempty"`

	// Product status
	ProductStatus string `json:"productStatus,omitempty"`

	// Product URL
	ProductURL string `json:"productURL,omitempty"`

	// Object Revision  of the model
	// Read Only: true
	Revision *ObjectRevision `json:"revision,omitempty"`

	// SysModel State which denotes the status of the model
	// Required: true
	State *SysModelState `json:"state"`

	// User defined title of the model. Title can be changed at any time.
	// Required: true
	// Max Length: 256
	// Min Length: 3
	// Pattern: [a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+
	Title *string `json:"title"`

	// Defines the Architecture type of the model
	// Required: true
	Type *ModelArchType `json:"type"`
}

// Validate validates this sys model
func (m *SysModel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePCRTemplates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAttr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBrandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIoMemberList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentDetail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysModel) validatePCRTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.PCRTemplates) { // not required
		return nil
	}

	for i := 0; i < len(m.PCRTemplates); i++ {
		if swag.IsZero(m.PCRTemplates[i]) { // not required
			continue
		}

		if m.PCRTemplates[i] != nil {
			if err := m.PCRTemplates[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PCRTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SysModel) validateAttr(formats strfmt.Registry) error {

	if err := validate.Required("attr", "body", m.Attr); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) validateBrandID(formats strfmt.Registry) error {

	if err := validate.Required("brandId", "body", m.BrandID); err != nil {
		return err
	}

	if err := validate.Pattern("brandId", "body", *m.BrandID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12`); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 256); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) validateID(formats strfmt.Registry) error {
	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.Pattern("id", "body", m.ID, `[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}`); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) validateIoMemberList(formats strfmt.Registry) error {
	if swag.IsZero(m.IoMemberList) { // not required
		return nil
	}

	for i := 0; i < len(m.IoMemberList); i++ {
		if swag.IsZero(m.IoMemberList[i]) { // not required
			continue
		}

		if m.IoMemberList[i] != nil {
			if err := m.IoMemberList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ioMemberList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SysModel) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 256); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", *m.Name, `[a-zA-Z0-9][a-zA-Z0-9_.-]+`); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) validateOriginType(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginType) { // not required
		return nil
	}

	if m.OriginType != nil {
		if err := m.OriginType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) validateParentDetail(formats strfmt.Registry) error {
	if swag.IsZero(m.ParentDetail) { // not required
		return nil
	}

	if m.ParentDetail != nil {
		if err := m.ParentDetail.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentDetail")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) validateRevision(formats strfmt.Registry) error {
	if swag.IsZero(m.Revision) { // not required
		return nil
	}

	if m.Revision != nil {
		if err := m.Revision.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	if err := validate.MinLength("title", "body", *m.Title, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("title", "body", *m.Title, 256); err != nil {
		return err
	}

	if err := validate.Pattern("title", "body", *m.Title, `[a-zA-Z0-9]+[a-zA-Z0-9!-~ ]+`); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this sys model based on the context it is used
func (m *SysModel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePCRTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIoMemberList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParentDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRevision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SysModel) contextValidatePCRTemplates(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PCRTemplates); i++ {

		if m.PCRTemplates[i] != nil {
			if err := m.PCRTemplates[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("PCRTemplates" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SysModel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *SysModel) contextValidateIoMemberList(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.IoMemberList); i++ {

		if m.IoMemberList[i] != nil {
			if err := m.IoMemberList[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ioMemberList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SysModel) contextValidateOriginType(ctx context.Context, formats strfmt.Registry) error {

	if m.OriginType != nil {
		if err := m.OriginType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("originType")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) contextValidateParentDetail(ctx context.Context, formats strfmt.Registry) error {

	if m.ParentDetail != nil {
		if err := m.ParentDetail.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parentDetail")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) contextValidateRevision(ctx context.Context, formats strfmt.Registry) error {

	if m.Revision != nil {
		if err := m.Revision.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("revision")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *SysModel) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SysModel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SysModel) UnmarshalBinary(b []byte) error {
	var res SysModel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}