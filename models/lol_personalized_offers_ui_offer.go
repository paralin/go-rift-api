// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPersonalizedOffersUIOffer lol personalized offers UI offer
// swagger:model LolPersonalizedOffersUIOffer
type LolPersonalizedOffersUIOffer struct {

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// discount price
	DiscountPrice int64 `json:"discountPrice,omitempty"`

	// expiration date
	ExpirationDate string `json:"expirationDate,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// original price
	OriginalPrice int64 `json:"originalPrice,omitempty"`

	// owned
	Owned bool `json:"owned,omitempty"`

	// revealed
	Revealed bool `json:"revealed,omitempty"`

	// skin Id
	SkinID int32 `json:"skinId,omitempty"`

	// skin name
	SkinName string `json:"skinName,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol personalized offers UI offer
func (m *LolPersonalizedOffersUIOffer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPersonalizedOffersUIOffer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPersonalizedOffersUIOffer) UnmarshalBinary(b []byte) error {
	var res LolPersonalizedOffersUIOffer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
