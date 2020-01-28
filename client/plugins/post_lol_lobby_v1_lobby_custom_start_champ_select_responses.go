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

// PostLolLobbyV1LobbyCustomStartChampSelectReader is a Reader for the PostLolLobbyV1LobbyCustomStartChampSelect structure.
type PostLolLobbyV1LobbyCustomStartChampSelectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV1LobbyCustomStartChampSelectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV1LobbyCustomStartChampSelectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV1LobbyCustomStartChampSelectOK creates a PostLolLobbyV1LobbyCustomStartChampSelectOK with default headers values
func NewPostLolLobbyV1LobbyCustomStartChampSelectOK() *PostLolLobbyV1LobbyCustomStartChampSelectOK {
	return &PostLolLobbyV1LobbyCustomStartChampSelectOK{}
}

/*PostLolLobbyV1LobbyCustomStartChampSelectOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV1LobbyCustomStartChampSelectOK struct {
	Payload *models.LolLobbyLobbyCustomChampSelectStartResponse
}

func (o *PostLolLobbyV1LobbyCustomStartChampSelectOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v1/lobby/custom/start-champ-select][%d] postLolLobbyV1LobbyCustomStartChampSelectOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV1LobbyCustomStartChampSelectOK) GetPayload() *models.LolLobbyLobbyCustomChampSelectStartResponse {
	return o.Payload
}

func (o *PostLolLobbyV1LobbyCustomStartChampSelectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLobbyLobbyCustomChampSelectStartResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}