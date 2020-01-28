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

// GetLolHighlightsV1HighlightsFolderPathDefaultReader is a Reader for the GetLolHighlightsV1HighlightsFolderPathDefault structure.
type GetLolHighlightsV1HighlightsFolderPathDefaultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolHighlightsV1HighlightsFolderPathDefaultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolHighlightsV1HighlightsFolderPathDefaultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolHighlightsV1HighlightsFolderPathDefaultOK creates a GetLolHighlightsV1HighlightsFolderPathDefaultOK with default headers values
func NewGetLolHighlightsV1HighlightsFolderPathDefaultOK() *GetLolHighlightsV1HighlightsFolderPathDefaultOK {
	return &GetLolHighlightsV1HighlightsFolderPathDefaultOK{}
}

/*GetLolHighlightsV1HighlightsFolderPathDefaultOK handles this case with default header values.

Successful response
*/
type GetLolHighlightsV1HighlightsFolderPathDefaultOK struct {
	Payload string
}

func (o *GetLolHighlightsV1HighlightsFolderPathDefaultOK) Error() string {
	return fmt.Sprintf("[GET /lol-highlights/v1/highlights-folder-path/default][%d] getLolHighlightsV1HighlightsFolderPathDefaultOK  %+v", 200, o.Payload)
}

func (o *GetLolHighlightsV1HighlightsFolderPathDefaultOK) GetPayload() string {
	return o.Payload
}

func (o *GetLolHighlightsV1HighlightsFolderPathDefaultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}