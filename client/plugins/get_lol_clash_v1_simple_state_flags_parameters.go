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

// NewGetLolClashV1SimpleStateFlagsParams creates a new GetLolClashV1SimpleStateFlagsParams object
// with the default values initialized.
func NewGetLolClashV1SimpleStateFlagsParams() *GetLolClashV1SimpleStateFlagsParams {

	return &GetLolClashV1SimpleStateFlagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1SimpleStateFlagsParamsWithTimeout creates a new GetLolClashV1SimpleStateFlagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1SimpleStateFlagsParamsWithTimeout(timeout time.Duration) *GetLolClashV1SimpleStateFlagsParams {

	return &GetLolClashV1SimpleStateFlagsParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1SimpleStateFlagsParamsWithContext creates a new GetLolClashV1SimpleStateFlagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1SimpleStateFlagsParamsWithContext(ctx context.Context) *GetLolClashV1SimpleStateFlagsParams {

	return &GetLolClashV1SimpleStateFlagsParams{

		Context: ctx,
	}
}

// NewGetLolClashV1SimpleStateFlagsParamsWithHTTPClient creates a new GetLolClashV1SimpleStateFlagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1SimpleStateFlagsParamsWithHTTPClient(client *http.Client) *GetLolClashV1SimpleStateFlagsParams {

	return &GetLolClashV1SimpleStateFlagsParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1SimpleStateFlagsParams contains all the parameters to send to the API endpoint
for the get lol clash v1 simple state flags operation typically these are written to a http.Request
*/
type GetLolClashV1SimpleStateFlagsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 simple state flags params
func (o *GetLolClashV1SimpleStateFlagsParams) WithTimeout(timeout time.Duration) *GetLolClashV1SimpleStateFlagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 simple state flags params
func (o *GetLolClashV1SimpleStateFlagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 simple state flags params
func (o *GetLolClashV1SimpleStateFlagsParams) WithContext(ctx context.Context) *GetLolClashV1SimpleStateFlagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 simple state flags params
func (o *GetLolClashV1SimpleStateFlagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 simple state flags params
func (o *GetLolClashV1SimpleStateFlagsParams) WithHTTPClient(client *http.Client) *GetLolClashV1SimpleStateFlagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 simple state flags params
func (o *GetLolClashV1SimpleStateFlagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1SimpleStateFlagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
