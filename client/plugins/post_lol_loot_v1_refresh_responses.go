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

// PostLolLootV1RefreshReader is a Reader for the PostLolLootV1Refresh structure.
type PostLolLootV1RefreshReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLootV1RefreshReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLootV1RefreshOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLootV1RefreshOK creates a PostLolLootV1RefreshOK with default headers values
func NewPostLolLootV1RefreshOK() *PostLolLootV1RefreshOK {
	return &PostLolLootV1RefreshOK{}
}

/*PostLolLootV1RefreshOK handles this case with default header values.

Successful response
*/
type PostLolLootV1RefreshOK struct {
	Payload string
}

func (o *PostLolLootV1RefreshOK) Error() string {
	return fmt.Sprintf("[POST /lol-loot/v1/refresh][%d] postLolLootV1RefreshOK  %+v", 200, o.Payload)
}

func (o *PostLolLootV1RefreshOK) GetPayload() string {
	return o.Payload
}

func (o *PostLolLootV1RefreshOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
