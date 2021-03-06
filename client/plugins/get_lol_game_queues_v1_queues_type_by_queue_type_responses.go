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

// GetLolGameQueuesV1QueuesTypeByQueueTypeReader is a Reader for the GetLolGameQueuesV1QueuesTypeByQueueType structure.
type GetLolGameQueuesV1QueuesTypeByQueueTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolGameQueuesV1QueuesTypeByQueueTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolGameQueuesV1QueuesTypeByQueueTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolGameQueuesV1QueuesTypeByQueueTypeOK creates a GetLolGameQueuesV1QueuesTypeByQueueTypeOK with default headers values
func NewGetLolGameQueuesV1QueuesTypeByQueueTypeOK() *GetLolGameQueuesV1QueuesTypeByQueueTypeOK {
	return &GetLolGameQueuesV1QueuesTypeByQueueTypeOK{}
}

/*GetLolGameQueuesV1QueuesTypeByQueueTypeOK handles this case with default header values.

Successful response
*/
type GetLolGameQueuesV1QueuesTypeByQueueTypeOK struct {
	Payload *models.LolGameQueuesQueue
}

func (o *GetLolGameQueuesV1QueuesTypeByQueueTypeOK) Error() string {
	return fmt.Sprintf("[GET /lol-game-queues/v1/queues/type/{queueType}][%d] getLolGameQueuesV1QueuesTypeByQueueTypeOK  %+v", 200, o.Payload)
}

func (o *GetLolGameQueuesV1QueuesTypeByQueueTypeOK) GetPayload() *models.LolGameQueuesQueue {
	return o.Payload
}

func (o *GetLolGameQueuesV1QueuesTypeByQueueTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolGameQueuesQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
