// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolLobbyV2LobbyMembersBySummonerIDPromoteReader is a Reader for the PostLolLobbyV2LobbyMembersBySummonerIDPromote structure.
type PostLolLobbyV2LobbyMembersBySummonerIDPromoteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV2LobbyMembersBySummonerIDPromoteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV2LobbyMembersBySummonerIDPromoteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV2LobbyMembersBySummonerIDPromoteOK creates a PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK with default headers values
func NewPostLolLobbyV2LobbyMembersBySummonerIDPromoteOK() *PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK {
	return &PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK{}
}

/*PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK struct {
	Payload int64
}

func (o *PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v2/lobby/members/{summonerId}/promote][%d] postLolLobbyV2LobbyMembersBySummonerIdPromoteOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK) GetPayload() int64 {
	return o.Payload
}

func (o *PostLolLobbyV2LobbyMembersBySummonerIDPromoteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
