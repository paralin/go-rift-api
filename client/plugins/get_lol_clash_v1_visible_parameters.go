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

// NewGetLolClashV1VisibleParams creates a new GetLolClashV1VisibleParams object
// with the default values initialized.
func NewGetLolClashV1VisibleParams() *GetLolClashV1VisibleParams {

	return &GetLolClashV1VisibleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1VisibleParamsWithTimeout creates a new GetLolClashV1VisibleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1VisibleParamsWithTimeout(timeout time.Duration) *GetLolClashV1VisibleParams {

	return &GetLolClashV1VisibleParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1VisibleParamsWithContext creates a new GetLolClashV1VisibleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1VisibleParamsWithContext(ctx context.Context) *GetLolClashV1VisibleParams {

	return &GetLolClashV1VisibleParams{

		Context: ctx,
	}
}

// NewGetLolClashV1VisibleParamsWithHTTPClient creates a new GetLolClashV1VisibleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1VisibleParamsWithHTTPClient(client *http.Client) *GetLolClashV1VisibleParams {

	return &GetLolClashV1VisibleParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1VisibleParams contains all the parameters to send to the API endpoint
for the get lol clash v1 visible operation typically these are written to a http.Request
*/
type GetLolClashV1VisibleParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 visible params
func (o *GetLolClashV1VisibleParams) WithTimeout(timeout time.Duration) *GetLolClashV1VisibleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 visible params
func (o *GetLolClashV1VisibleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 visible params
func (o *GetLolClashV1VisibleParams) WithContext(ctx context.Context) *GetLolClashV1VisibleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 visible params
func (o *GetLolClashV1VisibleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 visible params
func (o *GetLolClashV1VisibleParams) WithHTTPClient(client *http.Client) *GetLolClashV1VisibleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 visible params
func (o *GetLolClashV1VisibleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1VisibleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}