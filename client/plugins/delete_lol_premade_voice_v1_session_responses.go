// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLolPremadeVoiceV1SessionReader is a Reader for the DeleteLolPremadeVoiceV1Session structure.
type DeleteLolPremadeVoiceV1SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolPremadeVoiceV1SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLolPremadeVoiceV1SessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolPremadeVoiceV1SessionNoContent creates a DeleteLolPremadeVoiceV1SessionNoContent with default headers values
func NewDeleteLolPremadeVoiceV1SessionNoContent() *DeleteLolPremadeVoiceV1SessionNoContent {
	return &DeleteLolPremadeVoiceV1SessionNoContent{}
}

/*DeleteLolPremadeVoiceV1SessionNoContent handles this case with default header values.

No content
*/
type DeleteLolPremadeVoiceV1SessionNoContent struct {
}

func (o *DeleteLolPremadeVoiceV1SessionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lol-premade-voice/v1/session][%d] deleteLolPremadeVoiceV1SessionNoContent ", 204)
}

func (o *DeleteLolPremadeVoiceV1SessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}