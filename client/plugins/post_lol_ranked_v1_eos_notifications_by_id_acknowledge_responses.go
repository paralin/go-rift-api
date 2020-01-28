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

// PostLolRankedV1EosNotificationsByIDAcknowledgeReader is a Reader for the PostLolRankedV1EosNotificationsByIDAcknowledge structure.
type PostLolRankedV1EosNotificationsByIDAcknowledgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolRankedV1EosNotificationsByIDAcknowledgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolRankedV1EosNotificationsByIDAcknowledgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolRankedV1EosNotificationsByIDAcknowledgeOK creates a PostLolRankedV1EosNotificationsByIDAcknowledgeOK with default headers values
func NewPostLolRankedV1EosNotificationsByIDAcknowledgeOK() *PostLolRankedV1EosNotificationsByIDAcknowledgeOK {
	return &PostLolRankedV1EosNotificationsByIDAcknowledgeOK{}
}

/*PostLolRankedV1EosNotificationsByIDAcknowledgeOK handles this case with default header values.

Successful response
*/
type PostLolRankedV1EosNotificationsByIDAcknowledgeOK struct {
	Payload interface{}
}

func (o *PostLolRankedV1EosNotificationsByIDAcknowledgeOK) Error() string {
	return fmt.Sprintf("[POST /lol-ranked/v1/eos-notifications/{id}/acknowledge][%d] postLolRankedV1EosNotificationsByIdAcknowledgeOK  %+v", 200, o.Payload)
}

func (o *PostLolRankedV1EosNotificationsByIDAcknowledgeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolRankedV1EosNotificationsByIDAcknowledgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
