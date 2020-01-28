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

// GetLolGameflowV1SpectateReader is a Reader for the GetLolGameflowV1Spectate structure.
type GetLolGameflowV1SpectateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolGameflowV1SpectateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolGameflowV1SpectateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolGameflowV1SpectateOK creates a GetLolGameflowV1SpectateOK with default headers values
func NewGetLolGameflowV1SpectateOK() *GetLolGameflowV1SpectateOK {
	return &GetLolGameflowV1SpectateOK{}
}

/*GetLolGameflowV1SpectateOK handles this case with default header values.

Successful response
*/
type GetLolGameflowV1SpectateOK struct {
	Payload bool
}

func (o *GetLolGameflowV1SpectateOK) Error() string {
	return fmt.Sprintf("[GET /lol-gameflow/v1/spectate][%d] getLolGameflowV1SpectateOK  %+v", 200, o.Payload)
}

func (o *GetLolGameflowV1SpectateOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolGameflowV1SpectateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}