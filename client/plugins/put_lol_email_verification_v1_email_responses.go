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

// PutLolEmailVerificationV1EmailReader is a Reader for the PutLolEmailVerificationV1Email structure.
type PutLolEmailVerificationV1EmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolEmailVerificationV1EmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolEmailVerificationV1EmailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolEmailVerificationV1EmailOK creates a PutLolEmailVerificationV1EmailOK with default headers values
func NewPutLolEmailVerificationV1EmailOK() *PutLolEmailVerificationV1EmailOK {
	return &PutLolEmailVerificationV1EmailOK{}
}

/*PutLolEmailVerificationV1EmailOK handles this case with default header values.

Successful response
*/
type PutLolEmailVerificationV1EmailOK struct {
	Payload interface{}
}

func (o *PutLolEmailVerificationV1EmailOK) Error() string {
	return fmt.Sprintf("[PUT /lol-email-verification/v1/email][%d] putLolEmailVerificationV1EmailOK  %+v", 200, o.Payload)
}

func (o *PutLolEmailVerificationV1EmailOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PutLolEmailVerificationV1EmailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
