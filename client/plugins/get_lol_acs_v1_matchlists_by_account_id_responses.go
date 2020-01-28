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

// GetLolAcsV1MatchlistsByAccountIDReader is a Reader for the GetLolAcsV1MatchlistsByAccountID structure.
type GetLolAcsV1MatchlistsByAccountIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolAcsV1MatchlistsByAccountIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolAcsV1MatchlistsByAccountIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolAcsV1MatchlistsByAccountIDOK creates a GetLolAcsV1MatchlistsByAccountIDOK with default headers values
func NewGetLolAcsV1MatchlistsByAccountIDOK() *GetLolAcsV1MatchlistsByAccountIDOK {
	return &GetLolAcsV1MatchlistsByAccountIDOK{}
}

/*GetLolAcsV1MatchlistsByAccountIDOK handles this case with default header values.

Successful response
*/
type GetLolAcsV1MatchlistsByAccountIDOK struct {
	Payload interface{}
}

func (o *GetLolAcsV1MatchlistsByAccountIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-acs/v1/matchlists/{accountId}][%d] getLolAcsV1MatchlistsByAccountIdOK  %+v", 200, o.Payload)
}

func (o *GetLolAcsV1MatchlistsByAccountIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetLolAcsV1MatchlistsByAccountIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
