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

// GetLolGameSettingsV1GameSettingsReader is a Reader for the GetLolGameSettingsV1GameSettings structure.
type GetLolGameSettingsV1GameSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolGameSettingsV1GameSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolGameSettingsV1GameSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolGameSettingsV1GameSettingsOK creates a GetLolGameSettingsV1GameSettingsOK with default headers values
func NewGetLolGameSettingsV1GameSettingsOK() *GetLolGameSettingsV1GameSettingsOK {
	return &GetLolGameSettingsV1GameSettingsOK{}
}

/*GetLolGameSettingsV1GameSettingsOK handles this case with default header values.

Successful response
*/
type GetLolGameSettingsV1GameSettingsOK struct {
	Payload interface{}
}

func (o *GetLolGameSettingsV1GameSettingsOK) Error() string {
	return fmt.Sprintf("[GET /lol-game-settings/v1/game-settings][%d] getLolGameSettingsV1GameSettingsOK  %+v", 200, o.Payload)
}

func (o *GetLolGameSettingsV1GameSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolGameSettingsV1GameSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}