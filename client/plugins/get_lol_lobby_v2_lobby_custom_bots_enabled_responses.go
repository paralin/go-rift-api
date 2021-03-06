// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLolLobbyV2LobbyCustomBotsEnabledReader is a Reader for the GetLolLobbyV2LobbyCustomBotsEnabled structure.
type GetLolLobbyV2LobbyCustomBotsEnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyV2LobbyCustomBotsEnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyV2LobbyCustomBotsEnabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyV2LobbyCustomBotsEnabledOK creates a GetLolLobbyV2LobbyCustomBotsEnabledOK with default headers values
func NewGetLolLobbyV2LobbyCustomBotsEnabledOK() *GetLolLobbyV2LobbyCustomBotsEnabledOK {
	return &GetLolLobbyV2LobbyCustomBotsEnabledOK{}
}

/*GetLolLobbyV2LobbyCustomBotsEnabledOK handles this case with default header values.

Successful response
*/
type GetLolLobbyV2LobbyCustomBotsEnabledOK struct {
	Payload bool
}

func (o *GetLolLobbyV2LobbyCustomBotsEnabledOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby/v2/lobby/custom/bots-enabled][%d] getLolLobbyV2LobbyCustomBotsEnabledOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyV2LobbyCustomBotsEnabledOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolLobbyV2LobbyCustomBotsEnabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
