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

// GetLolPatchV1GameVersionReader is a Reader for the GetLolPatchV1GameVersion structure.
type GetLolPatchV1GameVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPatchV1GameVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPatchV1GameVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPatchV1GameVersionOK creates a GetLolPatchV1GameVersionOK with default headers values
func NewGetLolPatchV1GameVersionOK() *GetLolPatchV1GameVersionOK {
	return &GetLolPatchV1GameVersionOK{}
}

/*GetLolPatchV1GameVersionOK handles this case with default header values.

Successful response
*/
type GetLolPatchV1GameVersionOK struct {
	Payload string
}

func (o *GetLolPatchV1GameVersionOK) Error() string {
	return fmt.Sprintf("[GET /lol-patch/v1/game-version][%d] getLolPatchV1GameVersionOK  %+v", 200, o.Payload)
}

func (o *GetLolPatchV1GameVersionOK) GetPayload() string {
	return o.Payload
}

func (o *GetLolPatchV1GameVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
