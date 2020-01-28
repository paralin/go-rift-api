// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyAutoFillQueueDto lol lobby auto fill queue dto
// swagger:model LolLobbyAutoFillQueueDto
type LolLobbyAutoFillQueueDto struct {

	// auto fill eligible
	AutoFillEligible bool `json:"autoFillEligible,omitempty"`

	// auto fill protected for promos
	AutoFillProtectedForPromos bool `json:"autoFillProtectedForPromos,omitempty"`

	// auto fill protected for streaking
	AutoFillProtectedForStreaking bool `json:"autoFillProtectedForStreaking,omitempty"`

	// queue Id
	QueueID int32 `json:"queueId,omitempty"`
}

// Validate validates this lol lobby auto fill queue dto
func (m *LolLobbyAutoFillQueueDto) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyAutoFillQueueDto) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyAutoFillQueueDto) UnmarshalBinary(b []byte) error {
	var res LolLobbyAutoFillQueueDto
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
