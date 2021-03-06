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

// NewGetRecofrienderV1RegistrationsParams creates a new GetRecofrienderV1RegistrationsParams object
// with the default values initialized.
func NewGetRecofrienderV1RegistrationsParams() *GetRecofrienderV1RegistrationsParams {
	var ()
	return &GetRecofrienderV1RegistrationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecofrienderV1RegistrationsParamsWithTimeout creates a new GetRecofrienderV1RegistrationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecofrienderV1RegistrationsParamsWithTimeout(timeout time.Duration) *GetRecofrienderV1RegistrationsParams {
	var ()
	return &GetRecofrienderV1RegistrationsParams{

		timeout: timeout,
	}
}

// NewGetRecofrienderV1RegistrationsParamsWithContext creates a new GetRecofrienderV1RegistrationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecofrienderV1RegistrationsParamsWithContext(ctx context.Context) *GetRecofrienderV1RegistrationsParams {
	var ()
	return &GetRecofrienderV1RegistrationsParams{

		Context: ctx,
	}
}

// NewGetRecofrienderV1RegistrationsParamsWithHTTPClient creates a new GetRecofrienderV1RegistrationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecofrienderV1RegistrationsParamsWithHTTPClient(client *http.Client) *GetRecofrienderV1RegistrationsParams {
	var ()
	return &GetRecofrienderV1RegistrationsParams{
		HTTPClient: client,
	}
}

/*GetRecofrienderV1RegistrationsParams contains all the parameters to send to the API endpoint
for the get recofriender v1 registrations operation typically these are written to a http.Request
*/
type GetRecofrienderV1RegistrationsParams struct {

	/*Cb*/
	Cb *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) WithTimeout(timeout time.Duration) *GetRecofrienderV1RegistrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) WithContext(ctx context.Context) *GetRecofrienderV1RegistrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) WithHTTPClient(client *http.Client) *GetRecofrienderV1RegistrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCb adds the cb to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) WithCb(cb *string) *GetRecofrienderV1RegistrationsParams {
	o.SetCb(cb)
	return o
}

// SetCb adds the cb to the get recofriender v1 registrations params
func (o *GetRecofrienderV1RegistrationsParams) SetCb(cb *string) {
	o.Cb = cb
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecofrienderV1RegistrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cb != nil {

		// query param cb
		var qrCb string
		if o.Cb != nil {
			qrCb = *o.Cb
		}
		qCb := qrCb
		if qCb != "" {
			if err := r.SetQueryParam("cb", qCb); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
