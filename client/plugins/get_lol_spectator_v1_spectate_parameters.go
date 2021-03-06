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

// NewGetLolSpectatorV1SpectateParams creates a new GetLolSpectatorV1SpectateParams object
// with the default values initialized.
func NewGetLolSpectatorV1SpectateParams() *GetLolSpectatorV1SpectateParams {

	return &GetLolSpectatorV1SpectateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSpectatorV1SpectateParamsWithTimeout creates a new GetLolSpectatorV1SpectateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSpectatorV1SpectateParamsWithTimeout(timeout time.Duration) *GetLolSpectatorV1SpectateParams {

	return &GetLolSpectatorV1SpectateParams{

		timeout: timeout,
	}
}

// NewGetLolSpectatorV1SpectateParamsWithContext creates a new GetLolSpectatorV1SpectateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSpectatorV1SpectateParamsWithContext(ctx context.Context) *GetLolSpectatorV1SpectateParams {

	return &GetLolSpectatorV1SpectateParams{

		Context: ctx,
	}
}

// NewGetLolSpectatorV1SpectateParamsWithHTTPClient creates a new GetLolSpectatorV1SpectateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSpectatorV1SpectateParamsWithHTTPClient(client *http.Client) *GetLolSpectatorV1SpectateParams {

	return &GetLolSpectatorV1SpectateParams{
		HTTPClient: client,
	}
}

/*GetLolSpectatorV1SpectateParams contains all the parameters to send to the API endpoint
for the get lol spectator v1 spectate operation typically these are written to a http.Request
*/
type GetLolSpectatorV1SpectateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol spectator v1 spectate params
func (o *GetLolSpectatorV1SpectateParams) WithTimeout(timeout time.Duration) *GetLolSpectatorV1SpectateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol spectator v1 spectate params
func (o *GetLolSpectatorV1SpectateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol spectator v1 spectate params
func (o *GetLolSpectatorV1SpectateParams) WithContext(ctx context.Context) *GetLolSpectatorV1SpectateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol spectator v1 spectate params
func (o *GetLolSpectatorV1SpectateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol spectator v1 spectate params
func (o *GetLolSpectatorV1SpectateParams) WithHTTPClient(client *http.Client) *GetLolSpectatorV1SpectateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol spectator v1 spectate params
func (o *GetLolSpectatorV1SpectateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSpectatorV1SpectateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
