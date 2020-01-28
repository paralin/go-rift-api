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

// LolFeaturedModesQueueGameCategory lol featured modes queue game category
// swagger:model LolFeaturedModesQueueGameCategory
type LolFeaturedModesQueueGameCategory string

const (

	// LolFeaturedModesQueueGameCategoryNone captures enum value "None"
	LolFeaturedModesQueueGameCategoryNone LolFeaturedModesQueueGameCategory = "None"

	// LolFeaturedModesQueueGameCategoryCustom captures enum value "Custom"
	LolFeaturedModesQueueGameCategoryCustom LolFeaturedModesQueueGameCategory = "Custom"

	// LolFeaturedModesQueueGameCategoryPvP captures enum value "PvP"
	LolFeaturedModesQueueGameCategoryPvP LolFeaturedModesQueueGameCategory = "PvP"

	// LolFeaturedModesQueueGameCategoryVersusAi captures enum value "VersusAi"
	LolFeaturedModesQueueGameCategoryVersusAi LolFeaturedModesQueueGameCategory = "VersusAi"

	// LolFeaturedModesQueueGameCategoryAlpha captures enum value "Alpha"
	LolFeaturedModesQueueGameCategoryAlpha LolFeaturedModesQueueGameCategory = "Alpha"
)

// for schema
var lolFeaturedModesQueueGameCategoryEnum []interface{}

func init() {
	var res []LolFeaturedModesQueueGameCategory
	if err := json.Unmarshal([]byte(`["None","Custom","PvP","VersusAi","Alpha"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolFeaturedModesQueueGameCategoryEnum = append(lolFeaturedModesQueueGameCategoryEnum, v)
	}
}

func (m LolFeaturedModesQueueGameCategory) validateLolFeaturedModesQueueGameCategoryEnum(path, location string, value LolFeaturedModesQueueGameCategory) error {
	if err := validate.Enum(path, location, value, lolFeaturedModesQueueGameCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol featured modes queue game category
func (m LolFeaturedModesQueueGameCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolFeaturedModesQueueGameCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}