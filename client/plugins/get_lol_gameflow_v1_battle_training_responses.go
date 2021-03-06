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

// GetLolGameflowV1BattleTrainingReader is a Reader for the GetLolGameflowV1BattleTraining structure.
type GetLolGameflowV1BattleTrainingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolGameflowV1BattleTrainingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolGameflowV1BattleTrainingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolGameflowV1BattleTrainingOK creates a GetLolGameflowV1BattleTrainingOK with default headers values
func NewGetLolGameflowV1BattleTrainingOK() *GetLolGameflowV1BattleTrainingOK {
	return &GetLolGameflowV1BattleTrainingOK{}
}

/*GetLolGameflowV1BattleTrainingOK handles this case with default header values.

Successful response
*/
type GetLolGameflowV1BattleTrainingOK struct {
	Payload bool
}

func (o *GetLolGameflowV1BattleTrainingOK) Error() string {
	return fmt.Sprintf("[GET /lol-gameflow/v1/battle-training][%d] getLolGameflowV1BattleTrainingOK  %+v", 200, o.Payload)
}

func (o *GetLolGameflowV1BattleTrainingOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolGameflowV1BattleTrainingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
