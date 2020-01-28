// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// LolFeaturedModesQueue lol featured modes queue
// swagger:model LolFeaturedModesQueue
type LolFeaturedModesQueue struct {

	// category
	Category LolFeaturedModesQueueGameCategory `json:"category,omitempty"`

	// game mode
	GameMode string `json:"gameMode,omitempty"`

	// id
	ID int32 `json:"id,omitempty"`

	// last toggled off time
	LastToggledOffTime int64 `json:"lastToggledOffTime,omitempty"`

	// last toggled on time
	LastToggledOnTime int64 `json:"lastToggledOnTime,omitempty"`

	// map Id
	MapID int32 `json:"mapId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// queue availability
	QueueAvailability LolFeaturedModesQueueAvailability `json:"queueAvailability,omitempty"`
}

// Validate validates this lol featured modes queue
func (m *LolFeaturedModesQueue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueAvailability(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolFeaturedModesQueue) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	if err := m.Category.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("category")
		}
		return err
	}

	return nil
}

func (m *LolFeaturedModesQueue) validateQueueAvailability(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueAvailability) { // not required
		return nil
	}

	if err := m.QueueAvailability.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("queueAvailability")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolFeaturedModesQueue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolFeaturedModesQueue) UnmarshalBinary(b []byte) error {
	var res LolFeaturedModesQueue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
