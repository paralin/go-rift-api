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

// GetRiotMessagingServiceV1MessageByAReader is a Reader for the GetRiotMessagingServiceV1MessageByA structure.
type GetRiotMessagingServiceV1MessageByAReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRiotMessagingServiceV1MessageByAReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRiotMessagingServiceV1MessageByAOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRiotMessagingServiceV1MessageByAOK creates a GetRiotMessagingServiceV1MessageByAOK with default headers values
func NewGetRiotMessagingServiceV1MessageByAOK() *GetRiotMessagingServiceV1MessageByAOK {
	return &GetRiotMessagingServiceV1MessageByAOK{}
}

/*GetRiotMessagingServiceV1MessageByAOK handles this case with default header values.

Successful response
*/
type GetRiotMessagingServiceV1MessageByAOK struct {
	Payload *models.RmsMessage
}

func (o *GetRiotMessagingServiceV1MessageByAOK) Error() string {
	return fmt.Sprintf("[GET /riot-messaging-service/v1/message/{a}][%d] getRiotMessagingServiceV1MessageByAOK  %+v", 200, o.Payload)
}

func (o *GetRiotMessagingServiceV1MessageByAOK) GetPayload() *models.RmsMessage {
	return o.Payload
}

func (o *GetRiotMessagingServiceV1MessageByAOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RmsMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}