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

// GetLolLoginV1AccountStateReader is a Reader for the GetLolLoginV1AccountState structure.
type GetLolLoginV1AccountStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLoginV1AccountStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLoginV1AccountStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLoginV1AccountStateOK creates a GetLolLoginV1AccountStateOK with default headers values
func NewGetLolLoginV1AccountStateOK() *GetLolLoginV1AccountStateOK {
	return &GetLolLoginV1AccountStateOK{}
}

/*GetLolLoginV1AccountStateOK handles this case with default header values.

Successful response
*/
type GetLolLoginV1AccountStateOK struct {
	Payload *models.LolLoginAccountStateResource
}

func (o *GetLolLoginV1AccountStateOK) Error() string {
	return fmt.Sprintf("[GET /lol-login/v1/account-state][%d] getLolLoginV1AccountStateOK  %+v", 200, o.Payload)
}

func (o *GetLolLoginV1AccountStateOK) GetPayload() *models.LolLoginAccountStateResource {
	return o.Payload
}

func (o *GetLolLoginV1AccountStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolLoginAccountStateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
