// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLolGameClientChatV1IgnoredSummonersReader is a Reader for the DeleteLolGameClientChatV1IgnoredSummoners structure.
type DeleteLolGameClientChatV1IgnoredSummonersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolGameClientChatV1IgnoredSummonersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLolGameClientChatV1IgnoredSummonersNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolGameClientChatV1IgnoredSummonersNoContent creates a DeleteLolGameClientChatV1IgnoredSummonersNoContent with default headers values
func NewDeleteLolGameClientChatV1IgnoredSummonersNoContent() *DeleteLolGameClientChatV1IgnoredSummonersNoContent {
	return &DeleteLolGameClientChatV1IgnoredSummonersNoContent{}
}

/*DeleteLolGameClientChatV1IgnoredSummonersNoContent handles this case with default header values.

No content
*/
type DeleteLolGameClientChatV1IgnoredSummonersNoContent struct {
}

func (o *DeleteLolGameClientChatV1IgnoredSummonersNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lol-game-client-chat/v1/ignored-summoners][%d] deleteLolGameClientChatV1IgnoredSummonersNoContent ", 204)
}

func (o *DeleteLolGameClientChatV1IgnoredSummonersNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
