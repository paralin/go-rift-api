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

// GetRecofrienderV1FaqURLReader is a Reader for the GetRecofrienderV1FaqURL structure.
type GetRecofrienderV1FaqURLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecofrienderV1FaqURLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecofrienderV1FaqURLOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecofrienderV1FaqURLOK creates a GetRecofrienderV1FaqURLOK with default headers values
func NewGetRecofrienderV1FaqURLOK() *GetRecofrienderV1FaqURLOK {
	return &GetRecofrienderV1FaqURLOK{}
}

/*GetRecofrienderV1FaqURLOK handles this case with default header values.

Successful response
*/
type GetRecofrienderV1FaqURLOK struct {
	Payload *models.RecofrienderURLResource
}

func (o *GetRecofrienderV1FaqURLOK) Error() string {
	return fmt.Sprintf("[GET /recofriender/v1/faq-url][%d] getRecofrienderV1FaqUrlOK  %+v", 200, o.Payload)
}

func (o *GetRecofrienderV1FaqURLOK) GetPayload() *models.RecofrienderURLResource {
	return o.Payload
}

func (o *GetRecofrienderV1FaqURLOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecofrienderURLResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
