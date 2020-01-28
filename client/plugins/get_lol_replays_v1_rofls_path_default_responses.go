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

// GetLolReplaysV1RoflsPathDefaultReader is a Reader for the GetLolReplaysV1RoflsPathDefault structure.
type GetLolReplaysV1RoflsPathDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolReplaysV1RoflsPathDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolReplaysV1RoflsPathDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolReplaysV1RoflsPathDefaultOK creates a GetLolReplaysV1RoflsPathDefaultOK with default headers values
func NewGetLolReplaysV1RoflsPathDefaultOK() *GetLolReplaysV1RoflsPathDefaultOK {
	return &GetLolReplaysV1RoflsPathDefaultOK{}
}

/*GetLolReplaysV1RoflsPathDefaultOK handles this case with default header values.

Successful response
*/
type GetLolReplaysV1RoflsPathDefaultOK struct {
	Payload string
}

func (o *GetLolReplaysV1RoflsPathDefaultOK) Error() string {
	return fmt.Sprintf("[GET /lol-replays/v1/rofls/path/default][%d] getLolReplaysV1RoflsPathDefaultOK  %+v", 200, o.Payload)
}

func (o *GetLolReplaysV1RoflsPathDefaultOK) GetPayload() string {
	return o.Payload
}

func (o *GetLolReplaysV1RoflsPathDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}