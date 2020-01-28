// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolGameflowGameflowAvailability lol gameflow gameflow availability
// swagger:model LolGameflowGameflowAvailability
type LolGameflowGameflowAvailability struct {

	// is available
	IsAvailable bool `json:"isAvailable,omitempty"`

	// state
	State LolGameflowGameflowAvailabilityState `json:"state,omitempty"`
}

// Validate validates this lol gameflow gameflow availability
func (m *LolGameflowGameflowAvailability) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolGameflowGameflowAvailability) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolGameflowGameflowAvailability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGameflowGameflowAvailability) UnmarshalBinary(b []byte) error {
	var res LolGameflowGameflowAvailability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}