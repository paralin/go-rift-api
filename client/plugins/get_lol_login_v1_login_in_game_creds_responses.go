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

// GetLolLoginV1LoginInGameCredsReader is a Reader for the GetLolLoginV1LoginInGameCreds structure.
type GetLolLoginV1LoginInGameCredsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLoginV1LoginInGameCredsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLoginV1LoginInGameCredsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLoginV1LoginInGameCredsOK creates a GetLolLoginV1LoginInGameCredsOK with default headers values
func NewGetLolLoginV1LoginInGameCredsOK() *GetLolLoginV1LoginInGameCredsOK {
	return &GetLolLoginV1LoginInGameCredsOK{}
}

/*GetLolLoginV1LoginInGameCredsOK handles this case with default header values.

Successful response
*/
type GetLolLoginV1LoginInGameCredsOK struct {
	Payload interface{}
}

func (o *GetLolLoginV1LoginInGameCredsOK) Error() string {
	return fmt.Sprintf("[GET /lol-login/v1/login-in-game-creds][%d] getLolLoginV1LoginInGameCredsOK  %+v", 200, o.Payload)
}

func (o *GetLolLoginV1LoginInGameCredsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolLoginV1LoginInGameCredsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}