// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolHonorV2V1LevelChangeAckReader is a Reader for the PostLolHonorV2V1LevelChangeAck structure.
type PostLolHonorV2V1LevelChangeAckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolHonorV2V1LevelChangeAckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolHonorV2V1LevelChangeAckNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolHonorV2V1LevelChangeAckNoContent creates a PostLolHonorV2V1LevelChangeAckNoContent with default headers values
func NewPostLolHonorV2V1LevelChangeAckNoContent() *PostLolHonorV2V1LevelChangeAckNoContent {
	return &PostLolHonorV2V1LevelChangeAckNoContent{}
}

/*PostLolHonorV2V1LevelChangeAckNoContent handles this case with default header values.

No content
*/
type PostLolHonorV2V1LevelChangeAckNoContent struct {
}

func (o *PostLolHonorV2V1LevelChangeAckNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-honor-v2/v1/level-change/ack][%d] postLolHonorV2V1LevelChangeAckNoContent ", 204)
}

func (o *PostLolHonorV2V1LevelChangeAckNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
