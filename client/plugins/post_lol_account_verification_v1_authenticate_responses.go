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

// PostLolAccountVerificationV1AuthenticateReader is a Reader for the PostLolAccountVerificationV1Authenticate structure.
type PostLolAccountVerificationV1AuthenticateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolAccountVerificationV1AuthenticateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolAccountVerificationV1AuthenticateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolAccountVerificationV1AuthenticateOK creates a PostLolAccountVerificationV1AuthenticateOK with default headers values
func NewPostLolAccountVerificationV1AuthenticateOK() *PostLolAccountVerificationV1AuthenticateOK {
	return &PostLolAccountVerificationV1AuthenticateOK{}
}

/*PostLolAccountVerificationV1AuthenticateOK handles this case with default header values.

Successful response
*/
type PostLolAccountVerificationV1AuthenticateOK struct {
	Payload *models.LolAccountVerificationAuthenticateResponse
}

func (o *PostLolAccountVerificationV1AuthenticateOK) Error() string {
	return fmt.Sprintf("[POST /lol-account-verification/v1/authenticate][%d] postLolAccountVerificationV1AuthenticateOK  %+v", 200, o.Payload)
}

func (o *PostLolAccountVerificationV1AuthenticateOK) GetPayload() *models.LolAccountVerificationAuthenticateResponse {
	return o.Payload
}

func (o *PostLolAccountVerificationV1AuthenticateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolAccountVerificationAuthenticateResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}