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

// LolHonorV2MutualHonor lol honor v2 mutual honor
// swagger:model LolHonorV2MutualHonor
type LolHonorV2MutualHonor struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// summoners
	Summoners []*LolHonorV2MutualHonorPlayer `json:"summoners"`
}

// Validate validates this lol honor v2 mutual honor
func (m *LolHonorV2MutualHonor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSummoners(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolHonorV2MutualHonor) validateSummoners(formats strfmt.Registry) error {

	if swag.IsZero(m.Summoners) { // not required
		return nil
	}

	for i := 0; i < len(m.Summoners); i++ {
		if swag.IsZero(m.Summoners[i]) { // not required
			continue
		}

		if m.Summoners[i] != nil {
			if err := m.Summoners[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("summoners" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolHonorV2MutualHonor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHonorV2MutualHonor) UnmarshalBinary(b []byte) error {
	var res LolHonorV2MutualHonor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}