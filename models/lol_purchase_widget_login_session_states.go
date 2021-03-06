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

// LolPurchaseWidgetLoginSessionStates lol purchase widget login session states
// swagger:model LolPurchaseWidgetLoginSessionStates
type LolPurchaseWidgetLoginSessionStates string

const (

	// LolPurchaseWidgetLoginSessionStatesINPROGRESS captures enum value "IN_PROGRESS"
	LolPurchaseWidgetLoginSessionStatesINPROGRESS LolPurchaseWidgetLoginSessionStates = "IN_PROGRESS"

	// LolPurchaseWidgetLoginSessionStatesSUCCEEDED captures enum value "SUCCEEDED"
	LolPurchaseWidgetLoginSessionStatesSUCCEEDED LolPurchaseWidgetLoginSessionStates = "SUCCEEDED"

	// LolPurchaseWidgetLoginSessionStatesLOGGINGOUT captures enum value "LOGGING_OUT"
	LolPurchaseWidgetLoginSessionStatesLOGGINGOUT LolPurchaseWidgetLoginSessionStates = "LOGGING_OUT"

	// LolPurchaseWidgetLoginSessionStatesERROR captures enum value "ERROR"
	LolPurchaseWidgetLoginSessionStatesERROR LolPurchaseWidgetLoginSessionStates = "ERROR"
)

// for schema
var lolPurchaseWidgetLoginSessionStatesEnum []interface{}

func init() {
	var res []LolPurchaseWidgetLoginSessionStates
	if err := json.Unmarshal([]byte(`["IN_PROGRESS","SUCCEEDED","LOGGING_OUT","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolPurchaseWidgetLoginSessionStatesEnum = append(lolPurchaseWidgetLoginSessionStatesEnum, v)
	}
}

func (m LolPurchaseWidgetLoginSessionStates) validateLolPurchaseWidgetLoginSessionStatesEnum(path, location string, value LolPurchaseWidgetLoginSessionStates) error {
	if err := validate.Enum(path, location, value, lolPurchaseWidgetLoginSessionStatesEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol purchase widget login session states
func (m LolPurchaseWidgetLoginSessionStates) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolPurchaseWidgetLoginSessionStatesEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
