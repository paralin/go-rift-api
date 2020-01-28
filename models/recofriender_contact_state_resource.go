// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RecofrienderContactStateResource recofriender contact state resource
// swagger:model RecofrienderContactStateResource
type RecofrienderContactStateResource struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// action
	Action string `json:"action,omitempty"`

	// display state
	DisplayState string `json:"displayState,omitempty"`
}

// Validate validates this recofriender contact state resource
func (m *RecofrienderContactStateResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecofrienderContactStateResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecofrienderContactStateResource) UnmarshalBinary(b []byte) error {
	var res RecofrienderContactStateResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
