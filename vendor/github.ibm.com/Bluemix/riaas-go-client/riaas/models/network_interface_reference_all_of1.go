// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NetworkInterfaceReferenceAllOf1 network interface reference all of1
// swagger:model networkInterfaceReferenceAllOf1
type NetworkInterfaceReferenceAllOf1 struct {

	// The primary IPv4 address
	PrimaryIPV4Address string `json:"primary_ipv4_address,omitempty"`
}

// Validate validates this network interface reference all of1
func (m *NetworkInterfaceReferenceAllOf1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkInterfaceReferenceAllOf1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkInterfaceReferenceAllOf1) UnmarshalBinary(b []byte) error {
	var res NetworkInterfaceReferenceAllOf1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}