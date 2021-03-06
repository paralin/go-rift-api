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

// LolRevivalsMission lol revivals mission
// swagger:model LolRevivalsMission
type LolRevivalsMission struct {

	// completed date
	CompletedDate int64 `json:"completedDate,omitempty"`

	// display
	Display *LolRevivalsMissionDisplay `json:"display,omitempty"`

	// internal name
	InternalName string `json:"internalName,omitempty"`

	// objectives
	Objectives []*LolRevivalsObjective `json:"objectives"`

	// series name
	SeriesName string `json:"seriesName,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this lol revivals mission
func (m *LolRevivalsMission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectives(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolRevivalsMission) validateDisplay(formats strfmt.Registry) error {

	if swag.IsZero(m.Display) { // not required
		return nil
	}

	if m.Display != nil {
		if err := m.Display.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("display")
			}
			return err
		}
	}

	return nil
}

func (m *LolRevivalsMission) validateObjectives(formats strfmt.Registry) error {

	if swag.IsZero(m.Objectives) { // not required
		return nil
	}

	for i := 0; i < len(m.Objectives); i++ {
		if swag.IsZero(m.Objectives[i]) { // not required
			continue
		}

		if m.Objectives[i] != nil {
			if err := m.Objectives[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objectives" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolRevivalsMission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolRevivalsMission) UnmarshalBinary(b []byte) error {
	var res LolRevivalsMission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
