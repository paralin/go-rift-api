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

// GetLolLootV1LootItemsReader is a Reader for the GetLolLootV1LootItems structure.
type GetLolLootV1LootItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLootV1LootItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLootV1LootItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLootV1LootItemsOK creates a GetLolLootV1LootItemsOK with default headers values
func NewGetLolLootV1LootItemsOK() *GetLolLootV1LootItemsOK {
	return &GetLolLootV1LootItemsOK{}
}

/*GetLolLootV1LootItemsOK handles this case with default header values.

Successful response
*/
type GetLolLootV1LootItemsOK struct {
	Payload []*models.LolLootLootItem
}

func (o *GetLolLootV1LootItemsOK) Error() string {
	return fmt.Sprintf("[GET /lol-loot/v1/loot-items][%d] getLolLootV1LootItemsOK  %+v", 200, o.Payload)
}

func (o *GetLolLootV1LootItemsOK) GetPayload() []*models.LolLootLootItem {
	return o.Payload
}

func (o *GetLolLootV1LootItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
