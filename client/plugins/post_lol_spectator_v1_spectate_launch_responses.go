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

// PostLolSpectatorV1SpectateLaunchReader is a Reader for the PostLolSpectatorV1SpectateLaunch structure.
type PostLolSpectatorV1SpectateLaunchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolSpectatorV1SpectateLaunchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolSpectatorV1SpectateLaunchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolSpectatorV1SpectateLaunchOK creates a PostLolSpectatorV1SpectateLaunchOK with default headers values
func NewPostLolSpectatorV1SpectateLaunchOK() *PostLolSpectatorV1SpectateLaunchOK {
	return &PostLolSpectatorV1SpectateLaunchOK{}
}

/*PostLolSpectatorV1SpectateLaunchOK handles this case with default header values.

Successful response
*/
type PostLolSpectatorV1SpectateLaunchOK struct {
	Payload interface{}
}

func (o *PostLolSpectatorV1SpectateLaunchOK) Error() string {
	return fmt.Sprintf("[POST /lol-spectator/v1/spectate/launch][%d] postLolSpectatorV1SpectateLaunchOK  %+v", 200, o.Payload)
}

func (o *PostLolSpectatorV1SpectateLaunchOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolSpectatorV1SpectateLaunchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}