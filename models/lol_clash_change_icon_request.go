// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClashChangeIconRequest lol clash change icon request
// swagger:model LolClashChangeIconRequest
type LolClashChangeIconRequest struct {

	// icon color Id
	IconColorID int32 `json:"iconColorId,omitempty"`

	// icon Id
	IconID int32 `json:"iconId,omitempty"`
}

// Validate validates this lol clash change icon request
func (m *LolClashChangeIconRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClashChangeIconRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClashChangeIconRequest) UnmarshalBinary(b []byte) error {
	var res LolClashChangeIconRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
