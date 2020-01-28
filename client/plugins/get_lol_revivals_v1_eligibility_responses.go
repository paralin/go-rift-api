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

// GetLolRevivalsV1EligibilityReader is a Reader for the GetLolRevivalsV1Eligibility structure.
type GetLolRevivalsV1EligibilityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRevivalsV1EligibilityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRevivalsV1EligibilityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRevivalsV1EligibilityOK creates a GetLolRevivalsV1EligibilityOK with default headers values
func NewGetLolRevivalsV1EligibilityOK() *GetLolRevivalsV1EligibilityOK {
	return &GetLolRevivalsV1EligibilityOK{}
}

/*GetLolRevivalsV1EligibilityOK handles this case with default header values.

Successful response
*/
type GetLolRevivalsV1EligibilityOK struct {
	Payload *models.LolRevivalsEligibility
}

func (o *GetLolRevivalsV1EligibilityOK) Error() string {
	return fmt.Sprintf("[GET /lol-revivals/v1/eligibility][%d] getLolRevivalsV1EligibilityOK  %+v", 200, o.Payload)
}

func (o *GetLolRevivalsV1EligibilityOK) GetPayload() *models.LolRevivalsEligibility {
	return o.Payload
}

func (o *GetLolRevivalsV1EligibilityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRevivalsEligibility)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}