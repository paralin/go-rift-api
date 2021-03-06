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

// GetLolLobbyTeamBuilderV1MatchmakingReader is a Reader for the GetLolLobbyTeamBuilderV1Matchmaking structure.
type GetLolLobbyTeamBuilderV1MatchmakingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLobbyTeamBuilderV1MatchmakingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLobbyTeamBuilderV1MatchmakingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLobbyTeamBuilderV1MatchmakingOK creates a GetLolLobbyTeamBuilderV1MatchmakingOK with default headers values
func NewGetLolLobbyTeamBuilderV1MatchmakingOK() *GetLolLobbyTeamBuilderV1MatchmakingOK {
	return &GetLolLobbyTeamBuilderV1MatchmakingOK{}
}

/*GetLolLobbyTeamBuilderV1MatchmakingOK handles this case with default header values.

Successful response
*/
type GetLolLobbyTeamBuilderV1MatchmakingOK struct {
	Payload *models.LolLobbyTeamBuilderMatchmakingSearchResource
}

func (o *GetLolLobbyTeamBuilderV1MatchmakingOK) Error() string {
	return fmt.Sprintf("[GET /lol-lobby-team-builder/v1/matchmaking][%d] getLolLobbyTeamBuilderV1MatchmakingOK  %+v", 200, o.Payload)
}

func (o *GetLolLobbyTeamBuilderV1MatchmakingOK) GetPayload() *models.LolLobbyTeamBuilderMatchmakingSearchResource {
	return o.Payload
}

func (o *GetLolLobbyTeamBuilderV1MatchmakingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLobbyTeamBuilderMatchmakingSearchResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
