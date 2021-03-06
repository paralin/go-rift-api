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

// PatchLolChampSelectV1SessionActionsByIDReader is a Reader for the PatchLolChampSelectV1SessionActionsByID structure.
type PatchLolChampSelectV1SessionActionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLolChampSelectV1SessionActionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLolChampSelectV1SessionActionsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLolChampSelectV1SessionActionsByIDOK creates a PatchLolChampSelectV1SessionActionsByIDOK with default headers values
func NewPatchLolChampSelectV1SessionActionsByIDOK() *PatchLolChampSelectV1SessionActionsByIDOK {
	return &PatchLolChampSelectV1SessionActionsByIDOK{}
}

/*PatchLolChampSelectV1SessionActionsByIDOK handles this case with default header values.

Successful response
*/
type PatchLolChampSelectV1SessionActionsByIDOK struct {
	Payload interface{}
}

func (o *PatchLolChampSelectV1SessionActionsByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /lol-champ-select/v1/session/actions/{id}][%d] patchLolChampSelectV1SessionActionsByIdOK  %+v", 200, o.Payload)
}

func (o *PatchLolChampSelectV1SessionActionsByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchLolChampSelectV1SessionActionsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
