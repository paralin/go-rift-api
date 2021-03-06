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

// LolRegaliaRegaliaBannerType lol regalia regalia banner type
// swagger:model LolRegaliaRegaliaBannerType
type LolRegaliaRegaliaBannerType string

const (

	// LolRegaliaRegaliaBannerTypeBlank captures enum value "blank"
	LolRegaliaRegaliaBannerTypeBlank LolRegaliaRegaliaBannerType = "blank"

	// LolRegaliaRegaliaBannerTypeLastSeasonHighestRank captures enum value "lastSeasonHighestRank"
	LolRegaliaRegaliaBannerTypeLastSeasonHighestRank LolRegaliaRegaliaBannerType = "lastSeasonHighestRank"
)

// for schema
var lolRegaliaRegaliaBannerTypeEnum []interface{}

func init() {
	var res []LolRegaliaRegaliaBannerType
	if err := json.Unmarshal([]byte(`["blank","lastSeasonHighestRank"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lolRegaliaRegaliaBannerTypeEnum = append(lolRegaliaRegaliaBannerTypeEnum, v)
	}
}

func (m LolRegaliaRegaliaBannerType) validateLolRegaliaRegaliaBannerTypeEnum(path, location string, value LolRegaliaRegaliaBannerType) error {
	if err := validate.Enum(path, location, value, lolRegaliaRegaliaBannerTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this lol regalia regalia banner type
func (m LolRegaliaRegaliaBannerType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateLolRegaliaRegaliaBannerTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
