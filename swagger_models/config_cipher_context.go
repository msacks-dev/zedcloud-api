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

// ConfigCipherContext Cipher information to decrypt Sensitive Data
//
// swagger:model configCipherContext
type ConfigCipherContext struct {

	// cipher context id, key to this structure
	ContextID string `json:"contextId,omitempty"`

	// controller certificate hash
	// Format: byte
	ControllerCertHash strfmt.Base64 `json:"controllerCertHash,omitempty"`

	// device public certificate hash
	// Format: byte
	DeviceCertHash strfmt.Base64 `json:"deviceCertHash,omitempty"`

	// for encrypting sensitive data, like AES256 etc.
	EncryptionScheme *ConfigEncryptionScheme `json:"encryptionScheme,omitempty"`

	// algorithm used to compute hash for certificates
	HashScheme *CommonHashAlgorithm `json:"hashScheme,omitempty"`

	// for key exchange scheme, like ECDH etc.
	KeyExchangeScheme *ConfigKeyExchangeScheme `json:"keyExchangeScheme,omitempty"`
}

// Validate validates this config cipher context
func (m *ConfigCipherContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEncryptionScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHashScheme(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyExchangeScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigCipherContext) validateEncryptionScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.EncryptionScheme) { // not required
		return nil
	}

	if m.EncryptionScheme != nil {
		if err := m.EncryptionScheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionScheme")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigCipherContext) validateHashScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.HashScheme) { // not required
		return nil
	}

	if m.HashScheme != nil {
		if err := m.HashScheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hashScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hashScheme")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigCipherContext) validateKeyExchangeScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.KeyExchangeScheme) { // not required
		return nil
	}

	if m.KeyExchangeScheme != nil {
		if err := m.KeyExchangeScheme.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyExchangeScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyExchangeScheme")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config cipher context based on the context it is used
func (m *ConfigCipherContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEncryptionScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHashScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKeyExchangeScheme(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigCipherContext) contextValidateEncryptionScheme(ctx context.Context, formats strfmt.Registry) error {

	if m.EncryptionScheme != nil {
		if err := m.EncryptionScheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("encryptionScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("encryptionScheme")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigCipherContext) contextValidateHashScheme(ctx context.Context, formats strfmt.Registry) error {

	if m.HashScheme != nil {
		if err := m.HashScheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hashScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hashScheme")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigCipherContext) contextValidateKeyExchangeScheme(ctx context.Context, formats strfmt.Registry) error {

	if m.KeyExchangeScheme != nil {
		if err := m.KeyExchangeScheme.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keyExchangeScheme")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("keyExchangeScheme")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigCipherContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigCipherContext) UnmarshalBinary(b []byte) error {
	var res ConfigCipherContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
