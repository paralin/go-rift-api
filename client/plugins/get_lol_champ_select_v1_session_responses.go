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

// GetLolChampSelectV1SessionReader is a Reader for the GetLolChampSelectV1Session structure.
type GetLolChampSelectV1SessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolChampSelectV1SessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolChampSelectV1SessionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolChampSelectV1SessionOK creates a GetLolChampSelectV1SessionOK with default headers values
func NewGetLolChampSelectV1SessionOK() *GetLolChampSelectV1SessionOK {
	return &GetLolChampSelectV1SessionOK{}
}

/*GetLolChampSelectV1SessionOK handles this case with default header values.

Successful response
*/
type GetLolChampSelectV1SessionOK struct {
	Payload *models.LolChampSelectChampSelectSession
}

func (o *GetLolChampSelectV1SessionOK) Error() string {
	return fmt.Sprintf("[GET /lol-champ-select/v1/session][%d] getLolChampSelectV1SessionOK  %+v", 200, o.Payload)
}

func (o *GetLolChampSelectV1SessionOK) GetPayload() *models.LolChampSelectChampSelectSession {
	return o.Payload
}

func (o *GetLolChampSelectV1SessionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolChampSelectChampSelectSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
