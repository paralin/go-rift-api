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

// PostLolSummonerV1CurrentSummonerSummonerProfileReader is a Reader for the PostLolSummonerV1CurrentSummonerSummonerProfile structure.
type PostLolSummonerV1CurrentSummonerSummonerProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolSummonerV1CurrentSummonerSummonerProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolSummonerV1CurrentSummonerSummonerProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolSummonerV1CurrentSummonerSummonerProfileOK creates a PostLolSummonerV1CurrentSummonerSummonerProfileOK with default headers values
func NewPostLolSummonerV1CurrentSummonerSummonerProfileOK() *PostLolSummonerV1CurrentSummonerSummonerProfileOK {
	return &PostLolSummonerV1CurrentSummonerSummonerProfileOK{}
}

/*PostLolSummonerV1CurrentSummonerSummonerProfileOK handles this case with default header values.

Successful response
*/
type PostLolSummonerV1CurrentSummonerSummonerProfileOK struct {
	Payload interface{}
}

func (o *PostLolSummonerV1CurrentSummonerSummonerProfileOK) Error() string {
	return fmt.Sprintf("[POST /lol-summoner/v1/current-summoner/summoner-profile][%d] postLolSummonerV1CurrentSummonerSummonerProfileOK  %+v", 200, o.Payload)
}

func (o *PostLolSummonerV1CurrentSummonerSummonerProfileOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolSummonerV1CurrentSummonerSummonerProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
