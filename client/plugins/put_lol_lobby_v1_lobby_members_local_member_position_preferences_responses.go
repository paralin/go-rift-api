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

// PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesReader is a Reader for the PutLolLobbyV1LobbyMembersLocalMemberPositionPreferences structure.
type PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK creates a PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK with default headers values
func NewPutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK() *PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK {
	return &PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK{}
}

/*PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK handles this case with default header values.

Successful response
*/
type PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK struct {
	Payload interface{}
}

func (o *PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK) Error() string {
	return fmt.Sprintf("[PUT /lol-lobby/v1/lobby/members/localMember/position-preferences][%d] putLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK  %+v", 200, o.Payload)
}

func (o *PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PutLolLobbyV1LobbyMembersLocalMemberPositionPreferencesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
