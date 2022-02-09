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

// PolicyAccess policy access
//
// swagger:model PolicyAccess
type PolicyAccess string

func NewPolicyAccess(value PolicyAccess) *PolicyAccess {
	return &value
}

// Pointer returns a pointer to a freshly-allocated PolicyAccess.
func (m PolicyAccess) Pointer() *PolicyAccess {
	return &m
}

const (

	// PolicyAccessPermissionAccessINVALID captures enum value "PermissionAccess__INVALID__"
	PolicyAccessPermissionAccessINVALID PolicyAccess = "PermissionAccess__INVALID__"

	// PolicyAccessPermissionAccessCreate captures enum value "PermissionAccessCreate"
	PolicyAccessPermissionAccessCreate PolicyAccess = "PermissionAccessCreate"

	// PolicyAccessPermissionAccessRead captures enum value "PermissionAccessRead"
	PolicyAccessPermissionAccessRead PolicyAccess = "PermissionAccessRead"

	// PolicyAccessPermissionAccessCreateRead captures enum value "PermissionAccessCreateRead"
	PolicyAccessPermissionAccessCreateRead PolicyAccess = "PermissionAccessCreateRead"

	// PolicyAccessPermissionAccessUpdate captures enum value "PermissionAccessUpdate"
	PolicyAccessPermissionAccessUpdate PolicyAccess = "PermissionAccessUpdate"

	// PolicyAccessPermissionAccessCreateUpdate captures enum value "PermissionAccessCreateUpdate"
	PolicyAccessPermissionAccessCreateUpdate PolicyAccess = "PermissionAccessCreateUpdate"

	// PolicyAccessPermissionAccessReadUpdate captures enum value "PermissionAccessReadUpdate"
	PolicyAccessPermissionAccessReadUpdate PolicyAccess = "PermissionAccessReadUpdate"

	// PolicyAccessPermissionAccessCreateReadUpdate captures enum value "PermissionAccessCreateReadUpdate"
	PolicyAccessPermissionAccessCreateReadUpdate PolicyAccess = "PermissionAccessCreateReadUpdate"

	// PolicyAccessPermissionAccessDelete captures enum value "PermissionAccessDelete"
	PolicyAccessPermissionAccessDelete PolicyAccess = "PermissionAccessDelete"

	// PolicyAccessPermissionAccessCreateDelete captures enum value "PermissionAccessCreateDelete"
	PolicyAccessPermissionAccessCreateDelete PolicyAccess = "PermissionAccessCreateDelete"

	// PolicyAccessPermissionAccessReadDelete captures enum value "PermissionAccessReadDelete"
	PolicyAccessPermissionAccessReadDelete PolicyAccess = "PermissionAccessReadDelete"

	// PolicyAccessPermissionAccessCreateReadDelete captures enum value "PermissionAccessCreateReadDelete"
	PolicyAccessPermissionAccessCreateReadDelete PolicyAccess = "PermissionAccessCreateReadDelete"

	// PolicyAccessPermissionAccessUpdateDelete captures enum value "PermissionAccessUpdateDelete"
	PolicyAccessPermissionAccessUpdateDelete PolicyAccess = "PermissionAccessUpdateDelete"

	// PolicyAccessPermissionAccessCreateUpdateDelete captures enum value "PermissionAccessCreateUpdateDelete"
	PolicyAccessPermissionAccessCreateUpdateDelete PolicyAccess = "PermissionAccessCreateUpdateDelete"

	// PolicyAccessPermissionAccessReadUpdateDelete captures enum value "PermissionAccessReadUpdateDelete"
	PolicyAccessPermissionAccessReadUpdateDelete PolicyAccess = "PermissionAccessReadUpdateDelete"

	// PolicyAccessPermissionAccessCreateReadUpdateDelete captures enum value "PermissionAccessCreateReadUpdateDelete"
	PolicyAccessPermissionAccessCreateReadUpdateDelete PolicyAccess = "PermissionAccessCreateReadUpdateDelete"

	// PolicyAccessPermissionAccessExecute captures enum value "PermissionAccessExecute"
	PolicyAccessPermissionAccessExecute PolicyAccess = "PermissionAccessExecute"

	// PolicyAccessPermissionAccessCreateExecute captures enum value "PermissionAccessCreateExecute"
	PolicyAccessPermissionAccessCreateExecute PolicyAccess = "PermissionAccessCreateExecute"

	// PolicyAccessPermissionAccessReadExecute captures enum value "PermissionAccessReadExecute"
	PolicyAccessPermissionAccessReadExecute PolicyAccess = "PermissionAccessReadExecute"

	// PolicyAccessPermissionAccessCreateReadExecute captures enum value "PermissionAccessCreateReadExecute"
	PolicyAccessPermissionAccessCreateReadExecute PolicyAccess = "PermissionAccessCreateReadExecute"

	// PolicyAccessPermissionAccessUpdateExecute captures enum value "PermissionAccessUpdateExecute"
	PolicyAccessPermissionAccessUpdateExecute PolicyAccess = "PermissionAccessUpdateExecute"

	// PolicyAccessPermissionAccessCreateUpdateExecute captures enum value "PermissionAccessCreateUpdateExecute"
	PolicyAccessPermissionAccessCreateUpdateExecute PolicyAccess = "PermissionAccessCreateUpdateExecute"

	// PolicyAccessPermissionAccessReadUpdateExecute captures enum value "PermissionAccessReadUpdateExecute"
	PolicyAccessPermissionAccessReadUpdateExecute PolicyAccess = "PermissionAccessReadUpdateExecute"

	// PolicyAccessPermissionAccessCreateReadUpdateExecute captures enum value "PermissionAccessCreateReadUpdateExecute"
	PolicyAccessPermissionAccessCreateReadUpdateExecute PolicyAccess = "PermissionAccessCreateReadUpdateExecute"

	// PolicyAccessPermissionAccessDeleteExecute captures enum value "PermissionAccessDeleteExecute"
	PolicyAccessPermissionAccessDeleteExecute PolicyAccess = "PermissionAccessDeleteExecute"

	// PolicyAccessPermissionAccessCreateDeleteExecute captures enum value "PermissionAccessCreateDeleteExecute"
	PolicyAccessPermissionAccessCreateDeleteExecute PolicyAccess = "PermissionAccessCreateDeleteExecute"

	// PolicyAccessPermissionAccessReadDeleteExecute captures enum value "PermissionAccessReadDeleteExecute"
	PolicyAccessPermissionAccessReadDeleteExecute PolicyAccess = "PermissionAccessReadDeleteExecute"

	// PolicyAccessPermissionAccessCreateReadDeleteExecute captures enum value "PermissionAccessCreateReadDeleteExecute"
	PolicyAccessPermissionAccessCreateReadDeleteExecute PolicyAccess = "PermissionAccessCreateReadDeleteExecute"

	// PolicyAccessPermissionAccessUpdateDeleteExecute captures enum value "PermissionAccessUpdateDeleteExecute"
	PolicyAccessPermissionAccessUpdateDeleteExecute PolicyAccess = "PermissionAccessUpdateDeleteExecute"

	// PolicyAccessPermissionAccessCreateUpdateDeleteExecute captures enum value "PermissionAccessCreateUpdateDeleteExecute"
	PolicyAccessPermissionAccessCreateUpdateDeleteExecute PolicyAccess = "PermissionAccessCreateUpdateDeleteExecute"

	// PolicyAccessPermissionAccessReadUpdateDeleteExecute captures enum value "PermissionAccessReadUpdateDeleteExecute"
	PolicyAccessPermissionAccessReadUpdateDeleteExecute PolicyAccess = "PermissionAccessReadUpdateDeleteExecute"

	// PolicyAccessPermissionAccessCreateReadUpdateDeleteExecute captures enum value "PermissionAccessCreateReadUpdateDeleteExecute"
	PolicyAccessPermissionAccessCreateReadUpdateDeleteExecute PolicyAccess = "PermissionAccessCreateReadUpdateDeleteExecute"

	// PolicyAccessPermissionAccessQuery captures enum value "PermissionAccessQuery"
	PolicyAccessPermissionAccessQuery PolicyAccess = "PermissionAccessQuery"

	// PolicyAccessPermissionAccessCreateQuery captures enum value "PermissionAccessCreateQuery"
	PolicyAccessPermissionAccessCreateQuery PolicyAccess = "PermissionAccessCreateQuery"

	// PolicyAccessPermissionAccessReadQuery captures enum value "PermissionAccessReadQuery"
	PolicyAccessPermissionAccessReadQuery PolicyAccess = "PermissionAccessReadQuery"

	// PolicyAccessPermissionAccessCreateReadQuery captures enum value "PermissionAccessCreateReadQuery"
	PolicyAccessPermissionAccessCreateReadQuery PolicyAccess = "PermissionAccessCreateReadQuery"

	// PolicyAccessPermissionAccessUpdateQuery captures enum value "PermissionAccessUpdateQuery"
	PolicyAccessPermissionAccessUpdateQuery PolicyAccess = "PermissionAccessUpdateQuery"

	// PolicyAccessPermissionAccessCreateUpdateQuery captures enum value "PermissionAccessCreateUpdateQuery"
	PolicyAccessPermissionAccessCreateUpdateQuery PolicyAccess = "PermissionAccessCreateUpdateQuery"

	// PolicyAccessPermissionAccessReadUpdateQuery captures enum value "PermissionAccessReadUpdateQuery"
	PolicyAccessPermissionAccessReadUpdateQuery PolicyAccess = "PermissionAccessReadUpdateQuery"

	// PolicyAccessPermissionAccessCreateReadUpdateQuery captures enum value "PermissionAccessCreateReadUpdateQuery"
	PolicyAccessPermissionAccessCreateReadUpdateQuery PolicyAccess = "PermissionAccessCreateReadUpdateQuery"

	// PolicyAccessPermissionAccessDeleteQuery captures enum value "PermissionAccessDeleteQuery"
	PolicyAccessPermissionAccessDeleteQuery PolicyAccess = "PermissionAccessDeleteQuery"

	// PolicyAccessPermissionAccessCreateDeleteQuery captures enum value "PermissionAccessCreateDeleteQuery"
	PolicyAccessPermissionAccessCreateDeleteQuery PolicyAccess = "PermissionAccessCreateDeleteQuery"

	// PolicyAccessPermissionAccessReadDeleteQuery captures enum value "PermissionAccessReadDeleteQuery"
	PolicyAccessPermissionAccessReadDeleteQuery PolicyAccess = "PermissionAccessReadDeleteQuery"

	// PolicyAccessPermissionAccessCreateReadDeleteQuery captures enum value "PermissionAccessCreateReadDeleteQuery"
	PolicyAccessPermissionAccessCreateReadDeleteQuery PolicyAccess = "PermissionAccessCreateReadDeleteQuery"

	// PolicyAccessPermissionAccessUpdateDeleteQuery captures enum value "PermissionAccessUpdateDeleteQuery"
	PolicyAccessPermissionAccessUpdateDeleteQuery PolicyAccess = "PermissionAccessUpdateDeleteQuery"

	// PolicyAccessPermissionAccessCreateUpdateDeleteQuery captures enum value "PermissionAccessCreateUpdateDeleteQuery"
	PolicyAccessPermissionAccessCreateUpdateDeleteQuery PolicyAccess = "PermissionAccessCreateUpdateDeleteQuery"

	// PolicyAccessPermissionAccessReadUpdateDeleteQuery captures enum value "PermissionAccessReadUpdateDeleteQuery"
	PolicyAccessPermissionAccessReadUpdateDeleteQuery PolicyAccess = "PermissionAccessReadUpdateDeleteQuery"

	// PolicyAccessPermissionAccessCreateReadUpdateDeleteQuery captures enum value "PermissionAccessCreateReadUpdateDeleteQuery"
	PolicyAccessPermissionAccessCreateReadUpdateDeleteQuery PolicyAccess = "PermissionAccessCreateReadUpdateDeleteQuery"

	// PolicyAccessPermissionAccessExecuteQuery captures enum value "PermissionAccessExecuteQuery"
	PolicyAccessPermissionAccessExecuteQuery PolicyAccess = "PermissionAccessExecuteQuery"

	// PolicyAccessPermissionAccessCreateExecuteQuery captures enum value "PermissionAccessCreateExecuteQuery"
	PolicyAccessPermissionAccessCreateExecuteQuery PolicyAccess = "PermissionAccessCreateExecuteQuery"

	// PolicyAccessPermissionAccessReadExecuteQuery captures enum value "PermissionAccessReadExecuteQuery"
	PolicyAccessPermissionAccessReadExecuteQuery PolicyAccess = "PermissionAccessReadExecuteQuery"

	// PolicyAccessPermissionAccessCreateReadExecuteQuery captures enum value "PermissionAccessCreateReadExecuteQuery"
	PolicyAccessPermissionAccessCreateReadExecuteQuery PolicyAccess = "PermissionAccessCreateReadExecuteQuery"

	// PolicyAccessPermissionAccessUpdateExecuteQuery captures enum value "PermissionAccessUpdateExecuteQuery"
	PolicyAccessPermissionAccessUpdateExecuteQuery PolicyAccess = "PermissionAccessUpdateExecuteQuery"

	// PolicyAccessPermissionAccessCreateUpdateExecuteQuery captures enum value "PermissionAccessCreateUpdateExecuteQuery"
	PolicyAccessPermissionAccessCreateUpdateExecuteQuery PolicyAccess = "PermissionAccessCreateUpdateExecuteQuery"

	// PolicyAccessPermissionAccessReadUpdateExecuteQuery captures enum value "PermissionAccessReadUpdateExecuteQuery"
	PolicyAccessPermissionAccessReadUpdateExecuteQuery PolicyAccess = "PermissionAccessReadUpdateExecuteQuery"

	// PolicyAccessPermissionAccessCreateReadUpdateExecuteQuery captures enum value "PermissionAccessCreateReadUpdateExecuteQuery"
	PolicyAccessPermissionAccessCreateReadUpdateExecuteQuery PolicyAccess = "PermissionAccessCreateReadUpdateExecuteQuery"

	// PolicyAccessPermissionAccessDeleteExecuteQuery captures enum value "PermissionAccessDeleteExecuteQuery"
	PolicyAccessPermissionAccessDeleteExecuteQuery PolicyAccess = "PermissionAccessDeleteExecuteQuery"

	// PolicyAccessPermissionAccessCreateDeleteExecuteQuery captures enum value "PermissionAccessCreateDeleteExecuteQuery"
	PolicyAccessPermissionAccessCreateDeleteExecuteQuery PolicyAccess = "PermissionAccessCreateDeleteExecuteQuery"

	// PolicyAccessPermissionAccessReadDeleteExecuteQuery captures enum value "PermissionAccessReadDeleteExecuteQuery"
	PolicyAccessPermissionAccessReadDeleteExecuteQuery PolicyAccess = "PermissionAccessReadDeleteExecuteQuery"

	// PolicyAccessPermissionAccessCreateReadDeleteExecuteQuery captures enum value "PermissionAccessCreateReadDeleteExecuteQuery"
	PolicyAccessPermissionAccessCreateReadDeleteExecuteQuery PolicyAccess = "PermissionAccessCreateReadDeleteExecuteQuery"

	// PolicyAccessPermissionAccessUpdateDeleteExecuteQuery captures enum value "PermissionAccessUpdateDeleteExecuteQuery"
	PolicyAccessPermissionAccessUpdateDeleteExecuteQuery PolicyAccess = "PermissionAccessUpdateDeleteExecuteQuery"

	// PolicyAccessPermissionAccessCreateUpdateDeleteExecuteQuery captures enum value "PermissionAccessCreateUpdateDeleteExecuteQuery"
	PolicyAccessPermissionAccessCreateUpdateDeleteExecuteQuery PolicyAccess = "PermissionAccessCreateUpdateDeleteExecuteQuery"

	// PolicyAccessPermissionAccessReadUpdateDeleteExecuteQuery captures enum value "PermissionAccessReadUpdateDeleteExecuteQuery"
	PolicyAccessPermissionAccessReadUpdateDeleteExecuteQuery PolicyAccess = "PermissionAccessReadUpdateDeleteExecuteQuery"

	// PolicyAccessPermissionAccessCreateReadUpdateDeleteExecuteQuery captures enum value "PermissionAccessCreateReadUpdateDeleteExecuteQuery"
	PolicyAccessPermissionAccessCreateReadUpdateDeleteExecuteQuery PolicyAccess = "PermissionAccessCreateReadUpdateDeleteExecuteQuery"
)

