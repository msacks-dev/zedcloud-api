// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

// Code generated by go-swagger; DO NOT EDIT.

package swagger_models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DecsriptionCode description codes for internal use (Gilas and Niles)
//
// swagger:model DecsriptionCode
type DecsriptionCode string

func NewDecsriptionCode(value DecsriptionCode) *DecsriptionCode {
	v := value
	return &v
}

const (

	// DecsriptionCodeDESCRIPTIONCODEUNSPECIFIED captures enum value "DESCRIPTION_CODE_UNSPECIFIED"
	DecsriptionCodeDESCRIPTIONCODEUNSPECIFIED DecsriptionCode = "DESCRIPTION_CODE_UNSPECIFIED"

	// DecsriptionCodeDESCRIPTIONCODERECEIVED captures enum value "DESCRIPTION_CODE_RECEIVED"
	DecsriptionCodeDESCRIPTIONCODERECEIVED DecsriptionCode = "DESCRIPTION_CODE_RECEIVED"

	// DecsriptionCodeDESCRIPTIONCODEPACKETCORRUPT captures enum value "DESCRIPTION_CODE_PACKET_CORRUPT"
	DecsriptionCodeDESCRIPTIONCODEPACKETCORRUPT DecsriptionCode = "DESCRIPTION_CODE_PACKET_CORRUPT"

	// DecsriptionCodeDESCRIPTIONCODEFAILEDTORECEIVE captures enum value "DESCRIPTION_CODE_FAILED_TO_RECEIVE"
	DecsriptionCodeDESCRIPTIONCODEFAILEDTORECEIVE DecsriptionCode = "DESCRIPTION_CODE_FAILED_TO_RECEIVE"

	// DecsriptionCodeDESCRIPTIONCODENETWORKERROR captures enum value "DESCRIPTION_CODE_NETWORK_ERROR"
	DecsriptionCodeDESCRIPTIONCODENETWORKERROR DecsriptionCode = "DESCRIPTION_CODE_NETWORK_ERROR"

	// DecsriptionCodeDESCRIPTIONCODENOTFOUND captures enum value "DESCRIPTION_CODE_NOT_FOUND"
	DecsriptionCodeDESCRIPTIONCODENOTFOUND DecsriptionCode = "DESCRIPTION_CODE_NOT_FOUND"

	// DecsriptionCodeDESCRIPTIONCODENOTSUPPORTED captures enum value "DESCRIPTION_CODE_NOT_SUPPORTED"
	DecsriptionCodeDESCRIPTIONCODENOTSUPPORTED DecsriptionCode = "DESCRIPTION_CODE_NOT_SUPPORTED"

	// DecsriptionCodeDESCRIPTIONCODEINTERNALERROR captures enum value "DESCRIPTION_CODE_INTERNAL_ERROR"
	DecsriptionCodeDESCRIPTIONCODEINTERNALERROR DecsriptionCode = "DESCRIPTION_CODE_INTERNAL_ERROR"

	// DecsriptionCodeDESCRIPTIONCODENOTSPECIFIED captures enum value "DESCRIPTION_CODE_NOT_SPECIFIED"
	DecsriptionCodeDESCRIPTIONCODENOTSPECIFIED DecsriptionCode = "DESCRIPTION_CODE_NOT_SPECIFIED"

	// DecsriptionCodeDESCRIPTIONCODEUPLOADDONE captures enum value "DESCRIPTION_CODE_UPLOAD_DONE"
	DecsriptionCodeDESCRIPTIONCODEUPLOADDONE DecsriptionCode = "DESCRIPTION_CODE_UPLOAD_DONE"

	// DecsriptionCodeDESCRIPTIONCODEALREADYEXISTS captures enum value "DESCRIPTION_CODE_ALREADY_EXISTS"
	DecsriptionCodeDESCRIPTIONCODEALREADYEXISTS DecsriptionCode = "DESCRIPTION_CODE_ALREADY_EXISTS"

	// DecsriptionCodeDESCRIPTIONCODEINVALIDPARAMS captures enum value "DESCRIPTION_CODE_INVALID_PARAMS"
	DecsriptionCodeDESCRIPTIONCODEINVALIDPARAMS DecsriptionCode = "DESCRIPTION_CODE_INVALID_PARAMS"
)

// for schema
var decsriptionCodeEnum []interface{}

func init() {
	var res []DecsriptionCode
	if err := json.Unmarshal([]byte(`["DESCRIPTION_CODE_UNSPECIFIED","DESCRIPTION_CODE_RECEIVED","DESCRIPTION_CODE_PACKET_CORRUPT","DESCRIPTION_CODE_FAILED_TO_RECEIVE","DESCRIPTION_CODE_NETWORK_ERROR","DESCRIPTION_CODE_NOT_FOUND","DESCRIPTION_CODE_NOT_SUPPORTED","DESCRIPTION_CODE_INTERNAL_ERROR","DESCRIPTION_CODE_NOT_SPECIFIED","DESCRIPTION_CODE_UPLOAD_DONE","DESCRIPTION_CODE_ALREADY_EXISTS","DESCRIPTION_CODE_INVALID_PARAMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		decsriptionCodeEnum = append(decsriptionCodeEnum, v)
	}
}

func (m DecsriptionCode) validateDecsriptionCodeEnum(path, location string, value DecsriptionCode) error {
	if err := validate.EnumCase(path, location, value, decsriptionCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this decsription code
func (m DecsriptionCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDecsriptionCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this decsription code based on context it is used
func (m DecsriptionCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}