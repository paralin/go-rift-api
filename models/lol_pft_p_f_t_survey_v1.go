// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolPftPFTSurveyV1 lol pft p f t survey v1
// swagger:model LolPftPFTSurveyV1
type LolPftPFTSurveyV1 struct {

	// caption
	Caption string `json:"caption,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// url
	URL string `json:"url,omitempty"`
}

// Validate validates this lol pft p f t survey v1
func (m *LolPftPFTSurveyV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolPftPFTSurveyV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPftPFTSurveyV1) UnmarshalBinary(b []byte) error {
	var res LolPftPFTSurveyV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}