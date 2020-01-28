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

// GetLolStatstonesV1EogNotificationsReader is a Reader for the GetLolStatstonesV1EogNotifications structure.
type GetLolStatstonesV1EogNotificationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolStatstonesV1EogNotificationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolStatstonesV1EogNotificationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolStatstonesV1EogNotificationsOK creates a GetLolStatstonesV1EogNotificationsOK with default headers values
func NewGetLolStatstonesV1EogNotificationsOK() *GetLolStatstonesV1EogNotificationsOK {
	return &GetLolStatstonesV1EogNotificationsOK{}
}

/*GetLolStatstonesV1EogNotificationsOK handles this case with default header values.

Successful response
*/
type GetLolStatstonesV1EogNotificationsOK struct {
	Payload []interface{}
}

func (o *GetLolStatstonesV1EogNotificationsOK) Error() string {
	return fmt.Sprintf("[GET /lol-statstones/v1/eog-notifications][%d] getLolStatstonesV1EogNotificationsOK  %+v", 200, o.Payload)
}

func (o *GetLolStatstonesV1EogNotificationsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetLolStatstonesV1EogNotificationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
