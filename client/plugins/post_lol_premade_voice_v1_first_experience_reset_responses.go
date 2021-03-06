// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolPremadeVoiceV1FirstExperienceResetReader is a Reader for the PostLolPremadeVoiceV1FirstExperienceReset structure.
type PostLolPremadeVoiceV1FirstExperienceResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPremadeVoiceV1FirstExperienceResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolPremadeVoiceV1FirstExperienceResetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPremadeVoiceV1FirstExperienceResetNoContent creates a PostLolPremadeVoiceV1FirstExperienceResetNoContent with default headers values
func NewPostLolPremadeVoiceV1FirstExperienceResetNoContent() *PostLolPremadeVoiceV1FirstExperienceResetNoContent {
	return &PostLolPremadeVoiceV1FirstExperienceResetNoContent{}
}

/*PostLolPremadeVoiceV1FirstExperienceResetNoContent handles this case with default header values.

No content
*/
type PostLolPremadeVoiceV1FirstExperienceResetNoContent struct {
}

func (o *PostLolPremadeVoiceV1FirstExperienceResetNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-premade-voice/v1/first-experience/reset][%d] postLolPremadeVoiceV1FirstExperienceResetNoContent ", 204)
}

func (o *PostLolPremadeVoiceV1FirstExperienceResetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
