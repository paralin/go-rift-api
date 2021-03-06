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

// PostLolLoginV1NewPlayerFlowCompletedReader is a Reader for the PostLolLoginV1NewPlayerFlowCompleted structure.
type PostLolLoginV1NewPlayerFlowCompletedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLoginV1NewPlayerFlowCompletedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLoginV1NewPlayerFlowCompletedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLoginV1NewPlayerFlowCompletedOK creates a PostLolLoginV1NewPlayerFlowCompletedOK with default headers values
func NewPostLolLoginV1NewPlayerFlowCompletedOK() *PostLolLoginV1NewPlayerFlowCompletedOK {
	return &PostLolLoginV1NewPlayerFlowCompletedOK{}
}

/*PostLolLoginV1NewPlayerFlowCompletedOK handles this case with default header values.

Successful response
*/
type PostLolLoginV1NewPlayerFlowCompletedOK struct {
	Payload interface{}
}

func (o *PostLolLoginV1NewPlayerFlowCompletedOK) Error() string {
	return fmt.Sprintf("[POST /lol-login/v1/new-player-flow-completed][%d] postLolLoginV1NewPlayerFlowCompletedOK  %+v", 200, o.Payload)
}

func (o *PostLolLoginV1NewPlayerFlowCompletedOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLoginV1NewPlayerFlowCompletedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
