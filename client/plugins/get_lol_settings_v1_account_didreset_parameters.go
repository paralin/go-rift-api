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

// NewGetLolSettingsV1AccountDidresetParams creates a new GetLolSettingsV1AccountDidresetParams object
// with the default values initialized.
func NewGetLolSettingsV1AccountDidresetParams() *GetLolSettingsV1AccountDidresetParams {

	return &GetLolSettingsV1AccountDidresetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSettingsV1AccountDidresetParamsWithTimeout creates a new GetLolSettingsV1AccountDidresetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSettingsV1AccountDidresetParamsWithTimeout(timeout time.Duration) *GetLolSettingsV1AccountDidresetParams {

	return &GetLolSettingsV1AccountDidresetParams{

		timeout: timeout,
	}
}

// NewGetLolSettingsV1AccountDidresetParamsWithContext creates a new GetLolSettingsV1AccountDidresetParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSettingsV1AccountDidresetParamsWithContext(ctx context.Context) *GetLolSettingsV1AccountDidresetParams {

	return &GetLolSettingsV1AccountDidresetParams{

		Context: ctx,
	}
}

// NewGetLolSettingsV1AccountDidresetParamsWithHTTPClient creates a new GetLolSettingsV1AccountDidresetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSettingsV1AccountDidresetParamsWithHTTPClient(client *http.Client) *GetLolSettingsV1AccountDidresetParams {

	return &GetLolSettingsV1AccountDidresetParams{
		HTTPClient: client,
	}
}

/*GetLolSettingsV1AccountDidresetParams contains all the parameters to send to the API endpoint
for the get lol settings v1 account didreset operation typically these are written to a http.Request
*/
type GetLolSettingsV1AccountDidresetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol settings v1 account didreset params
func (o *GetLolSettingsV1AccountDidresetParams) WithTimeout(timeout time.Duration) *GetLolSettingsV1AccountDidresetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol settings v1 account didreset params
func (o *GetLolSettingsV1AccountDidresetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol settings v1 account didreset params
func (o *GetLolSettingsV1AccountDidresetParams) WithContext(ctx context.Context) *GetLolSettingsV1AccountDidresetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol settings v1 account didreset params
func (o *GetLolSettingsV1AccountDidresetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol settings v1 account didreset params
func (o *GetLolSettingsV1AccountDidresetParams) WithHTTPClient(client *http.Client) *GetLolSettingsV1AccountDidresetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol settings v1 account didreset params
func (o *GetLolSettingsV1AccountDidresetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSettingsV1AccountDidresetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
