// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// OAUTHProfile o a u t h profile
//
// swagger:model OAUTHProfile
type OAUTHProfile struct {

	// OIDC endpoint for oauth validation
	OIDCEndPoint string `json:"OIDCEndPoint,omitempty"`

	// OAUTH client ID
	ClientID string `json:"clientID,omitempty"`

	// OAUTH client secret
	ClientSecret string `json:"clientSecret,omitempty"`

	// crypto key
	CryptoKey string `json:"cryptoKey,omitempty"`

	// encrypted secrets
	EncryptedSecrets map[string]string `json:"encryptedSecrets,omitempty"`

	// OIDC scope to fetch application role
	RoleScope string `json:"roleScope,omitempty"`
}

// Validate validates this o a u t h profile
func (m *OAUTHProfile) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this o a u t h profile based on context it is used
func (m *OAUTHProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAUTHProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAUTHProfile) UnmarshalBinary(b []byte) error {
	var res OAUTHProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}