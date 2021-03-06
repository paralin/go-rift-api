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

// PostLolSummonerV2SummonersPuuidReader is a Reader for the PostLolSummonerV2SummonersPuuid structure.
type PostLolSummonerV2SummonersPuuidReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolSummonerV2SummonersPuuidReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolSummonerV2SummonersPuuidOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolSummonerV2SummonersPuuidOK creates a PostLolSummonerV2SummonersPuuidOK with default headers values
func NewPostLolSummonerV2SummonersPuuidOK() *PostLolSummonerV2SummonersPuuidOK {
	return &PostLolSummonerV2SummonersPuuidOK{}
}

/*PostLolSummonerV2SummonersPuuidOK handles this case with default header values.

Successful response
*/
type PostLolSummonerV2SummonersPuuidOK struct {
	Payload []*models.LolSummonerSummoner
}

func (o *PostLolSummonerV2SummonersPuuidOK) Error() string {
	return fmt.Sprintf("[POST /lol-summoner/v2/summoners/puuid][%d] postLolSummonerV2SummonersPuuidOK  %+v", 200, o.Payload)
}

func (o *PostLolSummonerV2SummonersPuuidOK) GetPayload() []*models.LolSummonerSummoner {
	return o.Payload
}

func (o *PostLolSummonerV2SummonersPuuidOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
