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

// PostLolLobbyV1TournamentsByIDJoinReader is a Reader for the PostLolLobbyV1TournamentsByIDJoin structure.
type PostLolLobbyV1TournamentsByIDJoinReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV1TournamentsByIDJoinReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV1TournamentsByIDJoinOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV1TournamentsByIDJoinOK creates a PostLolLobbyV1TournamentsByIDJoinOK with default headers values
func NewPostLolLobbyV1TournamentsByIDJoinOK() *PostLolLobbyV1TournamentsByIDJoinOK {
	return &PostLolLobbyV1TournamentsByIDJoinOK{}
}

/*PostLolLobbyV1TournamentsByIDJoinOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV1TournamentsByIDJoinOK struct {
	Payload interface{}
}

func (o *PostLolLobbyV1TournamentsByIDJoinOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v1/tournaments/{id}/join][%d] postLolLobbyV1TournamentsByIdJoinOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV1TournamentsByIDJoinOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLobbyV1TournamentsByIDJoinOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
