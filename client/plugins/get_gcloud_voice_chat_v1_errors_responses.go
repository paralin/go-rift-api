// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// GetGcloudVoiceChatV1ErrorsReader is a Reader for the GetGcloudVoiceChatV1Errors structure.
type GetGcloudVoiceChatV1ErrorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcloudVoiceChatV1ErrorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGcloudVoiceChatV1ErrorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGcloudVoiceChatV1ErrorsOK creates a GetGcloudVoiceChatV1ErrorsOK with default headers values
func NewGetGcloudVoiceChatV1ErrorsOK() *GetGcloudVoiceChatV1ErrorsOK {
	return &GetGcloudVoiceChatV1ErrorsOK{}
}

/*GetGcloudVoiceChatV1ErrorsOK handles this case with default header values.

Successful response
*/
type GetGcloudVoiceChatV1ErrorsOK struct {
	Payload []*models.GcloudVoiceChatVoiceErrorCounterResource
}

func (o *GetGcloudVoiceChatV1ErrorsOK) Error() string {
	return fmt.Sprintf("[GET /gcloud-voice-chat/v1/errors][%d] getGcloudVoiceChatV1ErrorsOK  %+v", 200, o.Payload)
}

func (o *GetGcloudVoiceChatV1ErrorsOK) GetPayload() []*models.GcloudVoiceChatVoiceErrorCounterResource {
	return o.Payload
}

func (o *GetGcloudVoiceChatV1ErrorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
