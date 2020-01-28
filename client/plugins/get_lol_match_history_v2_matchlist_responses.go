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

// GetLolMatchHistoryV2MatchlistReader is a Reader for the GetLolMatchHistoryV2Matchlist structure.
type GetLolMatchHistoryV2MatchlistReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolMatchHistoryV2MatchlistReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolMatchHistoryV2MatchlistOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolMatchHistoryV2MatchlistOK creates a GetLolMatchHistoryV2MatchlistOK with default headers values
func NewGetLolMatchHistoryV2MatchlistOK() *GetLolMatchHistoryV2MatchlistOK {
	return &GetLolMatchHistoryV2MatchlistOK{}
}

/*GetLolMatchHistoryV2MatchlistOK handles this case with default header values.

Successful response
*/
type GetLolMatchHistoryV2MatchlistOK struct {
	Payload *models.LolMatchHistoryMatchHistoryList
}

func (o *GetLolMatchHistoryV2MatchlistOK) Error() string {
	return fmt.Sprintf("[GET /lol-match-history/v2/matchlist][%d] getLolMatchHistoryV2MatchlistOK  %+v", 200, o.Payload)
}

func (o *GetLolMatchHistoryV2MatchlistOK) GetPayload() *models.LolMatchHistoryMatchHistoryList {
	return o.Payload
}

func (o *GetLolMatchHistoryV2MatchlistOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolMatchHistoryMatchHistoryList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
