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

// GetLolRsoAuthV1AuthorizationUserinfoReader is a Reader for the GetLolRsoAuthV1AuthorizationUserinfo structure.
type GetLolRsoAuthV1AuthorizationUserinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolRsoAuthV1AuthorizationUserinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolRsoAuthV1AuthorizationUserinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolRsoAuthV1AuthorizationUserinfoOK creates a GetLolRsoAuthV1AuthorizationUserinfoOK with default headers values
func NewGetLolRsoAuthV1AuthorizationUserinfoOK() *GetLolRsoAuthV1AuthorizationUserinfoOK {
	return &GetLolRsoAuthV1AuthorizationUserinfoOK{}
}

/*GetLolRsoAuthV1AuthorizationUserinfoOK handles this case with default header values.

Successful response
*/
type GetLolRsoAuthV1AuthorizationUserinfoOK struct {
	Payload *models.LolRsoAuthUserInfo
}

func (o *GetLolRsoAuthV1AuthorizationUserinfoOK) Error() string {
	return fmt.Sprintf("[GET /lol-rso-auth/v1/authorization/userinfo][%d] getLolRsoAuthV1AuthorizationUserinfoOK  %+v", 200, o.Payload)
}

func (o *GetLolRsoAuthV1AuthorizationUserinfoOK) GetPayload() *models.LolRsoAuthUserInfo {
	return o.Payload
}

func (o *GetLolRsoAuthV1AuthorizationUserinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolRsoAuthUserInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
