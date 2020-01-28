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

// PostLolLoginV1ChangeSummonerNameReader is a Reader for the PostLolLoginV1ChangeSummonerName structure.
type PostLolLoginV1ChangeSummonerNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLoginV1ChangeSummonerNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLoginV1ChangeSummonerNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLoginV1ChangeSummonerNameOK creates a PostLolLoginV1ChangeSummonerNameOK with default headers values
func NewPostLolLoginV1ChangeSummonerNameOK() *PostLolLoginV1ChangeSummonerNameOK {
	return &PostLolLoginV1ChangeSummonerNameOK{}
}

/*PostLolLoginV1ChangeSummonerNameOK handles this case with default header values.

Successful response
*/
type PostLolLoginV1ChangeSummonerNameOK struct {
	Payload interface{}
}

func (o *PostLolLoginV1ChangeSummonerNameOK) Error() string {
	return fmt.Sprintf("[POST /lol-login/v1/change-summoner-name][%d] postLolLoginV1ChangeSummonerNameOK  %+v", 200, o.Payload)
}

func (o *PostLolLoginV1ChangeSummonerNameOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLoginV1ChangeSummonerNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
