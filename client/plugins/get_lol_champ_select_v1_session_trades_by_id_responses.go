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

// GetLolChampSelectV1SessionTradesByIDReader is a Reader for the GetLolChampSelectV1SessionTradesByID structure.
type GetLolChampSelectV1SessionTradesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChampSelectV1SessionTradesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChampSelectV1SessionTradesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChampSelectV1SessionTradesByIDOK creates a GetLolChampSelectV1SessionTradesByIDOK with default headers values
func NewGetLolChampSelectV1SessionTradesByIDOK() *GetLolChampSelectV1SessionTradesByIDOK {
	return &GetLolChampSelectV1SessionTradesByIDOK{}
}

/*GetLolChampSelectV1SessionTradesByIDOK handles this case with default header values.

Successful response
*/
type GetLolChampSelectV1SessionTradesByIDOK struct {
	Payload *models.LolChampSelectChampSelectTradeContract
}

func (o *GetLolChampSelectV1SessionTradesByIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-champ-select/v1/session/trades/{id}][%d] getLolChampSelectV1SessionTradesByIdOK  %+v", 200, o.Payload)
}

func (o *GetLolChampSelectV1SessionTradesByIDOK) GetPayload() *models.LolChampSelectChampSelectTradeContract {
	return o.Payload
}

func (o *GetLolChampSelectV1SessionTradesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChampSelectChampSelectTradeContract)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}