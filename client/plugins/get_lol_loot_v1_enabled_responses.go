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

// GetLolLootV1EnabledReader is a Reader for the GetLolLootV1Enabled structure.
type GetLolLootV1EnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLootV1EnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLootV1EnabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLootV1EnabledOK creates a GetLolLootV1EnabledOK with default headers values
func NewGetLolLootV1EnabledOK() *GetLolLootV1EnabledOK {
	return &GetLolLootV1EnabledOK{}
}

/*GetLolLootV1EnabledOK handles this case with default header values.

Successful response
*/
type GetLolLootV1EnabledOK struct {
	Payload bool
}

func (o *GetLolLootV1EnabledOK) Error() string {
	return fmt.Sprintf("[GET /lol-loot/v1/enabled][%d] getLolLootV1EnabledOK  %+v", 200, o.Payload)
}

func (o *GetLolLootV1EnabledOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolLootV1EnabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
