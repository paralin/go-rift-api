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

// LolGameflowQueueGameCategory lol gameflow queue game category
// swagger:model LolGameflowQueueGameCategory
type LolGameflowQueueGameCategory string

const (

	// LolGameflowQueueGameCategoryNone captures enum value "None"
	LolGameflowQueueGameCategoryNone LolGameflowQueueGameCategory = "None"

	// LolGameflowQueueGameCategoryCustom captures enum value "Custom"
	LolGameflowQueueGameCategoryCustom LolGameflowQueueGameCategory = "Custom"

	// LolGameflowQueueGameCategoryPvP captures enum value "PvP"
	LolGameflowQueueGameCategoryPvP LolGameflowQueueGameCategory = "PvP"

	// LolGameflowQueueGameCategoryVersusAi captures enum value "VersusAi"
	LolGameflowQueueGameCategoryVersusAi LolGameflowQueueGameCategory = "VersusAi"

	// LolGameflowQueueGameCategoryAlpha captures enum value "Alpha"
	LolGameflowQueueGameCategoryAlpha LolGameflowQueueGameCategory = "Alpha"
)

// for schema
var lolGameflowQueueGameCategoryEnum []interface{}

func init() {
	var res []LolGameflowQueueGameCategory
	if err := json.Unmarshal([]byte(`["None","Custom","PvP","VersusAi","Alpha"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolGameflowQueueGameCategoryEnum = append(lolGameflowQueueGameCategoryEnum, v)
	}
}

func (m LolGameflowQueueGameCategory) validateLolGameflowQueueGameCategoryEnum(path, location string, value LolGameflowQueueGameCategory) error {
	if err := validate.Enum(path, location, value, lolGameflowQueueGameCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol gameflow queue game category
func (m LolGameflowQueueGameCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolGameflowQueueGameCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}