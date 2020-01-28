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

// PostLolGameflowV1ReconnectReader is a Reader for the PostLolGameflowV1Reconnect structure.
type PostLolGameflowV1ReconnectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1ReconnectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolGameflowV1ReconnectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1ReconnectOK creates a PostLolGameflowV1ReconnectOK with default headers values
func NewPostLolGameflowV1ReconnectOK() *PostLolGameflowV1ReconnectOK {
	return &PostLolGameflowV1ReconnectOK{}
}

/*PostLolGameflowV1ReconnectOK handles this case with default header values.

Successful response
*/
type PostLolGameflowV1ReconnectOK struct {
	Payload interface{}
}

func (o *PostLolGameflowV1ReconnectOK) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/reconnect][%d] postLolGameflowV1ReconnectOK  %+v", 200, o.Payload)
}

func (o *PostLolGameflowV1ReconnectOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolGameflowV1ReconnectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
