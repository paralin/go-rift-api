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

// LolChampionsGameDataChampionSkin lol champions game data champion skin
// swagger:model LolChampionsGameDataChampionSkin
type LolChampionsGameDataChampionSkin struct {

	// chroma path
	ChromaPath string `json:"chromaPath,omitempty"`

	// chromas
	Chromas []*LolChampionsGameDataChampionChroma `json:"chromas"`

	// emblems
	Emblems []*LolChampionsCollectionsChampionSkinEmblem `json:"emblems"`

	// features text
	FeaturesText string `json:"featuresText,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// load screen path
	LoadScreenPath string `json:"loadScreenPath,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// rarity gem path
	RarityGemPath string `json:"rarityGemPath,omitempty"`

	// skin type
	SkinType string `json:"skinType,omitempty"`

	// splash path
	SplashPath string `json:"splashPath,omitempty"`

	// splash video path
	SplashVideoPath string `json:"splashVideoPath,omitempty"`

	// tile path
	TilePath string `json:"tilePath,omitempty"`

	// uncentered splash path
	UncenteredSplashPath string `json:"uncenteredSplashPath,omitempty"`
}

// Validate validates this lol champions game data champion skin
func (m *LolChampionsGameDataChampionSkin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChromas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmblems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolChampionsGameDataChampionSkin) validateChromas(formats strfmt.Registry) error {

	if swag.IsZero(m.Chromas) { // not required
		return nil
	}

	for i := 0; i < len(m.Chromas); i++ {
		if swag.IsZero(m.Chromas[i]) { // not required
			continue
		}

		if m.Chromas[i] != nil {
			if err := m.Chromas[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("chromas" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolChampionsGameDataChampionSkin) validateEmblems(formats strfmt.Registry) error {

	if swag.IsZero(m.Emblems) { // not required
		return nil
	}

	for i := 0; i < len(m.Emblems); i++ {
		if swag.IsZero(m.Emblems[i]) { // not required
			continue
		}

		if m.Emblems[i] != nil {
			if err := m.Emblems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("emblems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolChampionsGameDataChampionSkin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampionsGameDataChampionSkin) UnmarshalBinary(b []byte) error {
	var res LolChampionsGameDataChampionSkin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
