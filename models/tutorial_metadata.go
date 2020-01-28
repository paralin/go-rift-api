// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TutorialMetadata tutorial metadata
// swagger:model TutorialMetadata
type TutorialMetadata struct {

	// display rewards
	DisplayRewards map[string]string `json:"displayRewards,omitempty"`

	// queue Id
	QueueID string `json:"queueId,omitempty"`

	// step number
	StepNumber int32 `json:"stepNumber,omitempty"`

	// use chosen champion
	UseChosenChampion bool `json:"useChosenChampion,omitempty"`

	// use quick search matchmaking
	UseQuickSearchMatchmaking bool `json:"useQuickSearchMatchmaking,omitempty"`
}

// Validate validates this tutorial metadata
func (m *TutorialMetadata) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TutorialMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TutorialMetadata) UnmarshalBinary(b []byte) error {
	var res TutorialMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
