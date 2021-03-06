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

// GetPluginManagerV2PluginsReader is a Reader for the GetPluginManagerV2Plugins structure.
type GetPluginManagerV2PluginsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginManagerV2PluginsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPluginManagerV2PluginsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginManagerV2PluginsOK creates a GetPluginManagerV2PluginsOK with default headers values
func NewGetPluginManagerV2PluginsOK() *GetPluginManagerV2PluginsOK {
	return &GetPluginManagerV2PluginsOK{}
}

/*GetPluginManagerV2PluginsOK handles this case with default header values.

Successful response
*/
type GetPluginManagerV2PluginsOK struct {
	Payload []*models.PluginResource
}

func (o *GetPluginManagerV2PluginsOK) Error() string {
	return fmt.Sprintf("[GET /plugin-manager/v2/plugins][%d] getPluginManagerV2PluginsOK  %+v", 200, o.Payload)
}

func (o *GetPluginManagerV2PluginsOK) GetPayload() []*models.PluginResource {
	return o.Payload
}

func (o *GetPluginManagerV2PluginsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
