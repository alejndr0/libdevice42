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

// Customers customers
//
// swagger:model Customers
type Customers struct {

	// contacts
	Contacts []*CustomersContactsItems0 `json:"Contacts"`

	// custom fields
	CustomFields []*CustomersCustomFieldsItems0 `json:"Custom Fields"`

	// contact info
	ContactInfo interface{} `json:"contact_info,omitempty"`

	// devices url
	DevicesURL interface{} `json:"devices_url,omitempty"`

	// groups
	Groups interface{} `json:"groups,omitempty"`

	// id
	ID interface{} `json:"id,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// subnets url
	SubnetsURL interface{} `json:"subnets_url,omitempty"`
}

// Validate validates this customers
func (m *Customers) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContacts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomFields(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Customers) validateContacts(formats strfmt.Registry) error {

	if swag.IsZero(m.Contacts) { // not required
		return nil
	}

	for i := 0; i < len(m.Contacts); i++ {
		if swag.IsZero(m.Contacts[i]) { // not required
			continue
		}

		if m.Contacts[i] != nil {
			if err := m.Contacts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Contacts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Customers) validateCustomFields(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomFields) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomFields); i++ {
		if swag.IsZero(m.CustomFields[i]) { // not required
			continue
		}

		if m.CustomFields[i] != nil {
			if err := m.CustomFields[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Custom Fields" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Customers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Customers) UnmarshalBinary(b []byte) error {
	var res Customers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CustomersContactsItems0 customers contacts items0
//
// swagger:model CustomersContactsItems0
type CustomersContactsItems0 struct {

	// address
	Address interface{} `json:"address,omitempty"`

	// email
	Email interface{} `json:"email,omitempty"`

	// name
	Name interface{} `json:"name,omitempty"`

	// phone
	Phone interface{} `json:"phone,omitempty"`

	// type
	Type interface{} `json:"type,omitempty"`
}

// Validate validates this customers contacts items0
func (m *CustomersContactsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomersContactsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomersContactsItems0) UnmarshalBinary(b []byte) error {
	var res CustomersContactsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CustomersCustomFieldsItems0 customers custom fields items0
//
// swagger:model CustomersCustomFieldsItems0
type CustomersCustomFieldsItems0 struct {

	// key
	Key interface{} `json:"key,omitempty"`

	// notes
	Notes interface{} `json:"notes,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`

	// value2
	Value2 interface{} `json:"value2,omitempty"`
}

// Validate validates this customers custom fields items0
func (m *CustomersCustomFieldsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomersCustomFieldsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomersCustomFieldsItems0) UnmarshalBinary(b []byte) error {
	var res CustomersCustomFieldsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
