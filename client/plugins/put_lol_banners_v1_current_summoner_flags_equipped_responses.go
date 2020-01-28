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

// PutLolBannersV1CurrentSummonerFlagsEquippedReader is a Reader for the PutLolBannersV1CurrentSummonerFlagsEquipped structure.
type PutLolBannersV1CurrentSummonerFlagsEquippedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolBannersV1CurrentSummonerFlagsEquippedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolBannersV1CurrentSummonerFlagsEquippedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolBannersV1CurrentSummonerFlagsEquippedOK creates a PutLolBannersV1CurrentSummonerFlagsEquippedOK with default headers values
func NewPutLolBannersV1CurrentSummonerFlagsEquippedOK() *PutLolBannersV1CurrentSummonerFlagsEquippedOK {
	return &PutLolBannersV1CurrentSummonerFlagsEquippedOK{}
}

/*PutLolBannersV1CurrentSummonerFlagsEquippedOK handles this case with default header values.

Successful response
*/
type PutLolBannersV1CurrentSummonerFlagsEquippedOK struct {
	Payload *models.LolBannersBannerFlag
}

func (o *PutLolBannersV1CurrentSummonerFlagsEquippedOK) Error() string {
	return fmt.Sprintf("[PUT /lol-banners/v1/current-summoner/flags/equipped][%d] putLolBannersV1CurrentSummonerFlagsEquippedOK  %+v", 200, o.Payload)
}

func (o *PutLolBannersV1CurrentSummonerFlagsEquippedOK) GetPayload() *models.LolBannersBannerFlag {
	return o.Payload
}

func (o *PutLolBannersV1CurrentSummonerFlagsEquippedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolBannersBannerFlag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}