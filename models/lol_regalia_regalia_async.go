// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRegaliaRegaliaAsync lol regalia regalia async
// swagger:model LolRegaliaRegaliaAsync
type LolRegaliaRegaliaAsync struct {

	// md5
	Md5 string `json:"md5,omitempty"`
}

// Validate validates this lol regalia regalia async
func (m *LolRegaliaRegaliaAsync) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRegaliaRegaliaAsync) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRegaliaRegaliaAsync) UnmarshalBinary(b []byte) error {
	var res LolRegaliaRegaliaAsync
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
