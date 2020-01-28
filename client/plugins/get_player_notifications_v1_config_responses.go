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

// GetPlayerNotificationsV1ConfigReader is a Reader for the GetPlayerNotificationsV1Config structure.
type GetPlayerNotificationsV1ConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPlayerNotificationsV1ConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPlayerNotificationsV1ConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPlayerNotificationsV1ConfigOK creates a GetPlayerNotificationsV1ConfigOK with default headers values
func NewGetPlayerNotificationsV1ConfigOK() *GetPlayerNotificationsV1ConfigOK {
	return &GetPlayerNotificationsV1ConfigOK{}
}

/*GetPlayerNotificationsV1ConfigOK handles this case with default header values.

Successful response
*/
type GetPlayerNotificationsV1ConfigOK struct {
	Payload *models.PlayerNotificationsPlayerNotificationConfigResource
}

func (o *GetPlayerNotificationsV1ConfigOK) Error() string {
	return fmt.Sprintf("[GET /player-notifications/v1/config][%d] getPlayerNotificationsV1ConfigOK  %+v", 200, o.Payload)
}

func (o *GetPlayerNotificationsV1ConfigOK) GetPayload() *models.PlayerNotificationsPlayerNotificationConfigResource {
	return o.Payload
}

func (o *GetPlayerNotificationsV1ConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlayerNotificationsPlayerNotificationConfigResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
