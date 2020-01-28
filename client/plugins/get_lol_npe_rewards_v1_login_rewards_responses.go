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

// GetLolNpeRewardsV1LoginRewardsReader is a Reader for the GetLolNpeRewardsV1LoginRewards structure.
type GetLolNpeRewardsV1LoginRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolNpeRewardsV1LoginRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolNpeRewardsV1LoginRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolNpeRewardsV1LoginRewardsOK creates a GetLolNpeRewardsV1LoginRewardsOK with default headers values
func NewGetLolNpeRewardsV1LoginRewardsOK() *GetLolNpeRewardsV1LoginRewardsOK {
	return &GetLolNpeRewardsV1LoginRewardsOK{}
}

/*GetLolNpeRewardsV1LoginRewardsOK handles this case with default header values.

Successful response
*/
type GetLolNpeRewardsV1LoginRewardsOK struct {
	Payload *models.LolNpeRewardsRewardSeries
}

func (o *GetLolNpeRewardsV1LoginRewardsOK) Error() string {
	return fmt.Sprintf("[GET /lol-npe-rewards/v1/login-rewards][%d] getLolNpeRewardsV1LoginRewardsOK  %+v", 200, o.Payload)
}

func (o *GetLolNpeRewardsV1LoginRewardsOK) GetPayload() *models.LolNpeRewardsRewardSeries {
	return o.Payload
}

func (o *GetLolNpeRewardsV1LoginRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolNpeRewardsRewardSeries)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}