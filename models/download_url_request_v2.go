// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DownloadURLRequestV2 download Url request v2
// swagger:model DownloadUrlRequestV2
type DownloadURLRequestV2 struct {

	// game Id
	GameID int64 `json:"gameId,omitempty"`

	// platform Id
	PlatformID string `json:"platformId,omitempty"`
}

// Validate validates this download Url request v2
func (m *DownloadURLRequestV2) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DownloadURLRequestV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DownloadURLRequestV2) UnmarshalBinary(b []byte) error {
	var res DownloadURLRequestV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
