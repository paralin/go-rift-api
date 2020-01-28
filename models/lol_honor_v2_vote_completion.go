// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolHonorV2VoteCompletion lol honor v2 vote completion
// swagger:model LolHonorV2VoteCompletion
type LolHonorV2VoteCompletion struct {

	// full team vote
	FullTeamVote bool `json:"fullTeamVote,omitempty"`

	// game Id
	GameID int64 `json:"gameId,omitempty"`
}

// Validate validates this lol honor v2 vote completion
func (m *LolHonorV2VoteCompletion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolHonorV2VoteCompletion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHonorV2VoteCompletion) UnmarshalBinary(b []byte) error {
	var res LolHonorV2VoteCompletion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}