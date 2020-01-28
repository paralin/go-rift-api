// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolStoreGiftingConfig lol store gifting config
// swagger:model LolStoreGiftingConfig
type LolStoreGiftingConfig struct {

	// gifting hextec max daily gifts receive
	GiftingHextecMaxDailyGiftsReceive int32 `json:"giftingHextecMaxDailyGiftsReceive,omitempty"`

	// gifting hextech max daily gifts send
	GiftingHextechMaxDailyGiftsSend int32 `json:"giftingHextechMaxDailyGiftsSend,omitempty"`

	// gifting item max daily gifts receive
	GiftingItemMaxDailyGiftsReceive int32 `json:"giftingItemMaxDailyGiftsReceive,omitempty"`

	// gifting item max daily gifts send
	GiftingItemMaxDailyGiftsSend int32 `json:"giftingItemMaxDailyGiftsSend,omitempty"`

	// gifting item min level send
	GiftingItemMinLevelSend int32 `json:"giftingItemMinLevelSend,omitempty"`

	// gifting restriction flag rioter
	GiftingRestrictionFlagRioter int32 `json:"giftingRestrictionFlagRioter,omitempty"`

	// gifting rp max daily gifts receive
	GiftingRpMaxDailyGiftsReceive int32 `json:"giftingRpMaxDailyGiftsReceive,omitempty"`

	// gifting rp max daily gifts send
	GiftingRpMaxDailyGiftsSend int32 `json:"giftingRpMaxDailyGiftsSend,omitempty"`

	// gifting rp min level send
	GiftingRpMinLevelSend int32 `json:"giftingRpMinLevelSend,omitempty"`

	// recipient level limit item
	RecipientLevelLimitItem int32 `json:"recipientLevelLimitItem,omitempty"`

	// recipient level limit rp
	RecipientLevelLimitRp int32 `json:"recipientLevelLimitRp,omitempty"`

	// requires identity verification
	RequiresIdentityVerification bool `json:"requiresIdentityVerification,omitempty"`
}

// Validate validates this lol store gifting config
func (m *LolStoreGiftingConfig) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolStoreGiftingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolStoreGiftingConfig) UnmarshalBinary(b []byte) error {
	var res LolStoreGiftingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}