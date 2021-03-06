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

// LolItemSetsItemSets lol item sets item sets
// swagger:model LolItemSetsItemSets
type LolItemSetsItemSets struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// item sets
	ItemSets []*LolItemSetsItemSet `json:"itemSets"`

	// timestamp
	Timestamp int64 `json:"timestamp,omitempty"`
}

// Validate validates this lol item sets item sets
func (m *LolItemSetsItemSets) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItemSets(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolItemSetsItemSets) validateItemSets(formats strfmt.Registry) error {

	if swag.IsZero(m.ItemSets) { // not required
		return nil
	}

	for i := 0; i < len(m.ItemSets); i++ {
		if swag.IsZero(m.ItemSets[i]) { // not required
			continue
		}

		if m.ItemSets[i] != nil {
			if err := m.ItemSets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("itemSets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolItemSetsItemSets) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolItemSetsItemSets) UnmarshalBinary(b []byte) error {
	var res LolItemSetsItemSets
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
