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

// PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonReader is a Reader for the PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandon structure.
type PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK creates a PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK with default headers values
func NewPostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK() *PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK {
	return &PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK{}
}

/*PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK handles this case with default header values.

Successful response
*/
type PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK struct {
	Payload interface{}
}

func (o *PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby-team-builder/v1/matchmaking/low-priority-queue/abandon][%d] postLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLobbyTeamBuilderV1MatchmakingLowPriorityQueueAbandonOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
