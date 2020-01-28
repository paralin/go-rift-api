// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ChampSelectLcdsTradeContractDTO champ select lcds trade contract d t o
// swagger:model ChampSelectLcdsTradeContractDTO
type ChampSelectLcdsTradeContractDTO struct {

	// requester champion Id
	RequesterChampionID int32 `json:"requesterChampionId,omitempty"`

	// requester internal summoner name
	RequesterInternalSummonerName string `json:"requesterInternalSummonerName,omitempty"`

	// responder champion Id
	ResponderChampionID int32 `json:"responderChampionId,omitempty"`

	// responder internal summoner name
	ResponderInternalSummonerName string `json:"responderInternalSummonerName,omitempty"`

	// state
	State string `json:"state,omitempty"`
}

// Validate validates this champ select lcds trade contract d t o
func (m *ChampSelectLcdsTradeContractDTO) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChampSelectLcdsTradeContractDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChampSelectLcdsTradeContractDTO) UnmarshalBinary(b []byte) error {
	var res ChampSelectLcdsTradeContractDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
