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

// LolPftPFTSurveyResults lol pft p f t survey results
// swagger:model LolPftPFTSurveyResults
type LolPftPFTSurveyResults struct {

	// actions
	Actions []*LolPftPFTEvent `json:"actions"`

	// question responses
	QuestionResponses []*LolPftPFTQuestionResponse `json:"questionResponses"`
}

// Validate validates this lol pft p f t survey results
func (m *LolPftPFTSurveyResults) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuestionResponses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LolPftPFTSurveyResults) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LolPftPFTSurveyResults) validateQuestionResponses(formats strfmt.Registry) error {

	if swag.IsZero(m.QuestionResponses) { // not required
		return nil
	}

	for i := 0; i < len(m.QuestionResponses); i++ {
		if swag.IsZero(m.QuestionResponses[i]) { // not required
			continue
		}

		if m.QuestionResponses[i] != nil {
			if err := m.QuestionResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("questionResponses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LolPftPFTSurveyResults) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolPftPFTSurveyResults) UnmarshalBinary(b []byte) error {
	var res LolPftPFTSurveyResults
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
