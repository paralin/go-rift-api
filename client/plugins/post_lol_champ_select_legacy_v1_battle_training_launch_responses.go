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

// PostLolChampSelectLegacyV1BattleTrainingLaunchReader is a Reader for the PostLolChampSelectLegacyV1BattleTrainingLaunch structure.
type PostLolChampSelectLegacyV1BattleTrainingLaunchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolChampSelectLegacyV1BattleTrainingLaunchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolChampSelectLegacyV1BattleTrainingLaunchOK creates a PostLolChampSelectLegacyV1BattleTrainingLaunchOK with default headers values
func NewPostLolChampSelectLegacyV1BattleTrainingLaunchOK() *PostLolChampSelectLegacyV1BattleTrainingLaunchOK {
	return &PostLolChampSelectLegacyV1BattleTrainingLaunchOK{}
}

/*PostLolChampSelectLegacyV1BattleTrainingLaunchOK handles this case with default header values.

Successful response
*/
type PostLolChampSelectLegacyV1BattleTrainingLaunchOK struct {
	Payload interface{}
}

func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchOK) Error() string {
	return fmt.Sprintf("[POST /lol-champ-select-legacy/v1/battle-training/launch][%d] postLolChampSelectLegacyV1BattleTrainingLaunchOK  %+v", 200, o.Payload)
}

func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolChampSelectLegacyV1BattleTrainingLaunchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}