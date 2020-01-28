// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolChampSelectChampSelectPickableSkins lol champ select champ select pickable skins
// swagger:model LolChampSelectChampSelectPickableSkins
type LolChampSelectChampSelectPickableSkins struct {

	// skin ids
	SkinIds []int32 `json:"skinIds"`
}

// Validate validates this lol champ select champ select pickable skins
func (m *LolChampSelectChampSelectPickableSkins) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolChampSelectChampSelectPickableSkins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolChampSelectChampSelectPickableSkins) UnmarshalBinary(b []byte) error {
	var res LolChampSelectChampSelectPickableSkins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
