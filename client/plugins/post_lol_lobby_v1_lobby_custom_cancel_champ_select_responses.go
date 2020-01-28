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

// PostLolLobbyV1LobbyCustomCancelChampSelectReader is a Reader for the PostLolLobbyV1LobbyCustomCancelChampSelect structure.
type PostLolLobbyV1LobbyCustomCancelChampSelectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV1LobbyCustomCancelChampSelectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyV1LobbyCustomCancelChampSelectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV1LobbyCustomCancelChampSelectOK creates a PostLolLobbyV1LobbyCustomCancelChampSelectOK with default headers values
func NewPostLolLobbyV1LobbyCustomCancelChampSelectOK() *PostLolLobbyV1LobbyCustomCancelChampSelectOK {
	return &PostLolLobbyV1LobbyCustomCancelChampSelectOK{}
}

/*PostLolLobbyV1LobbyCustomCancelChampSelectOK handles this case with default header values.

Successful response
*/
type PostLolLobbyV1LobbyCustomCancelChampSelectOK struct {
	Payload interface{}
}

func (o *PostLolLobbyV1LobbyCustomCancelChampSelectOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v1/lobby/custom/cancel-champ-select][%d] postLolLobbyV1LobbyCustomCancelChampSelectOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyV1LobbyCustomCancelChampSelectOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLobbyV1LobbyCustomCancelChampSelectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
