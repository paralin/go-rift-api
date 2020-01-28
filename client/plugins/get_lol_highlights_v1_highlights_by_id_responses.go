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

// GetLolHighlightsV1HighlightsByIDReader is a Reader for the GetLolHighlightsV1HighlightsByID structure.
type GetLolHighlightsV1HighlightsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLolHighlightsV1HighlightsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLolHighlightsV1HighlightsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLolHighlightsV1HighlightsByIDOK creates a GetLolHighlightsV1HighlightsByIDOK with default headers values
func NewGetLolHighlightsV1HighlightsByIDOK() *GetLolHighlightsV1HighlightsByIDOK {
	return &GetLolHighlightsV1HighlightsByIDOK{}
}

/*GetLolHighlightsV1HighlightsByIDOK handles this case with default header values.

Successful response
*/
type GetLolHighlightsV1HighlightsByIDOK struct {
	Payload *models.LolHighlightsHighlight
}

func (o *GetLolHighlightsV1HighlightsByIDOK) Error() string {
	return fmt.Sprintf("[GET /lol-highlights/v1/highlights/{id}][%d] getLolHighlightsV1HighlightsByIdOK  %+v", 200, o.Payload)
}

func (o *GetLolHighlightsV1HighlightsByIDOK) GetPayload() *models.LolHighlightsHighlight {
	return o.Payload
}

func (o *GetLolHighlightsV1HighlightsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LolHighlightsHighlight)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
