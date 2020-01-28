// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MultipleReplayMetadataResponseItemV2 multiple replay metadata response item v2
// swagger:model MultipleReplayMetadataResponseItemV2
type MultipleReplayMetadataResponseItemV2 struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// metadata
	Metadata *ReplayMetadataV2 `json:"metadata,omitempty"`

	// status
	Status ReplayResponseStatus `json:"status,omitempty"`
}

// Validate validates this multiple replay metadata response item v2
func (m *MultipleReplayMetadataResponseItemV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MultipleReplayMetadataResponseItemV2) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *MultipleReplayMetadataResponseItemV2) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MultipleReplayMetadataResponseItemV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MultipleReplayMetadataResponseItemV2) UnmarshalBinary(b []byte) error {
	var res MultipleReplayMetadataResponseItemV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}