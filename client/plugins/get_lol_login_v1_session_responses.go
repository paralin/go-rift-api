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

// GetLolLoginV1SessionReader is a Reader for the GetLolLoginV1Session structure.
type GetLolLoginV1SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLoginV1SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLoginV1SessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLoginV1SessionOK creates a GetLolLoginV1SessionOK with default headers values
func NewGetLolLoginV1SessionOK() *GetLolLoginV1SessionOK {
	return &GetLolLoginV1SessionOK{}
}

/*GetLolLoginV1SessionOK handles this case with default header values.

Successful response
*/
type GetLolLoginV1SessionOK struct {
	Payload *models.LolLoginLoginSession
}

func (o *GetLolLoginV1SessionOK) Error() string {
	return fmt.Sprintf("[GET /lol-login/v1/session][%d] getLolLoginV1SessionOK  %+v", 200, o.Payload)
}

func (o *GetLolLoginV1SessionOK) GetPayload() *models.LolLoginLoginSession {
	return o.Payload
}

func (o *GetLolLoginV1SessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLoginLoginSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
