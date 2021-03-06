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

// GetRecofrienderV2ContactsPageReader is a Reader for the GetRecofrienderV2ContactsPage structure.
type GetRecofrienderV2ContactsPageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecofrienderV2ContactsPageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecofrienderV2ContactsPageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecofrienderV2ContactsPageOK creates a GetRecofrienderV2ContactsPageOK with default headers values
func NewGetRecofrienderV2ContactsPageOK() *GetRecofrienderV2ContactsPageOK {
	return &GetRecofrienderV2ContactsPageOK{}
}

/*GetRecofrienderV2ContactsPageOK handles this case with default header values.

Successful response
*/
type GetRecofrienderV2ContactsPageOK struct {
	Payload *models.RecofrienderContactPaginationResource
}

func (o *GetRecofrienderV2ContactsPageOK) Error() string {
	return fmt.Sprintf("[GET /recofriender/v2/contacts/page][%d] getRecofrienderV2ContactsPageOK  %+v", 200, o.Payload)
}

func (o *GetRecofrienderV2ContactsPageOK) GetPayload() *models.RecofrienderContactPaginationResource {
	return o.Payload
}

func (o *GetRecofrienderV2ContactsPageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecofrienderContactPaginationResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
