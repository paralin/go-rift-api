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

// GetLolClashV1HistoryandwinnersReader is a Reader for the GetLolClashV1Historyandwinners structure.
type GetLolClashV1HistoryandwinnersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1HistoryandwinnersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1HistoryandwinnersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1HistoryandwinnersOK creates a GetLolClashV1HistoryandwinnersOK with default headers values
func NewGetLolClashV1HistoryandwinnersOK() *GetLolClashV1HistoryandwinnersOK {
	return &GetLolClashV1HistoryandwinnersOK{}
}

/*GetLolClashV1HistoryandwinnersOK handles this case with default header values.

Successful response
*/
type GetLolClashV1HistoryandwinnersOK struct {
	Payload *models.LolClashTournamentHistoryAndWinners
}

func (o *GetLolClashV1HistoryandwinnersOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/historyandwinners][%d] getLolClashV1HistoryandwinnersOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1HistoryandwinnersOK) GetPayload() *models.LolClashTournamentHistoryAndWinners {
	return o.Payload
}

func (o *GetLolClashV1HistoryandwinnersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolClashTournamentHistoryAndWinners)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}