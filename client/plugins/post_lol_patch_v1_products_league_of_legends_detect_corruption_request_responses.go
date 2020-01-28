// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestReader is a Reader for the PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequest structure.
type PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent creates a PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent with default headers values
func NewPostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent() *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent {
	return &PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent{}
}

/*PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent handles this case with default header values.

No content
*/
type PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent struct {
}

func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent) Error() string {
	return fmt.Sprintf("[POST /lol-patch/v1/products/league_of_legends/detect-corruption-request][%d] postLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent ", 204)
}

func (o *PostLolPatchV1ProductsLeagueOfLegendsDetectCorruptionRequestNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
