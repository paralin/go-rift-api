// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderReadyStateV1 lol lobby team builder ready state v1
// swagger:model LolLobbyTeamBuilderReadyStateV1
type LolLobbyTeamBuilderReadyStateV1 struct {

	// allowable premade sizes
	AllowablePremadeSizes []int32 `json:"allowablePremadeSizes"`

	// premade size allowed
	PremadeSizeAllowed bool `json:"premadeSizeAllowed,omitempty"`

	// ready to matchmake
	ReadyToMatchmake bool `json:"readyToMatchmake,omitempty"`

	// required position coverage met
	RequiredPositionCoverageMet bool `json:"requiredPositionCoverageMet,omitempty"`
}

// Validate validates this lol lobby team builder ready state v1
func (m *LolLobbyTeamBuilderReadyStateV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderReadyStateV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderReadyStateV1) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderReadyStateV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
