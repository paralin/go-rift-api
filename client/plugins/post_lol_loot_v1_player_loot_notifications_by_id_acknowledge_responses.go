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

// PostLolLootV1PlayerLootNotificationsByIDAcknowledgeReader is a Reader for the PostLolLootV1PlayerLootNotificationsByIDAcknowledge structure.
type PostLolLootV1PlayerLootNotificationsByIDAcknowledgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK creates a PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK with default headers values
func NewPostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK() *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK {
	return &PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK{}
}

/*PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK handles this case with default header values.

Successful response
*/
type PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK struct {
	Payload string
}

func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK) Error() string {
	return fmt.Sprintf("[POST /lol-loot/v1/player-loot-notifications/{id}/acknowledge][%d] postLolLootV1PlayerLootNotificationsByIdAcknowledgeOK  %+v", 200, o.Payload)
}

func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK) GetPayload() string {
	return o.Payload
}

func (o *PostLolLootV1PlayerLootNotificationsByIDAcknowledgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
