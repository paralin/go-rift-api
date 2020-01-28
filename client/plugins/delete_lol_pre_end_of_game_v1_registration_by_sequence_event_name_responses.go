// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameReader is a Reader for the DeleteLolPreEndOfGameV1RegistrationBySequenceEventName structure.
type DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent creates a DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent with default headers values
func NewDeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent() *DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent {
	return &DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent{}
}

/*DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent handles this case with default header values.

No content
*/
type DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent struct {
}

func (o *DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lol-pre-end-of-game/v1/registration/{sequenceEventName}][%d] deleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent ", 204)
}

func (o *DeleteLolPreEndOfGameV1RegistrationBySequenceEventNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
