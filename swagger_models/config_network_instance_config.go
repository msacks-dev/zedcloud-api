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

// ConfigNetworkInstanceConfig config network instance config
//
// swagger:model configNetworkInstanceConfig
type ConfigNetworkInstanceConfig struct {

	// activate
	//  - True by default. If set to false ( deactivate), the network instance
	//    configuration is downloaded to the device, but the network instance
	//    itself is not created on the device.
	Activate bool `json:"activate,omitempty"`

	// cfg - Used to pass some feature-specific configuration to the
	//       network instance. For Ex: Lisp, StriongSwan etc
	Cfg *ConfigNetworkInstanceOpaqueConfig `json:"cfg,omitempty"`

	// displayname
	Displayname string `json:"displayname,omitempty"`

	// static DNS entry, if we are running DNS/DHCP service
	DNS []*ConfigZnetStaticDNSEntry `json:"dns"`

	// instType - Type of network instance ( local, bridge etc )
	InstType *ConfigZNetworkInstType `json:"instType,omitempty"`

	// network ip specification
	IP *Configipspec `json:"ip,omitempty"`

	// type of ipSpec
	IPType *ConfigAddressType `json:"ipType,omitempty"`

	// port - Only a single port is supported.
	//    This is used as the external connection for the network instance.
	//    This can be a physical (eth0 ) or logical port (vlan 0).
	//    The port name comes from DeviceConfig ( When it is supported in future).
	//    If the user needs multiple physical ports, Device config should be
	//    used to create a label for multiple physical ports.
	Port *ConfigAdapter `json:"port,omitempty"`

	// uuidandversion
	Uuidandversion *ConfigUUIDandVersion `json:"uuidandversion,omitempty"`
}

// Validate validates this config network instance config
func (m *ConfigNetworkInstanceConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCfg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNS(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUuidandversion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigNetworkInstanceConfig) validateCfg(formats strfmt.Registry) error {
	if swag.IsZero(m.Cfg) { // not required
		return nil
	}

	if m.Cfg != nil {
		if err := m.Cfg.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cfg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cfg")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) validateDNS(formats strfmt.Registry) error {
	if swag.IsZero(m.DNS) { // not required
		return nil
	}

	for i := 0; i < len(m.DNS); i++ {
		if swag.IsZero(m.DNS[i]) { // not required
			continue
		}

		if m.DNS[i] != nil {
			if err := m.DNS[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) validateInstType(formats strfmt.Registry) error {
	if swag.IsZero(m.InstType) { // not required
		return nil
	}

	if m.InstType != nil {
		if err := m.InstType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) validateIP(formats strfmt.Registry) error {
	if swag.IsZero(m.IP) { // not required
		return nil
	}

	if m.IP != nil {
		if err := m.IP.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) validateIPType(formats strfmt.Registry) error {
	if swag.IsZero(m.IPType) { // not required
		return nil
	}

	if m.IPType != nil {
		if err := m.IPType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) validatePort(formats strfmt.Registry) error {
	if swag.IsZero(m.Port) { // not required
		return nil
	}

	if m.Port != nil {
		if err := m.Port.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) validateUuidandversion(formats strfmt.Registry) error {
	if swag.IsZero(m.Uuidandversion) { // not required
		return nil
	}

	if m.Uuidandversion != nil {
		if err := m.Uuidandversion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidandversion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uuidandversion")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this config network instance config based on the context it is used
func (m *ConfigNetworkInstanceConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCfg(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDNS(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInstType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIPType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUuidandversion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidateCfg(ctx context.Context, formats strfmt.Registry) error {

	if m.Cfg != nil {
		if err := m.Cfg.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cfg")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cfg")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidateDNS(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DNS); i++ {

		if m.DNS[i] != nil {
			if err := m.DNS[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dns" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidateInstType(ctx context.Context, formats strfmt.Registry) error {

	if m.InstType != nil {
		if err := m.InstType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("instType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("instType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidateIP(ctx context.Context, formats strfmt.Registry) error {

	if m.IP != nil {
		if err := m.IP.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ip")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ip")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidateIPType(ctx context.Context, formats strfmt.Registry) error {

	if m.IPType != nil {
		if err := m.IPType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipType")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidatePort(ctx context.Context, formats strfmt.Registry) error {

	if m.Port != nil {
		if err := m.Port.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("port")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("port")
			}
			return err
		}
	}

	return nil
}

func (m *ConfigNetworkInstanceConfig) contextValidateUuidandversion(ctx context.Context, formats strfmt.Registry) error {

	if m.Uuidandversion != nil {
		if err := m.Uuidandversion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uuidandversion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uuidandversion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConfigNetworkInstanceConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConfigNetworkInstanceConfig) UnmarshalBinary(b []byte) error {
	var res ConfigNetworkInstanceConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
