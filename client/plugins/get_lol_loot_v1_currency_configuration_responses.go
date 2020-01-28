// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetLolLootV1CurrencyConfigurationReader is a Reader for the GetLolLootV1CurrencyConfiguration structure.
type GetLolLootV1CurrencyConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolLootV1CurrencyConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewGetLolLootV1CurrencyConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolLootV1CurrencyConfigurationNoContent creates a GetLolLootV1CurrencyConfigurationNoContent with default headers values
func NewGetLolLootV1CurrencyConfigurationNoContent() *GetLolLootV1CurrencyConfigurationNoContent {
	return &GetLolLootV1CurrencyConfigurationNoContent{}
}

/*GetLolLootV1CurrencyConfigurationNoContent handles this case with default header values.

No content
*/
type GetLolLootV1CurrencyConfigurationNoContent struct {
}

func (o *GetLolLootV1CurrencyConfigurationNoContent) Error() string {
	return fmt.Sprintf("[GET /lol-loot/v1/currency-configuration][%d] getLolLootV1CurrencyConfigurationNoContent ", 204)
}

func (o *GetLolLootV1CurrencyConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
