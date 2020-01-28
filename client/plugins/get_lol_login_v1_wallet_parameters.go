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

// NewGetLolLoginV1WalletParams creates a new GetLolLoginV1WalletParams object
// with the default values initialized.
func NewGetLolLoginV1WalletParams() *GetLolLoginV1WalletParams {

	return &GetLolLoginV1WalletParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLoginV1WalletParamsWithTimeout creates a new GetLolLoginV1WalletParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLoginV1WalletParamsWithTimeout(timeout time.Duration) *GetLolLoginV1WalletParams {

	return &GetLolLoginV1WalletParams{

		timeout: timeout,
	}
}

// NewGetLolLoginV1WalletParamsWithContext creates a new GetLolLoginV1WalletParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLoginV1WalletParamsWithContext(ctx context.Context) *GetLolLoginV1WalletParams {

	return &GetLolLoginV1WalletParams{

		Context: ctx,
	}
}

// NewGetLolLoginV1WalletParamsWithHTTPClient creates a new GetLolLoginV1WalletParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLoginV1WalletParamsWithHTTPClient(client *http.Client) *GetLolLoginV1WalletParams {

	return &GetLolLoginV1WalletParams{
		HTTPClient: client,
	}
}

/*GetLolLoginV1WalletParams contains all the parameters to send to the API endpoint
for the get lol login v1 wallet operation typically these are written to a http.Request
*/
type GetLolLoginV1WalletParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol login v1 wallet params
func (o *GetLolLoginV1WalletParams) WithTimeout(timeout time.Duration) *GetLolLoginV1WalletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol login v1 wallet params
func (o *GetLolLoginV1WalletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol login v1 wallet params
func (o *GetLolLoginV1WalletParams) WithContext(ctx context.Context) *GetLolLoginV1WalletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol login v1 wallet params
func (o *GetLolLoginV1WalletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol login v1 wallet params
func (o *GetLolLoginV1WalletParams) WithHTTPClient(client *http.Client) *GetLolLoginV1WalletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol login v1 wallet params
func (o *GetLolLoginV1WalletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLoginV1WalletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
