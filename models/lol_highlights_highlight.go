// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolHighlightsHighlight lol highlights highlight
// swagger:model LolHighlightsHighlight
type LolHighlightsHighlight struct {

	// file size bytes
	FileSizeBytes int64 `json:"fileSizeBytes,omitempty"`

	// filepath
	Filepath string `json:"filepath,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// mtime iso8601
	MtimeIso8601 string `json:"mtimeIso8601,omitempty"`

	// mtime ms utc
	MtimeMsUtc int64 `json:"mtimeMsUtc,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this lol highlights highlight
func (m *LolHighlightsHighlight) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolHighlightsHighlight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolHighlightsHighlight) UnmarshalBinary(b []byte) error {
	var res LolHighlightsHighlight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}