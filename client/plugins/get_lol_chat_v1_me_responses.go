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

// GetLolChatV1MeReader is a Reader for the GetLolChatV1Me structure.
type GetLolChatV1MeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChatV1MeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChatV1MeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChatV1MeOK creates a GetLolChatV1MeOK with default headers values
func NewGetLolChatV1MeOK() *GetLolChatV1MeOK {
	return &GetLolChatV1MeOK{}
}

/*GetLolChatV1MeOK handles this case with default header values.

Successful response
*/
type GetLolChatV1MeOK struct {
	Payload *models.LolChatUserResource
}

func (o *GetLolChatV1MeOK) Error() string {
	return fmt.Sprintf("[GET /lol-chat/v1/me][%d] getLolChatV1MeOK  %+v", 200, o.Payload)
}

func (o *GetLolChatV1MeOK) GetPayload() *models.LolChatUserResource {
	return o.Payload
}

func (o *GetLolChatV1MeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChatUserResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
