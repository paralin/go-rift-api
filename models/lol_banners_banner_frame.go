// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolBannersBannerFrame lol banners banner frame
// swagger:model LolBannersBannerFrame
type LolBannersBannerFrame struct {

	// level
	Level int64 `json:"level,omitempty"`
}

// Validate validates this lol banners banner frame
func (m *LolBannersBannerFrame) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolBannersBannerFrame) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolBannersBannerFrame) UnmarshalBinary(b []byte) error {
	var res LolBannersBannerFrame
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
