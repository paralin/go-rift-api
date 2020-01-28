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

// GetLolPlayerBehaviorV1ReporterFeedbackByIDReader is a Reader for the GetLolPlayerBehaviorV1ReporterFeedbackByID structure.
type GetLolPlayerBehaviorV1ReporterFeedbackByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPlayerBehaviorV1ReporterFeedbackByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPlayerBehaviorV1ReporterFeedbackByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPlayerBehaviorV1ReporterFeedbackByIDOK creates a GetLolPlayerBehaviorV1ReporterFeedbackByIDOK with default headers values
func NewGetLolPlayerBehaviorV1ReporterFeedbackByIDOK() *GetLolPlayerBehaviorV1ReporterFeedbackByIDOK {
	return &GetLolPlayerBehaviorV1ReporterFeedbackByIDOK{}
}

/*GetLolPlayerBehaviorV1ReporterFeedbackByIDOK handles this case with default header values.

Successful response
*/
type GetLolPlayerBehaviorV1ReporterFeedbackByIDOK struct {
	Payload *models.LolPlayerBehaviorReporterFeedback
}

func (o *GetLolPlayerBehaviorV1ReporterFeedbackByIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-player-behavior/v1/reporter-feedback/{id}][%d] getLolPlayerBehaviorV1ReporterFeedbackByIdOK  %+v", 200, o.Payload)
}

func (o *GetLolPlayerBehaviorV1ReporterFeedbackByIDOK) GetPayload() *models.LolPlayerBehaviorReporterFeedback {
	return o.Payload
}

func (o *GetLolPlayerBehaviorV1ReporterFeedbackByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPlayerBehaviorReporterFeedback)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}