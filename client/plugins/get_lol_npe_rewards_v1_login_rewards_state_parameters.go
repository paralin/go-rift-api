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

// NewGetLolNpeRewardsV1LoginRewardsStateParams creates a new GetLolNpeRewardsV1LoginRewardsStateParams object
// with the default values initialized.
func NewGetLolNpeRewardsV1LoginRewardsStateParams() *GetLolNpeRewardsV1LoginRewardsStateParams {

	return &GetLolNpeRewardsV1LoginRewardsStateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolNpeRewardsV1LoginRewardsStateParamsWithTimeout creates a new GetLolNpeRewardsV1LoginRewardsStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolNpeRewardsV1LoginRewardsStateParamsWithTimeout(timeout time.Duration) *GetLolNpeRewardsV1LoginRewardsStateParams {

	return &GetLolNpeRewardsV1LoginRewardsStateParams{

		timeout: timeout,
	}
}

// NewGetLolNpeRewardsV1LoginRewardsStateParamsWithContext creates a new GetLolNpeRewardsV1LoginRewardsStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolNpeRewardsV1LoginRewardsStateParamsWithContext(ctx context.Context) *GetLolNpeRewardsV1LoginRewardsStateParams {

	return &GetLolNpeRewardsV1LoginRewardsStateParams{

		Context: ctx,
	}
}

// NewGetLolNpeRewardsV1LoginRewardsStateParamsWithHTTPClient creates a new GetLolNpeRewardsV1LoginRewardsStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolNpeRewardsV1LoginRewardsStateParamsWithHTTPClient(client *http.Client) *GetLolNpeRewardsV1LoginRewardsStateParams {

	return &GetLolNpeRewardsV1LoginRewardsStateParams{
		HTTPClient: client,
	}
}

/*GetLolNpeRewardsV1LoginRewardsStateParams contains all the parameters to send to the API endpoint
for the get lol npe rewards v1 login rewards state operation typically these are written to a http.Request
*/
type GetLolNpeRewardsV1LoginRewardsStateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol npe rewards v1 login rewards state params
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) WithTimeout(timeout time.Duration) *GetLolNpeRewardsV1LoginRewardsStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol npe rewards v1 login rewards state params
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol npe rewards v1 login rewards state params
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) WithContext(ctx context.Context) *GetLolNpeRewardsV1LoginRewardsStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol npe rewards v1 login rewards state params
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol npe rewards v1 login rewards state params
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) WithHTTPClient(client *http.Client) *GetLolNpeRewardsV1LoginRewardsStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol npe rewards v1 login rewards state params
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolNpeRewardsV1LoginRewardsStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}