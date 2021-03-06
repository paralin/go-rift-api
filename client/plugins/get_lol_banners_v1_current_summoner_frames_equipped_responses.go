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

// GetLolBannersV1CurrentSummonerFramesEquippedReader is a Reader for the GetLolBannersV1CurrentSummonerFramesEquipped structure.
type GetLolBannersV1CurrentSummonerFramesEquippedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolBannersV1CurrentSummonerFramesEquippedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolBannersV1CurrentSummonerFramesEquippedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolBannersV1CurrentSummonerFramesEquippedOK creates a GetLolBannersV1CurrentSummonerFramesEquippedOK with default headers values
func NewGetLolBannersV1CurrentSummonerFramesEquippedOK() *GetLolBannersV1CurrentSummonerFramesEquippedOK {
	return &GetLolBannersV1CurrentSummonerFramesEquippedOK{}
}

/*GetLolBannersV1CurrentSummonerFramesEquippedOK handles this case with default header values.

Successful response
*/
type GetLolBannersV1CurrentSummonerFramesEquippedOK struct {
	Payload *models.LolBannersBannerFrame
}

func (o *GetLolBannersV1CurrentSummonerFramesEquippedOK) Error() string {
	return fmt.Sprintf("[GET /lol-banners/v1/current-summoner/frames/equipped][%d] getLolBannersV1CurrentSummonerFramesEquippedOK  %+v", 200, o.Payload)
}

func (o *GetLolBannersV1CurrentSummonerFramesEquippedOK) GetPayload() *models.LolBannersBannerFrame {
	return o.Payload
}

func (o *GetLolBannersV1CurrentSummonerFramesEquippedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolBannersBannerFrame)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
