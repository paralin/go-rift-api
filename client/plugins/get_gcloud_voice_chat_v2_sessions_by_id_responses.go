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

// GetGcloudVoiceChatV2SessionsByIDReader is a Reader for the GetGcloudVoiceChatV2SessionsByID structure.
type GetGcloudVoiceChatV2SessionsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcloudVoiceChatV2SessionsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGcloudVoiceChatV2SessionsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGcloudVoiceChatV2SessionsByIDOK creates a GetGcloudVoiceChatV2SessionsByIDOK with default headers values
func NewGetGcloudVoiceChatV2SessionsByIDOK() *GetGcloudVoiceChatV2SessionsByIDOK {
	return &GetGcloudVoiceChatV2SessionsByIDOK{}
}

/*GetGcloudVoiceChatV2SessionsByIDOK handles this case with default header values.

Successful response
*/
type GetGcloudVoiceChatV2SessionsByIDOK struct {
	Payload *models.GcloudVoiceChatSessionResource
}

func (o *GetGcloudVoiceChatV2SessionsByIDOK) Error() string {
	return fmt.Sprintf("[GET /gcloud-voice-chat/v2/sessions/{id}][%d] getGcloudVoiceChatV2SessionsByIdOK  %+v", 200, o.Payload)
}

func (o *GetGcloudVoiceChatV2SessionsByIDOK) GetPayload() *models.GcloudVoiceChatSessionResource {
	return o.Payload
}

func (o *GetGcloudVoiceChatV2SessionsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GcloudVoiceChatSessionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
