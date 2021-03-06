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

// GetLolRsoAuthV1AuthorizationIDTokenReader is a Reader for the GetLolRsoAuthV1AuthorizationIDToken structure.
type GetLolRsoAuthV1AuthorizationIDTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRsoAuthV1AuthorizationIDTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRsoAuthV1AuthorizationIDTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRsoAuthV1AuthorizationIDTokenOK creates a GetLolRsoAuthV1AuthorizationIDTokenOK with default headers values
func NewGetLolRsoAuthV1AuthorizationIDTokenOK() *GetLolRsoAuthV1AuthorizationIDTokenOK {
	return &GetLolRsoAuthV1AuthorizationIDTokenOK{}
}

/*GetLolRsoAuthV1AuthorizationIDTokenOK handles this case with default header values.

Successful response
*/
type GetLolRsoAuthV1AuthorizationIDTokenOK struct {
	Payload *models.LolRsoAuthIDToken
}

func (o *GetLolRsoAuthV1AuthorizationIDTokenOK) Error() string {
	return fmt.Sprintf("[GET /lol-rso-auth/v1/authorization/id-token][%d] getLolRsoAuthV1AuthorizationIdTokenOK  %+v", 200, o.Payload)
}

func (o *GetLolRsoAuthV1AuthorizationIDTokenOK) GetPayload() *models.LolRsoAuthIDToken {
	return o.Payload
}

func (o *GetLolRsoAuthV1AuthorizationIDTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRsoAuthIDToken)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
