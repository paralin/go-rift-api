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

// PostLolSummonerV1SummonersReader is a Reader for the PostLolSummonerV1Summoners structure.
type PostLolSummonerV1SummonersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolSummonerV1SummonersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolSummonerV1SummonersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolSummonerV1SummonersOK creates a PostLolSummonerV1SummonersOK with default headers values
func NewPostLolSummonerV1SummonersOK() *PostLolSummonerV1SummonersOK {
	return &PostLolSummonerV1SummonersOK{}
}

/*PostLolSummonerV1SummonersOK handles this case with default header values.

Successful response
*/
type PostLolSummonerV1SummonersOK struct {
	Payload *models.LolSummonerInternalSummoner
}

func (o *PostLolSummonerV1SummonersOK) Error() string {
	return fmt.Sprintf("[POST /lol-summoner/v1/summoners][%d] postLolSummonerV1SummonersOK  %+v", 200, o.Payload)
}

func (o *PostLolSummonerV1SummonersOK) GetPayload() *models.LolSummonerInternalSummoner {
	return o.Payload
}

func (o *PostLolSummonerV1SummonersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolSummonerInternalSummoner)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
