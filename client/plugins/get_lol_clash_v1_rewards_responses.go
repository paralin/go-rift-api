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

// GetLolClashV1RewardsReader is a Reader for the GetLolClashV1Rewards structure.
type GetLolClashV1RewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1RewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1RewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1RewardsOK creates a GetLolClashV1RewardsOK with default headers values
func NewGetLolClashV1RewardsOK() *GetLolClashV1RewardsOK {
	return &GetLolClashV1RewardsOK{}
}

/*GetLolClashV1RewardsOK handles this case with default header values.

Successful response
*/
type GetLolClashV1RewardsOK struct {
	Payload *models.LolClashPlayerRewards
}

func (o *GetLolClashV1RewardsOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/rewards][%d] getLolClashV1RewardsOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1RewardsOK) GetPayload() *models.LolClashPlayerRewards {
	return o.Payload
}

func (o *GetLolClashV1RewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolClashPlayerRewards)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
