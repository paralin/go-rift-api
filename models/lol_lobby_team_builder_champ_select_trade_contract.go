// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolLobbyTeamBuilderChampSelectTradeContract lol lobby team builder champ select trade contract
// swagger:model LolLobbyTeamBuilderChampSelectTradeContract
type LolLobbyTeamBuilderChampSelectTradeContract struct {

	// cell Id
	CellID int64 `json:"cellId,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// state
	State LolLobbyTeamBuilderChampSelectTradeState `json:"state,omitempty"`
}

// Validate validates this lol lobby team builder champ select trade contract
func (m *LolLobbyTeamBuilderChampSelectTradeContract) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLobbyTeamBuilderChampSelectTradeContract) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.State.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("state")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLobbyTeamBuilderChampSelectTradeContract) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLobbyTeamBuilderChampSelectTradeContract) UnmarshalBinary(b []byte) error {
	var res LolLobbyTeamBuilderChampSelectTradeContract
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
