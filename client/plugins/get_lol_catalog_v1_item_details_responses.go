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

// GetLolCatalogV1ItemDetailsReader is a Reader for the GetLolCatalogV1ItemDetails structure.
type GetLolCatalogV1ItemDetailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCatalogV1ItemDetailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCatalogV1ItemDetailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCatalogV1ItemDetailsOK creates a GetLolCatalogV1ItemDetailsOK with default headers values
func NewGetLolCatalogV1ItemDetailsOK() *GetLolCatalogV1ItemDetailsOK {
	return &GetLolCatalogV1ItemDetailsOK{}
}

/*GetLolCatalogV1ItemDetailsOK handles this case with default header values.

Successful response
*/
type GetLolCatalogV1ItemDetailsOK struct {
	Payload *models.LolCatalogCatalogPluginItemWithDetails
}

func (o *GetLolCatalogV1ItemDetailsOK) Error() string {
	return fmt.Sprintf("[GET /lol-catalog/v1/item-details][%d] getLolCatalogV1ItemDetailsOK  %+v", 200, o.Payload)
}

func (o *GetLolCatalogV1ItemDetailsOK) GetPayload() *models.LolCatalogCatalogPluginItemWithDetails {
	return o.Payload
}

func (o *GetLolCatalogV1ItemDetailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolCatalogCatalogPluginItemWithDetails)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
