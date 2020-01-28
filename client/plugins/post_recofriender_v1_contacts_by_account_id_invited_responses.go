// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/paralin/go-rift-api/models"
)

// PostRecofrienderV1ContactsByAccountIDInvitedReader is a Reader for the PostRecofrienderV1ContactsByAccountIDInvited structure.
type PostRecofrienderV1ContactsByAccountIDInvitedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecofrienderV1ContactsByAccountIDInvitedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRecofrienderV1ContactsByAccountIDInvitedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRecofrienderV1ContactsByAccountIDInvitedOK creates a PostRecofrienderV1ContactsByAccountIDInvitedOK with default headers values
func NewPostRecofrienderV1ContactsByAccountIDInvitedOK() *PostRecofrienderV1ContactsByAccountIDInvitedOK {
	return &PostRecofrienderV1ContactsByAccountIDInvitedOK{}
}

/*PostRecofrienderV1ContactsByAccountIDInvitedOK handles this case with default header values.

Successful response
*/
type PostRecofrienderV1ContactsByAccountIDInvitedOK struct {
	Payload *models.RecofrienderContactStateResource
}

func (o *PostRecofrienderV1ContactsByAccountIDInvitedOK) Error() string {
	return fmt.Sprintf("[POST /recofriender/v1/contacts/{accountId}/invited][%d] postRecofrienderV1ContactsByAccountIdInvitedOK  %+v", 200, o.Payload)
}

func (o *PostRecofrienderV1ContactsByAccountIDInvitedOK) GetPayload() *models.RecofrienderContactStateResource {
	return o.Payload
}

func (o *PostRecofrienderV1ContactsByAccountIDInvitedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecofrienderContactStateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
