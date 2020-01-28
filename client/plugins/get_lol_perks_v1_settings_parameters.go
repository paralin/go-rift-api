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

// NewGetLolPerksV1SettingsParams creates a new GetLolPerksV1SettingsParams object
// with the default values initialized.
func NewGetLolPerksV1SettingsParams() *GetLolPerksV1SettingsParams {

	return &GetLolPerksV1SettingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPerksV1SettingsParamsWithTimeout creates a new GetLolPerksV1SettingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPerksV1SettingsParamsWithTimeout(timeout time.Duration) *GetLolPerksV1SettingsParams {

	return &GetLolPerksV1SettingsParams{

		timeout: timeout,
	}
}

// NewGetLolPerksV1SettingsParamsWithContext creates a new GetLolPerksV1SettingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPerksV1SettingsParamsWithContext(ctx context.Context) *GetLolPerksV1SettingsParams {

	return &GetLolPerksV1SettingsParams{

		Context: ctx,
	}
}

// NewGetLolPerksV1SettingsParamsWithHTTPClient creates a new GetLolPerksV1SettingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPerksV1SettingsParamsWithHTTPClient(client *http.Client) *GetLolPerksV1SettingsParams {

	return &GetLolPerksV1SettingsParams{
		HTTPClient: client,
	}
}

/*GetLolPerksV1SettingsParams contains all the parameters to send to the API endpoint
for the get lol perks v1 settings operation typically these are written to a http.Request
*/
type GetLolPerksV1SettingsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol perks v1 settings params
func (o *GetLolPerksV1SettingsParams) WithTimeout(timeout time.Duration) *GetLolPerksV1SettingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol perks v1 settings params
func (o *GetLolPerksV1SettingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol perks v1 settings params
func (o *GetLolPerksV1SettingsParams) WithContext(ctx context.Context) *GetLolPerksV1SettingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol perks v1 settings params
func (o *GetLolPerksV1SettingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol perks v1 settings params
func (o *GetLolPerksV1SettingsParams) WithHTTPClient(client *http.Client) *GetLolPerksV1SettingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol perks v1 settings params
func (o *GetLolPerksV1SettingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPerksV1SettingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}