// for schema
var policyAccessEnum []interface{}

func init() {
	var res []PolicyAccess
	if err := json.Unmarshal([]byte(`["PermissionAccess__INVALID__","PermissionAccessCreate","PermissionAccessRead","PermissionAccessCreateRead","PermissionAccessUpdate","PermissionAccessCreateUpdate","PermissionAccessReadUpdate","PermissionAccessCreateReadUpdate","PermissionAccessDelete","PermissionAccessCreateDelete","PermissionAccessReadDelete","PermissionAccessCreateReadDelete","PermissionAccessUpdateDelete","PermissionAccessCreateUpdateDelete","PermissionAccessReadUpdateDelete","PermissionAccessCreateReadUpdateDelete","PermissionAccessExecute","PermissionAccessCreateExecute","PermissionAccessReadExecute","PermissionAccessCreateReadExecute","PermissionAccessUpdateExecute","PermissionAccessCreateUpdateExecute","PermissionAccessReadUpdateExecute","PermissionAccessCreateReadUpdateExecute","PermissionAccessDeleteExecute","PermissionAccessCreateDeleteExecute","PermissionAccessReadDeleteExecute","PermissionAccessCreateReadDeleteExecute","PermissionAccessUpdateDeleteExecute","PermissionAccessCreateUpdateDeleteExecute","PermissionAccessReadUpdateDeleteExecute","PermissionAccessCreateReadUpdateDeleteExecute","PermissionAccessQuery","PermissionAccessCreateQuery","PermissionAccessReadQuery","PermissionAccessCreateReadQuery","PermissionAccessUpdateQuery","PermissionAccessCreateUpdateQuery","PermissionAccessReadUpdateQuery","PermissionAccessCreateReadUpdateQuery","PermissionAccessDeleteQuery","PermissionAccessCreateDeleteQuery","PermissionAccessReadDeleteQuery","PermissionAccessCreateReadDeleteQuery","PermissionAccessUpdateDeleteQuery","PermissionAccessCreateUpdateDeleteQuery","PermissionAccessReadUpdateDeleteQuery","PermissionAccessCreateReadUpdateDeleteQuery","PermissionAccessExecuteQuery","PermissionAccessCreateExecuteQuery","PermissionAccessReadExecuteQuery","PermissionAccessCreateReadExecuteQuery","PermissionAccessUpdateExecuteQuery","PermissionAccessCreateUpdateExecuteQuery","PermissionAccessReadUpdateExecuteQuery","PermissionAccessCreateReadUpdateExecuteQuery","PermissionAccessDeleteExecuteQuery","PermissionAccessCreateDeleteExecuteQuery","PermissionAccessReadDeleteExecuteQuery","PermissionAccessCreateReadDeleteExecuteQuery","PermissionAccessUpdateDeleteExecuteQuery","PermissionAccessCreateUpdateDeleteExecuteQuery","PermissionAccessReadUpdateDeleteExecuteQuery","PermissionAccessCreateReadUpdateDeleteExecuteQuery"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		policyAccessEnum = append(policyAccessEnum, v)
	}
}

func (m PolicyAccess) validatePolicyAccessEnum(path, location string, value PolicyAccess) error {
	if err := validate.EnumCase(path, location, value, policyAccessEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this policy access
func (m PolicyAccess) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validatePolicyAccessEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this policy access based on context it is used
func (m PolicyAccess) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
