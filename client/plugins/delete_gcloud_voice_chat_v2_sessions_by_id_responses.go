// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteGcloudVoiceChatV2SessionsByIDReader is a Reader for the DeleteGcloudVoiceChatV2SessionsByID structure.
type DeleteGcloudVoiceChatV2SessionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteGcloudVoiceChatV2SessionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteGcloudVoiceChatV2SessionsByIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteGcloudVoiceChatV2SessionsByIDNoContent creates a DeleteGcloudVoiceChatV2SessionsByIDNoContent with default headers values
func NewDeleteGcloudVoiceChatV2SessionsByIDNoContent() *DeleteGcloudVoiceChatV2SessionsByIDNoContent {
	return &DeleteGcloudVoiceChatV2SessionsByIDNoContent{}
}

/*DeleteGcloudVoiceChatV2SessionsByIDNoContent handles this case with default header values.

No content
*/
type DeleteGcloudVoiceChatV2SessionsByIDNoContent struct {
}

func (o *DeleteGcloudVoiceChatV2SessionsByIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /gcloud-voice-chat/v2/sessions/{id}][%d] deleteGcloudVoiceChatV2SessionsByIdNoContent ", 204)
}

func (o *DeleteGcloudVoiceChatV2SessionsByIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
