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

// GetLolStoreV1ByPageTypeReader is a Reader for the GetLolStoreV1ByPageType structure.
type GetLolStoreV1ByPageTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolStoreV1ByPageTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolStoreV1ByPageTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolStoreV1ByPageTypeOK creates a GetLolStoreV1ByPageTypeOK with default headers values
func NewGetLolStoreV1ByPageTypeOK() *GetLolStoreV1ByPageTypeOK {
	return &GetLolStoreV1ByPageTypeOK{}
}

/*GetLolStoreV1ByPageTypeOK handles this case with default header values.

Successful response
*/
type GetLolStoreV1ByPageTypeOK struct {
	Payload interface{}
}

func (o *GetLolStoreV1ByPageTypeOK) Error() string {
	return fmt.Sprintf("[GET /lol-store/v1/{pageType}][%d] getLolStoreV1ByPageTypeOK  %+v", 200, o.Payload)
}

func (o *GetLolStoreV1ByPageTypeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolStoreV1ByPageTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}