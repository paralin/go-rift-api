// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPremadeVoiceSummoner lol premade voice summoner
// swagger:model LolPremadeVoiceSummoner
type LolPremadeVoiceSummoner struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`
}

// Validate validates this lol premade voice summoner
func (m *LolPremadeVoiceSummoner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPremadeVoiceSummoner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPremadeVoiceSummoner) UnmarshalBinary(b []byte) error {
	var res LolPremadeVoiceSummoner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}