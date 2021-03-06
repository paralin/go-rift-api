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

// PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeReader is a Reader for the PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevoke structure.
type PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK creates a PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK with default headers values
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK() *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK {
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK{}
}

/*PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK handles this case with default header values.

Successful response
*/
type PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK struct {
	Payload interface{}
}

func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK) Error() string {
	return fmt.Sprintf("[POST /lol-clash/v1/roster/{rosterId}/ticket-offer/{summonerId}/revoke][%d] postLolClashV1RosterByRosterIdTicketOfferBySummonerIdRevokeOK  %+v", 200, o.Payload)
}

func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDRevokeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
