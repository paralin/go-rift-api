// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolGamhsMatchHistoryMetadata lol gamhs match history metadata
// swagger:model LolGamhsMatchHistoryMetadata
type LolGamhsMatchHistoryMetadata struct {

	// data version
	DataVersion int64 `json:"data_version,omitempty"`

	// info type
	InfoType string `json:"info_type,omitempty"`

	// match id
	MatchID string `json:"match_id,omitempty"`

	// participants
	Participants []string `json:"participants"`

	// product
	Product string `json:"product,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this lol gamhs match history metadata
func (m *LolGamhsMatchHistoryMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolGamhsMatchHistoryMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGamhsMatchHistoryMetadata) UnmarshalBinary(b []byte) error {
	var res LolGamhsMatchHistoryMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
