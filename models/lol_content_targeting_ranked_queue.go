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

// LolContentTargetingRankedQueue lol content targeting ranked queue
// swagger:model LolContentTargetingRankedQueue
type LolContentTargetingRankedQueue string

const (

	// LolContentTargetingRankedQueueNONE captures enum value "NONE"
	LolContentTargetingRankedQueueNONE LolContentTargetingRankedQueue = "NONE"

	// LolContentTargetingRankedQueueRANKEDSOLO5x5 captures enum value "RANKED_SOLO_5x5"
	LolContentTargetingRankedQueueRANKEDSOLO5x5 LolContentTargetingRankedQueue = "RANKED_SOLO_5x5"

	// LolContentTargetingRankedQueueRANKEDFLEXSR captures enum value "RANKED_FLEX_SR"
	LolContentTargetingRankedQueueRANKEDFLEXSR LolContentTargetingRankedQueue = "RANKED_FLEX_SR"

	// LolContentTargetingRankedQueueRANKEDFLEXTT captures enum value "RANKED_FLEX_TT"
	LolContentTargetingRankedQueueRANKEDFLEXTT LolContentTargetingRankedQueue = "RANKED_FLEX_TT"

	// LolContentTargetingRankedQueueRANKEDTFT captures enum value "RANKED_TFT"
	LolContentTargetingRankedQueueRANKEDTFT LolContentTargetingRankedQueue = "RANKED_TFT"
)

// for schema
var lolContentTargetingRankedQueueEnum []interface{}

func init() {
	var res []LolContentTargetingRankedQueue
	if err := json.Unmarshal([]byte(`["NONE","RANKED_SOLO_5x5","RANKED_FLEX_SR","RANKED_FLEX_TT","RANKED_TFT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolContentTargetingRankedQueueEnum = append(lolContentTargetingRankedQueueEnum, v)
	}
}

func (m LolContentTargetingRankedQueue) validateLolContentTargetingRankedQueueEnum(path, location string, value LolContentTargetingRankedQueue) error {
	if err := validate.Enum(path, location, value, lolContentTargetingRankedQueueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol content targeting ranked queue
func (m LolContentTargetingRankedQueue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolContentTargetingRankedQueueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
