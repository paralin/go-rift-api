// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolLobbyUserInfoToken lol lobby user info token
// swagger:model LolLobbyUserInfoToken
type LolLobbyUserInfoToken struct {

	// user info
	UserInfo string `json:"userInfo,omitempty"`
}

// Validate validates this lol lobby user info token
func (m *LolLobbyUserInfoToken) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyUserInfoToken) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyUserInfoToken) UnmarshalBinary(b []byte) error {
	var res LolLobbyUserInfoToken
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}