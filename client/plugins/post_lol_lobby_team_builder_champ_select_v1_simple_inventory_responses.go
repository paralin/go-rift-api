// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryReader is a Reader for the PostLolLobbyTeamBuilderChampSelectV1SimpleInventory structure.
type PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent creates a PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent with default headers values
func NewPostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent() *PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent {
	return &PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent{}
}

/*PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent handles this case with default header values.

No content
*/
type PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent struct {
}

func (o *PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-lobby-team-builder/champ-select/v1/simple-inventory][%d] postLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent ", 204)
}

func (o *PostLolLobbyTeamBuilderChampSelectV1SimpleInventoryNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}