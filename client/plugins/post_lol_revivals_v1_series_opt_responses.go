// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolRevivalsV1SeriesOptReader is a Reader for the PostLolRevivalsV1SeriesOpt structure.
type PostLolRevivalsV1SeriesOptReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolRevivalsV1SeriesOptReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolRevivalsV1SeriesOptNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolRevivalsV1SeriesOptNoContent creates a PostLolRevivalsV1SeriesOptNoContent with default headers values
func NewPostLolRevivalsV1SeriesOptNoContent() *PostLolRevivalsV1SeriesOptNoContent {
	return &PostLolRevivalsV1SeriesOptNoContent{}
}

/*PostLolRevivalsV1SeriesOptNoContent handles this case with default header values.

No content
*/
type PostLolRevivalsV1SeriesOptNoContent struct {
}

func (o *PostLolRevivalsV1SeriesOptNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-revivals/v1/series/opt][%d] postLolRevivalsV1SeriesOptNoContent ", 204)
}

func (o *PostLolRevivalsV1SeriesOptNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
