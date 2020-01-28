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

// GetGcloudVoiceChatV2DevicesCaptureReader is a Reader for the GetGcloudVoiceChatV2DevicesCapture structure.
type GetGcloudVoiceChatV2DevicesCaptureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcloudVoiceChatV2DevicesCaptureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGcloudVoiceChatV2DevicesCaptureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGcloudVoiceChatV2DevicesCaptureOK creates a GetGcloudVoiceChatV2DevicesCaptureOK with default headers values
func NewGetGcloudVoiceChatV2DevicesCaptureOK() *GetGcloudVoiceChatV2DevicesCaptureOK {
	return &GetGcloudVoiceChatV2DevicesCaptureOK{}
}

/*GetGcloudVoiceChatV2DevicesCaptureOK handles this case with default header values.

Successful response
*/
type GetGcloudVoiceChatV2DevicesCaptureOK struct {
	Payload []*models.GcloudVoiceChatDeviceResource
}

func (o *GetGcloudVoiceChatV2DevicesCaptureOK) Error() string {
	return fmt.Sprintf("[GET /gcloud-voice-chat/v2/devices/capture][%d] getGcloudVoiceChatV2DevicesCaptureOK  %+v", 200, o.Payload)
}

func (o *GetGcloudVoiceChatV2DevicesCaptureOK) GetPayload() []*models.GcloudVoiceChatDeviceResource {
	return o.Payload
}

func (o *GetGcloudVoiceChatV2DevicesCaptureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}