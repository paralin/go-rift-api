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

// GetLolTencentQtV1UIStatesReader is a Reader for the GetLolTencentQtV1UIStates structure.
type GetLolTencentQtV1UIStatesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolTencentQtV1UIStatesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolTencentQtV1UIStatesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolTencentQtV1UIStatesOK creates a GetLolTencentQtV1UIStatesOK with default headers values
func NewGetLolTencentQtV1UIStatesOK() *GetLolTencentQtV1UIStatesOK {
	return &GetLolTencentQtV1UIStatesOK{}
}

/*GetLolTencentQtV1UIStatesOK handles this case with default header values.

Successful response
*/
type GetLolTencentQtV1UIStatesOK struct {
	Payload map[string]models.LolTencentQtTencentQTNotification
}

func (o *GetLolTencentQtV1UIStatesOK) Error() string {
	return fmt.Sprintf("[GET /lol-tencent-qt/v1/ui-states][%d] getLolTencentQtV1UiStatesOK  %+v", 200, o.Payload)
}

func (o *GetLolTencentQtV1UIStatesOK) GetPayload() map[string]models.LolTencentQtTencentQTNotification {
	return o.Payload
}

func (o *GetLolTencentQtV1UIStatesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
