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

// GetLolEndOfGameV1TftEogStatsReader is a Reader for the GetLolEndOfGameV1TftEogStats structure.
type GetLolEndOfGameV1TftEogStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolEndOfGameV1TftEogStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolEndOfGameV1TftEogStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolEndOfGameV1TftEogStatsOK creates a GetLolEndOfGameV1TftEogStatsOK with default headers values
func NewGetLolEndOfGameV1TftEogStatsOK() *GetLolEndOfGameV1TftEogStatsOK {
	return &GetLolEndOfGameV1TftEogStatsOK{}
}

/*GetLolEndOfGameV1TftEogStatsOK handles this case with default header values.

Successful response
*/
type GetLolEndOfGameV1TftEogStatsOK struct {
	Payload *models.LolEndOfGameTFTEndOfGameViewModel
}

func (o *GetLolEndOfGameV1TftEogStatsOK) Error() string {
	return fmt.Sprintf("[GET /lol-end-of-game/v1/tft-eog-stats][%d] getLolEndOfGameV1TftEogStatsOK  %+v", 200, o.Payload)
}

func (o *GetLolEndOfGameV1TftEogStatsOK) GetPayload() *models.LolEndOfGameTFTEndOfGameViewModel {
	return o.Payload
}

func (o *GetLolEndOfGameV1TftEogStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolEndOfGameTFTEndOfGameViewModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
