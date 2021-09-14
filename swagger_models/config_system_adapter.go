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

// ConfigSystemAdapter systemAdapters, are the higher l2 concept built on physicalIOs.
// systemAdapters, gives all the required bits to turn the physical IOs
// into useful IP endpoints.
// These endpoints can be further used to connect to controller or
// can be shared between workload/services running on the node.
//
// swagger:model configSystemAdapter
type ConfigSystemAdapter struct {

	// addr - if its static network we need ip address
	// If this is specified, networkUUID must also be specified. addr
	// is expected to be in sync with the network object (same subnet etc ).
	Addr string `json:"addr,omitempty"`

	// alias - Device just reflects it back in status / Metrics back to
	// cloud.
	Alias string `json:"alias,omitempty"`

	// cost of using a port for EVE management traffic (which is determined
	// from PhysicalIO.usage)
	// 0 is the lowest cost (free); 255 the highest.
	// Load spreading will apply when multiple adapters have the same cost.
	// Higher cost adapters are only tried when none of the lower cost ones work.
	Cost int64 `json:"cost,omitempty"`

	// deprecated; need level 2 spec. sWAdapterParams allocDetails = 20;
	// this is part of the freelink group
	// DEPRECATED by cost below
	FreeUplink bool `json:"freeUplink,omitempty"`

	// lowerLayerName - For example, if lower layer is PhysicalAdapter
	// ( physical interface), this should point to PhyLabel of the
	// physicalIO.
	LowerLayerName string `json:"lowerLayerName,omitempty"`

	// name - Name of the Network Interface. This is the Port Name
	//  used in Info / Metrics / flowlog etc. Name cannot be changed.
	// This will be the Network Port name.
	Name string `json:"name,omitempty"`

	// networkUUID - attach this network config for this adapter
	// if not set, depending on Usage of Adapter, would be treated as
	// an L2 port
	NetworkUUID string `json:"networkUUID,omitempty"`

	// uplink - DEPRECATED by PhysicalIO.Usage / PhysicalIO.UsagePolicy
	// this is part of the uplink group
	// deprecate: have a separate device policy object in the API
	Uplink bool `json:"uplink,omitempty"`
}

// Validate validates this config system adapter
func (m *ConfigSystemAdapter) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this config system adapter based on context it is used
func (m *ConfigSystemAdapter) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConfigSystemAdapter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigSystemAdapter) UnmarshalBinary(b []byte) error {
	var res ConfigSystemAdapter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}