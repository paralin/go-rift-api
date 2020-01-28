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

// PutLolChatV1FriendGroupsOrderReader is a Reader for the PutLolChatV1FriendGroupsOrder structure.
type PutLolChatV1FriendGroupsOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolChatV1FriendGroupsOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolChatV1FriendGroupsOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolChatV1FriendGroupsOrderOK creates a PutLolChatV1FriendGroupsOrderOK with default headers values
func NewPutLolChatV1FriendGroupsOrderOK() *PutLolChatV1FriendGroupsOrderOK {
	return &PutLolChatV1FriendGroupsOrderOK{}
}

/*PutLolChatV1FriendGroupsOrderOK handles this case with default header values.

Successful response
*/
type PutLolChatV1FriendGroupsOrderOK struct {
	Payload interface{}
}

func (o *PutLolChatV1FriendGroupsOrderOK) Error() string {
	return fmt.Sprintf("[PUT /lol-chat/v1/friend-groups/order][%d] putLolChatV1FriendGroupsOrderOK  %+v", 200, o.Payload)
}

func (o *PutLolChatV1FriendGroupsOrderOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PutLolChatV1FriendGroupsOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
