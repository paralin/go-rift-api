// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteRiotMessagingServiceV1SessionReader is a Reader for the DeleteRiotMessagingServiceV1Session structure.
type DeleteRiotMessagingServiceV1SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRiotMessagingServiceV1SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteRiotMessagingServiceV1SessionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteRiotMessagingServiceV1SessionNoContent creates a DeleteRiotMessagingServiceV1SessionNoContent with default headers values
func NewDeleteRiotMessagingServiceV1SessionNoContent() *DeleteRiotMessagingServiceV1SessionNoContent {
	return &DeleteRiotMessagingServiceV1SessionNoContent{}
}

/*DeleteRiotMessagingServiceV1SessionNoContent handles this case with default header values.

No content
*/
type DeleteRiotMessagingServiceV1SessionNoContent struct {
}

func (o *DeleteRiotMessagingServiceV1SessionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /riot-messaging-service/v1/session][%d] deleteRiotMessagingServiceV1SessionNoContent ", 204)
}

func (o *DeleteRiotMessagingServiceV1SessionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
