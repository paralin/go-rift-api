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

// GetLolChatV1BlockedPlayersByIDReader is a Reader for the GetLolChatV1BlockedPlayersByID structure.
type GetLolChatV1BlockedPlayersByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChatV1BlockedPlayersByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChatV1BlockedPlayersByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChatV1BlockedPlayersByIDOK creates a GetLolChatV1BlockedPlayersByIDOK with default headers values
func NewGetLolChatV1BlockedPlayersByIDOK() *GetLolChatV1BlockedPlayersByIDOK {
	return &GetLolChatV1BlockedPlayersByIDOK{}
}

/*GetLolChatV1BlockedPlayersByIDOK handles this case with default header values.

Successful response
*/
type GetLolChatV1BlockedPlayersByIDOK struct {
	Payload *models.LolChatBlockedPlayerResource
}

func (o *GetLolChatV1BlockedPlayersByIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-chat/v1/blocked-players/{id}][%d] getLolChatV1BlockedPlayersByIdOK  %+v", 200, o.Payload)
}

func (o *GetLolChatV1BlockedPlayersByIDOK) GetPayload() *models.LolChatBlockedPlayerResource {
	return o.Payload
}

func (o *GetLolChatV1BlockedPlayersByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChatBlockedPlayerResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}