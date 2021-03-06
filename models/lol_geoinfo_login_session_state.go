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

// LolGeoinfoLoginSessionState lol geoinfo login session state
// swagger:model LolGeoinfoLoginSessionState
type LolGeoinfoLoginSessionState string

const (

	// LolGeoinfoLoginSessionStateINPROGRESS captures enum value "IN_PROGRESS"
	LolGeoinfoLoginSessionStateINPROGRESS LolGeoinfoLoginSessionState = "IN_PROGRESS"

	// LolGeoinfoLoginSessionStateSUCCEEDED captures enum value "SUCCEEDED"
	LolGeoinfoLoginSessionStateSUCCEEDED LolGeoinfoLoginSessionState = "SUCCEEDED"

	// LolGeoinfoLoginSessionStateLOGGINGOUT captures enum value "LOGGING_OUT"
	LolGeoinfoLoginSessionStateLOGGINGOUT LolGeoinfoLoginSessionState = "LOGGING_OUT"

	// LolGeoinfoLoginSessionStateERROR captures enum value "ERROR"
	LolGeoinfoLoginSessionStateERROR LolGeoinfoLoginSessionState = "ERROR"
)

// for schema
var lolGeoinfoLoginSessionStateEnum []interface{}

func init() {
	var res []LolGeoinfoLoginSessionState
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolGeoinfoLoginSessionStateEnum = append(lolGeoinfoLoginSessionStateEnum, v)
	}
}

func (m LolGeoinfoLoginSessionState) validateLolGeoinfoLoginSessionStateEnum(path, location string, value LolGeoinfoLoginSessionState) error {
	if err := validate.Enum(path, location, value, lolGeoinfoLoginSessionStateEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol geoinfo login session state
func (m LolGeoinfoLoginSessionState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolGeoinfoLoginSessionStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
