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

// PostLolChampSelectLegacyV1SessionTradesByIDCancelReader is a Reader for the PostLolChampSelectLegacyV1SessionTradesByIDCancel structure.
type PostLolChampSelectLegacyV1SessionTradesByIDCancelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolChampSelectLegacyV1SessionTradesByIDCancelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolChampSelectLegacyV1SessionTradesByIDCancelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolChampSelectLegacyV1SessionTradesByIDCancelOK creates a PostLolChampSelectLegacyV1SessionTradesByIDCancelOK with default headers values
func NewPostLolChampSelectLegacyV1SessionTradesByIDCancelOK() *PostLolChampSelectLegacyV1SessionTradesByIDCancelOK {
	return &PostLolChampSelectLegacyV1SessionTradesByIDCancelOK{}
}

/*PostLolChampSelectLegacyV1SessionTradesByIDCancelOK handles this case with default header values.

Successful response
*/
type PostLolChampSelectLegacyV1SessionTradesByIDCancelOK struct {
	Payload interface{}
}

func (o *PostLolChampSelectLegacyV1SessionTradesByIDCancelOK) Error() string {
	return fmt.Sprintf("[POST /lol-champ-select-legacy/v1/session/trades/{id}/cancel][%d] postLolChampSelectLegacyV1SessionTradesByIdCancelOK  %+v", 200, o.Payload)
}

func (o *PostLolChampSelectLegacyV1SessionTradesByIDCancelOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolChampSelectLegacyV1SessionTradesByIDCancelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
