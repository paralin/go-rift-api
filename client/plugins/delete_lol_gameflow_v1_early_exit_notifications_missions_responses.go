// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLolGameflowV1EarlyExitNotificationsMissionsReader is a Reader for the DeleteLolGameflowV1EarlyExitNotificationsMissions structure.
type DeleteLolGameflowV1EarlyExitNotificationsMissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolGameflowV1EarlyExitNotificationsMissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent creates a DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent with default headers values
func NewDeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent() *DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent {
	return &DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent{}
}

/*DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent handles this case with default header values.

No content
*/
type DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent struct {
}

func (o *DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lol-gameflow/v1/early-exit-notifications/missions][%d] deleteLolGameflowV1EarlyExitNotificationsMissionsNoContent ", 204)
}

func (o *DeleteLolGameflowV1EarlyExitNotificationsMissionsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
