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

// NewGetLolNpeRewardsV1LoginRewardsParams creates a new GetLolNpeRewardsV1LoginRewardsParams object
// with the default values initialized.
func NewGetLolNpeRewardsV1LoginRewardsParams() *GetLolNpeRewardsV1LoginRewardsParams {

	return &GetLolNpeRewardsV1LoginRewardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolNpeRewardsV1LoginRewardsParamsWithTimeout creates a new GetLolNpeRewardsV1LoginRewardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolNpeRewardsV1LoginRewardsParamsWithTimeout(timeout time.Duration) *GetLolNpeRewardsV1LoginRewardsParams {

	return &GetLolNpeRewardsV1LoginRewardsParams{

		timeout: timeout,
	}
}

// NewGetLolNpeRewardsV1LoginRewardsParamsWithContext creates a new GetLolNpeRewardsV1LoginRewardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolNpeRewardsV1LoginRewardsParamsWithContext(ctx context.Context) *GetLolNpeRewardsV1LoginRewardsParams {

	return &GetLolNpeRewardsV1LoginRewardsParams{

		Context: ctx,
	}
}

// NewGetLolNpeRewardsV1LoginRewardsParamsWithHTTPClient creates a new GetLolNpeRewardsV1LoginRewardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolNpeRewardsV1LoginRewardsParamsWithHTTPClient(client *http.Client) *GetLolNpeRewardsV1LoginRewardsParams {

	return &GetLolNpeRewardsV1LoginRewardsParams{
		HTTPClient: client,
	}
}

/*GetLolNpeRewardsV1LoginRewardsParams contains all the parameters to send to the API endpoint
for the get lol npe rewards v1 login rewards operation typically these are written to a http.Request
*/
type GetLolNpeRewardsV1LoginRewardsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol npe rewards v1 login rewards params
func (o *GetLolNpeRewardsV1LoginRewardsParams) WithTimeout(timeout time.Duration) *GetLolNpeRewardsV1LoginRewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol npe rewards v1 login rewards params
func (o *GetLolNpeRewardsV1LoginRewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol npe rewards v1 login rewards params
func (o *GetLolNpeRewardsV1LoginRewardsParams) WithContext(ctx context.Context) *GetLolNpeRewardsV1LoginRewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol npe rewards v1 login rewards params
func (o *GetLolNpeRewardsV1LoginRewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol npe rewards v1 login rewards params
func (o *GetLolNpeRewardsV1LoginRewardsParams) WithHTTPClient(client *http.Client) *GetLolNpeRewardsV1LoginRewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol npe rewards v1 login rewards params
func (o *GetLolNpeRewardsV1LoginRewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolNpeRewardsV1LoginRewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
