// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolChampionsGameDataChampion lol champions game data champion
// swagger:model LolChampionsGameDataChampion
type LolChampionsGameDataChampion struct {

	// alias
	Alias string `json:"alias,omitempty"`

	// ban vo path
	BanVoPath string `json:"banVoPath,omitempty"`

	// choose vo path
	ChooseVoPath string `json:"chooseVoPath,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// passive
	Passive *LolChampionsGameDataChampionSpell `json:"passive,omitempty"`

	// roles
	Roles []string `json:"roles"`

	// skins
	Skins []*LolChampionsGameDataChampionSkin `json:"skins"`

	// spells
	Spells []*LolChampionsGameDataChampionSpell `json:"spells"`

	// square portrait path
	SquarePortraitPath string `json:"squarePortraitPath,omitempty"`

	// stinger sfx path
	StingerSfxPath string `json:"stingerSfxPath,omitempty"`

	// tactical info
	TacticalInfo *LolChampionsGameDataChampionTacticalInfo `json:"tacticalInfo,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this lol champions game data champion
func (m *LolChampionsGameDataChampion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpells(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTacticalInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChampionsGameDataChampion) validatePassive(formats strfmt.Registry) error {

	if swag.IsZero(m.Passive) { // not required
		return nil
	}

	if m.Passive != nil {
		if err := m.Passive.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("passive")
			}
			return err
		}
	}

	return nil
}

func (m *LolChampionsGameDataChampion) validateSkins(formats strfmt.Registry) error {

	if swag.IsZero(m.Skins) { // not required
		return nil
	}

	for i := 0; i < len(m.Skins); i++ {
		if swag.IsZero(m.Skins[i]) { // not required
			continue
		}

		if m.Skins[i] != nil {
			if err := m.Skins[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skins" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolChampionsGameDataChampion) validateSpells(formats strfmt.Registry) error {

	if swag.IsZero(m.Spells) { // not required
		return nil
	}

	for i := 0; i < len(m.Spells); i++ {
		if swag.IsZero(m.Spells[i]) { // not required
			continue
		}

		if m.Spells[i] != nil {
			if err := m.Spells[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("spells" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolChampionsGameDataChampion) validateTacticalInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.TacticalInfo) { // not required
		return nil
	}

	if m.TacticalInfo != nil {
		if err := m.TacticalInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tacticalInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChampionsGameDataChampion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampionsGameDataChampion) UnmarshalBinary(b []byte) error {
	var res LolChampionsGameDataChampion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
