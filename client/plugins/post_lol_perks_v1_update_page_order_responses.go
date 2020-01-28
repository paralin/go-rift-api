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

// PostLolPerksV1UpdatePageOrderReader is a Reader for the PostLolPerksV1UpdatePageOrder structure.
type PostLolPerksV1UpdatePageOrderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPerksV1UpdatePageOrderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolPerksV1UpdatePageOrderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPerksV1UpdatePageOrderOK creates a PostLolPerksV1UpdatePageOrderOK with default headers values
func NewPostLolPerksV1UpdatePageOrderOK() *PostLolPerksV1UpdatePageOrderOK {
	return &PostLolPerksV1UpdatePageOrderOK{}
}

/*PostLolPerksV1UpdatePageOrderOK handles this case with default header values.

Successful response
*/
type PostLolPerksV1UpdatePageOrderOK struct {
	Payload interface{}
}

func (o *PostLolPerksV1UpdatePageOrderOK) Error() string {
	return fmt.Sprintf("[POST /lol-perks/v1/update-page-order][%d] postLolPerksV1UpdatePageOrderOK  %+v", 200, o.Payload)
}

func (o *PostLolPerksV1UpdatePageOrderOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolPerksV1UpdatePageOrderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
