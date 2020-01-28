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

// GetLolLobbyV1PartiesGamemodeReader is a Reader for the GetLolLobbyV1PartiesGamemode structure.
type GetLolLobbyV1PartiesGamemodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyV1PartiesGamemodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyV1PartiesGamemodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyV1PartiesGamemodeOK creates a GetLolLobbyV1PartiesGamemodeOK with default headers values
func NewGetLolLobbyV1PartiesGamemodeOK() *GetLolLobbyV1PartiesGamemodeOK {
	return &GetLolLobbyV1PartiesGamemodeOK{}
}

/*GetLolLobbyV1PartiesGamemodeOK handles this case with default header values.

Successful response
*/
type GetLolLobbyV1PartiesGamemodeOK struct {
	Payload *models.LolLobbyGameModeDto
}

func (o *GetLolLobbyV1PartiesGamemodeOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby/v1/parties/gamemode][%d] getLolLobbyV1PartiesGamemodeOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyV1PartiesGamemodeOK) GetPayload() *models.LolLobbyGameModeDto {
	return o.Payload
}

func (o *GetLolLobbyV1PartiesGamemodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLobbyGameModeDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}