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

// GetLolLobbyV2LobbyCustomAvailableBotsReader is a Reader for the GetLolLobbyV2LobbyCustomAvailableBots structure.
type GetLolLobbyV2LobbyCustomAvailableBotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyV2LobbyCustomAvailableBotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyV2LobbyCustomAvailableBotsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyV2LobbyCustomAvailableBotsOK creates a GetLolLobbyV2LobbyCustomAvailableBotsOK with default headers values
func NewGetLolLobbyV2LobbyCustomAvailableBotsOK() *GetLolLobbyV2LobbyCustomAvailableBotsOK {
	return &GetLolLobbyV2LobbyCustomAvailableBotsOK{}
}

/*GetLolLobbyV2LobbyCustomAvailableBotsOK handles this case with default header values.

Successful response
*/
type GetLolLobbyV2LobbyCustomAvailableBotsOK struct {
	Payload []*models.LolLobbyLobbyBotChampion
}

func (o *GetLolLobbyV2LobbyCustomAvailableBotsOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby/v2/lobby/custom/available-bots][%d] getLolLobbyV2LobbyCustomAvailableBotsOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyV2LobbyCustomAvailableBotsOK) GetPayload() []*models.LolLobbyLobbyBotChampion {
	return o.Payload
}

func (o *GetLolLobbyV2LobbyCustomAvailableBotsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
