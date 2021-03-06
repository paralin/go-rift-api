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

// GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionReader is a Reader for the GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPosition structure.
type GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK creates a GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK with default headers values
func NewGetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK() *GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK {
	return &GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK{}
}

/*GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK handles this case with default header values.

Successful response
*/
type GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK struct {
	Payload interface{}
}

func (o *GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK) Error() string {
	return fmt.Sprintf("[GET /lol-career-stats/v1/summoner-stats/{puuid}/{season}/{queue}/{position}][%d] getLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK  %+v", 200, o.Payload)
}

func (o *GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolCareerStatsV1SummonerStatsByPuuidBySeasonByQueueByPositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
