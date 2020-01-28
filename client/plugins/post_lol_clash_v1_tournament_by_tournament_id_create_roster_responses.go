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

// PostLolClashV1TournamentByTournamentIDCreateRosterReader is a Reader for the PostLolClashV1TournamentByTournamentIDCreateRoster structure.
type PostLolClashV1TournamentByTournamentIDCreateRosterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolClashV1TournamentByTournamentIDCreateRosterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolClashV1TournamentByTournamentIDCreateRosterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolClashV1TournamentByTournamentIDCreateRosterOK creates a PostLolClashV1TournamentByTournamentIDCreateRosterOK with default headers values
func NewPostLolClashV1TournamentByTournamentIDCreateRosterOK() *PostLolClashV1TournamentByTournamentIDCreateRosterOK {
	return &PostLolClashV1TournamentByTournamentIDCreateRosterOK{}
}

/*PostLolClashV1TournamentByTournamentIDCreateRosterOK handles this case with default header values.

Successful response
*/
type PostLolClashV1TournamentByTournamentIDCreateRosterOK struct {
	Payload interface{}
}

func (o *PostLolClashV1TournamentByTournamentIDCreateRosterOK) Error() string {
	return fmt.Sprintf("[POST /lol-clash/v1/tournament/{tournamentId}/create-roster][%d] postLolClashV1TournamentByTournamentIdCreateRosterOK  %+v", 200, o.Payload)
}

func (o *PostLolClashV1TournamentByTournamentIDCreateRosterOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolClashV1TournamentByTournamentIDCreateRosterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
