// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// LolPerksCustomizationLimits lol perks customization limits
// swagger:model LolPerksCustomizationLimits
type LolPerksCustomizationLimits string

const (

	// LolPerksCustomizationLimitsLocked captures enum value "Locked"
	LolPerksCustomizationLimitsLocked LolPerksCustomizationLimits = "Locked"

	// LolPerksCustomizationLimitsCanSelectPages captures enum value "CanSelectPages"
	LolPerksCustomizationLimitsCanSelectPages LolPerksCustomizationLimits = "CanSelectPages"

	// LolPerksCustomizationLimitsCanSelectKeystones captures enum value "CanSelectKeystones"
	LolPerksCustomizationLimitsCanSelectKeystones LolPerksCustomizationLimits = "CanSelectKeystones"

	// LolPerksCustomizationLimitsCanSelectPrimaries captures enum value "CanSelectPrimaries"
	LolPerksCustomizationLimitsCanSelectPrimaries LolPerksCustomizationLimits = "CanSelectPrimaries"

	// LolPerksCustomizationLimitsCanSelectSplash captures enum value "CanSelectSplash"
	LolPerksCustomizationLimitsCanSelectSplash LolPerksCustomizationLimits = "CanSelectSplash"

	// LolPerksCustomizationLimitsCanUseAdvancedStyles captures enum value "CanUseAdvancedStyles"
	LolPerksCustomizationLimitsCanUseAdvancedStyles LolPerksCustomizationLimits = "CanUseAdvancedStyles"
)

// for schema
var lolPerksCustomizationLimitsEnum []interface{}

func init() {
	var res []LolPerksCustomizationLimits
	if err := json.Unmarshal([]byte(`["Locked","CanSelectPages","CanSelectKeystones","CanSelectPrimaries","CanSelectSplash","CanUseAdvancedStyles"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolPerksCustomizationLimitsEnum = append(lolPerksCustomizationLimitsEnum, v)
	}
}

func (m LolPerksCustomizationLimits) validateLolPerksCustomizationLimitsEnum(path, location string, value LolPerksCustomizationLimits) error {
	if err := validate.Enum(path, location, value, lolPerksCustomizationLimitsEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol perks customization limits
func (m LolPerksCustomizationLimits) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolPerksCustomizationLimitsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}