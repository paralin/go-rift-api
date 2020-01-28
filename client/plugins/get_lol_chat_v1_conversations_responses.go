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

// GetLolChatV1ConversationsReader is a Reader for the GetLolChatV1Conversations structure.
type GetLolChatV1ConversationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChatV1ConversationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChatV1ConversationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChatV1ConversationsOK creates a GetLolChatV1ConversationsOK with default headers values
func NewGetLolChatV1ConversationsOK() *GetLolChatV1ConversationsOK {
	return &GetLolChatV1ConversationsOK{}
}

/*GetLolChatV1ConversationsOK handles this case with default header values.

Successful response
*/
type GetLolChatV1ConversationsOK struct {
	Payload []*models.LolChatConversationResource
}

func (o *GetLolChatV1ConversationsOK) Error() string {
	return fmt.Sprintf("[GET /lol-chat/v1/conversations][%d] getLolChatV1ConversationsOK  %+v", 200, o.Payload)
}

func (o *GetLolChatV1ConversationsOK) GetPayload() []*models.LolChatConversationResource {
	return o.Payload
}

func (o *GetLolChatV1ConversationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}