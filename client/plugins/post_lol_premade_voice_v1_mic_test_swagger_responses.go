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

// PostLolPremadeVoiceV1MicTestReader is a Reader for the PostLolPremadeVoiceV1MicTest structure.
type PostLolPremadeVoiceV1MicTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPremadeVoiceV1MicTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolPremadeVoiceV1MicTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPremadeVoiceV1MicTestOK creates a PostLolPremadeVoiceV1MicTestOK with default headers values
func NewPostLolPremadeVoiceV1MicTestOK() *PostLolPremadeVoiceV1MicTestOK {
	return &PostLolPremadeVoiceV1MicTestOK{}
}

/*PostLolPremadeVoiceV1MicTestOK handles this case with default header values.

Successful response
*/
type PostLolPremadeVoiceV1MicTestOK struct {
	Payload interface{}
}

func (o *PostLolPremadeVoiceV1MicTestOK) Error() string {
	return fmt.Sprintf("[POST /lol-premade-voice/v1/mic-test][%d] postLolPremadeVoiceV1MicTestOK  %+v", 200, o.Payload)
}

func (o *PostLolPremadeVoiceV1MicTestOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolPremadeVoiceV1MicTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}