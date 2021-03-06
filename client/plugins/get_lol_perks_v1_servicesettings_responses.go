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

// GetLolPerksV1ServicesettingsReader is a Reader for the GetLolPerksV1Servicesettings structure.
type GetLolPerksV1ServicesettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPerksV1ServicesettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPerksV1ServicesettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPerksV1ServicesettingsOK creates a GetLolPerksV1ServicesettingsOK with default headers values
func NewGetLolPerksV1ServicesettingsOK() *GetLolPerksV1ServicesettingsOK {
	return &GetLolPerksV1ServicesettingsOK{}
}

/*GetLolPerksV1ServicesettingsOK handles this case with default header values.

Successful response
*/
type GetLolPerksV1ServicesettingsOK struct {
	Payload *models.LolPerksServiceSettings
}

func (o *GetLolPerksV1ServicesettingsOK) Error() string {
	return fmt.Sprintf("[GET /lol-perks/v1/servicesettings][%d] getLolPerksV1ServicesettingsOK  %+v", 200, o.Payload)
}

func (o *GetLolPerksV1ServicesettingsOK) GetPayload() *models.LolPerksServiceSettings {
	return o.Payload
}

func (o *GetLolPerksV1ServicesettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPerksServiceSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
