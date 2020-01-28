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

// NewPutClientConfigV1RefreshConfigStatusParams creates a new PutClientConfigV1RefreshConfigStatusParams object
// with the default values initialized.
func NewPutClientConfigV1RefreshConfigStatusParams() *PutClientConfigV1RefreshConfigStatusParams {

	return &PutClientConfigV1RefreshConfigStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClientConfigV1RefreshConfigStatusParamsWithTimeout creates a new PutClientConfigV1RefreshConfigStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClientConfigV1RefreshConfigStatusParamsWithTimeout(timeout time.Duration) *PutClientConfigV1RefreshConfigStatusParams {

	return &PutClientConfigV1RefreshConfigStatusParams{

		timeout: timeout,
	}
}

// NewPutClientConfigV1RefreshConfigStatusParamsWithContext creates a new PutClientConfigV1RefreshConfigStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutClientConfigV1RefreshConfigStatusParamsWithContext(ctx context.Context) *PutClientConfigV1RefreshConfigStatusParams {

	return &PutClientConfigV1RefreshConfigStatusParams{

		Context: ctx,
	}
}

// NewPutClientConfigV1RefreshConfigStatusParamsWithHTTPClient creates a new PutClientConfigV1RefreshConfigStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClientConfigV1RefreshConfigStatusParamsWithHTTPClient(client *http.Client) *PutClientConfigV1RefreshConfigStatusParams {

	return &PutClientConfigV1RefreshConfigStatusParams{
		HTTPClient: client,
	}
}

/*PutClientConfigV1RefreshConfigStatusParams contains all the parameters to send to the API endpoint
for the put client config v1 refresh config status operation typically these are written to a http.Request
*/
type PutClientConfigV1RefreshConfigStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put client config v1 refresh config status params
func (o *PutClientConfigV1RefreshConfigStatusParams) WithTimeout(timeout time.Duration) *PutClientConfigV1RefreshConfigStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put client config v1 refresh config status params
func (o *PutClientConfigV1RefreshConfigStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put client config v1 refresh config status params
func (o *PutClientConfigV1RefreshConfigStatusParams) WithContext(ctx context.Context) *PutClientConfigV1RefreshConfigStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put client config v1 refresh config status params
func (o *PutClientConfigV1RefreshConfigStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put client config v1 refresh config status params
func (o *PutClientConfigV1RefreshConfigStatusParams) WithHTTPClient(client *http.Client) *PutClientConfigV1RefreshConfigStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put client config v1 refresh config status params
func (o *PutClientConfigV1RefreshConfigStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PutClientConfigV1RefreshConfigStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}