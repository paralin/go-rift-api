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

// PostLolGameflowV1SessionRequestLobbyReader is a Reader for the PostLolGameflowV1SessionRequestLobby structure.
type PostLolGameflowV1SessionRequestLobbyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1SessionRequestLobbyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolGameflowV1SessionRequestLobbyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1SessionRequestLobbyOK creates a PostLolGameflowV1SessionRequestLobbyOK with default headers values
func NewPostLolGameflowV1SessionRequestLobbyOK() *PostLolGameflowV1SessionRequestLobbyOK {
	return &PostLolGameflowV1SessionRequestLobbyOK{}
}

/*PostLolGameflowV1SessionRequestLobbyOK handles this case with default header values.

Successful response
*/
type PostLolGameflowV1SessionRequestLobbyOK struct {
	Payload bool
}

func (o *PostLolGameflowV1SessionRequestLobbyOK) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/session/request-lobby][%d] postLolGameflowV1SessionRequestLobbyOK  %+v", 200, o.Payload)
}

func (o *PostLolGameflowV1SessionRequestLobbyOK) GetPayload() bool {
	return o.Payload
}

func (o *PostLolGameflowV1SessionRequestLobbyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
