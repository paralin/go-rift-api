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

// GetLolPerksV1InventoryReader is a Reader for the GetLolPerksV1Inventory structure.
type GetLolPerksV1InventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPerksV1InventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPerksV1InventoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPerksV1InventoryOK creates a GetLolPerksV1InventoryOK with default headers values
func NewGetLolPerksV1InventoryOK() *GetLolPerksV1InventoryOK {
	return &GetLolPerksV1InventoryOK{}
}

/*GetLolPerksV1InventoryOK handles this case with default header values.

Successful response
*/
type GetLolPerksV1InventoryOK struct {
	Payload *models.LolPerksPlayerInventory
}

func (o *GetLolPerksV1InventoryOK) Error() string {
	return fmt.Sprintf("[GET /lol-perks/v1/inventory][%d] getLolPerksV1InventoryOK  %+v", 200, o.Payload)
}

func (o *GetLolPerksV1InventoryOK) GetPayload() *models.LolPerksPlayerInventory {
	return o.Payload
}

func (o *GetLolPerksV1InventoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPerksPlayerInventory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
