// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolKrShutdownLawRatingScreenInfo lol kr shutdown law rating screen info
// swagger:model LolKrShutdownLawRatingScreenInfo
type LolKrShutdownLawRatingScreenInfo struct {

	// shown
	Shown bool `json:"shown,omitempty"`
}

// Validate validates this lol kr shutdown law rating screen info
func (m *LolKrShutdownLawRatingScreenInfo) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolKrShutdownLawRatingScreenInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolKrShutdownLawRatingScreenInfo) UnmarshalBinary(b []byte) error {
	var res LolKrShutdownLawRatingScreenInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
