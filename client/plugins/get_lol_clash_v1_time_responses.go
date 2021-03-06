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

// GetLolClashV1TimeReader is a Reader for the GetLolClashV1Time structure.
type GetLolClashV1TimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1TimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1TimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1TimeOK creates a GetLolClashV1TimeOK with default headers values
func NewGetLolClashV1TimeOK() *GetLolClashV1TimeOK {
	return &GetLolClashV1TimeOK{}
}

/*GetLolClashV1TimeOK handles this case with default header values.

Successful response
*/
type GetLolClashV1TimeOK struct {
	Payload int64
}

func (o *GetLolClashV1TimeOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/time][%d] getLolClashV1TimeOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1TimeOK) GetPayload() int64 {
	return o.Payload
}

func (o *GetLolClashV1TimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
