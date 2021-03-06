// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathLobbyDto lol npe tutorial path lobby dto
// swagger:model LolNpeTutorialPathLobbyDto
type LolNpeTutorialPathLobbyDto struct {

	// game config
	GameConfig *LolNpeTutorialPathLobbyGameConfigDto `json:"gameConfig,omitempty"`

	// party Id
	PartyID string `json:"partyId,omitempty"`
}

// Validate validates this lol npe tutorial path lobby dto
func (m *LolNpeTutorialPathLobbyDto) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGameConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolNpeTutorialPathLobbyDto) validateGameConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.GameConfig) { // not required
		return nil
	}

	if m.GameConfig != nil {
		if err := m.GameConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gameConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathLobbyDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathLobbyDto) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathLobbyDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
