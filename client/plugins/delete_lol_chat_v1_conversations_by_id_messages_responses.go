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

// DeleteLolChatV1ConversationsByIDMessagesReader is a Reader for the DeleteLolChatV1ConversationsByIDMessages structure.
type DeleteLolChatV1ConversationsByIDMessagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolChatV1ConversationsByIDMessagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolChatV1ConversationsByIDMessagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolChatV1ConversationsByIDMessagesOK creates a DeleteLolChatV1ConversationsByIDMessagesOK with default headers values
func NewDeleteLolChatV1ConversationsByIDMessagesOK() *DeleteLolChatV1ConversationsByIDMessagesOK {
	return &DeleteLolChatV1ConversationsByIDMessagesOK{}
}

/*DeleteLolChatV1ConversationsByIDMessagesOK handles this case with default header values.

Successful response
*/
type DeleteLolChatV1ConversationsByIDMessagesOK struct {
	Payload interface{}
}

func (o *DeleteLolChatV1ConversationsByIDMessagesOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-chat/v1/conversations/{id}/messages][%d] deleteLolChatV1ConversationsByIdMessagesOK  %+v", 200, o.Payload)
}

func (o *DeleteLolChatV1ConversationsByIDMessagesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteLolChatV1ConversationsByIDMessagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
