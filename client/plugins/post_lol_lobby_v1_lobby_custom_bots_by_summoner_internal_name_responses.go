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

// PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameReader is a Reader for the PostLolLobbyV1LobbyCustomBotsBySummonerInternalName structure.
type PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK creates a PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK with default headers values
func NewPostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK() *PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK {
	return &PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK{}
}

/*PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK struct {
	Payload interface{}
}

func (o *PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v1/lobby/custom/bots/{summonerInternalName}][%d] postLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLobbyV1LobbyCustomBotsBySummonerInternalNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
