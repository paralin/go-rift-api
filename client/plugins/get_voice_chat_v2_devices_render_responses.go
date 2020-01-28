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

// GetVoiceChatV2DevicesRenderReader is a Reader for the GetVoiceChatV2DevicesRender structure.
type GetVoiceChatV2DevicesRenderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVoiceChatV2DevicesRenderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVoiceChatV2DevicesRenderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVoiceChatV2DevicesRenderOK creates a GetVoiceChatV2DevicesRenderOK with default headers values
func NewGetVoiceChatV2DevicesRenderOK() *GetVoiceChatV2DevicesRenderOK {
	return &GetVoiceChatV2DevicesRenderOK{}
}

/*GetVoiceChatV2DevicesRenderOK handles this case with default header values.

Successful response
*/
type GetVoiceChatV2DevicesRenderOK struct {
	Payload []*models.VoiceChatDeviceResource
}

func (o *GetVoiceChatV2DevicesRenderOK) Error() string {
	return fmt.Sprintf("[GET /voice-chat/v2/devices/render][%d] getVoiceChatV2DevicesRenderOK  %+v", 200, o.Payload)
}

func (o *GetVoiceChatV2DevicesRenderOK) GetPayload() []*models.VoiceChatDeviceResource {
	return o.Payload
}

func (o *GetVoiceChatV2DevicesRenderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}