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

// GetLolLootV1PlayerLootReader is a Reader for the GetLolLootV1PlayerLoot structure.
type GetLolLootV1PlayerLootReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLootV1PlayerLootReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLootV1PlayerLootOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLootV1PlayerLootOK creates a GetLolLootV1PlayerLootOK with default headers values
func NewGetLolLootV1PlayerLootOK() *GetLolLootV1PlayerLootOK {
	return &GetLolLootV1PlayerLootOK{}
}

/*GetLolLootV1PlayerLootOK handles this case with default header values.

Successful response
*/
type GetLolLootV1PlayerLootOK struct {
	Payload []*models.LolLootPlayerLoot
}

func (o *GetLolLootV1PlayerLootOK) Error() string {
	return fmt.Sprintf("[GET /lol-loot/v1/player-loot][%d] getLolLootV1PlayerLootOK  %+v", 200, o.Payload)
}

func (o *GetLolLootV1PlayerLootOK) GetPayload() []*models.LolLootPlayerLoot {
	return o.Payload
}

func (o *GetLolLootV1PlayerLootOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
