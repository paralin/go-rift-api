// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolGameflowV1SessionTournamentEndedReader is a Reader for the PostLolGameflowV1SessionTournamentEnded structure.
type PostLolGameflowV1SessionTournamentEndedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1SessionTournamentEndedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolGameflowV1SessionTournamentEndedNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1SessionTournamentEndedNoContent creates a PostLolGameflowV1SessionTournamentEndedNoContent with default headers values
func NewPostLolGameflowV1SessionTournamentEndedNoContent() *PostLolGameflowV1SessionTournamentEndedNoContent {
	return &PostLolGameflowV1SessionTournamentEndedNoContent{}
}

/*PostLolGameflowV1SessionTournamentEndedNoContent handles this case with default header values.

No content
*/
type PostLolGameflowV1SessionTournamentEndedNoContent struct {
}

func (o *PostLolGameflowV1SessionTournamentEndedNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/session/tournament-ended][%d] postLolGameflowV1SessionTournamentEndedNoContent ", 204)
}

func (o *PostLolGameflowV1SessionTournamentEndedNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
