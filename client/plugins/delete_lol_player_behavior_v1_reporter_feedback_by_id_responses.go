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

// DeleteLolPlayerBehaviorV1ReporterFeedbackByIDReader is a Reader for the DeleteLolPlayerBehaviorV1ReporterFeedbackByID structure.
type DeleteLolPlayerBehaviorV1ReporterFeedbackByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolPlayerBehaviorV1ReporterFeedbackByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK creates a DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK with default headers values
func NewDeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK() *DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK {
	return &DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK{}
}

/*DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK handles this case with default header values.

Successful response
*/
type DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK struct {
	Payload *models.LolPlayerBehaviorReporterFeedback
}

func (o *DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /lol-player-behavior/v1/reporter-feedback/{id}][%d] deleteLolPlayerBehaviorV1ReporterFeedbackByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK) GetPayload() *models.LolPlayerBehaviorReporterFeedback {
	return o.Payload
}

func (o *DeleteLolPlayerBehaviorV1ReporterFeedbackByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPlayerBehaviorReporterFeedback)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
