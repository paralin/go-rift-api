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

// GetLolGameQueuesV1QueuesByIDReader is a Reader for the GetLolGameQueuesV1QueuesByID structure.
type GetLolGameQueuesV1QueuesByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolGameQueuesV1QueuesByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolGameQueuesV1QueuesByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolGameQueuesV1QueuesByIDOK creates a GetLolGameQueuesV1QueuesByIDOK with default headers values
func NewGetLolGameQueuesV1QueuesByIDOK() *GetLolGameQueuesV1QueuesByIDOK {
	return &GetLolGameQueuesV1QueuesByIDOK{}
}

/*GetLolGameQueuesV1QueuesByIDOK handles this case with default header values.

Successful response
*/
type GetLolGameQueuesV1QueuesByIDOK struct {
	Payload *models.LolGameQueuesQueue
}

func (o *GetLolGameQueuesV1QueuesByIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-game-queues/v1/queues/{id}][%d] getLolGameQueuesV1QueuesByIdOK  %+v", 200, o.Payload)
}

func (o *GetLolGameQueuesV1QueuesByIDOK) GetPayload() *models.LolGameQueuesQueue {
	return o.Payload
}

func (o *GetLolGameQueuesV1QueuesByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolGameQueuesQueue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}