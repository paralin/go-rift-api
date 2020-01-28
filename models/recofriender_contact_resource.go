// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// RecofrienderContactResource recofriender contact resource
// swagger:model RecofrienderContactResource
type RecofrienderContactResource struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// action
	Action string `json:"action,omitempty"`

	// display state
	DisplayState string `json:"displayState,omitempty"`

	// friend state
	FriendState RecofrienderFriendState `json:"friendState,omitempty"`

	// image Url
	ImageURL string `json:"imageUrl,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// recommend score
	RecommendScore int64 `json:"recommendScore,omitempty"`

	// source
	Source string `json:"source,omitempty"`

	// summoner Id
	SummonerID int64 `json:"summonerId,omitempty"`
}

// Validate validates this recofriender contact resource
func (m *RecofrienderContactResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFriendState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecofrienderContactResource) validateFriendState(formats strfmt.Registry) error {

	if swag.IsZero(m.FriendState) { // not required
		return nil
	}

	if err := m.FriendState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("friendState")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecofrienderContactResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecofrienderContactResource) UnmarshalBinary(b []byte) error {
	var res RecofrienderContactResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}