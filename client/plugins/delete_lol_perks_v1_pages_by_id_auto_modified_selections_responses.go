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

// DeleteLolPerksV1PagesByIDAutoModifiedSelectionsReader is a Reader for the DeleteLolPerksV1PagesByIDAutoModifiedSelections structure.
type DeleteLolPerksV1PagesByIDAutoModifiedSelectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolPerksV1PagesByIDAutoModifiedSelectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK creates a DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK with default headers values
func NewDeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK() *DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK {
	return &DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK{}
}

/*DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK handles this case with default header values.

Successful response
*/
type DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK struct {
	Payload interface{}
}

func (o *DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-perks/v1/pages/{id}/auto-modified-selections][%d] deleteLolPerksV1PagesByIdAutoModifiedSelectionsOK  %+v", 200, o.Payload)
}

func (o *DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteLolPerksV1PagesByIDAutoModifiedSelectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
