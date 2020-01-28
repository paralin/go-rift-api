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

// PostLolClashV1RosterByRosterIDChangeNameReader is a Reader for the PostLolClashV1RosterByRosterIDChangeName structure.
type PostLolClashV1RosterByRosterIDChangeNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolClashV1RosterByRosterIDChangeNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolClashV1RosterByRosterIDChangeNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolClashV1RosterByRosterIDChangeNameOK creates a PostLolClashV1RosterByRosterIDChangeNameOK with default headers values
func NewPostLolClashV1RosterByRosterIDChangeNameOK() *PostLolClashV1RosterByRosterIDChangeNameOK {
	return &PostLolClashV1RosterByRosterIDChangeNameOK{}
}

/*PostLolClashV1RosterByRosterIDChangeNameOK handles this case with default header values.

Successful response
*/
type PostLolClashV1RosterByRosterIDChangeNameOK struct {
	Payload interface{}
}

func (o *PostLolClashV1RosterByRosterIDChangeNameOK) Error() string {
	return fmt.Sprintf("[POST /lol-clash/v1/roster/{rosterId}/change-name][%d] postLolClashV1RosterByRosterIdChangeNameOK  %+v", 200, o.Payload)
}

func (o *PostLolClashV1RosterByRosterIDChangeNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolClashV1RosterByRosterIDChangeNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
