// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolGameflowV1SessionEventReader is a Reader for the PostLolGameflowV1SessionEvent structure.
type PostLolGameflowV1SessionEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1SessionEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolGameflowV1SessionEventNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1SessionEventNoContent creates a PostLolGameflowV1SessionEventNoContent with default headers values
func NewPostLolGameflowV1SessionEventNoContent() *PostLolGameflowV1SessionEventNoContent {
	return &PostLolGameflowV1SessionEventNoContent{}
}

/*PostLolGameflowV1SessionEventNoContent handles this case with default header values.

No content
*/
type PostLolGameflowV1SessionEventNoContent struct {
}

func (o *PostLolGameflowV1SessionEventNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/session/event][%d] postLolGameflowV1SessionEventNoContent ", 204)
}

func (o *PostLolGameflowV1SessionEventNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
