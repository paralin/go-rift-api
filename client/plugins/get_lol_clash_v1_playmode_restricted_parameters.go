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

// NewGetLolClashV1PlaymodeRestrictedParams creates a new GetLolClashV1PlaymodeRestrictedParams object
// with the default values initialized.
func NewGetLolClashV1PlaymodeRestrictedParams() *GetLolClashV1PlaymodeRestrictedParams {

	return &GetLolClashV1PlaymodeRestrictedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1PlaymodeRestrictedParamsWithTimeout creates a new GetLolClashV1PlaymodeRestrictedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1PlaymodeRestrictedParamsWithTimeout(timeout time.Duration) *GetLolClashV1PlaymodeRestrictedParams {

	return &GetLolClashV1PlaymodeRestrictedParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1PlaymodeRestrictedParamsWithContext creates a new GetLolClashV1PlaymodeRestrictedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1PlaymodeRestrictedParamsWithContext(ctx context.Context) *GetLolClashV1PlaymodeRestrictedParams {

	return &GetLolClashV1PlaymodeRestrictedParams{

		Context: ctx,
	}
}

// NewGetLolClashV1PlaymodeRestrictedParamsWithHTTPClient creates a new GetLolClashV1PlaymodeRestrictedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1PlaymodeRestrictedParamsWithHTTPClient(client *http.Client) *GetLolClashV1PlaymodeRestrictedParams {

	return &GetLolClashV1PlaymodeRestrictedParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1PlaymodeRestrictedParams contains all the parameters to send to the API endpoint
for the get lol clash v1 playmode restricted operation typically these are written to a http.Request
*/
type GetLolClashV1PlaymodeRestrictedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 playmode restricted params
func (o *GetLolClashV1PlaymodeRestrictedParams) WithTimeout(timeout time.Duration) *GetLolClashV1PlaymodeRestrictedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 playmode restricted params
func (o *GetLolClashV1PlaymodeRestrictedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 playmode restricted params
func (o *GetLolClashV1PlaymodeRestrictedParams) WithContext(ctx context.Context) *GetLolClashV1PlaymodeRestrictedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 playmode restricted params
func (o *GetLolClashV1PlaymodeRestrictedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 playmode restricted params
func (o *GetLolClashV1PlaymodeRestrictedParams) WithHTTPClient(client *http.Client) *GetLolClashV1PlaymodeRestrictedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 playmode restricted params
func (o *GetLolClashV1PlaymodeRestrictedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1PlaymodeRestrictedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
