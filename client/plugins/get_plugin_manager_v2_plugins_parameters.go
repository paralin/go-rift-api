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

// NewGetPluginManagerV2PluginsParams creates a new GetPluginManagerV2PluginsParams object
// with the default values initialized.
func NewGetPluginManagerV2PluginsParams() *GetPluginManagerV2PluginsParams {

	return &GetPluginManagerV2PluginsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPluginManagerV2PluginsParamsWithTimeout creates a new GetPluginManagerV2PluginsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPluginManagerV2PluginsParamsWithTimeout(timeout time.Duration) *GetPluginManagerV2PluginsParams {

	return &GetPluginManagerV2PluginsParams{

		timeout: timeout,
	}
}

// NewGetPluginManagerV2PluginsParamsWithContext creates a new GetPluginManagerV2PluginsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPluginManagerV2PluginsParamsWithContext(ctx context.Context) *GetPluginManagerV2PluginsParams {

	return &GetPluginManagerV2PluginsParams{

		Context: ctx,
	}
}

// NewGetPluginManagerV2PluginsParamsWithHTTPClient creates a new GetPluginManagerV2PluginsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPluginManagerV2PluginsParamsWithHTTPClient(client *http.Client) *GetPluginManagerV2PluginsParams {

	return &GetPluginManagerV2PluginsParams{
		HTTPClient: client,
	}
}

/*GetPluginManagerV2PluginsParams contains all the parameters to send to the API endpoint
for the get plugin manager v2 plugins operation typically these are written to a http.Request
*/
type GetPluginManagerV2PluginsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get plugin manager v2 plugins params
func (o *GetPluginManagerV2PluginsParams) WithTimeout(timeout time.Duration) *GetPluginManagerV2PluginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get plugin manager v2 plugins params
func (o *GetPluginManagerV2PluginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get plugin manager v2 plugins params
func (o *GetPluginManagerV2PluginsParams) WithContext(ctx context.Context) *GetPluginManagerV2PluginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get plugin manager v2 plugins params
func (o *GetPluginManagerV2PluginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get plugin manager v2 plugins params
func (o *GetPluginManagerV2PluginsParams) WithHTTPClient(client *http.Client) *GetPluginManagerV2PluginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get plugin manager v2 plugins params
func (o *GetPluginManagerV2PluginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPluginManagerV2PluginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
