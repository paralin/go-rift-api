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

// GetLolHonorV2V1LateRecognitionReader is a Reader for the GetLolHonorV2V1LateRecognition structure.
type GetLolHonorV2V1LateRecognitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolHonorV2V1LateRecognitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolHonorV2V1LateRecognitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolHonorV2V1LateRecognitionOK creates a GetLolHonorV2V1LateRecognitionOK with default headers values
func NewGetLolHonorV2V1LateRecognitionOK() *GetLolHonorV2V1LateRecognitionOK {
	return &GetLolHonorV2V1LateRecognitionOK{}
}

/*GetLolHonorV2V1LateRecognitionOK handles this case with default header values.

Successful response
*/
type GetLolHonorV2V1LateRecognitionOK struct {
	Payload []*models.LolHonorV2Honor
}

func (o *GetLolHonorV2V1LateRecognitionOK) Error() string {
	return fmt.Sprintf("[GET /lol-honor-v2/v1/late-recognition][%d] getLolHonorV2V1LateRecognitionOK  %+v", 200, o.Payload)
}

func (o *GetLolHonorV2V1LateRecognitionOK) GetPayload() []*models.LolHonorV2Honor {
	return o.Payload
}

func (o *GetLolHonorV2V1LateRecognitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
