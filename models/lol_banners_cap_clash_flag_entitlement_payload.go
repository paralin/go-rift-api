// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolBannersCapClashFlagEntitlementPayload lol banners cap clash flag entitlement payload
// swagger:model LolBannersCapClashFlagEntitlementPayload
type LolBannersCapClashFlagEntitlementPayload struct {

	// reward spec
	RewardSpec *LolBannersClashV2FlagRewardSpec `json:"rewardSpec,omitempty"`

	// reward type
	RewardType string `json:"rewardType,omitempty"`
}

// Validate validates this lol banners cap clash flag entitlement payload
func (m *LolBannersCapClashFlagEntitlementPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRewardSpec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolBannersCapClashFlagEntitlementPayload) validateRewardSpec(formats strfmt.Registry) error {

	if swag.IsZero(m.RewardSpec) { // not required
		return nil
	}

	if m.RewardSpec != nil {
		if err := m.RewardSpec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewardSpec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolBannersCapClashFlagEntitlementPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolBannersCapClashFlagEntitlementPayload) UnmarshalBinary(b []byte) error {
	var res LolBannersCapClashFlagEntitlementPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}