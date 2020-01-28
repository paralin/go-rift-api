// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolHonorV2ServiceProxyRetrieveProfileResponse lol honor v2 service proxy retrieve profile response
// swagger:model LolHonorV2ServiceProxyRetrieveProfileResponse
type LolHonorV2ServiceProxyRetrieveProfileResponse struct {

	// checkpoint
	Checkpoint int32 `json:"checkpoint,omitempty"`

	// honor level
	HonorLevel int32 `json:"honorLevel,omitempty"`

	// rewards locked
	RewardsLocked bool `json:"rewardsLocked,omitempty"`
}

// Validate validates this lol honor v2 service proxy retrieve profile response
func (m *LolHonorV2ServiceProxyRetrieveProfileResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolHonorV2ServiceProxyRetrieveProfileResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHonorV2ServiceProxyRetrieveProfileResponse) UnmarshalBinary(b []byte) error {
	var res LolHonorV2ServiceProxyRetrieveProfileResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
