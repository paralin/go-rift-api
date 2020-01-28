// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameReader is a Reader for the DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodName structure.
type DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent creates a DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent with default headers values
func NewDeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent() *DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent {
	return &DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent{}
}

/*DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent handles this case with default header values.

No content
*/
type DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent struct {
}

func (o *DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /lol-login/v1/service-proxy-async-requests/{serviceName}/{methodName}][%d] deleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent ", 204)
}

func (o *DeleteLolLoginV1ServiceProxyAsyncRequestsByServiceNameByMethodNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
