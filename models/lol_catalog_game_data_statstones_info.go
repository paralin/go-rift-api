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

// LolCatalogGameDataStatstonesInfo lol catalog game data statstones info
// swagger:model LolCatalogGameDataStatstonesInfo
type LolCatalogGameDataStatstonesInfo struct {

	// champ Id to pack ids
	ChampIDToPackIds interface{} `json:"champIdToPackIds,omitempty"`

	// collection Id to stat stone ids
	CollectionIDToStatStoneIds interface{} `json:"collectionIdToStatStoneIds,omitempty"`

	// pack data
	PackData []*LolCatalogGameDataStatstonePack `json:"packData"`

	// pack Id to champ ids
	PackIDToChampIds interface{} `json:"packIdToChampIds,omitempty"`

	// pack Id to stat stones ids
	PackIDToStatStonesIds interface{} `json:"packIdToStatStonesIds,omitempty"`

	// pack Id to sub pack ids
	PackIDToSubPackIds interface{} `json:"packIdToSubPackIds,omitempty"`

	// series Id to stat stone ids
	SeriesIDToStatStoneIds interface{} `json:"seriesIdToStatStoneIds,omitempty"`

	// statstone data
	StatstoneData []*LolCatalogGameDataStatstoneSet `json:"statstoneData"`
}

// Validate validates this lol catalog game data statstones info
func (m *LolCatalogGameDataStatstonesInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePackData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatstoneData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolCatalogGameDataStatstonesInfo) validatePackData(formats strfmt.Registry) error {

	if swag.IsZero(m.PackData) { // not required
		return nil
	}

	for i := 0; i < len(m.PackData); i++ {
		if swag.IsZero(m.PackData[i]) { // not required
			continue
		}

		if m.PackData[i] != nil {
			if err := m.PackData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("packData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolCatalogGameDataStatstonesInfo) validateStatstoneData(formats strfmt.Registry) error {

	if swag.IsZero(m.StatstoneData) { // not required
		return nil
	}

	for i := 0; i < len(m.StatstoneData); i++ {
		if swag.IsZero(m.StatstoneData[i]) { // not required
			continue
		}

		if m.StatstoneData[i] != nil {
			if err := m.StatstoneData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statstoneData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolCatalogGameDataStatstonesInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolCatalogGameDataStatstonesInfo) UnmarshalBinary(b []byte) error {
	var res LolCatalogGameDataStatstonesInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}