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

// LolLootRecipe lol loot recipe
// swagger:model LolLootRecipe
type LolLootRecipe struct {

	// context menu text
	ContextMenuText string `json:"contextMenuText,omitempty"`

	// crafter name
	CrafterName string `json:"crafterName,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// display categories
	DisplayCategories string `json:"displayCategories,omitempty"`

	// image path
	ImagePath string `json:"imagePath,omitempty"`

	// intro video path
	IntroVideoPath string `json:"introVideoPath,omitempty"`

	// loop video path
	LoopVideoPath string `json:"loopVideoPath,omitempty"`

	// metadata
	Metadata *LolLootRecipeMetadata `json:"metadata,omitempty"`

	// outputs
	Outputs []*LolLootRecipeOutput `json:"outputs"`

	// outro video path
	OutroVideoPath string `json:"outroVideoPath,omitempty"`

	// recipe name
	RecipeName string `json:"recipeName,omitempty"`

	// requirement text
	RequirementText string `json:"requirementText,omitempty"`

	// slots
	Slots []*LolLootRecipeSlot `json:"slots"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this lol loot recipe
func (m *LolLootRecipe) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutputs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlots(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolLootRecipe) validateMetadata(formats strfmt.Registry) error {

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

func (m *LolLootRecipe) validateOutputs(formats strfmt.Registry) error {

	if swag.IsZero(m.Outputs) { // not required
		return nil
	}

	for i := 0; i < len(m.Outputs); i++ {
		if swag.IsZero(m.Outputs[i]) { // not required
			continue
		}

		if m.Outputs[i] != nil {
			if err := m.Outputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("outputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolLootRecipe) validateSlots(formats strfmt.Registry) error {

	if swag.IsZero(m.Slots) { // not required
		return nil
	}

	for i := 0; i < len(m.Slots); i++ {
		if swag.IsZero(m.Slots[i]) { // not required
			continue
		}

		if m.Slots[i] != nil {
			if err := m.Slots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("slots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolLootRecipe) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolLootRecipe) UnmarshalBinary(b []byte) error {
	var res LolLootRecipe
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
