// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolTftV1TftSeriesOptInReader is a Reader for the PostLolTftV1TftSeriesOptIn structure.
type PostLolTftV1TftSeriesOptInReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolTftV1TftSeriesOptInReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolTftV1TftSeriesOptInNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolTftV1TftSeriesOptInNoContent creates a PostLolTftV1TftSeriesOptInNoContent with default headers values
func NewPostLolTftV1TftSeriesOptInNoContent() *PostLolTftV1TftSeriesOptInNoContent {
	return &PostLolTftV1TftSeriesOptInNoContent{}
}

/*PostLolTftV1TftSeriesOptInNoContent handles this case with default header values.

No content
*/
type PostLolTftV1TftSeriesOptInNoContent struct {
}

func (o *PostLolTftV1TftSeriesOptInNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-tft/v1/tft/series/opt-in][%d] postLolTftV1TftSeriesOptInNoContent ", 204)
}

func (o *PostLolTftV1TftSeriesOptInNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
