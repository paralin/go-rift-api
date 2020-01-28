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

// GetLolLobbyV2ReceivedInvitationsReader is a Reader for the GetLolLobbyV2ReceivedInvitations structure.
type GetLolLobbyV2ReceivedInvitationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyV2ReceivedInvitationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyV2ReceivedInvitationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyV2ReceivedInvitationsOK creates a GetLolLobbyV2ReceivedInvitationsOK with default headers values
func NewGetLolLobbyV2ReceivedInvitationsOK() *GetLolLobbyV2ReceivedInvitationsOK {
	return &GetLolLobbyV2ReceivedInvitationsOK{}
}

/*GetLolLobbyV2ReceivedInvitationsOK handles this case with default header values.

Successful response
*/
type GetLolLobbyV2ReceivedInvitationsOK struct {
	Payload []*models.LolLobbyReceivedInvitationDto
}

func (o *GetLolLobbyV2ReceivedInvitationsOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby/v2/received-invitations][%d] getLolLobbyV2ReceivedInvitationsOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyV2ReceivedInvitationsOK) GetPayload() []*models.LolLobbyReceivedInvitationDto {
	return o.Payload
}

func (o *GetLolLobbyV2ReceivedInvitationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}