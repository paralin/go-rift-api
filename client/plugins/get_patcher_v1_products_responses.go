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

// GetPatcherV1ProductsReader is a Reader for the GetPatcherV1Products structure.
type GetPatcherV1ProductsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatcherV1ProductsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatcherV1ProductsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatcherV1ProductsOK creates a GetPatcherV1ProductsOK with default headers values
func NewGetPatcherV1ProductsOK() *GetPatcherV1ProductsOK {
	return &GetPatcherV1ProductsOK{}
}

/*GetPatcherV1ProductsOK handles this case with default header values.

Successful response
*/
type GetPatcherV1ProductsOK struct {
	Payload []string
}

func (o *GetPatcherV1ProductsOK) Error() string {
	return fmt.Sprintf("[GET /patcher/v1/products][%d] getPatcherV1ProductsOK  %+v", 200, o.Payload)
}

func (o *GetPatcherV1ProductsOK) GetPayload() []string {
	return o.Payload
}

func (o *GetPatcherV1ProductsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
