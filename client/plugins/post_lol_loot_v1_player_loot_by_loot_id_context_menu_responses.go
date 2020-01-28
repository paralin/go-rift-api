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

// PostLolLootV1PlayerLootByLootIDContextMenuReader is a Reader for the PostLolLootV1PlayerLootByLootIDContextMenu structure.
type PostLolLootV1PlayerLootByLootIDContextMenuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLootV1PlayerLootByLootIDContextMenuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLootV1PlayerLootByLootIDContextMenuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLootV1PlayerLootByLootIDContextMenuOK creates a PostLolLootV1PlayerLootByLootIDContextMenuOK with default headers values
func NewPostLolLootV1PlayerLootByLootIDContextMenuOK() *PostLolLootV1PlayerLootByLootIDContextMenuOK {
	return &PostLolLootV1PlayerLootByLootIDContextMenuOK{}
}

/*PostLolLootV1PlayerLootByLootIDContextMenuOK handles this case with default header values.

Successful response
*/
type PostLolLootV1PlayerLootByLootIDContextMenuOK struct {
	Payload []*models.LolLootContextMenu
}

func (o *PostLolLootV1PlayerLootByLootIDContextMenuOK) Error() string {
	return fmt.Sprintf("[POST /lol-loot/v1/player-loot/{lootId}/context-menu][%d] postLolLootV1PlayerLootByLootIdContextMenuOK  %+v", 200, o.Payload)
}

func (o *PostLolLootV1PlayerLootByLootIDContextMenuOK) GetPayload() []*models.LolLootContextMenu {
	return o.Payload
}

func (o *PostLolLootV1PlayerLootByLootIDContextMenuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}