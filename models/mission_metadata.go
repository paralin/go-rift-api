// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MissionMetadata mission metadata
// swagger:model MissionMetadata
type MissionMetadata struct {

	// npe reward pack
	NpeRewardPack *NpeRewardPackMetadata `json:"npeRewardPack,omitempty"`

	// tutorial
	Tutorial *TutorialMetadata `json:"tutorial,omitempty"`
}

// Validate validates this mission metadata
func (m *MissionMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNpeRewardPack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTutorial(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MissionMetadata) validateNpeRewardPack(formats strfmt.Registry) error {

	if swag.IsZero(m.NpeRewardPack) { // not required
		return nil
	}

	if m.NpeRewardPack != nil {
		if err := m.NpeRewardPack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("npeRewardPack")
			}
			return err
		}
	}

	return nil
}

func (m *MissionMetadata) validateTutorial(formats strfmt.Registry) error {

	if swag.IsZero(m.Tutorial) { // not required
		return nil
	}

	if m.Tutorial != nil {
		if err := m.Tutorial.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tutorial")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MissionMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MissionMetadata) UnmarshalBinary(b []byte) error {
	var res MissionMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
