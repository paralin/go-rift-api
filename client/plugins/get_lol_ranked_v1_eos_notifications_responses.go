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

// GetLolRankedV1EosNotificationsReader is a Reader for the GetLolRankedV1EosNotifications structure.
type GetLolRankedV1EosNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRankedV1EosNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRankedV1EosNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRankedV1EosNotificationsOK creates a GetLolRankedV1EosNotificationsOK with default headers values
func NewGetLolRankedV1EosNotificationsOK() *GetLolRankedV1EosNotificationsOK {
	return &GetLolRankedV1EosNotificationsOK{}
}

/*GetLolRankedV1EosNotificationsOK handles this case with default header values.

Successful response
*/
type GetLolRankedV1EosNotificationsOK struct {
	Payload []*models.LolRankedEosNotificationResource
}

func (o *GetLolRankedV1EosNotificationsOK) Error() string {
	return fmt.Sprintf("[GET /lol-ranked/v1/eos-notifications][%d] getLolRankedV1EosNotificationsOK  %+v", 200, o.Payload)
}

func (o *GetLolRankedV1EosNotificationsOK) GetPayload() []*models.LolRankedEosNotificationResource {
	return o.Payload
}

func (o *GetLolRankedV1EosNotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
