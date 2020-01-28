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

// GetLolChampSelectLegacyV1ImplementationActiveReader is a Reader for the GetLolChampSelectLegacyV1ImplementationActive structure.
type GetLolChampSelectLegacyV1ImplementationActiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChampSelectLegacyV1ImplementationActiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChampSelectLegacyV1ImplementationActiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChampSelectLegacyV1ImplementationActiveOK creates a GetLolChampSelectLegacyV1ImplementationActiveOK with default headers values
func NewGetLolChampSelectLegacyV1ImplementationActiveOK() *GetLolChampSelectLegacyV1ImplementationActiveOK {
	return &GetLolChampSelectLegacyV1ImplementationActiveOK{}
}

/*GetLolChampSelectLegacyV1ImplementationActiveOK handles this case with default header values.

Successful response
*/
type GetLolChampSelectLegacyV1ImplementationActiveOK struct {
	Payload bool
}

func (o *GetLolChampSelectLegacyV1ImplementationActiveOK) Error() string {
	return fmt.Sprintf("[GET /lol-champ-select-legacy/v1/implementation-active][%d] getLolChampSelectLegacyV1ImplementationActiveOK  %+v", 200, o.Payload)
}

func (o *GetLolChampSelectLegacyV1ImplementationActiveOK) GetPayload() bool {
	return o.Payload
}

func (o *GetLolChampSelectLegacyV1ImplementationActiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
