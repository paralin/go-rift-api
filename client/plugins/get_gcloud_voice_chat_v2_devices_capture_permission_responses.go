// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetGcloudVoiceChatV2DevicesCapturePermissionReader is a Reader for the GetGcloudVoiceChatV2DevicesCapturePermission structure.
type GetGcloudVoiceChatV2DevicesCapturePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGcloudVoiceChatV2DevicesCapturePermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGcloudVoiceChatV2DevicesCapturePermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetGcloudVoiceChatV2DevicesCapturePermissionOK creates a GetGcloudVoiceChatV2DevicesCapturePermissionOK with default headers values
func NewGetGcloudVoiceChatV2DevicesCapturePermissionOK() *GetGcloudVoiceChatV2DevicesCapturePermissionOK {
	return &GetGcloudVoiceChatV2DevicesCapturePermissionOK{}
}

/*GetGcloudVoiceChatV2DevicesCapturePermissionOK handles this case with default header values.

Successful response
*/
type GetGcloudVoiceChatV2DevicesCapturePermissionOK struct {
	Payload string
}

func (o *GetGcloudVoiceChatV2DevicesCapturePermissionOK) Error() string {
	return fmt.Sprintf("[GET /gcloud-voice-chat/v2/devices/capture/permission][%d] getGcloudVoiceChatV2DevicesCapturePermissionOK  %+v", 200, o.Payload)
}

func (o *GetGcloudVoiceChatV2DevicesCapturePermissionOK) GetPayload() string {
	return o.Payload
}

func (o *GetGcloudVoiceChatV2DevicesCapturePermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
