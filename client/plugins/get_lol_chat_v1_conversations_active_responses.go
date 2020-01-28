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

// GetLolChatV1ConversationsActiveReader is a Reader for the GetLolChatV1ConversationsActive structure.
type GetLolChatV1ConversationsActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChatV1ConversationsActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChatV1ConversationsActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChatV1ConversationsActiveOK creates a GetLolChatV1ConversationsActiveOK with default headers values
func NewGetLolChatV1ConversationsActiveOK() *GetLolChatV1ConversationsActiveOK {
	return &GetLolChatV1ConversationsActiveOK{}
}

/*GetLolChatV1ConversationsActiveOK handles this case with default header values.

Successful response
*/
type GetLolChatV1ConversationsActiveOK struct {
	Payload *models.LolChatActiveConversationResource
}

func (o *GetLolChatV1ConversationsActiveOK) Error() string {
	return fmt.Sprintf("[GET /lol-chat/v1/conversations/active][%d] getLolChatV1ConversationsActiveOK  %+v", 200, o.Payload)
}

func (o *GetLolChatV1ConversationsActiveOK) GetPayload() *models.LolChatActiveConversationResource {
	return o.Payload
}

func (o *GetLolChatV1ConversationsActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChatActiveConversationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
