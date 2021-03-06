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

// PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferReader is a Reader for the PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOffer structure.
type PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK creates a PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK with default headers values
func NewPostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK() *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK {
	return &PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK{}
}

/*PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK handles this case with default header values.

Successful response
*/
type PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK struct {
	Payload interface{}
}

func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK) Error() string {
	return fmt.Sprintf("[POST /lol-clash/v1/roster/{rosterId}/ticket-offer/{summonerId}/offer][%d] postLolClashV1RosterByRosterIdTicketOfferBySummonerIdOfferOK  %+v", 200, o.Payload)
}

func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolClashV1RosterByRosterIDTicketOfferBySummonerIDOfferOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
