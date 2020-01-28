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

// GetLolPerksV1SettingsReader is a Reader for the GetLolPerksV1Settings structure.
type GetLolPerksV1SettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPerksV1SettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPerksV1SettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPerksV1SettingsOK creates a GetLolPerksV1SettingsOK with default headers values
func NewGetLolPerksV1SettingsOK() *GetLolPerksV1SettingsOK {
	return &GetLolPerksV1SettingsOK{}
}

/*GetLolPerksV1SettingsOK handles this case with default header values.

Successful response
*/
type GetLolPerksV1SettingsOK struct {
	Payload *models.LolPerksUISettings
}

func (o *GetLolPerksV1SettingsOK) Error() string {
	return fmt.Sprintf("[GET /lol-perks/v1/settings][%d] getLolPerksV1SettingsOK  %+v", 200, o.Payload)
}

func (o *GetLolPerksV1SettingsOK) GetPayload() *models.LolPerksUISettings {
	return o.Payload
}

func (o *GetLolPerksV1SettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPerksUISettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
