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

// GetLolLobbyV1PartiesPlayerReader is a Reader for the GetLolLobbyV1PartiesPlayer structure.
type GetLolLobbyV1PartiesPlayerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyV1PartiesPlayerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyV1PartiesPlayerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyV1PartiesPlayerOK creates a GetLolLobbyV1PartiesPlayerOK with default headers values
func NewGetLolLobbyV1PartiesPlayerOK() *GetLolLobbyV1PartiesPlayerOK {
	return &GetLolLobbyV1PartiesPlayerOK{}
}

/*GetLolLobbyV1PartiesPlayerOK handles this case with default header values.

Successful response
*/
type GetLolLobbyV1PartiesPlayerOK struct {
	Payload *models.LolLobbyPlayerDto
}

func (o *GetLolLobbyV1PartiesPlayerOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby/v1/parties/player][%d] getLolLobbyV1PartiesPlayerOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyV1PartiesPlayerOK) GetPayload() *models.LolLobbyPlayerDto {
	return o.Payload
}

func (o *GetLolLobbyV1PartiesPlayerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLobbyPlayerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
