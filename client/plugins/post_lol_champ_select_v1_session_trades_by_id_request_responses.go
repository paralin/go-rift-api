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

// PostLolChampSelectV1SessionTradesByIDRequestReader is a Reader for the PostLolChampSelectV1SessionTradesByIDRequest structure.
type PostLolChampSelectV1SessionTradesByIDRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolChampSelectV1SessionTradesByIDRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolChampSelectV1SessionTradesByIDRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolChampSelectV1SessionTradesByIDRequestOK creates a PostLolChampSelectV1SessionTradesByIDRequestOK with default headers values
func NewPostLolChampSelectV1SessionTradesByIDRequestOK() *PostLolChampSelectV1SessionTradesByIDRequestOK {
	return &PostLolChampSelectV1SessionTradesByIDRequestOK{}
}

/*PostLolChampSelectV1SessionTradesByIDRequestOK handles this case with default header values.

Successful response
*/
type PostLolChampSelectV1SessionTradesByIDRequestOK struct {
	Payload *models.LolChampSelectChampSelectTradeContract
}

func (o *PostLolChampSelectV1SessionTradesByIDRequestOK) Error() string {
	return fmt.Sprintf("[POST /lol-champ-select/v1/session/trades/{id}/request][%d] postLolChampSelectV1SessionTradesByIdRequestOK  %+v", 200, o.Payload)
}

func (o *PostLolChampSelectV1SessionTradesByIDRequestOK) GetPayload() *models.LolChampSelectChampSelectTradeContract {
	return o.Payload
}

func (o *PostLolChampSelectV1SessionTradesByIDRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChampSelectChampSelectTradeContract)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}