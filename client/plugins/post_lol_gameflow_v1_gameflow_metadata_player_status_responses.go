// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolGameflowV1GameflowMetadataPlayerStatusReader is a Reader for the PostLolGameflowV1GameflowMetadataPlayerStatus structure.
type PostLolGameflowV1GameflowMetadataPlayerStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1GameflowMetadataPlayerStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolGameflowV1GameflowMetadataPlayerStatusNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1GameflowMetadataPlayerStatusNoContent creates a PostLolGameflowV1GameflowMetadataPlayerStatusNoContent with default headers values
func NewPostLolGameflowV1GameflowMetadataPlayerStatusNoContent() *PostLolGameflowV1GameflowMetadataPlayerStatusNoContent {
	return &PostLolGameflowV1GameflowMetadataPlayerStatusNoContent{}
}

/*PostLolGameflowV1GameflowMetadataPlayerStatusNoContent handles this case with default header values.

No content
*/
type PostLolGameflowV1GameflowMetadataPlayerStatusNoContent struct {
}

func (o *PostLolGameflowV1GameflowMetadataPlayerStatusNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/gameflow-metadata/player-status][%d] postLolGameflowV1GameflowMetadataPlayerStatusNoContent ", 204)
}

func (o *PostLolGameflowV1GameflowMetadataPlayerStatusNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
