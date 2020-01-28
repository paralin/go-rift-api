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

// GetLolLobbyTeamBuilderChampSelectV1SessionReader is a Reader for the GetLolLobbyTeamBuilderChampSelectV1Session structure.
type GetLolLobbyTeamBuilderChampSelectV1SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyTeamBuilderChampSelectV1SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyTeamBuilderChampSelectV1SessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyTeamBuilderChampSelectV1SessionOK creates a GetLolLobbyTeamBuilderChampSelectV1SessionOK with default headers values
func NewGetLolLobbyTeamBuilderChampSelectV1SessionOK() *GetLolLobbyTeamBuilderChampSelectV1SessionOK {
	return &GetLolLobbyTeamBuilderChampSelectV1SessionOK{}
}

/*GetLolLobbyTeamBuilderChampSelectV1SessionOK handles this case with default header values.

Successful response
*/
type GetLolLobbyTeamBuilderChampSelectV1SessionOK struct {
	Payload *models.LolLobbyTeamBuilderChampSelectSession
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SessionOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby-team-builder/champ-select/v1/session][%d] getLolLobbyTeamBuilderChampSelectV1SessionOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SessionOK) GetPayload() *models.LolLobbyTeamBuilderChampSelectSession {
	return o.Payload
}

func (o *GetLolLobbyTeamBuilderChampSelectV1SessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLobbyTeamBuilderChampSelectSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
