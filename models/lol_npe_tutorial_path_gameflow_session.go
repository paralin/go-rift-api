// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathGameflowSession lol npe tutorial path gameflow session
// swagger:model LolNpeTutorialPathGameflowSession
type LolNpeTutorialPathGameflowSession struct {

	// phase
	Phase LolNpeTutorialPathGameflowPhase `json:"phase,omitempty"`
}

// Validate validates this lol npe tutorial path gameflow session
func (m *LolNpeTutorialPathGameflowSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolNpeTutorialPathGameflowSession) validatePhase(formats strfmt.Registry) error {

	if swag.IsZero(m.Phase) { // not required
		return nil
	}

	if err := m.Phase.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("phase")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathGameflowSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathGameflowSession) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathGameflowSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}