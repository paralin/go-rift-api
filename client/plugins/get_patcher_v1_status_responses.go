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

// GetPatcherV1StatusReader is a Reader for the GetPatcherV1Status structure.
type GetPatcherV1StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPatcherV1StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPatcherV1StatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPatcherV1StatusOK creates a GetPatcherV1StatusOK with default headers values
func NewGetPatcherV1StatusOK() *GetPatcherV1StatusOK {
	return &GetPatcherV1StatusOK{}
}

/*GetPatcherV1StatusOK handles this case with default header values.

Successful response
*/
type GetPatcherV1StatusOK struct {
	Payload *models.PatcherStatus
}

func (o *GetPatcherV1StatusOK) Error() string {
	return fmt.Sprintf("[GET /patcher/v1/status][%d] getPatcherV1StatusOK  %+v", 200, o.Payload)
}

func (o *GetPatcherV1StatusOK) GetPayload() *models.PatcherStatus {
	return o.Payload
}

func (o *GetPatcherV1StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PatcherStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
