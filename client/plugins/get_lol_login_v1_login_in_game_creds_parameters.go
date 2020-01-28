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

// NewGetLolLoginV1LoginInGameCredsParams creates a new GetLolLoginV1LoginInGameCredsParams object
// with the default values initialized.
func NewGetLolLoginV1LoginInGameCredsParams() *GetLolLoginV1LoginInGameCredsParams {

	return &GetLolLoginV1LoginInGameCredsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLoginV1LoginInGameCredsParamsWithTimeout creates a new GetLolLoginV1LoginInGameCredsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLoginV1LoginInGameCredsParamsWithTimeout(timeout time.Duration) *GetLolLoginV1LoginInGameCredsParams {

	return &GetLolLoginV1LoginInGameCredsParams{

		timeout: timeout,
	}
}

// NewGetLolLoginV1LoginInGameCredsParamsWithContext creates a new GetLolLoginV1LoginInGameCredsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLoginV1LoginInGameCredsParamsWithContext(ctx context.Context) *GetLolLoginV1LoginInGameCredsParams {

	return &GetLolLoginV1LoginInGameCredsParams{

		Context: ctx,
	}
}

// NewGetLolLoginV1LoginInGameCredsParamsWithHTTPClient creates a new GetLolLoginV1LoginInGameCredsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLoginV1LoginInGameCredsParamsWithHTTPClient(client *http.Client) *GetLolLoginV1LoginInGameCredsParams {

	return &GetLolLoginV1LoginInGameCredsParams{
		HTTPClient: client,
	}
}

/*GetLolLoginV1LoginInGameCredsParams contains all the parameters to send to the API endpoint
for the get lol login v1 login in game creds operation typically these are written to a http.Request
*/
type GetLolLoginV1LoginInGameCredsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol login v1 login in game creds params
func (o *GetLolLoginV1LoginInGameCredsParams) WithTimeout(timeout time.Duration) *GetLolLoginV1LoginInGameCredsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol login v1 login in game creds params
func (o *GetLolLoginV1LoginInGameCredsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol login v1 login in game creds params
func (o *GetLolLoginV1LoginInGameCredsParams) WithContext(ctx context.Context) *GetLolLoginV1LoginInGameCredsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol login v1 login in game creds params
func (o *GetLolLoginV1LoginInGameCredsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol login v1 login in game creds params
func (o *GetLolLoginV1LoginInGameCredsParams) WithHTTPClient(client *http.Client) *GetLolLoginV1LoginInGameCredsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol login v1 login in game creds params
func (o *GetLolLoginV1LoginInGameCredsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLoginV1LoginInGameCredsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}