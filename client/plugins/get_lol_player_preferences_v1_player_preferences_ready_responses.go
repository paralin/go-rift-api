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

// GetLolPlayerPreferencesV1PlayerPreferencesReadyReader is a Reader for the GetLolPlayerPreferencesV1PlayerPreferencesReady structure.
type GetLolPlayerPreferencesV1PlayerPreferencesReadyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPlayerPreferencesV1PlayerPreferencesReadyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPlayerPreferencesV1PlayerPreferencesReadyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPlayerPreferencesV1PlayerPreferencesReadyOK creates a GetLolPlayerPreferencesV1PlayerPreferencesReadyOK with default headers values
func NewGetLolPlayerPreferencesV1PlayerPreferencesReadyOK() *GetLolPlayerPreferencesV1PlayerPreferencesReadyOK {
	return &GetLolPlayerPreferencesV1PlayerPreferencesReadyOK{}
}

/*GetLolPlayerPreferencesV1PlayerPreferencesReadyOK handles this case with default header values.

Successful response
*/
type GetLolPlayerPreferencesV1PlayerPreferencesReadyOK struct {
	Payload bool
}

func (o *GetLolPlayerPreferencesV1PlayerPreferencesReadyOK) Error() string {
	return fmt.Sprintf("[GET /lol-player-preferences/v1/player-preferences-ready][%d] getLolPlayerPreferencesV1PlayerPreferencesReadyOK  %+v", 200, o.Payload)
}

func (o *GetLolPlayerPreferencesV1PlayerPreferencesReadyOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolPlayerPreferencesV1PlayerPreferencesReadyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
