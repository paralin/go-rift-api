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

// GetLolPerksV1PerksDisabledReader is a Reader for the GetLolPerksV1PerksDisabled structure.
type GetLolPerksV1PerksDisabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPerksV1PerksDisabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPerksV1PerksDisabledOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPerksV1PerksDisabledOK creates a GetLolPerksV1PerksDisabledOK with default headers values
func NewGetLolPerksV1PerksDisabledOK() *GetLolPerksV1PerksDisabledOK {
	return &GetLolPerksV1PerksDisabledOK{}
}

/*GetLolPerksV1PerksDisabledOK handles this case with default header values.

Successful response
*/
type GetLolPerksV1PerksDisabledOK struct {
	Payload []int32
}

func (o *GetLolPerksV1PerksDisabledOK) Error() string {
	return fmt.Sprintf("[GET /lol-perks/v1/perks/disabled][%d] getLolPerksV1PerksDisabledOK  %+v", 200, o.Payload)
}

func (o *GetLolPerksV1PerksDisabledOK) GetPayload() []int32 {
	return o.Payload
}

func (o *GetLolPerksV1PerksDisabledOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
