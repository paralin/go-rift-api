// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEsportStreamNotificationsESportsAPIStreamgroups lol esport stream notifications e sports API streamgroups
// swagger:model LolEsportStreamNotificationsESportsAPI_streamgroups
type LolEsportStreamNotificationsESportsAPIStreamgroups struct {

	// content
	Content string `json:"content,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// live
	Live bool `json:"live,omitempty"`

	// slug
	Slug string `json:"slug,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this lol esport stream notifications e sports API streamgroups
func (m *LolEsportStreamNotificationsESportsAPIStreamgroups) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEsportStreamNotificationsESportsAPIStreamgroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEsportStreamNotificationsESportsAPIStreamgroups) UnmarshalBinary(b []byte) error {
	var res LolEsportStreamNotificationsESportsAPIStreamgroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}