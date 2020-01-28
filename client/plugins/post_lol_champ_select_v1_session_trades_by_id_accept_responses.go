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

// PostLolChampSelectV1SessionTradesByIDAcceptReader is a Reader for the PostLolChampSelectV1SessionTradesByIDAccept structure.
type PostLolChampSelectV1SessionTradesByIDAcceptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolChampSelectV1SessionTradesByIDAcceptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolChampSelectV1SessionTradesByIDAcceptOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolChampSelectV1SessionTradesByIDAcceptOK creates a PostLolChampSelectV1SessionTradesByIDAcceptOK with default headers values
func NewPostLolChampSelectV1SessionTradesByIDAcceptOK() *PostLolChampSelectV1SessionTradesByIDAcceptOK {
	return &PostLolChampSelectV1SessionTradesByIDAcceptOK{}
}

/*PostLolChampSelectV1SessionTradesByIDAcceptOK handles this case with default header values.

Successful response
*/
type PostLolChampSelectV1SessionTradesByIDAcceptOK struct {
	Payload interface{}
}

func (o *PostLolChampSelectV1SessionTradesByIDAcceptOK) Error() string {
	return fmt.Sprintf("[POST /lol-champ-select/v1/session/trades/{id}/accept][%d] postLolChampSelectV1SessionTradesByIdAcceptOK  %+v", 200, o.Payload)
}

func (o *PostLolChampSelectV1SessionTradesByIDAcceptOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolChampSelectV1SessionTradesByIDAcceptOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}