// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPremadeVoiceGameEventHotkeys lol premade voice game event hotkeys
// swagger:model LolPremadeVoiceGameEventHotkeys
type LolPremadeVoiceGameEventHotkeys struct {

	// evt push to talk
	EvtPushToTalk string `json:"evtPushToTalk,omitempty"`
}

// Validate validates this lol premade voice game event hotkeys
func (m *LolPremadeVoiceGameEventHotkeys) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPremadeVoiceGameEventHotkeys) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPremadeVoiceGameEventHotkeys) UnmarshalBinary(b []byte) error {
	var res LolPremadeVoiceGameEventHotkeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
