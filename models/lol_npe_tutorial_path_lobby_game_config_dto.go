// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolNpeTutorialPathLobbyGameConfigDto lol npe tutorial path lobby game config dto
// swagger:model LolNpeTutorialPathLobbyGameConfigDto
type LolNpeTutorialPathLobbyGameConfigDto struct {

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`
}

// Validate validates this lol npe tutorial path lobby game config dto
func (m *LolNpeTutorialPathLobbyGameConfigDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolNpeTutorialPathLobbyGameConfigDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolNpeTutorialPathLobbyGameConfigDto) UnmarshalBinary(b []byte) error {
	var res LolNpeTutorialPathLobbyGameConfigDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
