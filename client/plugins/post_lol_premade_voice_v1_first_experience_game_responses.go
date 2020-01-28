// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolPremadeVoiceV1FirstExperienceGameReader is a Reader for the PostLolPremadeVoiceV1FirstExperienceGame structure.
type PostLolPremadeVoiceV1FirstExperienceGameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPremadeVoiceV1FirstExperienceGameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolPremadeVoiceV1FirstExperienceGameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPremadeVoiceV1FirstExperienceGameNoContent creates a PostLolPremadeVoiceV1FirstExperienceGameNoContent with default headers values
func NewPostLolPremadeVoiceV1FirstExperienceGameNoContent() *PostLolPremadeVoiceV1FirstExperienceGameNoContent {
	return &PostLolPremadeVoiceV1FirstExperienceGameNoContent{}
}

/*PostLolPremadeVoiceV1FirstExperienceGameNoContent handles this case with default header values.

No content
*/
type PostLolPremadeVoiceV1FirstExperienceGameNoContent struct {
}

func (o *PostLolPremadeVoiceV1FirstExperienceGameNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-premade-voice/v1/first-experience/game][%d] postLolPremadeVoiceV1FirstExperienceGameNoContent ", 204)
}

func (o *PostLolPremadeVoiceV1FirstExperienceGameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
