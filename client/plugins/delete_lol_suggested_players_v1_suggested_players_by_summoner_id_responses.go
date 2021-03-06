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

// DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDReader is a Reader for the DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerID structure.
type DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK creates a DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK with default headers values
func NewDeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK() *DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK {
	return &DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK{}
}

/*DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK handles this case with default header values.

Successful response
*/
type DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK struct {
	Payload interface{}
}

func (o *DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-suggested-players/v1/suggested-players/{summonerId}][%d] deleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIdOK  %+v", 200, o.Payload)
}

func (o *DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteLolSuggestedPlayersV1SuggestedPlayersBySummonerIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
