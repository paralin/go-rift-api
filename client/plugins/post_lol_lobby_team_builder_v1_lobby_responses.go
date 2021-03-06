// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// PostLolLobbyTeamBuilderV1LobbyReader is a Reader for the PostLolLobbyTeamBuilderV1Lobby structure.
type PostLolLobbyTeamBuilderV1LobbyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyTeamBuilderV1LobbyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyTeamBuilderV1LobbyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyTeamBuilderV1LobbyOK creates a PostLolLobbyTeamBuilderV1LobbyOK with default headers values
func NewPostLolLobbyTeamBuilderV1LobbyOK() *PostLolLobbyTeamBuilderV1LobbyOK {
	return &PostLolLobbyTeamBuilderV1LobbyOK{}
}

/*PostLolLobbyTeamBuilderV1LobbyOK handles this case with default header values.

Successful response
*/
type PostLolLobbyTeamBuilderV1LobbyOK struct {
	Payload *models.LolLobbyTeamBuilderLobby
}

func (o *PostLolLobbyTeamBuilderV1LobbyOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby-team-builder/v1/lobby][%d] postLolLobbyTeamBuilderV1LobbyOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyTeamBuilderV1LobbyOK) GetPayload() *models.LolLobbyTeamBuilderLobby {
	return o.Payload
}

func (o *PostLolLobbyTeamBuilderV1LobbyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLobbyTeamBuilderLobby)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
