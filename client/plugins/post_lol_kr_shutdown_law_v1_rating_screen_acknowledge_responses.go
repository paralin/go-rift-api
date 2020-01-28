// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolKrShutdownLawV1RatingScreenAcknowledgeReader is a Reader for the PostLolKrShutdownLawV1RatingScreenAcknowledge structure.
type PostLolKrShutdownLawV1RatingScreenAcknowledgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolKrShutdownLawV1RatingScreenAcknowledgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent creates a PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent with default headers values
func NewPostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent() *PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent {
	return &PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent{}
}

/*PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent handles this case with default header values.

No content
*/
type PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent struct {
}

func (o *PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-kr-shutdown-law/v1/rating-screen/acknowledge][%d] postLolKrShutdownLawV1RatingScreenAcknowledgeNoContent ", 204)
}

func (o *PostLolKrShutdownLawV1RatingScreenAcknowledgeNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}