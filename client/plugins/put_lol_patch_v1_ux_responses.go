// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutLolPatchV1UxReader is a Reader for the PutLolPatchV1Ux structure.
type PutLolPatchV1UxReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolPatchV1UxReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLolPatchV1UxNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolPatchV1UxNoContent creates a PutLolPatchV1UxNoContent with default headers values
func NewPutLolPatchV1UxNoContent() *PutLolPatchV1UxNoContent {
	return &PutLolPatchV1UxNoContent{}
}

/*PutLolPatchV1UxNoContent handles this case with default header values.

No content
*/
type PutLolPatchV1UxNoContent struct {
}

func (o *PutLolPatchV1UxNoContent) Error() string {
	return fmt.Sprintf("[PUT /lol-patch/v1/ux][%d] putLolPatchV1UxNoContent ", 204)
}

func (o *PutLolPatchV1UxNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
