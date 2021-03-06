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

// GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueReader is a Reader for the GetLolCareerStatsV1PositionAveragesByPositionByTierByQueue structure.
type GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK creates a GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK with default headers values
func NewGetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK() *GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK {
	return &GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK{}
}

/*GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK handles this case with default header values.

Successful response
*/
type GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK struct {
	Payload *models.LolCareerStatsChampionQueueStatsResponse
}

func (o *GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK) Error() string {
	return fmt.Sprintf("[GET /lol-career-stats/v1/position-averages/{position}/{tier}/{queue}][%d] getLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK  %+v", 200, o.Payload)
}

func (o *GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK) GetPayload() *models.LolCareerStatsChampionQueueStatsResponse {
	return o.Payload
}

func (o *GetLolCareerStatsV1PositionAveragesByPositionByTierByQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolCareerStatsChampionQueueStatsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
