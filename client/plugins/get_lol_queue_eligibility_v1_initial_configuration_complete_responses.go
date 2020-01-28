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

// GetLolQueueEligibilityV1InitialConfigurationCompleteReader is a Reader for the GetLolQueueEligibilityV1InitialConfigurationComplete structure.
type GetLolQueueEligibilityV1InitialConfigurationCompleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolQueueEligibilityV1InitialConfigurationCompleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolQueueEligibilityV1InitialConfigurationCompleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolQueueEligibilityV1InitialConfigurationCompleteOK creates a GetLolQueueEligibilityV1InitialConfigurationCompleteOK with default headers values
func NewGetLolQueueEligibilityV1InitialConfigurationCompleteOK() *GetLolQueueEligibilityV1InitialConfigurationCompleteOK {
	return &GetLolQueueEligibilityV1InitialConfigurationCompleteOK{}
}

/*GetLolQueueEligibilityV1InitialConfigurationCompleteOK handles this case with default header values.

Successful response
*/
type GetLolQueueEligibilityV1InitialConfigurationCompleteOK struct {
	Payload bool
}

func (o *GetLolQueueEligibilityV1InitialConfigurationCompleteOK) Error() string {
	return fmt.Sprintf("[GET /lol-queue-eligibility/v1/initial-configuration-complete][%d] getLolQueueEligibilityV1InitialConfigurationCompleteOK  %+v", 200, o.Payload)
}

func (o *GetLolQueueEligibilityV1InitialConfigurationCompleteOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolQueueEligibilityV1InitialConfigurationCompleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
