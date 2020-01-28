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

// NewGetLolPersonalizedOffersV1OffersParams creates a new GetLolPersonalizedOffersV1OffersParams object
// with the default values initialized.
func NewGetLolPersonalizedOffersV1OffersParams() *GetLolPersonalizedOffersV1OffersParams {

	return &GetLolPersonalizedOffersV1OffersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPersonalizedOffersV1OffersParamsWithTimeout creates a new GetLolPersonalizedOffersV1OffersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPersonalizedOffersV1OffersParamsWithTimeout(timeout time.Duration) *GetLolPersonalizedOffersV1OffersParams {

	return &GetLolPersonalizedOffersV1OffersParams{

		timeout: timeout,
	}
}

// NewGetLolPersonalizedOffersV1OffersParamsWithContext creates a new GetLolPersonalizedOffersV1OffersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPersonalizedOffersV1OffersParamsWithContext(ctx context.Context) *GetLolPersonalizedOffersV1OffersParams {

	return &GetLolPersonalizedOffersV1OffersParams{

		Context: ctx,
	}
}

// NewGetLolPersonalizedOffersV1OffersParamsWithHTTPClient creates a new GetLolPersonalizedOffersV1OffersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPersonalizedOffersV1OffersParamsWithHTTPClient(client *http.Client) *GetLolPersonalizedOffersV1OffersParams {

	return &GetLolPersonalizedOffersV1OffersParams{
		HTTPClient: client,
	}
}

/*GetLolPersonalizedOffersV1OffersParams contains all the parameters to send to the API endpoint
for the get lol personalized offers v1 offers operation typically these are written to a http.Request
*/
type GetLolPersonalizedOffersV1OffersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol personalized offers v1 offers params
func (o *GetLolPersonalizedOffersV1OffersParams) WithTimeout(timeout time.Duration) *GetLolPersonalizedOffersV1OffersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol personalized offers v1 offers params
func (o *GetLolPersonalizedOffersV1OffersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol personalized offers v1 offers params
func (o *GetLolPersonalizedOffersV1OffersParams) WithContext(ctx context.Context) *GetLolPersonalizedOffersV1OffersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol personalized offers v1 offers params
func (o *GetLolPersonalizedOffersV1OffersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol personalized offers v1 offers params
func (o *GetLolPersonalizedOffersV1OffersParams) WithHTTPClient(client *http.Client) *GetLolPersonalizedOffersV1OffersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol personalized offers v1 offers params
func (o *GetLolPersonalizedOffersV1OffersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPersonalizedOffersV1OffersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
