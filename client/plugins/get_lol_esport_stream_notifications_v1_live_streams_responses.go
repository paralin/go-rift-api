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

// GetLolEsportStreamNotificationsV1LiveStreamsReader is a Reader for the GetLolEsportStreamNotificationsV1LiveStreams structure.
type GetLolEsportStreamNotificationsV1LiveStreamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolEsportStreamNotificationsV1LiveStreamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolEsportStreamNotificationsV1LiveStreamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolEsportStreamNotificationsV1LiveStreamsOK creates a GetLolEsportStreamNotificationsV1LiveStreamsOK with default headers values
func NewGetLolEsportStreamNotificationsV1LiveStreamsOK() *GetLolEsportStreamNotificationsV1LiveStreamsOK {
	return &GetLolEsportStreamNotificationsV1LiveStreamsOK{}
}

/*GetLolEsportStreamNotificationsV1LiveStreamsOK handles this case with default header values.

Successful response
*/
type GetLolEsportStreamNotificationsV1LiveStreamsOK struct {
	Payload *models.LolEsportStreamNotificationsESportsLiveStreams
}

func (o *GetLolEsportStreamNotificationsV1LiveStreamsOK) Error() string {
	return fmt.Sprintf("[GET /lol-esport-stream-notifications/v1/live-streams][%d] getLolEsportStreamNotificationsV1LiveStreamsOK  %+v", 200, o.Payload)
}

func (o *GetLolEsportStreamNotificationsV1LiveStreamsOK) GetPayload() *models.LolEsportStreamNotificationsESportsLiveStreams {
	return o.Payload
}

func (o *GetLolEsportStreamNotificationsV1LiveStreamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolEsportStreamNotificationsESportsLiveStreams)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
