// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PutLolPremadeVoiceV1SelfInputModeReader is a Reader for the PutLolPremadeVoiceV1SelfInputMode structure.
type PutLolPremadeVoiceV1SelfInputModeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolPremadeVoiceV1SelfInputModeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLolPremadeVoiceV1SelfInputModeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolPremadeVoiceV1SelfInputModeNoContent creates a PutLolPremadeVoiceV1SelfInputModeNoContent with default headers values
func NewPutLolPremadeVoiceV1SelfInputModeNoContent() *PutLolPremadeVoiceV1SelfInputModeNoContent {
	return &PutLolPremadeVoiceV1SelfInputModeNoContent{}
}

/*PutLolPremadeVoiceV1SelfInputModeNoContent handles this case with default header values.

No content
*/
type PutLolPremadeVoiceV1SelfInputModeNoContent struct {
}

func (o *PutLolPremadeVoiceV1SelfInputModeNoContent) Error() string {
	return fmt.Sprintf("[PUT /lol-premade-voice/v1/self/inputMode][%d] putLolPremadeVoiceV1SelfInputModeNoContent ", 204)
}

func (o *PutLolPremadeVoiceV1SelfInputModeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
