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

// GetLolLoginV1LeagueSessionLogoutOnFailureReader is a Reader for the GetLolLoginV1LeagueSessionLogoutOnFailure structure.
type GetLolLoginV1LeagueSessionLogoutOnFailureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLoginV1LeagueSessionLogoutOnFailureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolLoginV1LeagueSessionLogoutOnFailureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLoginV1LeagueSessionLogoutOnFailureOK creates a GetLolLoginV1LeagueSessionLogoutOnFailureOK with default headers values
func NewGetLolLoginV1LeagueSessionLogoutOnFailureOK() *GetLolLoginV1LeagueSessionLogoutOnFailureOK {
	return &GetLolLoginV1LeagueSessionLogoutOnFailureOK{}
}

/*GetLolLoginV1LeagueSessionLogoutOnFailureOK handles this case with default header values.

Successful response
*/
type GetLolLoginV1LeagueSessionLogoutOnFailureOK struct {
	Payload bool
}

func (o *GetLolLoginV1LeagueSessionLogoutOnFailureOK) Error() string {
	return fmt.Sprintf("[GET /lol-login/v1/league-session-logout-on-failure][%d] getLolLoginV1LeagueSessionLogoutOnFailureOK  %+v", 200, o.Payload)
}

func (o *GetLolLoginV1LeagueSessionLogoutOnFailureOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolLoginV1LeagueSessionLogoutOnFailureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
