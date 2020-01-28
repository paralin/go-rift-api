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

// GetLolClubsPublicV1ClubsPublicBySummonerIDTagReader is a Reader for the GetLolClubsPublicV1ClubsPublicBySummonerIDTag structure.
type GetLolClubsPublicV1ClubsPublicBySummonerIDTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClubsPublicV1ClubsPublicBySummonerIDTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClubsPublicV1ClubsPublicBySummonerIDTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClubsPublicV1ClubsPublicBySummonerIDTagOK creates a GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK with default headers values
func NewGetLolClubsPublicV1ClubsPublicBySummonerIDTagOK() *GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK {
	return &GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK{}
}

/*GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK handles this case with default header values.

Successful response
*/
type GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK struct {
	Payload *models.LolClubsPublicClubTag
}

func (o *GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK) Error() string {
	return fmt.Sprintf("[GET /lol-clubs-public/v1/clubs/public/{summonerId}/tag][%d] getLolClubsPublicV1ClubsPublicBySummonerIdTagOK  %+v", 200, o.Payload)
}

func (o *GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK) GetPayload() *models.LolClubsPublicClubTag {
	return o.Payload
}

func (o *GetLolClubsPublicV1ClubsPublicBySummonerIDTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolClubsPublicClubTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
