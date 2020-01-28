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

// NewGetLolLoginV1LoginPlatformCredentialsParams creates a new GetLolLoginV1LoginPlatformCredentialsParams object
// with the default values initialized.
func NewGetLolLoginV1LoginPlatformCredentialsParams() *GetLolLoginV1LoginPlatformCredentialsParams {

	return &GetLolLoginV1LoginPlatformCredentialsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLoginV1LoginPlatformCredentialsParamsWithTimeout creates a new GetLolLoginV1LoginPlatformCredentialsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLoginV1LoginPlatformCredentialsParamsWithTimeout(timeout time.Duration) *GetLolLoginV1LoginPlatformCredentialsParams {

	return &GetLolLoginV1LoginPlatformCredentialsParams{

		timeout: timeout,
	}
}

// NewGetLolLoginV1LoginPlatformCredentialsParamsWithContext creates a new GetLolLoginV1LoginPlatformCredentialsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLoginV1LoginPlatformCredentialsParamsWithContext(ctx context.Context) *GetLolLoginV1LoginPlatformCredentialsParams {

	return &GetLolLoginV1LoginPlatformCredentialsParams{

		Context: ctx,
	}
}

// NewGetLolLoginV1LoginPlatformCredentialsParamsWithHTTPClient creates a new GetLolLoginV1LoginPlatformCredentialsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLoginV1LoginPlatformCredentialsParamsWithHTTPClient(client *http.Client) *GetLolLoginV1LoginPlatformCredentialsParams {

	return &GetLolLoginV1LoginPlatformCredentialsParams{
		HTTPClient: client,
	}
}

/*GetLolLoginV1LoginPlatformCredentialsParams contains all the parameters to send to the API endpoint
for the get lol login v1 login platform credentials operation typically these are written to a http.Request
*/
type GetLolLoginV1LoginPlatformCredentialsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol login v1 login platform credentials params
func (o *GetLolLoginV1LoginPlatformCredentialsParams) WithTimeout(timeout time.Duration) *GetLolLoginV1LoginPlatformCredentialsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol login v1 login platform credentials params
func (o *GetLolLoginV1LoginPlatformCredentialsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol login v1 login platform credentials params
func (o *GetLolLoginV1LoginPlatformCredentialsParams) WithContext(ctx context.Context) *GetLolLoginV1LoginPlatformCredentialsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol login v1 login platform credentials params
func (o *GetLolLoginV1LoginPlatformCredentialsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol login v1 login platform credentials params
func (o *GetLolLoginV1LoginPlatformCredentialsParams) WithHTTPClient(client *http.Client) *GetLolLoginV1LoginPlatformCredentialsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol login v1 login platform credentials params
func (o *GetLolLoginV1LoginPlatformCredentialsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLoginV1LoginPlatformCredentialsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
