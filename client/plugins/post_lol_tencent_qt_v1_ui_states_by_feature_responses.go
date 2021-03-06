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

// PostLolTencentQtV1UIStatesByFeatureReader is a Reader for the PostLolTencentQtV1UIStatesByFeature structure.
type PostLolTencentQtV1UIStatesByFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLolTencentQtV1UIStatesByFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLolTencentQtV1UIStatesByFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostLolTencentQtV1UIStatesByFeatureOK creates a PostLolTencentQtV1UIStatesByFeatureOK with default headers values
func NewPostLolTencentQtV1UIStatesByFeatureOK() *PostLolTencentQtV1UIStatesByFeatureOK {
	return &PostLolTencentQtV1UIStatesByFeatureOK{}
}

/*PostLolTencentQtV1UIStatesByFeatureOK handles this case with default header values.

Successful response
*/
type PostLolTencentQtV1UIStatesByFeatureOK struct {
	Payload interface{}
}

func (o *PostLolTencentQtV1UIStatesByFeatureOK) Error() string {
	return fmt.Sprintf("[POST /lol-tencent-qt/v1/ui-states/{feature}][%d] postLolTencentQtV1UiStatesByFeatureOK  %+v", 200, o.Payload)
}

func (o *PostLolTencentQtV1UIStatesByFeatureOK) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLolTencentQtV1UIStatesByFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
