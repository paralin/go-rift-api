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

// PostLolClashV1RosterByRosterIDSubstituteDeclineReader is a Reader for the PostLolClashV1RosterByRosterIDSubstituteDecline structure.
type PostLolClashV1RosterByRosterIDSubstituteDeclineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolClashV1RosterByRosterIDSubstituteDeclineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteDeclineOK creates a PostLolClashV1RosterByRosterIDSubstituteDeclineOK with default headers values
func NewPostLolClashV1RosterByRosterIDSubstituteDeclineOK() *PostLolClashV1RosterByRosterIDSubstituteDeclineOK {
	return &PostLolClashV1RosterByRosterIDSubstituteDeclineOK{}
}

/*PostLolClashV1RosterByRosterIDSubstituteDeclineOK handles this case with default header values.

Successful response
*/
type PostLolClashV1RosterByRosterIDSubstituteDeclineOK struct {
	Payload interface{}
}

func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineOK) Error() string {
	return fmt.Sprintf("[POST /lol-clash/v1/roster/{rosterId}/substitute/decline][%d] postLolClashV1RosterByRosterIdSubstituteDeclineOK  %+v", 200, o.Payload)
}

func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolClashV1RosterByRosterIDSubstituteDeclineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
