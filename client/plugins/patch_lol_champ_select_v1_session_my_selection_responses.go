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

// PatchLolChampSelectV1SessionMySelectionReader is a Reader for the PatchLolChampSelectV1SessionMySelection structure.
type PatchLolChampSelectV1SessionMySelectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLolChampSelectV1SessionMySelectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLolChampSelectV1SessionMySelectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLolChampSelectV1SessionMySelectionOK creates a PatchLolChampSelectV1SessionMySelectionOK with default headers values
func NewPatchLolChampSelectV1SessionMySelectionOK() *PatchLolChampSelectV1SessionMySelectionOK {
	return &PatchLolChampSelectV1SessionMySelectionOK{}
}

/*PatchLolChampSelectV1SessionMySelectionOK handles this case with default header values.

Successful response
*/
type PatchLolChampSelectV1SessionMySelectionOK struct {
	Payload interface{}
}

func (o *PatchLolChampSelectV1SessionMySelectionOK) Error() string {
	return fmt.Sprintf("[PATCH /lol-champ-select/v1/session/my-selection][%d] patchLolChampSelectV1SessionMySelectionOK  %+v", 200, o.Payload)
}

func (o *PatchLolChampSelectV1SessionMySelectionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchLolChampSelectV1SessionMySelectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
