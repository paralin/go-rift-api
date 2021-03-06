// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolLobbyV2PartiesOverridesEnabledReader is a Reader for the PostLolLobbyV2PartiesOverridesEnabled structure.
type PostLolLobbyV2PartiesOverridesEnabledReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLobbyV2PartiesOverridesEnabledReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolLobbyV2PartiesOverridesEnabledNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLobbyV2PartiesOverridesEnabledNoContent creates a PostLolLobbyV2PartiesOverridesEnabledNoContent with default headers values
func NewPostLolLobbyV2PartiesOverridesEnabledNoContent() *PostLolLobbyV2PartiesOverridesEnabledNoContent {
	return &PostLolLobbyV2PartiesOverridesEnabledNoContent{}
}

/*PostLolLobbyV2PartiesOverridesEnabledNoContent handles this case with default header values.

No content
*/
type PostLolLobbyV2PartiesOverridesEnabledNoContent struct {
}

func (o *PostLolLobbyV2PartiesOverridesEnabledNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-lobby/v2/parties/overrides/Enabled][%d] postLolLobbyV2PartiesOverridesEnabledNoContent ", 204)
}

func (o *PostLolLobbyV2PartiesOverridesEnabledNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
