// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ListenerConnectionStats listener connection stats
//
// swagger:model Listener_connection_stats
type ListenerConnectionStats struct {

	// limit
	Limit interface{} `json:"limit,omitempty"`

	// offset
	Offset interface{} `json:"offset,omitempty"`

	// service ports
	ServicePorts []*ListenerConnectionStatsServicePortsItems0 `json:"service_ports"`

	// total count
	TotalCount interface{} `json:"total_count,omitempty"`
}

// Validate validates this listener connection stats
func (m *ListenerConnectionStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServicePorts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ListenerConnectionStats) validateServicePorts(formats strfmt.Registry) error {

	if swag.IsZero(m.ServicePorts) { // not required
		return nil
	}

	for i := 0; i < len(m.ServicePorts); i++ {
		if swag.IsZero(m.ServicePorts[i]) { // not required
			continue
		}

		if m.ServicePorts[i] != nil {
			if err := m.ServicePorts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("service_ports" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ListenerConnectionStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerConnectionStats) UnmarshalBinary(b []byte) error {
	var res ListenerConnectionStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ListenerConnectionStatsServicePortsItems0 listener connection stats service ports items0
//
// swagger:model ListenerConnectionStatsServicePortsItems0
type ListenerConnectionStatsServicePortsItems0 struct {

	// appcomp ids
	AppcompIds string `json:"appcomp_ids,omitempty"`

	// client ips
	ClientIps string `json:"client_ips,omitempty"`

	// client ips ids
	ClientIpsIds string `json:"client_ips_ids,omitempty"`

	// client stats
	ClientStats []string `json:"client_stats"`

	// id
	ID int64 `json:"id,omitempty"`

	// listener device id
	ListenerDeviceID int64 `json:"listener_device_id,omitempty"`

	// listener device name
	ListenerDeviceName string `json:"listener_device_name,omitempty"`

	// listener service
	ListenerService string `json:"listener_service,omitempty"`

	// listener service id
	ListenerServiceID int64 `json:"listener_service_id,omitempty"`

	// listening ip
	ListeningIP string `json:"listening_ip,omitempty"`

	// port
	Port int64 `json:"port,omitempty"`

	// protocol
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this listener connection stats service ports items0
func (m *ListenerConnectionStatsServicePortsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ListenerConnectionStatsServicePortsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ListenerConnectionStatsServicePortsItems0) UnmarshalBinary(b []byte) error {
	var res ListenerConnectionStatsServicePortsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}