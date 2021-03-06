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

// GetLolPlatformConfigV1NamespacesByNsReader is a Reader for the GetLolPlatformConfigV1NamespacesByNs structure.
type GetLolPlatformConfigV1NamespacesByNsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPlatformConfigV1NamespacesByNsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPlatformConfigV1NamespacesByNsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPlatformConfigV1NamespacesByNsOK creates a GetLolPlatformConfigV1NamespacesByNsOK with default headers values
func NewGetLolPlatformConfigV1NamespacesByNsOK() *GetLolPlatformConfigV1NamespacesByNsOK {
	return &GetLolPlatformConfigV1NamespacesByNsOK{}
}

/*GetLolPlatformConfigV1NamespacesByNsOK handles this case with default header values.

Successful response
*/
type GetLolPlatformConfigV1NamespacesByNsOK struct {
	Payload interface{}
}

func (o *GetLolPlatformConfigV1NamespacesByNsOK) Error() string {
	return fmt.Sprintf("[GET /lol-platform-config/v1/namespaces/{ns}][%d] getLolPlatformConfigV1NamespacesByNsOK  %+v", 200, o.Payload)
}

func (o *GetLolPlatformConfigV1NamespacesByNsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolPlatformConfigV1NamespacesByNsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
