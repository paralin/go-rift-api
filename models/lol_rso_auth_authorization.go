// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolRsoAuthAuthorization lol rso auth authorization
// swagger:model LolRsoAuthAuthorization
type LolRsoAuthAuthorization struct {

	// current account Id
	CurrentAccountID int64 `json:"currentAccountId,omitempty"`

	// current platform Id
	CurrentPlatformID string `json:"currentPlatformId,omitempty"`

	// subject
	Subject string `json:"subject,omitempty"`
}

// Validate validates this lol rso auth authorization
func (m *LolRsoAuthAuthorization) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolRsoAuthAuthorization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRsoAuthAuthorization) UnmarshalBinary(b []byte) error {
	var res LolRsoAuthAuthorization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}