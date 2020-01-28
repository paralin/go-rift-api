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

// LolLeaverBusterLeaverBusterNotificationType lol leaver buster leaver buster notification type
// swagger:model LolLeaverBusterLeaverBusterNotificationType
type LolLeaverBusterLeaverBusterNotificationType string

const (

	// LolLeaverBusterLeaverBusterNotificationTypeInvalid captures enum value "Invalid"
	LolLeaverBusterLeaverBusterNotificationTypeInvalid LolLeaverBusterLeaverBusterNotificationType = "Invalid"

	// LolLeaverBusterLeaverBusterNotificationTypeTaintedWarning captures enum value "TaintedWarning"
	LolLeaverBusterLeaverBusterNotificationTypeTaintedWarning LolLeaverBusterLeaverBusterNotificationType = "TaintedWarning"

	// LolLeaverBusterLeaverBusterNotificationTypePunishmentIncurred captures enum value "PunishmentIncurred"
	LolLeaverBusterLeaverBusterNotificationTypePunishmentIncurred LolLeaverBusterLeaverBusterNotificationType = "PunishmentIncurred"

	// LolLeaverBusterLeaverBusterNotificationTypePunishedGamesRemaining captures enum value "PunishedGamesRemaining"
	LolLeaverBusterLeaverBusterNotificationTypePunishedGamesRemaining LolLeaverBusterLeaverBusterNotificationType = "PunishedGamesRemaining"

	// LolLeaverBusterLeaverBusterNotificationTypeReforming captures enum value "Reforming"
	LolLeaverBusterLeaverBusterNotificationTypeReforming LolLeaverBusterLeaverBusterNotificationType = "Reforming"
)

// for schema
var lolLeaverBusterLeaverBusterNotificationTypeEnum []interface{}

func init() {
	var res []LolLeaverBusterLeaverBusterNotificationType
	if err := json.Unmarshal([]byte(`["Invalid","TaintedWarning","PunishmentIncurred","PunishedGamesRemaining","Reforming"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLeaverBusterLeaverBusterNotificationTypeEnum = append(lolLeaverBusterLeaverBusterNotificationTypeEnum, v)
	}
}

func (m LolLeaverBusterLeaverBusterNotificationType) validateLolLeaverBusterLeaverBusterNotificationTypeEnum(path, location string, value LolLeaverBusterLeaverBusterNotificationType) error {
	if err := validate.Enum(path, location, value, lolLeaverBusterLeaverBusterNotificationTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol leaver buster leaver buster notification type
func (m LolLeaverBusterLeaverBusterNotificationType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLeaverBusterLeaverBusterNotificationTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}