// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPlayerBehaviorPlayerNotificationResource lol player behavior player notification resource
// swagger:model LolPlayerBehaviorPlayerNotificationResource
type LolPlayerBehaviorPlayerNotificationResource struct {

	// background Url
	BackgroundURL string `json:"backgroundUrl,omitempty"`

	// created
	Created string `json:"created,omitempty"`

	// critical
	Critical bool `json:"critical,omitempty"`

	// data
	Data map[string]string `json:"data,omitempty"`

	// detail key
	DetailKey string `json:"detailKey,omitempty"`

	// expires
	Expires string `json:"expires,omitempty"`

	// icon Url
	IconURL string `json:"iconUrl,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// title key
	TitleKey string `json:"titleKey,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol player behavior player notification resource
func (m *LolPlayerBehaviorPlayerNotificationResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPlayerBehaviorPlayerNotificationResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPlayerBehaviorPlayerNotificationResource) UnmarshalBinary(b []byte) error {
	var res LolPlayerBehaviorPlayerNotificationResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
