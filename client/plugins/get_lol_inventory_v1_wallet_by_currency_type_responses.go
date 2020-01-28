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

// GetLolInventoryV1WalletByCurrencyTypeReader is a Reader for the GetLolInventoryV1WalletByCurrencyType structure.
type GetLolInventoryV1WalletByCurrencyTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolInventoryV1WalletByCurrencyTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolInventoryV1WalletByCurrencyTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolInventoryV1WalletByCurrencyTypeOK creates a GetLolInventoryV1WalletByCurrencyTypeOK with default headers values
func NewGetLolInventoryV1WalletByCurrencyTypeOK() *GetLolInventoryV1WalletByCurrencyTypeOK {
	return &GetLolInventoryV1WalletByCurrencyTypeOK{}
}

/*GetLolInventoryV1WalletByCurrencyTypeOK handles this case with default header values.

Successful response
*/
type GetLolInventoryV1WalletByCurrencyTypeOK struct {
	Payload map[string]int32
}

func (o *GetLolInventoryV1WalletByCurrencyTypeOK) Error() string {
	return fmt.Sprintf("[GET /lol-inventory/v1/wallet/{currencyType}][%d] getLolInventoryV1WalletByCurrencyTypeOK  %+v", 200, o.Payload)
}

func (o *GetLolInventoryV1WalletByCurrencyTypeOK) GetPayload() map[string]int32 {
	return o.Payload
}

func (o *GetLolInventoryV1WalletByCurrencyTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
