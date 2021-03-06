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

// PutLolSettingsV2AccountByPpTypeByCategoryReader is a Reader for the PutLolSettingsV2AccountByPpTypeByCategory structure.
type PutLolSettingsV2AccountByPpTypeByCategoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLolSettingsV2AccountByPpTypeByCategoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutLolSettingsV2AccountByPpTypeByCategoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutLolSettingsV2AccountByPpTypeByCategoryOK creates a PutLolSettingsV2AccountByPpTypeByCategoryOK with default headers values
func NewPutLolSettingsV2AccountByPpTypeByCategoryOK() *PutLolSettingsV2AccountByPpTypeByCategoryOK {
	return &PutLolSettingsV2AccountByPpTypeByCategoryOK{}
}

/*PutLolSettingsV2AccountByPpTypeByCategoryOK handles this case with default header values.

Successful response
*/
type PutLolSettingsV2AccountByPpTypeByCategoryOK struct {
	Payload interface{}
}

func (o *PutLolSettingsV2AccountByPpTypeByCategoryOK) Error() string {
	return fmt.Sprintf("[PUT /lol-settings/v2/account/{ppType}/{category}][%d] putLolSettingsV2AccountByPpTypeByCategoryOK  %+v", 200, o.Payload)
}

func (o *PutLolSettingsV2AccountByPpTypeByCategoryOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PutLolSettingsV2AccountByPpTypeByCategoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
