// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVoiceChatV2DevicesCaptureParams creates a new GetVoiceChatV2DevicesCaptureParams object
// with the default values initialized.
func NewGetVoiceChatV2DevicesCaptureParams() *GetVoiceChatV2DevicesCaptureParams {

	return &GetVoiceChatV2DevicesCaptureParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoiceChatV2DevicesCaptureParamsWithTimeout creates a new GetVoiceChatV2DevicesCaptureParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoiceChatV2DevicesCaptureParamsWithTimeout(timeout time.Duration) *GetVoiceChatV2DevicesCaptureParams {

	return &GetVoiceChatV2DevicesCaptureParams{

		timeout: timeout,
	}
}

// NewGetVoiceChatV2DevicesCaptureParamsWithContext creates a new GetVoiceChatV2DevicesCaptureParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoiceChatV2DevicesCaptureParamsWithContext(ctx context.Context) *GetVoiceChatV2DevicesCaptureParams {

	return &GetVoiceChatV2DevicesCaptureParams{

		Context: ctx,
	}
}

// NewGetVoiceChatV2DevicesCaptureParamsWithHTTPClient creates a new GetVoiceChatV2DevicesCaptureParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoiceChatV2DevicesCaptureParamsWithHTTPClient(client *http.Client) *GetVoiceChatV2DevicesCaptureParams {

	return &GetVoiceChatV2DevicesCaptureParams{
		HTTPClient: client,
	}
}

/*GetVoiceChatV2DevicesCaptureParams contains all the parameters to send to the API endpoint
for the get voice chat v2 devices capture operation typically these are written to a http.Request
*/
type GetVoiceChatV2DevicesCaptureParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voice chat v2 devices capture params
func (o *GetVoiceChatV2DevicesCaptureParams) WithTimeout(timeout time.Duration) *GetVoiceChatV2DevicesCaptureParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voice chat v2 devices capture params
func (o *GetVoiceChatV2DevicesCaptureParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voice chat v2 devices capture params
func (o *GetVoiceChatV2DevicesCaptureParams) WithContext(ctx context.Context) *GetVoiceChatV2DevicesCaptureParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voice chat v2 devices capture params
func (o *GetVoiceChatV2DevicesCaptureParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voice chat v2 devices capture params
func (o *GetVoiceChatV2DevicesCaptureParams) WithHTTPClient(client *http.Client) *GetVoiceChatV2DevicesCaptureParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voice chat v2 devices capture params
func (o *GetVoiceChatV2DevicesCaptureParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoiceChatV2DevicesCaptureParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
