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

// PostLolHighlightsV1FileBrowserByHighlightIDReader is a Reader for the PostLolHighlightsV1FileBrowserByHighlightID structure.
type PostLolHighlightsV1FileBrowserByHighlightIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolHighlightsV1FileBrowserByHighlightIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolHighlightsV1FileBrowserByHighlightIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolHighlightsV1FileBrowserByHighlightIDOK creates a PostLolHighlightsV1FileBrowserByHighlightIDOK with default headers values
func NewPostLolHighlightsV1FileBrowserByHighlightIDOK() *PostLolHighlightsV1FileBrowserByHighlightIDOK {
	return &PostLolHighlightsV1FileBrowserByHighlightIDOK{}
}

/*PostLolHighlightsV1FileBrowserByHighlightIDOK handles this case with default header values.

Successful response
*/
type PostLolHighlightsV1FileBrowserByHighlightIDOK struct {
	Payload interface{}
}

func (o *PostLolHighlightsV1FileBrowserByHighlightIDOK) Error() string {
	return fmt.Sprintf("[POST /lol-highlights/v1/file-browser/{highlightId}][%d] postLolHighlightsV1FileBrowserByHighlightIdOK  %+v", 200, o.Payload)
}

func (o *PostLolHighlightsV1FileBrowserByHighlightIDOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolHighlightsV1FileBrowserByHighlightIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
