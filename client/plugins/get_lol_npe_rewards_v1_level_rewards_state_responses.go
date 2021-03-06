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

// GetLolNpeRewardsV1LevelRewardsStateReader is a Reader for the GetLolNpeRewardsV1LevelRewardsState structure.
type GetLolNpeRewardsV1LevelRewardsStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolNpeRewardsV1LevelRewardsStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolNpeRewardsV1LevelRewardsStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolNpeRewardsV1LevelRewardsStateOK creates a GetLolNpeRewardsV1LevelRewardsStateOK with default headers values
func NewGetLolNpeRewardsV1LevelRewardsStateOK() *GetLolNpeRewardsV1LevelRewardsStateOK {
	return &GetLolNpeRewardsV1LevelRewardsStateOK{}
}

/*GetLolNpeRewardsV1LevelRewardsStateOK handles this case with default header values.

Successful response
*/
type GetLolNpeRewardsV1LevelRewardsStateOK struct {
	Payload *models.LolNpeRewardsRewardSeriesState
}

func (o *GetLolNpeRewardsV1LevelRewardsStateOK) Error() string {
	return fmt.Sprintf("[GET /lol-npe-rewards/v1/level-rewards/state][%d] getLolNpeRewardsV1LevelRewardsStateOK  %+v", 200, o.Payload)
}

func (o *GetLolNpeRewardsV1LevelRewardsStateOK) GetPayload() *models.LolNpeRewardsRewardSeriesState {
	return o.Payload
}

func (o *GetLolNpeRewardsV1LevelRewardsStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolNpeRewardsRewardSeriesState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
