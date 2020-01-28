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

// PatchLolChampSelectLegacyV1SessionActionsByIDReader is a Reader for the PatchLolChampSelectLegacyV1SessionActionsByID structure.
type PatchLolChampSelectLegacyV1SessionActionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLolChampSelectLegacyV1SessionActionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLolChampSelectLegacyV1SessionActionsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLolChampSelectLegacyV1SessionActionsByIDOK creates a PatchLolChampSelectLegacyV1SessionActionsByIDOK with default headers values
func NewPatchLolChampSelectLegacyV1SessionActionsByIDOK() *PatchLolChampSelectLegacyV1SessionActionsByIDOK {
	return &PatchLolChampSelectLegacyV1SessionActionsByIDOK{}
}

/*PatchLolChampSelectLegacyV1SessionActionsByIDOK handles this case with default header values.

Successful response
*/
type PatchLolChampSelectLegacyV1SessionActionsByIDOK struct {
	Payload interface{}
}

func (o *PatchLolChampSelectLegacyV1SessionActionsByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /lol-champ-select-legacy/v1/session/actions/{id}][%d] patchLolChampSelectLegacyV1SessionActionsByIdOK  %+v", 200, o.Payload)
}

func (o *PatchLolChampSelectLegacyV1SessionActionsByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchLolChampSelectLegacyV1SessionActionsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}