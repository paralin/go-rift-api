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

// GetLolPurchaseWidgetV1PurchasableItemReader is a Reader for the GetLolPurchaseWidgetV1PurchasableItem structure.
type GetLolPurchaseWidgetV1PurchasableItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolPurchaseWidgetV1PurchasableItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolPurchaseWidgetV1PurchasableItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolPurchaseWidgetV1PurchasableItemOK creates a GetLolPurchaseWidgetV1PurchasableItemOK with default headers values
func NewGetLolPurchaseWidgetV1PurchasableItemOK() *GetLolPurchaseWidgetV1PurchasableItemOK {
	return &GetLolPurchaseWidgetV1PurchasableItemOK{}
}

/*GetLolPurchaseWidgetV1PurchasableItemOK handles this case with default header values.

Successful response
*/
type GetLolPurchaseWidgetV1PurchasableItemOK struct {
	Payload *models.LolPurchaseWidgetPurchasableItem
}

func (o *GetLolPurchaseWidgetV1PurchasableItemOK) Error() string {
	return fmt.Sprintf("[GET /lol-purchase-widget/v1/purchasable-item][%d] getLolPurchaseWidgetV1PurchasableItemOK  %+v", 200, o.Payload)
}

func (o *GetLolPurchaseWidgetV1PurchasableItemOK) GetPayload() *models.LolPurchaseWidgetPurchasableItem {
	return o.Payload
}

func (o *GetLolPurchaseWidgetV1PurchasableItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolPurchaseWidgetPurchasableItem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
