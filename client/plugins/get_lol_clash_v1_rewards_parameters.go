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

// NewGetLolClashV1RewardsParams creates a new GetLolClashV1RewardsParams object
// with the default values initialized.
func NewGetLolClashV1RewardsParams() *GetLolClashV1RewardsParams {

	return &GetLolClashV1RewardsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1RewardsParamsWithTimeout creates a new GetLolClashV1RewardsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1RewardsParamsWithTimeout(timeout time.Duration) *GetLolClashV1RewardsParams {

	return &GetLolClashV1RewardsParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1RewardsParamsWithContext creates a new GetLolClashV1RewardsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1RewardsParamsWithContext(ctx context.Context) *GetLolClashV1RewardsParams {

	return &GetLolClashV1RewardsParams{

		Context: ctx,
	}
}

// NewGetLolClashV1RewardsParamsWithHTTPClient creates a new GetLolClashV1RewardsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1RewardsParamsWithHTTPClient(client *http.Client) *GetLolClashV1RewardsParams {

	return &GetLolClashV1RewardsParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1RewardsParams contains all the parameters to send to the API endpoint
for the get lol clash v1 rewards operation typically these are written to a http.Request
*/
type GetLolClashV1RewardsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 rewards params
func (o *GetLolClashV1RewardsParams) WithTimeout(timeout time.Duration) *GetLolClashV1RewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 rewards params
func (o *GetLolClashV1RewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 rewards params
func (o *GetLolClashV1RewardsParams) WithContext(ctx context.Context) *GetLolClashV1RewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 rewards params
func (o *GetLolClashV1RewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 rewards params
func (o *GetLolClashV1RewardsParams) WithHTTPClient(client *http.Client) *GetLolClashV1RewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 rewards params
func (o *GetLolClashV1RewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1RewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
