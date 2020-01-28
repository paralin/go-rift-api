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

// PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptReader is a Reader for the PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAccept structure.
type PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK creates a PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK with default headers values
func NewPostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK() *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK {
	return &PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK{}
}

/*PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK handles this case with default header values.

Successful response
*/
type PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK struct {
	Payload interface{}
}

func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK) Error() string {
	return fmt.Sprintf("[POST /lol-lobby-team-builder/champ-select/v1/session/trades/{id}/accept][%d] postLolLobbyTeamBuilderChampSelectV1SessionTradesByIdAcceptOK  %+v", 200, o.Payload)
}

func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLobbyTeamBuilderChampSelectV1SessionTradesByIDAcceptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
