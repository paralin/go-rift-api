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

// GetLolInventoryV1SignedInventorySimpleReader is a Reader for the GetLolInventoryV1SignedInventorySimple structure.
type GetLolInventoryV1SignedInventorySimpleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolInventoryV1SignedInventorySimpleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolInventoryV1SignedInventorySimpleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolInventoryV1SignedInventorySimpleOK creates a GetLolInventoryV1SignedInventorySimpleOK with default headers values
func NewGetLolInventoryV1SignedInventorySimpleOK() *GetLolInventoryV1SignedInventorySimpleOK {
	return &GetLolInventoryV1SignedInventorySimpleOK{}
}

/*GetLolInventoryV1SignedInventorySimpleOK handles this case with default header values.

Successful response
*/
type GetLolInventoryV1SignedInventorySimpleOK struct {
	Payload string
}

func (o *GetLolInventoryV1SignedInventorySimpleOK) Error() string {
	return fmt.Sprintf("[GET /lol-inventory/v1/signedInventory/simple][%d] getLolInventoryV1SignedInventorySimpleOK  %+v", 200, o.Payload)
}

func (o *GetLolInventoryV1SignedInventorySimpleOK) GetPayload() string {
	return o.Payload
}

func (o *GetLolInventoryV1SignedInventorySimpleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
