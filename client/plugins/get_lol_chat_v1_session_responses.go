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

// GetLolChatV1SessionReader is a Reader for the GetLolChatV1Session structure.
type GetLolChatV1SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChatV1SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChatV1SessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChatV1SessionOK creates a GetLolChatV1SessionOK with default headers values
func NewGetLolChatV1SessionOK() *GetLolChatV1SessionOK {
	return &GetLolChatV1SessionOK{}
}

/*GetLolChatV1SessionOK handles this case with default header values.

Successful response
*/
type GetLolChatV1SessionOK struct {
	Payload *models.LolChatSessionResource
}

func (o *GetLolChatV1SessionOK) Error() string {
	return fmt.Sprintf("[GET /lol-chat/v1/session][%d] getLolChatV1SessionOK  %+v", 200, o.Payload)
}

func (o *GetLolChatV1SessionOK) GetPayload() *models.LolChatSessionResource {
	return o.Payload
}

func (o *GetLolChatV1SessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChatSessionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
