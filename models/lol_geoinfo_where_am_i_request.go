// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolGeoinfoWhereAmIRequest lol geoinfo where am i request
// swagger:model LolGeoinfoWhereAmIRequest
type LolGeoinfoWhereAmIRequest struct {

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`
}

// Validate validates this lol geoinfo where am i request
func (m *LolGeoinfoWhereAmIRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolGeoinfoWhereAmIRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGeoinfoWhereAmIRequest) UnmarshalBinary(b []byte) error {
	var res LolGeoinfoWhereAmIRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}