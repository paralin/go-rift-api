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

// PatchLolGameSettingsV1GameSettingsReader is a Reader for the PatchLolGameSettingsV1GameSettings structure.
type PatchLolGameSettingsV1GameSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchLolGameSettingsV1GameSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchLolGameSettingsV1GameSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchLolGameSettingsV1GameSettingsOK creates a PatchLolGameSettingsV1GameSettingsOK with default headers values
func NewPatchLolGameSettingsV1GameSettingsOK() *PatchLolGameSettingsV1GameSettingsOK {
	return &PatchLolGameSettingsV1GameSettingsOK{}
}

/*PatchLolGameSettingsV1GameSettingsOK handles this case with default header values.

Successful response
*/
type PatchLolGameSettingsV1GameSettingsOK struct {
	Payload interface{}
}

func (o *PatchLolGameSettingsV1GameSettingsOK) Error() string {
	return fmt.Sprintf("[PATCH /lol-game-settings/v1/game-settings][%d] patchLolGameSettingsV1GameSettingsOK  %+v", 200, o.Payload)
}

func (o *PatchLolGameSettingsV1GameSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PatchLolGameSettingsV1GameSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
