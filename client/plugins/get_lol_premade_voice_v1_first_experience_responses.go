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

// GetLolPremadeVoiceV1FirstExperienceReader is a Reader for the GetLolPremadeVoiceV1FirstExperience structure.
type GetLolPremadeVoiceV1FirstExperienceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPremadeVoiceV1FirstExperienceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPremadeVoiceV1FirstExperienceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPremadeVoiceV1FirstExperienceOK creates a GetLolPremadeVoiceV1FirstExperienceOK with default headers values
func NewGetLolPremadeVoiceV1FirstExperienceOK() *GetLolPremadeVoiceV1FirstExperienceOK {
	return &GetLolPremadeVoiceV1FirstExperienceOK{}
}

/*GetLolPremadeVoiceV1FirstExperienceOK handles this case with default header values.

Successful response
*/
type GetLolPremadeVoiceV1FirstExperienceOK struct {
	Payload *models.LolPremadeVoiceFirstExperience
}

func (o *GetLolPremadeVoiceV1FirstExperienceOK) Error() string {
	return fmt.Sprintf("[GET /lol-premade-voice/v1/first-experience][%d] getLolPremadeVoiceV1FirstExperienceOK  %+v", 200, o.Payload)
}

func (o *GetLolPremadeVoiceV1FirstExperienceOK) GetPayload() *models.LolPremadeVoiceFirstExperience {
	return o.Payload
}

func (o *GetLolPremadeVoiceV1FirstExperienceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPremadeVoiceFirstExperience)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
