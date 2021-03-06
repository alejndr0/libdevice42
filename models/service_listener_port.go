// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceListenerPort service listener port
//
// swagger:model Service_listener_port
type ServiceListenerPort struct {

	// service ports
	ServicePorts *ServiceListenerPortServicePorts `json:"service_ports,omitempty"`
}

// Validate validates this service listener port
func (m *ServiceListenerPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServicePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceListenerPort) validateServicePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicePorts) { // not required
		return nil
	}

	if m.ServicePorts != nil {
		if err := m.ServicePorts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("service_ports")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceListenerPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceListenerPort) UnmarshalBinary(b []byte) error {
	var res ServiceListenerPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ServiceListenerPortServicePorts service listener port service ports
//
// swagger:model ServiceListenerPortServicePorts
type ServiceListenerPortServicePorts struct {

	// appcomp ids
	AppcompIds string `json:"appcomp_ids,omitempty"`

	// device name
	DeviceName string `json:"device_name,omitempty"`

	// discovered service
	DiscoveredService string `json:"discovered_service,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// listening ip
	ListeningIP string `json:"listening_ip,omitempty"`

	// mapped service
	MappedService string `json:"mapped_service,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`

	// remote ips
	RemoteIps string `json:"remote_ips,omitempty"`

	// remote ips ids
	RemoteIpsIds string `json:"remote_ips_ids,omitempty"`
}

// Validate validates this service listener port service ports
func (m *ServiceListenerPortServicePorts) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ServiceListenerPortServicePorts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceListenerPortServicePorts) UnmarshalBinary(b []byte) error {
	var res ServiceListenerPortServicePorts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
