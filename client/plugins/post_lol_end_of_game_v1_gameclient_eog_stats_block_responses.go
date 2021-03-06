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

// PostLolEndOfGameV1GameclientEogStatsBlockReader is a Reader for the PostLolEndOfGameV1GameclientEogStatsBlock structure.
type PostLolEndOfGameV1GameclientEogStatsBlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolEndOfGameV1GameclientEogStatsBlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolEndOfGameV1GameclientEogStatsBlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolEndOfGameV1GameclientEogStatsBlockOK creates a PostLolEndOfGameV1GameclientEogStatsBlockOK with default headers values
func NewPostLolEndOfGameV1GameclientEogStatsBlockOK() *PostLolEndOfGameV1GameclientEogStatsBlockOK {
	return &PostLolEndOfGameV1GameclientEogStatsBlockOK{}
}

/*PostLolEndOfGameV1GameclientEogStatsBlockOK handles this case with default header values.

Successful response
*/
type PostLolEndOfGameV1GameclientEogStatsBlockOK struct {
	Payload interface{}
}

func (o *PostLolEndOfGameV1GameclientEogStatsBlockOK) Error() string {
	return fmt.Sprintf("[POST /lol-end-of-game/v1/gameclient-eog-stats-block][%d] postLolEndOfGameV1GameclientEogStatsBlockOK  %+v", 200, o.Payload)
}

func (o *PostLolEndOfGameV1GameclientEogStatsBlockOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolEndOfGameV1GameclientEogStatsBlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
