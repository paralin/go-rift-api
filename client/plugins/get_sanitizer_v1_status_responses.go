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

// GetSanitizerV1StatusReader is a Reader for the GetSanitizerV1Status structure.
type GetSanitizerV1StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSanitizerV1StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSanitizerV1StatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSanitizerV1StatusOK creates a GetSanitizerV1StatusOK with default headers values
func NewGetSanitizerV1StatusOK() *GetSanitizerV1StatusOK {
	return &GetSanitizerV1StatusOK{}
}

/*GetSanitizerV1StatusOK handles this case with default header values.

Successful response
*/
type GetSanitizerV1StatusOK struct {
	Payload *models.SanitizerSanitizerStatus
}

func (o *GetSanitizerV1StatusOK) Error() string {
	return fmt.Sprintf("[GET /sanitizer/v1/status][%d] getSanitizerV1StatusOK  %+v", 200, o.Payload)
}

func (o *GetSanitizerV1StatusOK) GetPayload() *models.SanitizerSanitizerStatus {
	return o.Payload
}

func (o *GetSanitizerV1StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SanitizerSanitizerStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
