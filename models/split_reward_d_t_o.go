// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// SplitRewardDTO split reward d t o
// swagger:model SplitRewardDTO
type SplitRewardDTO struct {

	// default reward Id
	DefaultRewardID string `json:"defaultRewardId,omitempty"`

	// metadata
	Metadata *SplitRewardsMetaData `json:"metadata,omitempty"`

	// reward type
	RewardType string `json:"rewardType,omitempty"`

	// tiered reward ids
	TieredRewardIds map[string]string `json:"tieredRewardIds,omitempty"`
}

// Validate validates this split reward d t o
func (m *SplitRewardDTO) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SplitRewardDTO) validateMetadata(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *SplitRewardDTO) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SplitRewardDTO) UnmarshalBinary(b []byte) error {
	var res SplitRewardDTO
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
