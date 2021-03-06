// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolGameflowV1ExtraGameClientArgsReader is a Reader for the PostLolGameflowV1ExtraGameClientArgs structure.
type PostLolGameflowV1ExtraGameClientArgsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1ExtraGameClientArgsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolGameflowV1ExtraGameClientArgsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1ExtraGameClientArgsNoContent creates a PostLolGameflowV1ExtraGameClientArgsNoContent with default headers values
func NewPostLolGameflowV1ExtraGameClientArgsNoContent() *PostLolGameflowV1ExtraGameClientArgsNoContent {
	return &PostLolGameflowV1ExtraGameClientArgsNoContent{}
}

/*PostLolGameflowV1ExtraGameClientArgsNoContent handles this case with default header values.

No content
*/
type PostLolGameflowV1ExtraGameClientArgsNoContent struct {
}

func (o *PostLolGameflowV1ExtraGameClientArgsNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/extra-game-client-args][%d] postLolGameflowV1ExtraGameClientArgsNoContent ", 204)
}

func (o *PostLolGameflowV1ExtraGameClientArgsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
