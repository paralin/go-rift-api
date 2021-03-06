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

// GetLolClashV1TournamentByTournamentIDReader is a Reader for the GetLolClashV1TournamentByTournamentID structure.
type GetLolClashV1TournamentByTournamentIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1TournamentByTournamentIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1TournamentByTournamentIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1TournamentByTournamentIDOK creates a GetLolClashV1TournamentByTournamentIDOK with default headers values
func NewGetLolClashV1TournamentByTournamentIDOK() *GetLolClashV1TournamentByTournamentIDOK {
	return &GetLolClashV1TournamentByTournamentIDOK{}
}

/*GetLolClashV1TournamentByTournamentIDOK handles this case with default header values.

Successful response
*/
type GetLolClashV1TournamentByTournamentIDOK struct {
	Payload *models.LolClashTournament
}

func (o *GetLolClashV1TournamentByTournamentIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/tournament/{tournamentId}][%d] getLolClashV1TournamentByTournamentIdOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1TournamentByTournamentIDOK) GetPayload() *models.LolClashTournament {
	return o.Payload
}

func (o *GetLolClashV1TournamentByTournamentIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolClashTournament)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
