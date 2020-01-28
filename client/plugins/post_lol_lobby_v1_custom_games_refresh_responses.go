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

// PostLolLobbyV1CustomGamesRefreshReader is a Reader for the PostLolLobbyV1CustomGamesRefresh structure.
type PostLolLobbyV1CustomGamesRefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV1CustomGamesRefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV1CustomGamesRefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV1CustomGamesRefreshOK creates a PostLolLobbyV1CustomGamesRefreshOK with default headers values
func NewPostLolLobbyV1CustomGamesRefreshOK() *PostLolLobbyV1CustomGamesRefreshOK {
	return &PostLolLobbyV1CustomGamesRefreshOK{}
}

/*PostLolLobbyV1CustomGamesRefreshOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV1CustomGamesRefreshOK struct {
	Payload interface{}
}

func (o *PostLolLobbyV1CustomGamesRefreshOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v1/custom-games/refresh][%d] postLolLobbyV1CustomGamesRefreshOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV1CustomGamesRefreshOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLobbyV1CustomGamesRefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
