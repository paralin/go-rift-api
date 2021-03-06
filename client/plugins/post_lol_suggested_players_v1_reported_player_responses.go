// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolSuggestedPlayersV1ReportedPlayerReader is a Reader for the PostLolSuggestedPlayersV1ReportedPlayer structure.
type PostLolSuggestedPlayersV1ReportedPlayerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolSuggestedPlayersV1ReportedPlayerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolSuggestedPlayersV1ReportedPlayerNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolSuggestedPlayersV1ReportedPlayerNoContent creates a PostLolSuggestedPlayersV1ReportedPlayerNoContent with default headers values
func NewPostLolSuggestedPlayersV1ReportedPlayerNoContent() *PostLolSuggestedPlayersV1ReportedPlayerNoContent {
	return &PostLolSuggestedPlayersV1ReportedPlayerNoContent{}
}

/*PostLolSuggestedPlayersV1ReportedPlayerNoContent handles this case with default header values.

No content
*/
type PostLolSuggestedPlayersV1ReportedPlayerNoContent struct {
}

func (o *PostLolSuggestedPlayersV1ReportedPlayerNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-suggested-players/v1/reported-player][%d] postLolSuggestedPlayersV1ReportedPlayerNoContent ", 204)
}

func (o *PostLolSuggestedPlayersV1ReportedPlayerNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
