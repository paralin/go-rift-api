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

// NewGetPluginManagerV2DescriptionsParams creates a new GetPluginManagerV2DescriptionsParams object
// with the default values initialized.
func NewGetPluginManagerV2DescriptionsParams() *GetPluginManagerV2DescriptionsParams {

	return &GetPluginManagerV2DescriptionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginManagerV2DescriptionsParamsWithTimeout creates a new GetPluginManagerV2DescriptionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPluginManagerV2DescriptionsParamsWithTimeout(timeout time.Duration) *GetPluginManagerV2DescriptionsParams {

	return &GetPluginManagerV2DescriptionsParams{

		timeout: timeout,
	}
}

// NewGetPluginManagerV2DescriptionsParamsWithContext creates a new GetPluginManagerV2DescriptionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPluginManagerV2DescriptionsParamsWithContext(ctx context.Context) *GetPluginManagerV2DescriptionsParams {

	return &GetPluginManagerV2DescriptionsParams{

		Context: ctx,
	}
}

// NewGetPluginManagerV2DescriptionsParamsWithHTTPClient creates a new GetPluginManagerV2DescriptionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPluginManagerV2DescriptionsParamsWithHTTPClient(client *http.Client) *GetPluginManagerV2DescriptionsParams {

	return &GetPluginManagerV2DescriptionsParams{
		HTTPClient: client,
	}
}

/*GetPluginManagerV2DescriptionsParams contains all the parameters to send to the API endpoint
for the get plugin manager v2 descriptions operation typically these are written to a http.Request
*/
type GetPluginManagerV2DescriptionsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get plugin manager v2 descriptions params
func (o *GetPluginManagerV2DescriptionsParams) WithTimeout(timeout time.Duration) *GetPluginManagerV2DescriptionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin manager v2 descriptions params
func (o *GetPluginManagerV2DescriptionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin manager v2 descriptions params
func (o *GetPluginManagerV2DescriptionsParams) WithContext(ctx context.Context) *GetPluginManagerV2DescriptionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin manager v2 descriptions params
func (o *GetPluginManagerV2DescriptionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin manager v2 descriptions params
func (o *GetPluginManagerV2DescriptionsParams) WithHTTPClient(client *http.Client) *GetPluginManagerV2DescriptionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin manager v2 descriptions params
func (o *GetPluginManagerV2DescriptionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginManagerV2DescriptionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
