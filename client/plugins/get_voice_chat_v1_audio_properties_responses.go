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

// GetVoiceChatV1AudioPropertiesReader is a Reader for the GetVoiceChatV1AudioProperties structure.
type GetVoiceChatV1AudioPropertiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoiceChatV1AudioPropertiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoiceChatV1AudioPropertiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVoiceChatV1AudioPropertiesOK creates a GetVoiceChatV1AudioPropertiesOK with default headers values
func NewGetVoiceChatV1AudioPropertiesOK() *GetVoiceChatV1AudioPropertiesOK {
	return &GetVoiceChatV1AudioPropertiesOK{}
}

/*GetVoiceChatV1AudioPropertiesOK handles this case with default header values.

Successful response
*/
type GetVoiceChatV1AudioPropertiesOK struct {
	Payload *models.VoiceChatAudioPropertiesResource
}

func (o *GetVoiceChatV1AudioPropertiesOK) Error() string {
	return fmt.Sprintf("[GET /voice-chat/v1/audio-properties][%d] getVoiceChatV1AudioPropertiesOK  %+v", 200, o.Payload)
}

func (o *GetVoiceChatV1AudioPropertiesOK) GetPayload() *models.VoiceChatAudioPropertiesResource {
	return o.Payload
}

func (o *GetVoiceChatV1AudioPropertiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VoiceChatAudioPropertiesResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
