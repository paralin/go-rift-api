// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RecofrienderLinkResource recofriender link resource
// swagger:model RecofrienderLinkResource
type RecofrienderLinkResource struct {

	// linked
	Linked bool `json:"linked,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// reason
	Reason string `json:"reason,omitempty"`
}

// Validate validates this recofriender link resource
func (m *RecofrienderLinkResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecofrienderLinkResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecofrienderLinkResource) UnmarshalBinary(b []byte) error {
	var res RecofrienderLinkResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
