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

// GetLolPftV2SurveyReader is a Reader for the GetLolPftV2Survey structure.
type GetLolPftV2SurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPftV2SurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPftV2SurveyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPftV2SurveyOK creates a GetLolPftV2SurveyOK with default headers values
func NewGetLolPftV2SurveyOK() *GetLolPftV2SurveyOK {
	return &GetLolPftV2SurveyOK{}
}

/*GetLolPftV2SurveyOK handles this case with default header values.

Successful response
*/
type GetLolPftV2SurveyOK struct {
	Payload *models.LolPftPFTSurvey
}

func (o *GetLolPftV2SurveyOK) Error() string {
	return fmt.Sprintf("[GET /lol-pft/v2/survey][%d] getLolPftV2SurveyOK  %+v", 200, o.Payload)
}

func (o *GetLolPftV2SurveyOK) GetPayload() *models.LolPftPFTSurvey {
	return o.Payload
}

func (o *GetLolPftV2SurveyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPftPFTSurvey)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}