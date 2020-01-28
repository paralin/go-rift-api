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

// NewGetLolSettingsV2ReadyParams creates a new GetLolSettingsV2ReadyParams object
// with the default values initialized.
func NewGetLolSettingsV2ReadyParams() *GetLolSettingsV2ReadyParams {

	return &GetLolSettingsV2ReadyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSettingsV2ReadyParamsWithTimeout creates a new GetLolSettingsV2ReadyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSettingsV2ReadyParamsWithTimeout(timeout time.Duration) *GetLolSettingsV2ReadyParams {

	return &GetLolSettingsV2ReadyParams{

		timeout: timeout,
	}
}

// NewGetLolSettingsV2ReadyParamsWithContext creates a new GetLolSettingsV2ReadyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSettingsV2ReadyParamsWithContext(ctx context.Context) *GetLolSettingsV2ReadyParams {

	return &GetLolSettingsV2ReadyParams{

		Context: ctx,
	}
}

// NewGetLolSettingsV2ReadyParamsWithHTTPClient creates a new GetLolSettingsV2ReadyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSettingsV2ReadyParamsWithHTTPClient(client *http.Client) *GetLolSettingsV2ReadyParams {

	return &GetLolSettingsV2ReadyParams{
		HTTPClient: client,
	}
}

/*GetLolSettingsV2ReadyParams contains all the parameters to send to the API endpoint
for the get lol settings v2 ready operation typically these are written to a http.Request
*/
type GetLolSettingsV2ReadyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol settings v2 ready params
func (o *GetLolSettingsV2ReadyParams) WithTimeout(timeout time.Duration) *GetLolSettingsV2ReadyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol settings v2 ready params
func (o *GetLolSettingsV2ReadyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol settings v2 ready params
func (o *GetLolSettingsV2ReadyParams) WithContext(ctx context.Context) *GetLolSettingsV2ReadyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol settings v2 ready params
func (o *GetLolSettingsV2ReadyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol settings v2 ready params
func (o *GetLolSettingsV2ReadyParams) WithHTTPClient(client *http.Client) *GetLolSettingsV2ReadyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol settings v2 ready params
func (o *GetLolSettingsV2ReadyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSettingsV2ReadyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}