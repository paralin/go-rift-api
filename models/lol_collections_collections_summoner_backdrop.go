// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolCollectionsCollectionsSummonerBackdrop lol collections collections summoner backdrop
// swagger:model LolCollectionsCollectionsSummonerBackdrop
type LolCollectionsCollectionsSummonerBackdrop struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// backdrop image
	BackdropImage string `json:"backdropImage,omitempty"`

	// backdrop mask color
	BackdropMaskColor string `json:"backdropMaskColor,omitempty"`

	// backdrop type
	BackdropType LolCollectionsCollectionsSummonerBackdropType `json:"backdropType,omitempty"`

	// backdrop video
	BackdropVideo string `json:"backdropVideo,omitempty"`

	// champion Id
	ChampionID int32 `json:"championId,omitempty"`

	// profile icon Id
	ProfileIconID int32 `json:"profileIconId,omitempty"`

	// puuid
	Puuid string `json:"puuid,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this lol collections collections summoner backdrop
func (m *LolCollectionsCollectionsSummonerBackdrop) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBackdropType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCollectionsCollectionsSummonerBackdrop) validateBackdropType(formats strfmt.Registry) error {

	if swag.IsZero(m.BackdropType) { // not required
		return nil
	}

	if err := m.BackdropType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("backdropType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCollectionsCollectionsSummonerBackdrop) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCollectionsCollectionsSummonerBackdrop) UnmarshalBinary(b []byte) error {
	var res LolCollectionsCollectionsSummonerBackdrop
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
