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

// GetLolPlayerPreferencesV1PreferenceByTypeReader is a Reader for the GetLolPlayerPreferencesV1PreferenceByType structure.
type GetLolPlayerPreferencesV1PreferenceByTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPlayerPreferencesV1PreferenceByTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPlayerPreferencesV1PreferenceByTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPlayerPreferencesV1PreferenceByTypeOK creates a GetLolPlayerPreferencesV1PreferenceByTypeOK with default headers values
func NewGetLolPlayerPreferencesV1PreferenceByTypeOK() *GetLolPlayerPreferencesV1PreferenceByTypeOK {
	return &GetLolPlayerPreferencesV1PreferenceByTypeOK{}
}

/*GetLolPlayerPreferencesV1PreferenceByTypeOK handles this case with default header values.

Successful response
*/
type GetLolPlayerPreferencesV1PreferenceByTypeOK struct {
	Payload interface{}
}

func (o *GetLolPlayerPreferencesV1PreferenceByTypeOK) Error() string {
	return fmt.Sprintf("[GET /lol-player-preferences/v1/preference/{type}][%d] getLolPlayerPreferencesV1PreferenceByTypeOK  %+v", 200, o.Payload)
}

func (o *GetLolPlayerPreferencesV1PreferenceByTypeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolPlayerPreferencesV1PreferenceByTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
