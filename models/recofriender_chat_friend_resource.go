// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RecofrienderChatFriendResource recofriender chat friend resource
// swagger:model RecofrienderChatFriendResource
type RecofrienderChatFriendResource struct {

	// name
	Name string `json:"name,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this recofriender chat friend resource
func (m *RecofrienderChatFriendResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecofrienderChatFriendResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecofrienderChatFriendResource) UnmarshalBinary(b []byte) error {
	var res RecofrienderChatFriendResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
