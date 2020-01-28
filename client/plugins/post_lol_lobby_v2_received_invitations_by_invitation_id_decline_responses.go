// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineReader is a Reader for the PostLolLobbyV2ReceivedInvitationsByInvitationIDDecline structure.
type PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent creates a PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent with default headers values
func NewPostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent() *PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent {
	return &PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent{}
}

/*PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent handles this case with default header values.

No content
*/
type PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent struct {
}

func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v2/received-invitations/{invitationId}/decline][%d] postLolLobbyV2ReceivedInvitationsByInvitationIdDeclineNoContent ", 204)
}

func (o *PostLolLobbyV2ReceivedInvitationsByInvitationIDDeclineNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}