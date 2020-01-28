// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectLegacyChampSelectTimer lol champ select legacy champ select timer
// swagger:model LolChampSelectLegacyChampSelectTimer
type LolChampSelectLegacyChampSelectTimer struct {

	// adjusted time left in phase
	AdjustedTimeLeftInPhase int64 `json:"adjustedTimeLeftInPhase,omitempty"`

	// adjusted time left in phase in sec
	AdjustedTimeLeftInPhaseInSec int32 `json:"adjustedTimeLeftInPhaseInSec,omitempty"`

	// internal now in epoch ms
	InternalNowInEpochMs int64 `json:"internalNowInEpochMs,omitempty"`

	// is infinite
	IsInfinite bool `json:"isInfinite,omitempty"`

	// phase
	Phase string `json:"phase,omitempty"`

	// time left in phase
	TimeLeftInPhase int64 `json:"timeLeftInPhase,omitempty"`

	// time left in phase in sec
	TimeLeftInPhaseInSec int32 `json:"timeLeftInPhaseInSec,omitempty"`

	// total time in phase
	TotalTimeInPhase int64 `json:"totalTimeInPhase,omitempty"`
}

// Validate validates this lol champ select legacy champ select timer
func (m *LolChampSelectLegacyChampSelectTimer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectLegacyChampSelectTimer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectLegacyChampSelectTimer) UnmarshalBinary(b []byte) error {
	var res LolChampSelectLegacyChampSelectTimer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}