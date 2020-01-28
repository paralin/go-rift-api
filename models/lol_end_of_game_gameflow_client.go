// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolEndOfGameGameflowClient lol end of game gameflow client
// swagger:model LolEndOfGameGameflowClient
type LolEndOfGameGameflowClient struct {

	// observer server Ip
	ObserverServerIP string `json:"observerServerIp,omitempty"`

	// observer server port
	ObserverServerPort int64 `json:"observerServerPort,omitempty"`

	// running
	Running bool `json:"running,omitempty"`
}

// Validate validates this lol end of game gameflow client
func (m *LolEndOfGameGameflowClient) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolEndOfGameGameflowClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolEndOfGameGameflowClient) UnmarshalBinary(b []byte) error {
	var res LolEndOfGameGameflowClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}