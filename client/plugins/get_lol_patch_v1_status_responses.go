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

// GetLolPatchV1StatusReader is a Reader for the GetLolPatchV1Status structure.
type GetLolPatchV1StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPatchV1StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPatchV1StatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPatchV1StatusOK creates a GetLolPatchV1StatusOK with default headers values
func NewGetLolPatchV1StatusOK() *GetLolPatchV1StatusOK {
	return &GetLolPatchV1StatusOK{}
}

/*GetLolPatchV1StatusOK handles this case with default header values.

Successful response
*/
type GetLolPatchV1StatusOK struct {
	Payload *models.LolPatchStatus
}

func (o *GetLolPatchV1StatusOK) Error() string {
	return fmt.Sprintf("[GET /lol-patch/v1/status][%d] getLolPatchV1StatusOK  %+v", 200, o.Payload)
}

func (o *GetLolPatchV1StatusOK) GetPayload() *models.LolPatchStatus {
	return o.Payload
}

func (o *GetLolPatchV1StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPatchStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
