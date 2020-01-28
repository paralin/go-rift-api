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

// GetLolKrPlaytimeReminderV1PlaytimeReader is a Reader for the GetLolKrPlaytimeReminderV1Playtime structure.
type GetLolKrPlaytimeReminderV1PlaytimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolKrPlaytimeReminderV1PlaytimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolKrPlaytimeReminderV1PlaytimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolKrPlaytimeReminderV1PlaytimeOK creates a GetLolKrPlaytimeReminderV1PlaytimeOK with default headers values
func NewGetLolKrPlaytimeReminderV1PlaytimeOK() *GetLolKrPlaytimeReminderV1PlaytimeOK {
	return &GetLolKrPlaytimeReminderV1PlaytimeOK{}
}

/*GetLolKrPlaytimeReminderV1PlaytimeOK handles this case with default header values.

Successful response
*/
type GetLolKrPlaytimeReminderV1PlaytimeOK struct {
	Payload *models.LolKrPlaytimeReminderPlaytimeReminder
}

func (o *GetLolKrPlaytimeReminderV1PlaytimeOK) Error() string {
	return fmt.Sprintf("[GET /lol-kr-playtime-reminder/v1/playtime][%d] getLolKrPlaytimeReminderV1PlaytimeOK  %+v", 200, o.Payload)
}

func (o *GetLolKrPlaytimeReminderV1PlaytimeOK) GetPayload() *models.LolKrPlaytimeReminderPlaytimeReminder {
	return o.Payload
}

func (o *GetLolKrPlaytimeReminderV1PlaytimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolKrPlaytimeReminderPlaytimeReminder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}