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

// NewGetPatcherV1ProductsParams creates a new GetPatcherV1ProductsParams object
// with the default values initialized.
func NewGetPatcherV1ProductsParams() *GetPatcherV1ProductsParams {

	return &GetPatcherV1ProductsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatcherV1ProductsParamsWithTimeout creates a new GetPatcherV1ProductsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatcherV1ProductsParamsWithTimeout(timeout time.Duration) *GetPatcherV1ProductsParams {

	return &GetPatcherV1ProductsParams{

		timeout: timeout,
	}
}

// NewGetPatcherV1ProductsParamsWithContext creates a new GetPatcherV1ProductsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatcherV1ProductsParamsWithContext(ctx context.Context) *GetPatcherV1ProductsParams {

	return &GetPatcherV1ProductsParams{

		Context: ctx,
	}
}

// NewGetPatcherV1ProductsParamsWithHTTPClient creates a new GetPatcherV1ProductsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatcherV1ProductsParamsWithHTTPClient(client *http.Client) *GetPatcherV1ProductsParams {

	return &GetPatcherV1ProductsParams{
		HTTPClient: client,
	}
}

/*GetPatcherV1ProductsParams contains all the parameters to send to the API endpoint
for the get patcher v1 products operation typically these are written to a http.Request
*/
type GetPatcherV1ProductsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patcher v1 products params
func (o *GetPatcherV1ProductsParams) WithTimeout(timeout time.Duration) *GetPatcherV1ProductsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patcher v1 products params
func (o *GetPatcherV1ProductsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patcher v1 products params
func (o *GetPatcherV1ProductsParams) WithContext(ctx context.Context) *GetPatcherV1ProductsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patcher v1 products params
func (o *GetPatcherV1ProductsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patcher v1 products params
func (o *GetPatcherV1ProductsParams) WithHTTPClient(client *http.Client) *GetPatcherV1ProductsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patcher v1 products params
func (o *GetPatcherV1ProductsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatcherV1ProductsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
