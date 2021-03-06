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

// NewGetVoiceChatV2DevicesRenderParams creates a new GetVoiceChatV2DevicesRenderParams object
// with the default values initialized.
func NewGetVoiceChatV2DevicesRenderParams() *GetVoiceChatV2DevicesRenderParams {

	return &GetVoiceChatV2DevicesRenderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoiceChatV2DevicesRenderParamsWithTimeout creates a new GetVoiceChatV2DevicesRenderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoiceChatV2DevicesRenderParamsWithTimeout(timeout time.Duration) *GetVoiceChatV2DevicesRenderParams {

	return &GetVoiceChatV2DevicesRenderParams{

		timeout: timeout,
	}
}

// NewGetVoiceChatV2DevicesRenderParamsWithContext creates a new GetVoiceChatV2DevicesRenderParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoiceChatV2DevicesRenderParamsWithContext(ctx context.Context) *GetVoiceChatV2DevicesRenderParams {

	return &GetVoiceChatV2DevicesRenderParams{

		Context: ctx,
	}
}

// NewGetVoiceChatV2DevicesRenderParamsWithHTTPClient creates a new GetVoiceChatV2DevicesRenderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoiceChatV2DevicesRenderParamsWithHTTPClient(client *http.Client) *GetVoiceChatV2DevicesRenderParams {

	return &GetVoiceChatV2DevicesRenderParams{
		HTTPClient: client,
	}
}

/*GetVoiceChatV2DevicesRenderParams contains all the parameters to send to the API endpoint
for the get voice chat v2 devices render operation typically these are written to a http.Request
*/
type GetVoiceChatV2DevicesRenderParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get voice chat v2 devices render params
func (o *GetVoiceChatV2DevicesRenderParams) WithTimeout(timeout time.Duration) *GetVoiceChatV2DevicesRenderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voice chat v2 devices render params
func (o *GetVoiceChatV2DevicesRenderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voice chat v2 devices render params
func (o *GetVoiceChatV2DevicesRenderParams) WithContext(ctx context.Context) *GetVoiceChatV2DevicesRenderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voice chat v2 devices render params
func (o *GetVoiceChatV2DevicesRenderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voice chat v2 devices render params
func (o *GetVoiceChatV2DevicesRenderParams) WithHTTPClient(client *http.Client) *GetVoiceChatV2DevicesRenderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voice chat v2 devices render params
func (o *GetVoiceChatV2DevicesRenderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoiceChatV2DevicesRenderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
