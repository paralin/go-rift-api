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

// GetLolQueueEligibilityV3CustomReader is a Reader for the GetLolQueueEligibilityV3Custom structure.
type GetLolQueueEligibilityV3CustomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolQueueEligibilityV3CustomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolQueueEligibilityV3CustomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolQueueEligibilityV3CustomOK creates a GetLolQueueEligibilityV3CustomOK with default headers values
func NewGetLolQueueEligibilityV3CustomOK() *GetLolQueueEligibilityV3CustomOK {
	return &GetLolQueueEligibilityV3CustomOK{}
}

/*GetLolQueueEligibilityV3CustomOK handles this case with default header values.

Successful response
*/
type GetLolQueueEligibilityV3CustomOK struct {
	Payload []*models.LolQueueEligibilityEligibility
}

func (o *GetLolQueueEligibilityV3CustomOK) Error() string {
	return fmt.Sprintf("[GET /lol-queue-eligibility/v3/custom][%d] getLolQueueEligibilityV3CustomOK  %+v", 200, o.Payload)
}

func (o *GetLolQueueEligibilityV3CustomOK) GetPayload() []*models.LolQueueEligibilityEligibility {
	return o.Payload
}

func (o *GetLolQueueEligibilityV3CustomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
