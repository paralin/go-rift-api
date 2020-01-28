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

// PostLolPlayerPreferencesV1HashReader is a Reader for the PostLolPlayerPreferencesV1Hash structure.
type PostLolPlayerPreferencesV1HashReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPlayerPreferencesV1HashReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolPlayerPreferencesV1HashOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPlayerPreferencesV1HashOK creates a PostLolPlayerPreferencesV1HashOK with default headers values
func NewPostLolPlayerPreferencesV1HashOK() *PostLolPlayerPreferencesV1HashOK {
	return &PostLolPlayerPreferencesV1HashOK{}
}

/*PostLolPlayerPreferencesV1HashOK handles this case with default header values.

Successful response
*/
type PostLolPlayerPreferencesV1HashOK struct {
	Payload string
}

func (o *PostLolPlayerPreferencesV1HashOK) Error() string {
	return fmt.Sprintf("[POST /lol-player-preferences/v1/hash][%d] postLolPlayerPreferencesV1HashOK  %+v", 200, o.Payload)
}

func (o *PostLolPlayerPreferencesV1HashOK) GetPayload() string {
	return o.Payload
}

func (o *PostLolPlayerPreferencesV1HashOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
