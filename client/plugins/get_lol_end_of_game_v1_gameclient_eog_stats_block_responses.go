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

// GetLolEndOfGameV1GameclientEogStatsBlockReader is a Reader for the GetLolEndOfGameV1GameclientEogStatsBlock structure.
type GetLolEndOfGameV1GameclientEogStatsBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolEndOfGameV1GameclientEogStatsBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolEndOfGameV1GameclientEogStatsBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolEndOfGameV1GameclientEogStatsBlockOK creates a GetLolEndOfGameV1GameclientEogStatsBlockOK with default headers values
func NewGetLolEndOfGameV1GameclientEogStatsBlockOK() *GetLolEndOfGameV1GameclientEogStatsBlockOK {
	return &GetLolEndOfGameV1GameclientEogStatsBlockOK{}
}

/*GetLolEndOfGameV1GameclientEogStatsBlockOK handles this case with default header values.

Successful response
*/
type GetLolEndOfGameV1GameclientEogStatsBlockOK struct {
	Payload *models.LolEndOfGameGameClientEndOfGameStats
}

func (o *GetLolEndOfGameV1GameclientEogStatsBlockOK) Error() string {
	return fmt.Sprintf("[GET /lol-end-of-game/v1/gameclient-eog-stats-block][%d] getLolEndOfGameV1GameclientEogStatsBlockOK  %+v", 200, o.Payload)
}

func (o *GetLolEndOfGameV1GameclientEogStatsBlockOK) GetPayload() *models.LolEndOfGameGameClientEndOfGameStats {
	return o.Payload
}

func (o *GetLolEndOfGameV1GameclientEogStatsBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolEndOfGameGameClientEndOfGameStats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}