// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginSettings PluginSettings Settings that can be modified by users.
// swagger:model PluginSettings
type PluginSettings struct {

	// args
	// Required: true
	Args []string `json:"Args"`

	// devices
	// Required: true
	Devices []*PluginDevice `json:"Devices"`

	// env
	// Required: true
	Env []string `json:"Env"`

	// mounts
	// Required: true
	Mounts []*PluginMount `json:"Mounts"`
}

// Validate validates this plugin settings
func (m *PluginSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArgs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnv(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginSettings) validateArgs(formats strfmt.Registry) error {

	if err := validate.Required("Args", "body", m.Args); err != nil {
		return err
	}

	return nil
}

func (m *PluginSettings) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("Devices", "body", m.Devices); err != nil {
		return err
	}

	for i := 0; i < len(m.Devices); i++ {
		if swag.IsZero(m.Devices[i]) { // not required
			continue
		}

		if m.Devices[i] != nil {
			if err := m.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PluginSettings) validateEnv(formats strfmt.Registry) error {

	if err := validate.Required("Env", "body", m.Env); err != nil {
		return err
	}

	return nil
}

func (m *PluginSettings) validateMounts(formats strfmt.Registry) error {

	if err := validate.Required("Mounts", "body", m.Mounts); err != nil {
		return err
	}

	for i := 0; i < len(m.Mounts); i++ {
		if swag.IsZero(m.Mounts[i]) { // not required
			continue
		}

		if m.Mounts[i] != nil {
			if err := m.Mounts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Mounts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PluginSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginSettings) UnmarshalBinary(b []byte) error {
	var res PluginSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
