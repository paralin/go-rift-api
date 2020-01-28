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

// LolLobbyTeamBuilderQueueGameCategory lol lobby team builder queue game category
// swagger:model LolLobbyTeamBuilderQueueGameCategory
type LolLobbyTeamBuilderQueueGameCategory string

const (

	// LolLobbyTeamBuilderQueueGameCategoryNone captures enum value "None"
	LolLobbyTeamBuilderQueueGameCategoryNone LolLobbyTeamBuilderQueueGameCategory = "None"

	// LolLobbyTeamBuilderQueueGameCategoryCustom captures enum value "Custom"
	LolLobbyTeamBuilderQueueGameCategoryCustom LolLobbyTeamBuilderQueueGameCategory = "Custom"

	// LolLobbyTeamBuilderQueueGameCategoryPvP captures enum value "PvP"
	LolLobbyTeamBuilderQueueGameCategoryPvP LolLobbyTeamBuilderQueueGameCategory = "PvP"

	// LolLobbyTeamBuilderQueueGameCategoryVersusAi captures enum value "VersusAi"
	LolLobbyTeamBuilderQueueGameCategoryVersusAi LolLobbyTeamBuilderQueueGameCategory = "VersusAi"

	// LolLobbyTeamBuilderQueueGameCategoryAlpha captures enum value "Alpha"
	LolLobbyTeamBuilderQueueGameCategoryAlpha LolLobbyTeamBuilderQueueGameCategory = "Alpha"
)

// for schema
var lolLobbyTeamBuilderQueueGameCategoryEnum []interface{}

func init() {
	var res []LolLobbyTeamBuilderQueueGameCategory
	if err := json.Unmarshal([]byte(`["None","Custom","PvP","VersusAi","Alpha"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLobbyTeamBuilderQueueGameCategoryEnum = append(lolLobbyTeamBuilderQueueGameCategoryEnum, v)
	}
}

func (m LolLobbyTeamBuilderQueueGameCategory) validateLolLobbyTeamBuilderQueueGameCategoryEnum(path, location string, value LolLobbyTeamBuilderQueueGameCategory) error {
	if err := validate.Enum(path, location, value, lolLobbyTeamBuilderQueueGameCategoryEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol lobby team builder queue game category
func (m LolLobbyTeamBuilderQueueGameCategory) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLobbyTeamBuilderQueueGameCategoryEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}