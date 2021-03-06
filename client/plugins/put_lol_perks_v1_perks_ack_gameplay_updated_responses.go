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

// PutLolPerksV1PerksAckGameplayUpdatedReader is a Reader for the PutLolPerksV1PerksAckGameplayUpdated structure.
type PutLolPerksV1PerksAckGameplayUpdatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolPerksV1PerksAckGameplayUpdatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolPerksV1PerksAckGameplayUpdatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolPerksV1PerksAckGameplayUpdatedOK creates a PutLolPerksV1PerksAckGameplayUpdatedOK with default headers values
func NewPutLolPerksV1PerksAckGameplayUpdatedOK() *PutLolPerksV1PerksAckGameplayUpdatedOK {
	return &PutLolPerksV1PerksAckGameplayUpdatedOK{}
}

/*PutLolPerksV1PerksAckGameplayUpdatedOK handles this case with default header values.

Successful response
*/
type PutLolPerksV1PerksAckGameplayUpdatedOK struct {
	Payload interface{}
}

func (o *PutLolPerksV1PerksAckGameplayUpdatedOK) Error() string {
	return fmt.Sprintf("[PUT /lol-perks/v1/perks/ack-gameplay-updated][%d] putLolPerksV1PerksAckGameplayUpdatedOK  %+v", 200, o.Payload)
}

func (o *PutLolPerksV1PerksAckGameplayUpdatedOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PutLolPerksV1PerksAckGameplayUpdatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
