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

// GetLolPerksV1CustomizationlimitsReader is a Reader for the GetLolPerksV1Customizationlimits structure.
type GetLolPerksV1CustomizationlimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPerksV1CustomizationlimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPerksV1CustomizationlimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPerksV1CustomizationlimitsOK creates a GetLolPerksV1CustomizationlimitsOK with default headers values
func NewGetLolPerksV1CustomizationlimitsOK() *GetLolPerksV1CustomizationlimitsOK {
	return &GetLolPerksV1CustomizationlimitsOK{}
}

/*GetLolPerksV1CustomizationlimitsOK handles this case with default header values.

Successful response
*/
type GetLolPerksV1CustomizationlimitsOK struct {
	Payload models.LolPerksCustomizationLimits
}

func (o *GetLolPerksV1CustomizationlimitsOK) Error() string {
	return fmt.Sprintf("[GET /lol-perks/v1/customizationlimits][%d] getLolPerksV1CustomizationlimitsOK  %+v", 200, o.Payload)
}

func (o *GetLolPerksV1CustomizationlimitsOK) GetPayload() models.LolPerksCustomizationLimits {
	return o.Payload
}

func (o *GetLolPerksV1CustomizationlimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}