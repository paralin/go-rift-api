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

// GetRiotMessagingServiceV1MessageByAByBByCReader is a Reader for the GetRiotMessagingServiceV1MessageByAByBByC structure.
type GetRiotMessagingServiceV1MessageByAByBByCReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRiotMessagingServiceV1MessageByAByBByCReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRiotMessagingServiceV1MessageByAByBByCOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRiotMessagingServiceV1MessageByAByBByCOK creates a GetRiotMessagingServiceV1MessageByAByBByCOK with default headers values
func NewGetRiotMessagingServiceV1MessageByAByBByCOK() *GetRiotMessagingServiceV1MessageByAByBByCOK {
	return &GetRiotMessagingServiceV1MessageByAByBByCOK{}
}

/*GetRiotMessagingServiceV1MessageByAByBByCOK handles this case with default header values.

Successful response
*/
type GetRiotMessagingServiceV1MessageByAByBByCOK struct {
	Payload *models.RmsMessage
}

func (o *GetRiotMessagingServiceV1MessageByAByBByCOK) Error() string {
	return fmt.Sprintf("[GET /riot-messaging-service/v1/message/{a}/{b}/{c}][%d] getRiotMessagingServiceV1MessageByAByBByCOK  %+v", 200, o.Payload)
}

func (o *GetRiotMessagingServiceV1MessageByAByBByCOK) GetPayload() *models.RmsMessage {
	return o.Payload
}

func (o *GetRiotMessagingServiceV1MessageByAByBByCOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RmsMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
