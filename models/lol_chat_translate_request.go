// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChatTranslateRequest lol chat translate request
// swagger:model LolChatTranslateRequest
type LolChatTranslateRequest struct {

	// blocking
	Blocking bool `json:"blocking,omitempty"`

	// keys
	Keys []string `json:"keys"`

	// locale
	Locale string `json:"locale,omitempty"`

	// patchline
	Patchline string `json:"patchline,omitempty"`

	// product id
	ProductID string `json:"product_id,omitempty"`
}

// Validate validates this lol chat translate request
func (m *LolChatTranslateRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChatTranslateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChatTranslateRequest) UnmarshalBinary(b []byte) error {
	var res LolChatTranslateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
