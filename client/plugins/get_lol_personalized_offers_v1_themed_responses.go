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

// GetLolPersonalizedOffersV1ThemedReader is a Reader for the GetLolPersonalizedOffersV1Themed structure.
type GetLolPersonalizedOffersV1ThemedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPersonalizedOffersV1ThemedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPersonalizedOffersV1ThemedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPersonalizedOffersV1ThemedOK creates a GetLolPersonalizedOffersV1ThemedOK with default headers values
func NewGetLolPersonalizedOffersV1ThemedOK() *GetLolPersonalizedOffersV1ThemedOK {
	return &GetLolPersonalizedOffersV1ThemedOK{}
}

/*GetLolPersonalizedOffersV1ThemedOK handles this case with default header values.

Successful response
*/
type GetLolPersonalizedOffersV1ThemedOK struct {
	Payload bool
}

func (o *GetLolPersonalizedOffersV1ThemedOK) Error() string {
	return fmt.Sprintf("[GET /lol-personalized-offers/v1/themed][%d] getLolPersonalizedOffersV1ThemedOK  %+v", 200, o.Payload)
}

func (o *GetLolPersonalizedOffersV1ThemedOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolPersonalizedOffersV1ThemedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}