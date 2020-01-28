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

// GetLolPersonalizedOffersV1StatusReader is a Reader for the GetLolPersonalizedOffersV1Status structure.
type GetLolPersonalizedOffersV1StatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPersonalizedOffersV1StatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPersonalizedOffersV1StatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPersonalizedOffersV1StatusOK creates a GetLolPersonalizedOffersV1StatusOK with default headers values
func NewGetLolPersonalizedOffersV1StatusOK() *GetLolPersonalizedOffersV1StatusOK {
	return &GetLolPersonalizedOffersV1StatusOK{}
}

/*GetLolPersonalizedOffersV1StatusOK handles this case with default header values.

Successful response
*/
type GetLolPersonalizedOffersV1StatusOK struct {
	Payload *models.LolPersonalizedOffersUIStatus
}

func (o *GetLolPersonalizedOffersV1StatusOK) Error() string {
	return fmt.Sprintf("[GET /lol-personalized-offers/v1/status][%d] getLolPersonalizedOffersV1StatusOK  %+v", 200, o.Payload)
}

func (o *GetLolPersonalizedOffersV1StatusOK) GetPayload() *models.LolPersonalizedOffersUIStatus {
	return o.Payload
}

func (o *GetLolPersonalizedOffersV1StatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPersonalizedOffersUIStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}