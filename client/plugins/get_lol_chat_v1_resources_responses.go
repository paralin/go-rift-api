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

// GetLolChatV1ResourcesReader is a Reader for the GetLolChatV1Resources structure.
type GetLolChatV1ResourcesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChatV1ResourcesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChatV1ResourcesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChatV1ResourcesOK creates a GetLolChatV1ResourcesOK with default headers values
func NewGetLolChatV1ResourcesOK() *GetLolChatV1ResourcesOK {
	return &GetLolChatV1ResourcesOK{}
}

/*GetLolChatV1ResourcesOK handles this case with default header values.

Successful response
*/
type GetLolChatV1ResourcesOK struct {
	Payload *models.LolChatProductMetadataMap
}

func (o *GetLolChatV1ResourcesOK) Error() string {
	return fmt.Sprintf("[GET /lol-chat/v1/resources][%d] getLolChatV1ResourcesOK  %+v", 200, o.Payload)
}

func (o *GetLolChatV1ResourcesOK) GetPayload() *models.LolChatProductMetadataMap {
	return o.Payload
}

func (o *GetLolChatV1ResourcesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChatProductMetadataMap)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
