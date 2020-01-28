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

// GetPatcherV1ProductsByProductIDTagsReader is a Reader for the GetPatcherV1ProductsByProductIDTags structure.
type GetPatcherV1ProductsByProductIDTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatcherV1ProductsByProductIDTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatcherV1ProductsByProductIDTagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatcherV1ProductsByProductIDTagsOK creates a GetPatcherV1ProductsByProductIDTagsOK with default headers values
func NewGetPatcherV1ProductsByProductIDTagsOK() *GetPatcherV1ProductsByProductIDTagsOK {
	return &GetPatcherV1ProductsByProductIDTagsOK{}
}

/*GetPatcherV1ProductsByProductIDTagsOK handles this case with default header values.

Successful response
*/
type GetPatcherV1ProductsByProductIDTagsOK struct {
	Payload map[string]string
}

func (o *GetPatcherV1ProductsByProductIDTagsOK) Error() string {
	return fmt.Sprintf("[GET /patcher/v1/products/{product-id}/tags][%d] getPatcherV1ProductsByProductIdTagsOK  %+v", 200, o.Payload)
}

func (o *GetPatcherV1ProductsByProductIDTagsOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *GetPatcherV1ProductsByProductIDTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
