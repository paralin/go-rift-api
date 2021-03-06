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

// LolClashTournamentNotifyReason lol clash tournament notify reason
// swagger:model LolClashTournamentNotifyReason
type LolClashTournamentNotifyReason string

const (

	// LolClashTournamentNotifyReasonNEWTOURNAMENT captures enum value "NEW_TOURNAMENT"
	LolClashTournamentNotifyReasonNEWTOURNAMENT LolClashTournamentNotifyReason = "NEW_TOURNAMENT"

	// LolClashTournamentNotifyReasonUPDATETOURNAMENT captures enum value "UPDATE_TOURNAMENT"
	LolClashTournamentNotifyReasonUPDATETOURNAMENT LolClashTournamentNotifyReason = "UPDATE_TOURNAMENT"

	// LolClashTournamentNotifyReasonCANCELTOURNAMENT captures enum value "CANCEL_TOURNAMENT"
	LolClashTournamentNotifyReasonCANCELTOURNAMENT LolClashTournamentNotifyReason = "CANCEL_TOURNAMENT"

	// LolClashTournamentNotifyReasonCANCELPERIOD captures enum value "CANCEL_PERIOD"
	LolClashTournamentNotifyReasonCANCELPERIOD LolClashTournamentNotifyReason = "CANCEL_PERIOD"

	// LolClashTournamentNotifyReasonADDPHASE captures enum value "ADD_PHASE"
	LolClashTournamentNotifyReasonADDPHASE LolClashTournamentNotifyReason = "ADD_PHASE"

	// LolClashTournamentNotifyReasonUPDATEPHASE captures enum value "UPDATE_PHASE"
	LolClashTournamentNotifyReasonUPDATEPHASE LolClashTournamentNotifyReason = "UPDATE_PHASE"

	// LolClashTournamentNotifyReasonREVERTPHASE captures enum value "REVERT_PHASE"
	LolClashTournamentNotifyReasonREVERTPHASE LolClashTournamentNotifyReason = "REVERT_PHASE"

	// LolClashTournamentNotifyReasonUPDATESTATUS captures enum value "UPDATE_STATUS"
	LolClashTournamentNotifyReasonUPDATESTATUS LolClashTournamentNotifyReason = "UPDATE_STATUS"
)

// for schema
var lolClashTournamentNotifyReasonEnum []interface{}

func init() {
	var res []LolClashTournamentNotifyReason
	if err := json.Unmarshal([]byte(`["NEW_TOURNAMENT","UPDATE_TOURNAMENT","CANCEL_TOURNAMENT","CANCEL_PERIOD","ADD_PHASE","UPDATE_PHASE","REVERT_PHASE","UPDATE_STATUS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolClashTournamentNotifyReasonEnum = append(lolClashTournamentNotifyReasonEnum, v)
	}
}

func (m LolClashTournamentNotifyReason) validateLolClashTournamentNotifyReasonEnum(path, location string, value LolClashTournamentNotifyReason) error {
	if err := validate.Enum(path, location, value, lolClashTournamentNotifyReasonEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol clash tournament notify reason
func (m LolClashTournamentNotifyReason) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolClashTournamentNotifyReasonEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
