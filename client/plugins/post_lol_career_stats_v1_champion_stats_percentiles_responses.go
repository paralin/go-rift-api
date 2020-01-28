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

// PostLolCareerStatsV1ChampionStatsPercentilesReader is a Reader for the PostLolCareerStatsV1ChampionStatsPercentiles structure.
type PostLolCareerStatsV1ChampionStatsPercentilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolCareerStatsV1ChampionStatsPercentilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolCareerStatsV1ChampionStatsPercentilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolCareerStatsV1ChampionStatsPercentilesOK creates a PostLolCareerStatsV1ChampionStatsPercentilesOK with default headers values
func NewPostLolCareerStatsV1ChampionStatsPercentilesOK() *PostLolCareerStatsV1ChampionStatsPercentilesOK {
	return &PostLolCareerStatsV1ChampionStatsPercentilesOK{}
}

/*PostLolCareerStatsV1ChampionStatsPercentilesOK handles this case with default header values.

Successful response
*/
type PostLolCareerStatsV1ChampionStatsPercentilesOK struct {
	Payload []*models.LolCareerStatsStatisticsPercentilesResponse
}

func (o *PostLolCareerStatsV1ChampionStatsPercentilesOK) Error() string {
	return fmt.Sprintf("[POST /lol-career-stats/v1/champion-stats-percentiles][%d] postLolCareerStatsV1ChampionStatsPercentilesOK  %+v", 200, o.Payload)
}

func (o *PostLolCareerStatsV1ChampionStatsPercentilesOK) GetPayload() []*models.LolCareerStatsStatisticsPercentilesResponse {
	return o.Payload
}

func (o *PostLolCareerStatsV1ChampionStatsPercentilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}