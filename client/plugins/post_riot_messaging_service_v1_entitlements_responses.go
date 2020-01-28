// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PostRiotMessagingServiceV1EntitlementsReader is a Reader for the PostRiotMessagingServiceV1Entitlements structure.
type PostRiotMessagingServiceV1EntitlementsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRiotMessagingServiceV1EntitlementsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostRiotMessagingServiceV1EntitlementsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRiotMessagingServiceV1EntitlementsNoContent creates a PostRiotMessagingServiceV1EntitlementsNoContent with default headers values
func NewPostRiotMessagingServiceV1EntitlementsNoContent() *PostRiotMessagingServiceV1EntitlementsNoContent {
	return &PostRiotMessagingServiceV1EntitlementsNoContent{}
}

/*PostRiotMessagingServiceV1EntitlementsNoContent handles this case with default header values.

No content
*/
type PostRiotMessagingServiceV1EntitlementsNoContent struct {
}

func (o *PostRiotMessagingServiceV1EntitlementsNoContent) Error() string {
	return fmt.Sprintf("[POST /riot-messaging-service/v1/entitlements][%d] postRiotMessagingServiceV1EntitlementsNoContent ", 204)
}

func (o *PostRiotMessagingServiceV1EntitlementsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
