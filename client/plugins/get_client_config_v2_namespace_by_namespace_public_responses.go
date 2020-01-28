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

// GetClientConfigV2NamespaceByNamespacePublicReader is a Reader for the GetClientConfigV2NamespaceByNamespacePublic structure.
type GetClientConfigV2NamespaceByNamespacePublicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClientConfigV2NamespaceByNamespacePublicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetClientConfigV2NamespaceByNamespacePublicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClientConfigV2NamespaceByNamespacePublicOK creates a GetClientConfigV2NamespaceByNamespacePublicOK with default headers values
func NewGetClientConfigV2NamespaceByNamespacePublicOK() *GetClientConfigV2NamespaceByNamespacePublicOK {
	return &GetClientConfigV2NamespaceByNamespacePublicOK{}
}

/*GetClientConfigV2NamespaceByNamespacePublicOK handles this case with default header values.

Successful response
*/
type GetClientConfigV2NamespaceByNamespacePublicOK struct {
	Payload map[string]interface{}
}

func (o *GetClientConfigV2NamespaceByNamespacePublicOK) Error() string {
	return fmt.Sprintf("[GET /client-config/v2/namespace/{namespace}/public][%d] getClientConfigV2NamespaceByNamespacePublicOK  %+v", 200, o.Payload)
}

func (o *GetClientConfigV2NamespaceByNamespacePublicOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *GetClientConfigV2NamespaceByNamespacePublicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
