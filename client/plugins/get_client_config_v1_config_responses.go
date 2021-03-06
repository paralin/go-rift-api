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

// GetClientConfigV1ConfigReader is a Reader for the GetClientConfigV1Config structure.
type GetClientConfigV1ConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientConfigV1ConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClientConfigV1ConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClientConfigV1ConfigOK creates a GetClientConfigV1ConfigOK with default headers values
func NewGetClientConfigV1ConfigOK() *GetClientConfigV1ConfigOK {
	return &GetClientConfigV1ConfigOK{}
}

/*GetClientConfigV1ConfigOK handles this case with default header values.

Successful response
*/
type GetClientConfigV1ConfigOK struct {
	Payload map[string]interface{}
}

func (o *GetClientConfigV1ConfigOK) Error() string {
	return fmt.Sprintf("[GET /client-config/v1/config][%d] getClientConfigV1ConfigOK  %+v", 200, o.Payload)
}

func (o *GetClientConfigV1ConfigOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetClientConfigV1ConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
