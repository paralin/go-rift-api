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

// LolRsoAuthAuthHintType lol rso auth auth hint type
// swagger:model LolRsoAuthAuthHintType
type LolRsoAuthAuthHintType string

const (

	// LolRsoAuthAuthHintTypeEmailVerification captures enum value "email_verification"
	LolRsoAuthAuthHintTypeEmailVerification LolRsoAuthAuthHintType = "email_verification"

	// LolRsoAuthAuthHintTypePasswordReset captures enum value "password_reset"
	LolRsoAuthAuthHintTypePasswordReset LolRsoAuthAuthHintType = "password_reset"

	// LolRsoAuthAuthHintTypeParentalConsent captures enum value "parental_consent"
	LolRsoAuthAuthHintTypeParentalConsent LolRsoAuthAuthHintType = "parental_consent"

	// LolRsoAuthAuthHintTypeAmbiguousUsername captures enum value "ambiguous_username"
	LolRsoAuthAuthHintTypeAmbiguousUsername LolRsoAuthAuthHintType = "ambiguous_username"
)

// for schema
var lolRsoAuthAuthHintTypeEnum []interface{}

func init() {
	var res []LolRsoAuthAuthHintType
	if err := json.Unmarshal([]byte(`["email_verification","password_reset","parental_consent","ambiguous_username"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolRsoAuthAuthHintTypeEnum = append(lolRsoAuthAuthHintTypeEnum, v)
	}
}

func (m LolRsoAuthAuthHintType) validateLolRsoAuthAuthHintTypeEnum(path, location string, value LolRsoAuthAuthHintType) error {
	if err := validate.Enum(path, location, value, lolRsoAuthAuthHintTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol rso auth auth hint type
func (m LolRsoAuthAuthHintType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolRsoAuthAuthHintTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
