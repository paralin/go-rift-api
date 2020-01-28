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

// PostLolLobbyV2LobbyInvitationsReader is a Reader for the PostLolLobbyV2LobbyInvitations structure.
type PostLolLobbyV2LobbyInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV2LobbyInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV2LobbyInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV2LobbyInvitationsOK creates a PostLolLobbyV2LobbyInvitationsOK with default headers values
func NewPostLolLobbyV2LobbyInvitationsOK() *PostLolLobbyV2LobbyInvitationsOK {
	return &PostLolLobbyV2LobbyInvitationsOK{}
}

/*PostLolLobbyV2LobbyInvitationsOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV2LobbyInvitationsOK struct {
	Payload []*models.LolLobbyLobbyInvitationDto
}

func (o *PostLolLobbyV2LobbyInvitationsOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v2/lobby/invitations][%d] postLolLobbyV2LobbyInvitationsOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV2LobbyInvitationsOK) GetPayload() []*models.LolLobbyLobbyInvitationDto {
	return o.Payload
}

func (o *PostLolLobbyV2LobbyInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}