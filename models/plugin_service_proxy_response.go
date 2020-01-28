// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PluginServiceProxyResponse plugin service proxy response
// swagger:model PluginServiceProxyResponse
type PluginServiceProxyResponse struct {

	// error
	Error string `json:"error,omitempty"`

	// method name
	MethodName string `json:"methodName,omitempty"`

	// payload
	Payload string `json:"payload,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this plugin service proxy response
func (m *PluginServiceProxyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginServiceProxyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginServiceProxyResponse) UnmarshalBinary(b []byte) error {
	var res PluginServiceProxyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}