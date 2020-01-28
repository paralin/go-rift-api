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

// GetLolHonorV2V1VoteCompletionReader is a Reader for the GetLolHonorV2V1VoteCompletion structure.
type GetLolHonorV2V1VoteCompletionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolHonorV2V1VoteCompletionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolHonorV2V1VoteCompletionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolHonorV2V1VoteCompletionOK creates a GetLolHonorV2V1VoteCompletionOK with default headers values
func NewGetLolHonorV2V1VoteCompletionOK() *GetLolHonorV2V1VoteCompletionOK {
	return &GetLolHonorV2V1VoteCompletionOK{}
}

/*GetLolHonorV2V1VoteCompletionOK handles this case with default header values.

Successful response
*/
type GetLolHonorV2V1VoteCompletionOK struct {
	Payload *models.LolHonorV2VoteCompletion
}

func (o *GetLolHonorV2V1VoteCompletionOK) Error() string {
	return fmt.Sprintf("[GET /lol-honor-v2/v1/vote-completion][%d] getLolHonorV2V1VoteCompletionOK  %+v", 200, o.Payload)
}

func (o *GetLolHonorV2V1VoteCompletionOK) GetPayload() *models.LolHonorV2VoteCompletion {
	return o.Payload
}

func (o *GetLolHonorV2V1VoteCompletionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolHonorV2VoteCompletion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}