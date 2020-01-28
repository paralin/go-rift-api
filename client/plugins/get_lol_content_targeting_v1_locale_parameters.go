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

// NewGetLolContentTargetingV1LocaleParams creates a new GetLolContentTargetingV1LocaleParams object
// with the default values initialized.
func NewGetLolContentTargetingV1LocaleParams() *GetLolContentTargetingV1LocaleParams {

	return &GetLolContentTargetingV1LocaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolContentTargetingV1LocaleParamsWithTimeout creates a new GetLolContentTargetingV1LocaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolContentTargetingV1LocaleParamsWithTimeout(timeout time.Duration) *GetLolContentTargetingV1LocaleParams {

	return &GetLolContentTargetingV1LocaleParams{

		timeout: timeout,
	}
}

// NewGetLolContentTargetingV1LocaleParamsWithContext creates a new GetLolContentTargetingV1LocaleParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolContentTargetingV1LocaleParamsWithContext(ctx context.Context) *GetLolContentTargetingV1LocaleParams {

	return &GetLolContentTargetingV1LocaleParams{

		Context: ctx,
	}
}

// NewGetLolContentTargetingV1LocaleParamsWithHTTPClient creates a new GetLolContentTargetingV1LocaleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolContentTargetingV1LocaleParamsWithHTTPClient(client *http.Client) *GetLolContentTargetingV1LocaleParams {

	return &GetLolContentTargetingV1LocaleParams{
		HTTPClient: client,
	}
}

/*GetLolContentTargetingV1LocaleParams contains all the parameters to send to the API endpoint
for the get lol content targeting v1 locale operation typically these are written to a http.Request
*/
type GetLolContentTargetingV1LocaleParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol content targeting v1 locale params
func (o *GetLolContentTargetingV1LocaleParams) WithTimeout(timeout time.Duration) *GetLolContentTargetingV1LocaleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol content targeting v1 locale params
func (o *GetLolContentTargetingV1LocaleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol content targeting v1 locale params
func (o *GetLolContentTargetingV1LocaleParams) WithContext(ctx context.Context) *GetLolContentTargetingV1LocaleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol content targeting v1 locale params
func (o *GetLolContentTargetingV1LocaleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol content targeting v1 locale params
func (o *GetLolContentTargetingV1LocaleParams) WithHTTPClient(client *http.Client) *GetLolContentTargetingV1LocaleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol content targeting v1 locale params
func (o *GetLolContentTargetingV1LocaleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolContentTargetingV1LocaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}