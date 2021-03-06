// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DeviceCustomerIDMacAddresses device customer Id mac addresses
//
// swagger:model deviceCustomerIdMac_addresses
type DeviceCustomerIDMacAddresses struct {

	// mac
	Mac interface{} `json:"mac,omitempty"`
}

// Validate validates this device customer Id mac addresses
func (m *DeviceCustomerIDMacAddresses) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeviceCustomerIDMacAddresses) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeviceCustomerIDMacAddresses) UnmarshalBinary(b []byte) error {
	var res DeviceCustomerIDMacAddresses
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
