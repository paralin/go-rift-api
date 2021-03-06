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

// GetLolTrophiesV1CurrentSummonerTrophiesProfileReader is a Reader for the GetLolTrophiesV1CurrentSummonerTrophiesProfile structure.
type GetLolTrophiesV1CurrentSummonerTrophiesProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolTrophiesV1CurrentSummonerTrophiesProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolTrophiesV1CurrentSummonerTrophiesProfileOK creates a GetLolTrophiesV1CurrentSummonerTrophiesProfileOK with default headers values
func NewGetLolTrophiesV1CurrentSummonerTrophiesProfileOK() *GetLolTrophiesV1CurrentSummonerTrophiesProfileOK {
	return &GetLolTrophiesV1CurrentSummonerTrophiesProfileOK{}
}

/*GetLolTrophiesV1CurrentSummonerTrophiesProfileOK handles this case with default header values.

Successful response
*/
type GetLolTrophiesV1CurrentSummonerTrophiesProfileOK struct {
	Payload *models.LolTrophiesTrophyProfileData
}

func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileOK) Error() string {
	return fmt.Sprintf("[GET /lol-trophies/v1/current-summoner/trophies/profile][%d] getLolTrophiesV1CurrentSummonerTrophiesProfileOK  %+v", 200, o.Payload)
}

func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileOK) GetPayload() *models.LolTrophiesTrophyProfileData {
	return o.Payload
}

func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolTrophiesTrophyProfileData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
