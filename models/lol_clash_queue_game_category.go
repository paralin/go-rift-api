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

// LolClashQueueGameCategory lol clash queue game category
// swagger:model LolClashQueueGameCategory
type LolClashQueueGameCategory string

const (

	// LolClashQueueGameCategoryNone captures enum value "None"
	LolClashQueueGameCategoryNone LolClashQueueGameCategory = "None"

	// LolClashQueueGameCategoryCustom captures enum value "Custom"
	LolClashQueueGameCategoryCustom LolClashQueueGameCategory = "Custom"

	// LolClashQueueGameCategoryPvP captures enum value "PvP"
	LolClashQueueGameCategoryPvP LolClashQueueGameCategory = "PvP"

	// LolClashQueueGameCategoryVersusAi captures enum value "VersusAi"
	LolClashQueueGameCategoryVersusAi LolClashQueueGameCategory = "VersusAi"

	// LolClashQueueGameCategoryAlpha captures enum value "Alpha"
	LolClashQueueGameCategoryAlpha LolClashQueueGameCategory = "Alpha"
)

// for schema
var lolClashQueueGameCategoryEnum []interface{}

func init() {
	var res []LolClashQueueGameCategory
	if err := json.Unmarshal([]byte(`["None","Custom","PvP","VersusAi","Alpha"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolClashQueueGameCategoryEnum = append(lolClashQueueGameCategoryEnum, v)
	}
}

func (m LolClashQueueGameCategory) validateLolClashQueueGameCategoryEnum(path, location string, value LolClashQueueGameCategory) error {
	if err := validate.Enum(path, location, value, lolClashQueueGameCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol clash queue game category
func (m LolClashQueueGameCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolClashQueueGameCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
