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

// GetLolRsoAuthV1StatusByPlatformIDReader is a Reader for the GetLolRsoAuthV1StatusByPlatformID structure.
type GetLolRsoAuthV1StatusByPlatformIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRsoAuthV1StatusByPlatformIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRsoAuthV1StatusByPlatformIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRsoAuthV1StatusByPlatformIDOK creates a GetLolRsoAuthV1StatusByPlatformIDOK with default headers values
func NewGetLolRsoAuthV1StatusByPlatformIDOK() *GetLolRsoAuthV1StatusByPlatformIDOK {
	return &GetLolRsoAuthV1StatusByPlatformIDOK{}
}

/*GetLolRsoAuthV1StatusByPlatformIDOK handles this case with default header values.

Successful response
*/
type GetLolRsoAuthV1StatusByPlatformIDOK struct {
	Payload *models.LolRsoAuthRegionStatus
}

func (o *GetLolRsoAuthV1StatusByPlatformIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-rso-auth/v1/status/{platformId}][%d] getLolRsoAuthV1StatusByPlatformIdOK  %+v", 200, o.Payload)
}

func (o *GetLolRsoAuthV1StatusByPlatformIDOK) GetPayload() *models.LolRsoAuthRegionStatus {
	return o.Payload
}

func (o *GetLolRsoAuthV1StatusByPlatformIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRsoAuthRegionStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}