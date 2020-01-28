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

// GetRecofrienderV2DismissedReader is a Reader for the GetRecofrienderV2Dismissed structure.
type GetRecofrienderV2DismissedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecofrienderV2DismissedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecofrienderV2DismissedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecofrienderV2DismissedOK creates a GetRecofrienderV2DismissedOK with default headers values
func NewGetRecofrienderV2DismissedOK() *GetRecofrienderV2DismissedOK {
	return &GetRecofrienderV2DismissedOK{}
}

/*GetRecofrienderV2DismissedOK handles this case with default header values.

Successful response
*/
type GetRecofrienderV2DismissedOK struct {
	Payload []*models.RecofrienderContactResource
}

func (o *GetRecofrienderV2DismissedOK) Error() string {
	return fmt.Sprintf("[GET /recofriender/v2/dismissed][%d] getRecofrienderV2DismissedOK  %+v", 200, o.Payload)
}

func (o *GetRecofrienderV2DismissedOK) GetPayload() []*models.RecofrienderContactResource {
	return o.Payload
}

func (o *GetRecofrienderV2DismissedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}