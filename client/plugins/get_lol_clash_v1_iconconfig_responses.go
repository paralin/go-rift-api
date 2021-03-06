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

// GetLolClashV1IconconfigReader is a Reader for the GetLolClashV1Iconconfig structure.
type GetLolClashV1IconconfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolClashV1IconconfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolClashV1IconconfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolClashV1IconconfigOK creates a GetLolClashV1IconconfigOK with default headers values
func NewGetLolClashV1IconconfigOK() *GetLolClashV1IconconfigOK {
	return &GetLolClashV1IconconfigOK{}
}

/*GetLolClashV1IconconfigOK handles this case with default header values.

Successful response
*/
type GetLolClashV1IconconfigOK struct {
	Payload interface{}
}

func (o *GetLolClashV1IconconfigOK) Error() string {
	return fmt.Sprintf("[GET /lol-clash/v1/iconconfig][%d] getLolClashV1IconconfigOK  %+v", 200, o.Payload)
}

func (o *GetLolClashV1IconconfigOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolClashV1IconconfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
