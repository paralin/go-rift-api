// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PatcherP2PStatus patcher p2 p status
// swagger:model PatcherP2PStatus
type PatcherP2PStatus struct {

	// is allowed by user
	IsAllowedByUser bool `json:"isAllowedByUser,omitempty"`

	// is enabled for patchline
	IsEnabledForPatchline bool `json:"isEnabledForPatchline,omitempty"`

	// requires restart
	RequiresRestart bool `json:"requiresRestart,omitempty"`
}

// Validate validates this patcher p2 p status
func (m *PatcherP2PStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatcherP2PStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatcherP2PStatus) UnmarshalBinary(b []byte) error {
	var res PatcherP2PStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
