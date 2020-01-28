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

// PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeReader is a Reader for the PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevoke structure.
type PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK creates a PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK with default headers values
func NewPostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK() *PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK {
	return &PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK{}
}

/*PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK handles this case with default header values.

Successful response
*/
type PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK struct {
	Payload interface{}
}

func (o *PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK) Error() string {
	return fmt.Sprintf("[POST /lol-clash/v1/roster/{rosterId}/substitute/{summonerId}/revoke][%d] postLolClashV1RosterByRosterIdSubstituteBySummonerIdRevokeOK  %+v", 200, o.Payload)
}

func (o *PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolClashV1RosterByRosterIDSubstituteBySummonerIDRevokeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
