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

// PutLolPerksV1SettingsReader is a Reader for the PutLolPerksV1Settings structure.
type PutLolPerksV1SettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolPerksV1SettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolPerksV1SettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolPerksV1SettingsOK creates a PutLolPerksV1SettingsOK with default headers values
func NewPutLolPerksV1SettingsOK() *PutLolPerksV1SettingsOK {
	return &PutLolPerksV1SettingsOK{}
}

/*PutLolPerksV1SettingsOK handles this case with default header values.

Successful response
*/
type PutLolPerksV1SettingsOK struct {
	Payload interface{}
}

func (o *PutLolPerksV1SettingsOK) Error() string {
	return fmt.Sprintf("[PUT /lol-perks/v1/settings][%d] putLolPerksV1SettingsOK  %+v", 200, o.Payload)
}

func (o *PutLolPerksV1SettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PutLolPerksV1SettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
