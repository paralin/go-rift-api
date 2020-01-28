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

// NewGetLolAcsV1DeltaParams creates a new GetLolAcsV1DeltaParams object
// with the default values initialized.
func NewGetLolAcsV1DeltaParams() *GetLolAcsV1DeltaParams {

	return &GetLolAcsV1DeltaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolAcsV1DeltaParamsWithTimeout creates a new GetLolAcsV1DeltaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolAcsV1DeltaParamsWithTimeout(timeout time.Duration) *GetLolAcsV1DeltaParams {

	return &GetLolAcsV1DeltaParams{

		timeout: timeout,
	}
}

// NewGetLolAcsV1DeltaParamsWithContext creates a new GetLolAcsV1DeltaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolAcsV1DeltaParamsWithContext(ctx context.Context) *GetLolAcsV1DeltaParams {

	return &GetLolAcsV1DeltaParams{

		Context: ctx,
	}
}

// NewGetLolAcsV1DeltaParamsWithHTTPClient creates a new GetLolAcsV1DeltaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolAcsV1DeltaParamsWithHTTPClient(client *http.Client) *GetLolAcsV1DeltaParams {

	return &GetLolAcsV1DeltaParams{
		HTTPClient: client,
	}
}

/*GetLolAcsV1DeltaParams contains all the parameters to send to the API endpoint
for the get lol acs v1 delta operation typically these are written to a http.Request
*/
type GetLolAcsV1DeltaParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol acs v1 delta params
func (o *GetLolAcsV1DeltaParams) WithTimeout(timeout time.Duration) *GetLolAcsV1DeltaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol acs v1 delta params
func (o *GetLolAcsV1DeltaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol acs v1 delta params
func (o *GetLolAcsV1DeltaParams) WithContext(ctx context.Context) *GetLolAcsV1DeltaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol acs v1 delta params
func (o *GetLolAcsV1DeltaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol acs v1 delta params
func (o *GetLolAcsV1DeltaParams) WithHTTPClient(client *http.Client) *GetLolAcsV1DeltaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol acs v1 delta params
func (o *GetLolAcsV1DeltaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolAcsV1DeltaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}