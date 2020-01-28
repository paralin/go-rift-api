// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolPlayerBehaviorPlayerBehaviorGameflowSessionResource lol player behavior player behavior gameflow session resource
// swagger:model LolPlayerBehaviorPlayerBehavior_GameflowSessionResource
type LolPlayerBehaviorPlayerBehaviorGameflowSessionResource struct {

	// phase
	Phase LolPlayerBehaviorGameflowPhase `json:"phase,omitempty"`
}

// Validate validates this lol player behavior player behavior gameflow session resource
func (m *LolPlayerBehaviorPlayerBehaviorGameflowSessionResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePhase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPlayerBehaviorPlayerBehaviorGameflowSessionResource) validatePhase(formats strfmt.Registry) error {

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
func (m *LolPlayerBehaviorPlayerBehaviorGameflowSessionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPlayerBehaviorPlayerBehaviorGameflowSessionResource) UnmarshalBinary(b []byte) error {
	var res LolPlayerBehaviorPlayerBehaviorGameflowSessionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}