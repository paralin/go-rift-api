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

// GetPlayerNotificationsV1NotificationsReader is a Reader for the GetPlayerNotificationsV1Notifications structure.
type GetPlayerNotificationsV1NotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlayerNotificationsV1NotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlayerNotificationsV1NotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlayerNotificationsV1NotificationsOK creates a GetPlayerNotificationsV1NotificationsOK with default headers values
func NewGetPlayerNotificationsV1NotificationsOK() *GetPlayerNotificationsV1NotificationsOK {
	return &GetPlayerNotificationsV1NotificationsOK{}
}

/*GetPlayerNotificationsV1NotificationsOK handles this case with default header values.

Successful response
*/
type GetPlayerNotificationsV1NotificationsOK struct {
	Payload []*models.PlayerNotificationsPlayerNotificationResource
}

func (o *GetPlayerNotificationsV1NotificationsOK) Error() string {
	return fmt.Sprintf("[GET /player-notifications/v1/notifications][%d] getPlayerNotificationsV1NotificationsOK  %+v", 200, o.Payload)
}

func (o *GetPlayerNotificationsV1NotificationsOK) GetPayload() []*models.PlayerNotificationsPlayerNotificationResource {
	return o.Payload
}

func (o *GetPlayerNotificationsV1NotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
