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

// PostRecofrienderV1ContactsByAccountIDAvailableReader is a Reader for the PostRecofrienderV1ContactsByAccountIDAvailable structure.
type PostRecofrienderV1ContactsByAccountIDAvailableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRecofrienderV1ContactsByAccountIDAvailableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostRecofrienderV1ContactsByAccountIDAvailableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRecofrienderV1ContactsByAccountIDAvailableOK creates a PostRecofrienderV1ContactsByAccountIDAvailableOK with default headers values
func NewPostRecofrienderV1ContactsByAccountIDAvailableOK() *PostRecofrienderV1ContactsByAccountIDAvailableOK {
	return &PostRecofrienderV1ContactsByAccountIDAvailableOK{}
}

/*PostRecofrienderV1ContactsByAccountIDAvailableOK handles this case with default header values.

Successful response
*/
type PostRecofrienderV1ContactsByAccountIDAvailableOK struct {
	Payload *models.RecofrienderContactStateResource
}

func (o *PostRecofrienderV1ContactsByAccountIDAvailableOK) Error() string {
	return fmt.Sprintf("[POST /recofriender/v1/contacts/{accountId}/available][%d] postRecofrienderV1ContactsByAccountIdAvailableOK  %+v", 200, o.Payload)
}

func (o *PostRecofrienderV1ContactsByAccountIDAvailableOK) GetPayload() *models.RecofrienderContactStateResource {
	return o.Payload
}

func (o *PostRecofrienderV1ContactsByAccountIDAvailableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecofrienderContactStateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
