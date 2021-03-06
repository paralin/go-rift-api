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

// DeleteLolChatV1ConversationsByIDReader is a Reader for the DeleteLolChatV1ConversationsByID structure.
type DeleteLolChatV1ConversationsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolChatV1ConversationsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolChatV1ConversationsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolChatV1ConversationsByIDOK creates a DeleteLolChatV1ConversationsByIDOK with default headers values
func NewDeleteLolChatV1ConversationsByIDOK() *DeleteLolChatV1ConversationsByIDOK {
	return &DeleteLolChatV1ConversationsByIDOK{}
}

/*DeleteLolChatV1ConversationsByIDOK handles this case with default header values.

Successful response
*/
type DeleteLolChatV1ConversationsByIDOK struct {
	Payload interface{}
}

func (o *DeleteLolChatV1ConversationsByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-chat/v1/conversations/{id}][%d] deleteLolChatV1ConversationsByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteLolChatV1ConversationsByIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteLolChatV1ConversationsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
