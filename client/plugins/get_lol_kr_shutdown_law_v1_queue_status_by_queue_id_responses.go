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

// GetLolKrShutdownLawV1QueueStatusByQueueIDReader is a Reader for the GetLolKrShutdownLawV1QueueStatusByQueueID structure.
type GetLolKrShutdownLawV1QueueStatusByQueueIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolKrShutdownLawV1QueueStatusByQueueIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolKrShutdownLawV1QueueStatusByQueueIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolKrShutdownLawV1QueueStatusByQueueIDOK creates a GetLolKrShutdownLawV1QueueStatusByQueueIDOK with default headers values
func NewGetLolKrShutdownLawV1QueueStatusByQueueIDOK() *GetLolKrShutdownLawV1QueueStatusByQueueIDOK {
	return &GetLolKrShutdownLawV1QueueStatusByQueueIDOK{}
}

/*GetLolKrShutdownLawV1QueueStatusByQueueIDOK handles this case with default header values.

Successful response
*/
type GetLolKrShutdownLawV1QueueStatusByQueueIDOK struct {
	Payload *models.LolKrShutdownLawQueueShutdownStatus
}

func (o *GetLolKrShutdownLawV1QueueStatusByQueueIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-kr-shutdown-law/v1/queue-status/{queue_id}][%d] getLolKrShutdownLawV1QueueStatusByQueueIdOK  %+v", 200, o.Payload)
}

func (o *GetLolKrShutdownLawV1QueueStatusByQueueIDOK) GetPayload() *models.LolKrShutdownLawQueueShutdownStatus {
	return o.Payload
}

func (o *GetLolKrShutdownLawV1QueueStatusByQueueIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolKrShutdownLawQueueShutdownStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}