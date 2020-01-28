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

// PostLolLoginV1SummonerCreatedReader is a Reader for the PostLolLoginV1SummonerCreated structure.
type PostLolLoginV1SummonerCreatedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolLoginV1SummonerCreatedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolLoginV1SummonerCreatedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolLoginV1SummonerCreatedOK creates a PostLolLoginV1SummonerCreatedOK with default headers values
func NewPostLolLoginV1SummonerCreatedOK() *PostLolLoginV1SummonerCreatedOK {
	return &PostLolLoginV1SummonerCreatedOK{}
}

/*PostLolLoginV1SummonerCreatedOK handles this case with default header values.

Successful response
*/
type PostLolLoginV1SummonerCreatedOK struct {
	Payload interface{}
}

func (o *PostLolLoginV1SummonerCreatedOK) Error() string {
	return fmt.Sprintf("[POST /lol-login/v1/summoner-created][%d] postLolLoginV1SummonerCreatedOK  %+v", 200, o.Payload)
}

func (o *PostLolLoginV1SummonerCreatedOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolLoginV1SummonerCreatedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}