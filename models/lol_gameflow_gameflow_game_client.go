// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LolGameflowGameflowGameClient lol gameflow gameflow game client
// swagger:model LolGameflowGameflowGameClient
type LolGameflowGameflowGameClient struct {

	// observer server Ip
	ObserverServerIP string `json:"observerServerIp,omitempty"`

	// observer server port
	ObserverServerPort int64 `json:"observerServerPort,omitempty"`

	// running
	Running bool `json:"running,omitempty"`

	// server Ip
	ServerIP string `json:"serverIp,omitempty"`

	// server port
	ServerPort int64 `json:"serverPort,omitempty"`

	// visible
	Visible bool `json:"visible,omitempty"`
}

// Validate validates this lol gameflow gameflow game client
func (m *LolGameflowGameflowGameClient) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LolGameflowGameflowGameClient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LolGameflowGameflowGameClient) UnmarshalBinary(b []byte) error {
	var res LolGameflowGameflowGameClient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}