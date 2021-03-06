// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TracingEventV1 tracing event v1
// swagger:model TracingEventV1
type TracingEventV1 struct {

	// destination module name
	Dest string `json:"dest,omitempty"`

	// event specific details
	Details string `json:"details,omitempty"`

	// event name
	Name string `json:"name,omitempty"`

	// source module name
	Src string `json:"src,omitempty"`

	// a list of tags associated with this event
	Tags string `json:"tags,omitempty"`

	// timestamp in microseconds of when the event occurred
	When int64 `json:"when,omitempty"`
}

// Validate validates this tracing event v1
func (m *TracingEventV1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TracingEventV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TracingEventV1) UnmarshalBinary(b []byte) error {
	var res TracingEventV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
