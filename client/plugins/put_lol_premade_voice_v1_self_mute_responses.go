// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutLolPremadeVoiceV1SelfMuteReader is a Reader for the PutLolPremadeVoiceV1SelfMute structure.
type PutLolPremadeVoiceV1SelfMuteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolPremadeVoiceV1SelfMuteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLolPremadeVoiceV1SelfMuteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolPremadeVoiceV1SelfMuteNoContent creates a PutLolPremadeVoiceV1SelfMuteNoContent with default headers values
func NewPutLolPremadeVoiceV1SelfMuteNoContent() *PutLolPremadeVoiceV1SelfMuteNoContent {
	return &PutLolPremadeVoiceV1SelfMuteNoContent{}
}

/*PutLolPremadeVoiceV1SelfMuteNoContent handles this case with default header values.

No content
*/
type PutLolPremadeVoiceV1SelfMuteNoContent struct {
}

func (o *PutLolPremadeVoiceV1SelfMuteNoContent) Error() string {
	return fmt.Sprintf("[PUT /lol-premade-voice/v1/self/mute][%d] putLolPremadeVoiceV1SelfMuteNoContent ", 204)
}

func (o *PutLolPremadeVoiceV1SelfMuteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
