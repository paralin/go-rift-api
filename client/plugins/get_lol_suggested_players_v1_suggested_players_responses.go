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

// GetLolSuggestedPlayersV1SuggestedPlayersReader is a Reader for the GetLolSuggestedPlayersV1SuggestedPlayers structure.
type GetLolSuggestedPlayersV1SuggestedPlayersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolSuggestedPlayersV1SuggestedPlayersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolSuggestedPlayersV1SuggestedPlayersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolSuggestedPlayersV1SuggestedPlayersOK creates a GetLolSuggestedPlayersV1SuggestedPlayersOK with default headers values
func NewGetLolSuggestedPlayersV1SuggestedPlayersOK() *GetLolSuggestedPlayersV1SuggestedPlayersOK {
	return &GetLolSuggestedPlayersV1SuggestedPlayersOK{}
}

/*GetLolSuggestedPlayersV1SuggestedPlayersOK handles this case with default header values.

Successful response
*/
type GetLolSuggestedPlayersV1SuggestedPlayersOK struct {
	Payload []*models.LolSuggestedPlayersSuggestedPlayersSuggestedPlayer
}

func (o *GetLolSuggestedPlayersV1SuggestedPlayersOK) Error() string {
	return fmt.Sprintf("[GET /lol-suggested-players/v1/suggested-players][%d] getLolSuggestedPlayersV1SuggestedPlayersOK  %+v", 200, o.Payload)
}

func (o *GetLolSuggestedPlayersV1SuggestedPlayersOK) GetPayload() []*models.LolSuggestedPlayersSuggestedPlayersSuggestedPlayer {
	return o.Payload
}

func (o *GetLolSuggestedPlayersV1SuggestedPlayersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}