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

// LolLicenseAgreementLicenseAgreementType lol license agreement license agreement type
// swagger:model LolLicenseAgreementLicenseAgreementType
type LolLicenseAgreementLicenseAgreementType string

const (

	// LolLicenseAgreementLicenseAgreementTypeEula captures enum value "Eula"
	LolLicenseAgreementLicenseAgreementTypeEula LolLicenseAgreementLicenseAgreementType = "Eula"

	// LolLicenseAgreementLicenseAgreementTypeTermsOfUse captures enum value "TermsOfUse"
	LolLicenseAgreementLicenseAgreementTypeTermsOfUse LolLicenseAgreementLicenseAgreementType = "TermsOfUse"
)

// for schema
var lolLicenseAgreementLicenseAgreementTypeEnum []interface{}

func init() {
	var res []LolLicenseAgreementLicenseAgreementType
	if err := json.Unmarshal([]byte(`["Eula","TermsOfUse"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolLicenseAgreementLicenseAgreementTypeEnum = append(lolLicenseAgreementLicenseAgreementTypeEnum, v)
	}
}

func (m LolLicenseAgreementLicenseAgreementType) validateLolLicenseAgreementLicenseAgreementTypeEnum(path, location string, value LolLicenseAgreementLicenseAgreementType) error {
	if err := validate.Enum(path, location, value, lolLicenseAgreementLicenseAgreementTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol license agreement license agreement type
func (m LolLicenseAgreementLicenseAgreementType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolLicenseAgreementLicenseAgreementTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}