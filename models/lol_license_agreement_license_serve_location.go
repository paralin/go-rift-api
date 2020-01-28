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

// LolLicenseAgreementLicenseServeLocation lol license agreement license serve location
// swagger:model LolLicenseAgreementLicenseServeLocation
type LolLicenseAgreementLicenseServeLocation string

const (

	// LolLicenseAgreementLicenseServeLocationPreparing captures enum value "Preparing"
	LolLicenseAgreementLicenseServeLocationPreparing LolLicenseAgreementLicenseServeLocation = "Preparing"

	// LolLicenseAgreementLicenseServeLocationLocal captures enum value "Local"
	LolLicenseAgreementLicenseServeLocationLocal LolLicenseAgreementLicenseServeLocation = "Local"

	// LolLicenseAgreementLicenseServeLocationExternal captures enum value "External"
	LolLicenseAgreementLicenseServeLocationExternal LolLicenseAgreementLicenseServeLocation = "External"
)

// for schema
var lolLicenseAgreementLicenseServeLocationEnum []interface{}

func init() {
	var res []LolLicenseAgreementLicenseServeLocation
	if err := json.Unmarshal([]byte(`["Preparing","Local","External"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLicenseAgreementLicenseServeLocationEnum = append(lolLicenseAgreementLicenseServeLocationEnum, v)
	}
}

func (m LolLicenseAgreementLicenseServeLocation) validateLolLicenseAgreementLicenseServeLocationEnum(path, location string, value LolLicenseAgreementLicenseServeLocation) error {
	if err := validate.Enum(path, location, value, lolLicenseAgreementLicenseServeLocationEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol license agreement license serve location
func (m LolLicenseAgreementLicenseServeLocation) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLicenseAgreementLicenseServeLocationEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}