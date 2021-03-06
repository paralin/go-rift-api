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

// GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueReader is a Reader for the GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueue structure.
type GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK creates a GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK with default headers values
func NewGetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK() *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK {
	return &GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK{}
}

/*GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK handles this case with default header values.

Successful response
*/
type GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK struct {
	Payload *models.LolCareerStatsChampionQueueStatsResponse
}

func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK) Error() string {
	return fmt.Sprintf("[GET /lol-career-stats/v1/champion-averages/season/{season}/{championId}/{position}/{tier}/{queue}][%d] getLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIdByPositionByTierByQueueOK  %+v", 200, o.Payload)
}

func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK) GetPayload() *models.LolCareerStatsChampionQueueStatsResponse {
	return o.Payload
}

func (o *GetLolCareerStatsV1ChampionAveragesSeasonBySeasonByChampionIDByPositionByTierByQueueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolCareerStatsChampionQueueStatsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
