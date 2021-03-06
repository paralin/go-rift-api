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

// GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedReader is a Reader for the GetLolClashV1TournamentByTournamentIDPlayerHonorRestricted structure.
type GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK creates a GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK with default headers values
func NewGetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK() *GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK {
	return &GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK{}
}

/*GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK handles this case with default header values.

Successful response
*/
type GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK struct {
	Payload bool
}

func (o *GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/tournament/{tournamentId}/player-honor-restricted][%d] getLolClashV1TournamentByTournamentIdPlayerHonorRestrictedOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolClashV1TournamentByTournamentIDPlayerHonorRestrictedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
