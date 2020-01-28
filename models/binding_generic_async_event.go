// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// BindingGenericAsyncEvent Represents generic data for an asynchronous event.
// swagger:model BindingGenericAsyncEvent
type BindingGenericAsyncEvent struct {

	// Asynchronous operation token
	AsyncToken int32 `json:"asyncToken,omitempty"`

	// Event data
	Data interface{} `json:"data,omitempty"`
}

// Validate validates this binding generic async event
func (m *BindingGenericAsyncEvent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BindingGenericAsyncEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BindingGenericAsyncEvent) UnmarshalBinary(b []byte) error {
	var res BindingGenericAsyncEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}