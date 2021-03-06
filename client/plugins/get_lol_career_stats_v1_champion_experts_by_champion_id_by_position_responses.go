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

// GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionReader is a Reader for the GetLolCareerStatsV1ChampionExpertsByChampionIDByPosition structure.
type GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK creates a GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK with default headers values
func NewGetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK() *GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK {
	return &GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK{}
}

/*GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK handles this case with default header values.

Successful response
*/
type GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK struct {
	Payload []*models.LolCareerStatsExpertPlayer
}

func (o *GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK) Error() string {
	return fmt.Sprintf("[GET /lol-career-stats/v1/champion-experts/{championId}/{position}][%d] getLolCareerStatsV1ChampionExpertsByChampionIdByPositionOK  %+v", 200, o.Payload)
}

func (o *GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK) GetPayload() []*models.LolCareerStatsExpertPlayer {
	return o.Payload
}

func (o *GetLolCareerStatsV1ChampionExpertsByChampionIDByPositionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
