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

// PostLolPersonalizedOffersV1OffersPurchaseReader is a Reader for the PostLolPersonalizedOffersV1OffersPurchase structure.
type PostLolPersonalizedOffersV1OffersPurchaseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPersonalizedOffersV1OffersPurchaseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolPersonalizedOffersV1OffersPurchaseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPersonalizedOffersV1OffersPurchaseOK creates a PostLolPersonalizedOffersV1OffersPurchaseOK with default headers values
func NewPostLolPersonalizedOffersV1OffersPurchaseOK() *PostLolPersonalizedOffersV1OffersPurchaseOK {
	return &PostLolPersonalizedOffersV1OffersPurchaseOK{}
}

/*PostLolPersonalizedOffersV1OffersPurchaseOK handles this case with default header values.

Successful response
*/
type PostLolPersonalizedOffersV1OffersPurchaseOK struct {
	Payload *models.LolPersonalizedOffersPurchaseResponse
}

func (o *PostLolPersonalizedOffersV1OffersPurchaseOK) Error() string {
	return fmt.Sprintf("[POST /lol-personalized-offers/v1/offers/purchase][%d] postLolPersonalizedOffersV1OffersPurchaseOK  %+v", 200, o.Payload)
}

func (o *PostLolPersonalizedOffersV1OffersPurchaseOK) GetPayload() *models.LolPersonalizedOffersPurchaseResponse {
	return o.Payload
}

func (o *PostLolPersonalizedOffersV1OffersPurchaseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPersonalizedOffersPurchaseResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
