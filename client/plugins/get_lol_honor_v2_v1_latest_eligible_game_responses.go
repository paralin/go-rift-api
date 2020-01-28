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

// GetLolHonorV2V1LatestEligibleGameReader is a Reader for the GetLolHonorV2V1LatestEligibleGame structure.
type GetLolHonorV2V1LatestEligibleGameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolHonorV2V1LatestEligibleGameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolHonorV2V1LatestEligibleGameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolHonorV2V1LatestEligibleGameOK creates a GetLolHonorV2V1LatestEligibleGameOK with default headers values
func NewGetLolHonorV2V1LatestEligibleGameOK() *GetLolHonorV2V1LatestEligibleGameOK {
	return &GetLolHonorV2V1LatestEligibleGameOK{}
}

/*GetLolHonorV2V1LatestEligibleGameOK handles this case with default header values.

Successful response
*/
type GetLolHonorV2V1LatestEligibleGameOK struct {
	Payload int64
}

func (o *GetLolHonorV2V1LatestEligibleGameOK) Error() string {
	return fmt.Sprintf("[GET /lol-honor-v2/v1/latest-eligible-game][%d] getLolHonorV2V1LatestEligibleGameOK  %+v", 200, o.Payload)
}

func (o *GetLolHonorV2V1LatestEligibleGameOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetLolHonorV2V1LatestEligibleGameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
