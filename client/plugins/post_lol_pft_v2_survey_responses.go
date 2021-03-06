// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolPftV2SurveyReader is a Reader for the PostLolPftV2Survey structure.
type PostLolPftV2SurveyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPftV2SurveyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolPftV2SurveyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPftV2SurveyNoContent creates a PostLolPftV2SurveyNoContent with default headers values
func NewPostLolPftV2SurveyNoContent() *PostLolPftV2SurveyNoContent {
	return &PostLolPftV2SurveyNoContent{}
}

/*PostLolPftV2SurveyNoContent handles this case with default header values.

No content
*/
type PostLolPftV2SurveyNoContent struct {
}

func (o *PostLolPftV2SurveyNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-pft/v2/survey][%d] postLolPftV2SurveyNoContent ", 204)
}

func (o *PostLolPftV2SurveyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
