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

// ZsrvErrorCode ZedCould internal error code
//
// - zMsgErrorNone: common validation errors
//  - IncompleteData: message had fields that weren't filled in
//  - InvalidData: message contained the field that wasn't expected
//  - FunctionUnsupported: this feature unavailable on this version of device
//  - InvalidFieldFormat: message contained the field that wasn't correctly formatted
//  - JsonFmtError: Marshal / Unmarshal errors
//  - DataBaseConnection: generic DB error
//
// swagger:model ZsrvErrorCode
type ZsrvErrorCode string

func NewZsrvErrorCode(value ZsrvErrorCode) *ZsrvErrorCode {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ZsrvErrorCode.
func (m ZsrvErrorCode) Pointer() *ZsrvErrorCode {
	return &m
}

const (

	// ZsrvErrorCodeZMsgErrorNone captures enum value "zMsgErrorNone"
	ZsrvErrorCodeZMsgErrorNone ZsrvErrorCode = "zMsgErrorNone"

	// ZsrvErrorCodeZMsgSucess captures enum value "zMsgSucess"
	ZsrvErrorCodeZMsgSucess ZsrvErrorCode = "zMsgSucess"

	// ZsrvErrorCodeNotFound captures enum value "NotFound"
	ZsrvErrorCodeNotFound ZsrvErrorCode = "NotFound"

	// ZsrvErrorCodeAlreadyExists captures enum value "AlreadyExists"
	ZsrvErrorCodeAlreadyExists ZsrvErrorCode = "AlreadyExists"

	// ZsrvErrorCodeVersionMismatch captures enum value "VersionMismatch"
	ZsrvErrorCodeVersionMismatch ZsrvErrorCode = "VersionMismatch"

	// ZsrvErrorCodeRangeError captures enum value "RangeError"
	ZsrvErrorCodeRangeError ZsrvErrorCode = "RangeError"

	// ZsrvErrorCodeLargeResult captures enum value "LargeResult"
	ZsrvErrorCodeLargeResult ZsrvErrorCode = "LargeResult"

	// ZsrvErrorCodeIncompleteData captures enum value "IncompleteData"
	ZsrvErrorCodeIncompleteData ZsrvErrorCode = "IncompleteData"

	// ZsrvErrorCodeInvalidData captures enum value "InvalidData"
	ZsrvErrorCodeInvalidData ZsrvErrorCode = "InvalidData"

	// ZsrvErrorCodeFunctionUnsupported captures enum value "FunctionUnsupported"
	ZsrvErrorCodeFunctionUnsupported ZsrvErrorCode = "FunctionUnsupported"

	// ZsrvErrorCodeNoMemory captures enum value "NoMemory"
	ZsrvErrorCodeNoMemory ZsrvErrorCode = "NoMemory"

	// ZsrvErrorCodeSendFailure captures enum value "SendFailure"
	ZsrvErrorCodeSendFailure ZsrvErrorCode = "SendFailure"

	// ZsrvErrorCodeTimeout captures enum value "Timeout"
	ZsrvErrorCodeTimeout ZsrvErrorCode = "Timeout"

	// ZsrvErrorCodeBadReqBody captures enum value "BadReqBody"
	ZsrvErrorCodeBadReqBody ZsrvErrorCode = "BadReqBody"

	// ZsrvErrorCodeBadReqParam captures enum value "BadReqParam"
	ZsrvErrorCodeBadReqParam ZsrvErrorCode = "BadReqParam"

	// ZsrvErrorCodeInvalidFieldFormat captures enum value "InvalidFieldFormat"
	ZsrvErrorCodeInvalidFieldFormat ZsrvErrorCode = "InvalidFieldFormat"

	// ZsrvErrorCodeURLNotFound captures enum value "UrlNotFound"
	ZsrvErrorCodeURLNotFound ZsrvErrorCode = "UrlNotFound"

	// ZsrvErrorCodeAPIVersionNotSupported captures enum value "ApiVersionNotSupported"
	ZsrvErrorCodeAPIVersionNotSupported ZsrvErrorCode = "ApiVersionNotSupported"

	// ZsrvErrorCodeUnauthorized captures enum value "Unauthorized"
	ZsrvErrorCodeUnauthorized ZsrvErrorCode = "Unauthorized"

	// ZsrvErrorCodeForbidden captures enum value "Forbidden"
	ZsrvErrorCodeForbidden ZsrvErrorCode = "Forbidden"

	// ZsrvErrorCodeConflict captures enum value "Conflict"
	ZsrvErrorCodeConflict ZsrvErrorCode = "Conflict"

	// ZsrvErrorCodeNotModified captures enum value "NotModified"
	ZsrvErrorCodeNotModified ZsrvErrorCode = "NotModified"

	// ZsrvErrorCodeDependencyConflict captures enum value "DependencyConflict"
	ZsrvErrorCodeDependencyConflict ZsrvErrorCode = "DependencyConflict"

	// ZsrvErrorCodeJSONFmtError captures enum value "JsonFmtError"
	ZsrvErrorCodeJSONFmtError ZsrvErrorCode = "JsonFmtError"

	// ZsrvErrorCodeProtoFmtError captures enum value "ProtoFmtError"
	ZsrvErrorCodeProtoFmtError ZsrvErrorCode = "ProtoFmtError"

	// ZsrvErrorCodeCertError captures enum value "CertError"
	ZsrvErrorCodeCertError ZsrvErrorCode = "CertError"

	// ZsrvErrorCodeDataBaseConnection captures enum value "DataBaseConnection"
	ZsrvErrorCodeDataBaseConnection ZsrvErrorCode = "DataBaseConnection"

	// ZsrvErrorCodeDBError captures enum value "DBError"
	ZsrvErrorCodeDBError ZsrvErrorCode = "DBError"

	// ZsrvErrorCodeZMsgAccepted captures enum value "zMsgAccepted"
	ZsrvErrorCodeZMsgAccepted ZsrvErrorCode = "zMsgAccepted"

	// ZsrvErrorCodeZMsgCreated captures enum value "zMsgCreated"
	ZsrvErrorCodeZMsgCreated ZsrvErrorCode = "zMsgCreated"

	// ZsrvErrorCodePreConditionFailed captures enum value "PreConditionFailed"
	ZsrvErrorCodePreConditionFailed ZsrvErrorCode = "PreConditionFailed"
)

// for schema
var zsrvErrorCodeEnum []interface{}

func init() {
	var res []ZsrvErrorCode
	if err := json.Unmarshal([]byte(`["zMsgErrorNone","zMsgSucess","NotFound","AlreadyExists","VersionMismatch","RangeError","LargeResult","IncompleteData","InvalidData","FunctionUnsupported","NoMemory","SendFailure","Timeout","BadReqBody","BadReqParam","InvalidFieldFormat","UrlNotFound","ApiVersionNotSupported","Unauthorized","Forbidden","Conflict","NotModified","DependencyConflict","JsonFmtError","ProtoFmtError","CertError","DataBaseConnection","DBError","zMsgAccepted","zMsgCreated","PreConditionFailed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		zsrvErrorCodeEnum = append(zsrvErrorCodeEnum, v)
	}
}

func (m ZsrvErrorCode) validateZsrvErrorCodeEnum(path, location string, value ZsrvErrorCode) error {
	if err := validate.EnumCase(path, location, value, zsrvErrorCodeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this zsrv error code
func (m ZsrvErrorCode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateZsrvErrorCodeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this zsrv error code based on context it is used
func (m ZsrvErrorCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
