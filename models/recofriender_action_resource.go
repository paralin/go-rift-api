// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RecofrienderActionResource recofriender action resource
// swagger:model RecofrienderActionResource
type RecofrienderActionResource struct {

	// account Id
	AccountID int64 `json:"accountId,omitempty"`

	// action
	Action string `json:"action,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`
}

// Validate validates this recofriender action resource
func (m *RecofrienderActionResource) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecofrienderActionResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecofrienderActionResource) UnmarshalBinary(b []byte) error {
	var res RecofrienderActionResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
