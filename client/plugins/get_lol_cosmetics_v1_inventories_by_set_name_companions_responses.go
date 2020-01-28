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

// GetLolCosmeticsV1InventoriesBySetNameCompanionsReader is a Reader for the GetLolCosmeticsV1InventoriesBySetNameCompanions structure.
type GetLolCosmeticsV1InventoriesBySetNameCompanionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolCosmeticsV1InventoriesBySetNameCompanionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolCosmeticsV1InventoriesBySetNameCompanionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolCosmeticsV1InventoriesBySetNameCompanionsOK creates a GetLolCosmeticsV1InventoriesBySetNameCompanionsOK with default headers values
func NewGetLolCosmeticsV1InventoriesBySetNameCompanionsOK() *GetLolCosmeticsV1InventoriesBySetNameCompanionsOK {
	return &GetLolCosmeticsV1InventoriesBySetNameCompanionsOK{}
}

/*GetLolCosmeticsV1InventoriesBySetNameCompanionsOK handles this case with default header values.

Successful response
*/
type GetLolCosmeticsV1InventoriesBySetNameCompanionsOK struct {
	Payload *models.LolCosmeticsCompanionsGroupedViewModel
}

func (o *GetLolCosmeticsV1InventoriesBySetNameCompanionsOK) Error() string {
	return fmt.Sprintf("[GET /lol-cosmetics/v1/inventories/{setName}/companions][%d] getLolCosmeticsV1InventoriesBySetNameCompanionsOK  %+v", 200, o.Payload)
}

func (o *GetLolCosmeticsV1InventoriesBySetNameCompanionsOK) GetPayload() *models.LolCosmeticsCompanionsGroupedViewModel {
	return o.Payload
}

func (o *GetLolCosmeticsV1InventoriesBySetNameCompanionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolCosmeticsCompanionsGroupedViewModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
