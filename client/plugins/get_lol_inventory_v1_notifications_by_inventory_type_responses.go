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

// GetLolInventoryV1NotificationsByInventoryTypeReader is a Reader for the GetLolInventoryV1NotificationsByInventoryType structure.
type GetLolInventoryV1NotificationsByInventoryTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolInventoryV1NotificationsByInventoryTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolInventoryV1NotificationsByInventoryTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolInventoryV1NotificationsByInventoryTypeOK creates a GetLolInventoryV1NotificationsByInventoryTypeOK with default headers values
func NewGetLolInventoryV1NotificationsByInventoryTypeOK() *GetLolInventoryV1NotificationsByInventoryTypeOK {
	return &GetLolInventoryV1NotificationsByInventoryTypeOK{}
}

/*GetLolInventoryV1NotificationsByInventoryTypeOK handles this case with default header values.

Successful response
*/
type GetLolInventoryV1NotificationsByInventoryTypeOK struct {
	Payload []*models.LolInventoryInventoryNotification
}

func (o *GetLolInventoryV1NotificationsByInventoryTypeOK) Error() string {
	return fmt.Sprintf("[GET /lol-inventory/v1/notifications/{inventoryType}][%d] getLolInventoryV1NotificationsByInventoryTypeOK  %+v", 200, o.Payload)
}

func (o *GetLolInventoryV1NotificationsByInventoryTypeOK) GetPayload() []*models.LolInventoryInventoryNotification {
	return o.Payload
}

func (o *GetLolInventoryV1NotificationsByInventoryTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}