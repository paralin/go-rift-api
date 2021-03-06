// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolClubsPublicSessionResource lol clubs public session resource
// swagger:model LolClubsPublicSessionResource
type LolClubsPublicSessionResource struct {

	// session expire
	SessionExpire int32 `json:"sessionExpire,omitempty"`

	// session state
	SessionState string `json:"sessionState,omitempty"`
}

// Validate validates this lol clubs public session resource
func (m *LolClubsPublicSessionResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolClubsPublicSessionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolClubsPublicSessionResource) UnmarshalBinary(b []byte) error {
	var res LolClubsPublicSessionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
