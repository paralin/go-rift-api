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

// GetLolStoreV1WalletReader is a Reader for the GetLolStoreV1Wallet structure.
type GetLolStoreV1WalletReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolStoreV1WalletReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolStoreV1WalletOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolStoreV1WalletOK creates a GetLolStoreV1WalletOK with default headers values
func NewGetLolStoreV1WalletOK() *GetLolStoreV1WalletOK {
	return &GetLolStoreV1WalletOK{}
}

/*GetLolStoreV1WalletOK handles this case with default header values.

Successful response
*/
type GetLolStoreV1WalletOK struct {
	Payload *models.LolStoreWallet
}

func (o *GetLolStoreV1WalletOK) Error() string {
	return fmt.Sprintf("[GET /lol-store/v1/wallet][%d] getLolStoreV1WalletOK  %+v", 200, o.Payload)
}

func (o *GetLolStoreV1WalletOK) GetPayload() *models.LolStoreWallet {
	return o.Payload
}

func (o *GetLolStoreV1WalletOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolStoreWallet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
