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

// DeleteLolLootV1PlayerLootByLootIDNewNotificationReader is a Reader for the DeleteLolLootV1PlayerLootByLootIDNewNotification structure.
type DeleteLolLootV1PlayerLootByLootIDNewNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolLootV1PlayerLootByLootIDNewNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolLootV1PlayerLootByLootIDNewNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolLootV1PlayerLootByLootIDNewNotificationOK creates a DeleteLolLootV1PlayerLootByLootIDNewNotificationOK with default headers values
func NewDeleteLolLootV1PlayerLootByLootIDNewNotificationOK() *DeleteLolLootV1PlayerLootByLootIDNewNotificationOK {
	return &DeleteLolLootV1PlayerLootByLootIDNewNotificationOK{}
}

/*DeleteLolLootV1PlayerLootByLootIDNewNotificationOK handles this case with default header values.

Successful response
*/
type DeleteLolLootV1PlayerLootByLootIDNewNotificationOK struct {
	Payload interface{}
}

func (o *DeleteLolLootV1PlayerLootByLootIDNewNotificationOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-loot/v1/player-loot/{lootId}/new-notification][%d] deleteLolLootV1PlayerLootByLootIdNewNotificationOK  %+v", 200, o.Payload)
}

func (o *DeleteLolLootV1PlayerLootByLootIDNewNotificationOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteLolLootV1PlayerLootByLootIDNewNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}