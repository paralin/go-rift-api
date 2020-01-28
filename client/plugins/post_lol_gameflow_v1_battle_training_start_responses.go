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

// PostLolGameflowV1BattleTrainingStartReader is a Reader for the PostLolGameflowV1BattleTrainingStart structure.
type PostLolGameflowV1BattleTrainingStartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolGameflowV1BattleTrainingStartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolGameflowV1BattleTrainingStartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolGameflowV1BattleTrainingStartOK creates a PostLolGameflowV1BattleTrainingStartOK with default headers values
func NewPostLolGameflowV1BattleTrainingStartOK() *PostLolGameflowV1BattleTrainingStartOK {
	return &PostLolGameflowV1BattleTrainingStartOK{}
}

/*PostLolGameflowV1BattleTrainingStartOK handles this case with default header values.

Successful response
*/
type PostLolGameflowV1BattleTrainingStartOK struct {
	Payload interface{}
}

func (o *PostLolGameflowV1BattleTrainingStartOK) Error() string {
	return fmt.Sprintf("[POST /lol-gameflow/v1/battle-training/start][%d] postLolGameflowV1BattleTrainingStartOK  %+v", 200, o.Payload)
}

func (o *PostLolGameflowV1BattleTrainingStartOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolGameflowV1BattleTrainingStartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}