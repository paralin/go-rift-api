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

// GetLolRsoAuthV1AuthorizationReader is a Reader for the GetLolRsoAuthV1Authorization structure.
type GetLolRsoAuthV1AuthorizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRsoAuthV1AuthorizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRsoAuthV1AuthorizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRsoAuthV1AuthorizationOK creates a GetLolRsoAuthV1AuthorizationOK with default headers values
func NewGetLolRsoAuthV1AuthorizationOK() *GetLolRsoAuthV1AuthorizationOK {
	return &GetLolRsoAuthV1AuthorizationOK{}
}

/*GetLolRsoAuthV1AuthorizationOK handles this case with default header values.

Successful response
*/
type GetLolRsoAuthV1AuthorizationOK struct {
	Payload *models.LolRsoAuthAuthorization
}

func (o *GetLolRsoAuthV1AuthorizationOK) Error() string {
	return fmt.Sprintf("[GET /lol-rso-auth/v1/authorization][%d] getLolRsoAuthV1AuthorizationOK  %+v", 200, o.Payload)
}

func (o *GetLolRsoAuthV1AuthorizationOK) GetPayload() *models.LolRsoAuthAuthorization {
	return o.Payload
}

func (o *GetLolRsoAuthV1AuthorizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRsoAuthAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}