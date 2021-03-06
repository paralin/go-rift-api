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

// DeleteLolChatV1SettingsByKeyReader is a Reader for the DeleteLolChatV1SettingsByKey structure.
type DeleteLolChatV1SettingsByKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolChatV1SettingsByKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolChatV1SettingsByKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolChatV1SettingsByKeyOK creates a DeleteLolChatV1SettingsByKeyOK with default headers values
func NewDeleteLolChatV1SettingsByKeyOK() *DeleteLolChatV1SettingsByKeyOK {
	return &DeleteLolChatV1SettingsByKeyOK{}
}

/*DeleteLolChatV1SettingsByKeyOK handles this case with default header values.

Successful response
*/
type DeleteLolChatV1SettingsByKeyOK struct {
	Payload interface{}
}

func (o *DeleteLolChatV1SettingsByKeyOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-chat/v1/settings/{key}][%d] deleteLolChatV1SettingsByKeyOK  %+v", 200, o.Payload)
}

func (o *DeleteLolChatV1SettingsByKeyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteLolChatV1SettingsByKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
