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

// NewGetLolLoadoutsV4LoadoutsScopeAccountParams creates a new GetLolLoadoutsV4LoadoutsScopeAccountParams object
// with the default values initialized.
func NewGetLolLoadoutsV4LoadoutsScopeAccountParams() *GetLolLoadoutsV4LoadoutsScopeAccountParams {

	return &GetLolLoadoutsV4LoadoutsScopeAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLoadoutsV4LoadoutsScopeAccountParamsWithTimeout creates a new GetLolLoadoutsV4LoadoutsScopeAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLoadoutsV4LoadoutsScopeAccountParamsWithTimeout(timeout time.Duration) *GetLolLoadoutsV4LoadoutsScopeAccountParams {

	return &GetLolLoadoutsV4LoadoutsScopeAccountParams{

		timeout: timeout,
	}
}

// NewGetLolLoadoutsV4LoadoutsScopeAccountParamsWithContext creates a new GetLolLoadoutsV4LoadoutsScopeAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLoadoutsV4LoadoutsScopeAccountParamsWithContext(ctx context.Context) *GetLolLoadoutsV4LoadoutsScopeAccountParams {

	return &GetLolLoadoutsV4LoadoutsScopeAccountParams{

		Context: ctx,
	}
}

// NewGetLolLoadoutsV4LoadoutsScopeAccountParamsWithHTTPClient creates a new GetLolLoadoutsV4LoadoutsScopeAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLoadoutsV4LoadoutsScopeAccountParamsWithHTTPClient(client *http.Client) *GetLolLoadoutsV4LoadoutsScopeAccountParams {

	return &GetLolLoadoutsV4LoadoutsScopeAccountParams{
		HTTPClient: client,
	}
}

/*GetLolLoadoutsV4LoadoutsScopeAccountParams contains all the parameters to send to the API endpoint
for the get lol loadouts v4 loadouts scope account operation typically these are written to a http.Request
*/
type GetLolLoadoutsV4LoadoutsScopeAccountParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol loadouts v4 loadouts scope account params
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) WithTimeout(timeout time.Duration) *GetLolLoadoutsV4LoadoutsScopeAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol loadouts v4 loadouts scope account params
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol loadouts v4 loadouts scope account params
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) WithContext(ctx context.Context) *GetLolLoadoutsV4LoadoutsScopeAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol loadouts v4 loadouts scope account params
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol loadouts v4 loadouts scope account params
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) WithHTTPClient(client *http.Client) *GetLolLoadoutsV4LoadoutsScopeAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol loadouts v4 loadouts scope account params
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLoadoutsV4LoadoutsScopeAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}