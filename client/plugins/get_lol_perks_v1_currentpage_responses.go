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

// GetLolPerksV1CurrentpageReader is a Reader for the GetLolPerksV1Currentpage structure.
type GetLolPerksV1CurrentpageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPerksV1CurrentpageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPerksV1CurrentpageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPerksV1CurrentpageOK creates a GetLolPerksV1CurrentpageOK with default headers values
func NewGetLolPerksV1CurrentpageOK() *GetLolPerksV1CurrentpageOK {
	return &GetLolPerksV1CurrentpageOK{}
}

/*GetLolPerksV1CurrentpageOK handles this case with default header values.

Successful response
*/
type GetLolPerksV1CurrentpageOK struct {
	Payload *models.LolPerksPerkPageResource
}

func (o *GetLolPerksV1CurrentpageOK) Error() string {
	return fmt.Sprintf("[GET /lol-perks/v1/currentpage][%d] getLolPerksV1CurrentpageOK  %+v", 200, o.Payload)
}

func (o *GetLolPerksV1CurrentpageOK) GetPayload() *models.LolPerksPerkPageResource {
	return o.Payload
}

func (o *GetLolPerksV1CurrentpageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPerksPerkPageResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}