// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolGameflowV1GameflowMetadataRegistrationStatusReader is a Reader for the PostLolGameflowV1GameflowMetadataRegistrationStatus structure.
type PostLolGameflowV1GameflowMetadataRegistrationStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1GameflowMetadataRegistrationStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolGameflowV1GameflowMetadataRegistrationStatusNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1GameflowMetadataRegistrationStatusNoContent creates a PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent with default headers values
func NewPostLolGameflowV1GameflowMetadataRegistrationStatusNoContent() *PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent {
	return &PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent{}
}

/*PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent handles this case with default header values.

No content
*/
type PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent struct {
}

func (o *PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/gameflow-metadata/registration-status][%d] postLolGameflowV1GameflowMetadataRegistrationStatusNoContent ", 204)
}

func (o *PostLolGameflowV1GameflowMetadataRegistrationStatusNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
