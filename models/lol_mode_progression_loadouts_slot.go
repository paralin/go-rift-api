// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolModeProgressionLoadoutsSlot lol mode progression loadouts slot
// swagger:model LolModeProgressionLoadoutsSlot
type LolModeProgressionLoadoutsSlot struct {

	// content Id
	ContentID string `json:"contentId,omitempty"`

	// inventory type
	InventoryType string `json:"inventoryType,omitempty"`

	// item Id
	ItemID int32 `json:"itemId,omitempty"`
}

// Validate validates this lol mode progression loadouts slot
func (m *LolModeProgressionLoadoutsSlot) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolModeProgressionLoadoutsSlot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolModeProgressionLoadoutsSlot) UnmarshalBinary(b []byte) error {
	var res LolModeProgressionLoadoutsSlot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
