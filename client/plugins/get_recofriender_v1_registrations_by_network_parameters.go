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

// NewGetRecofrienderV1RegistrationsByNetworkParams creates a new GetRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized.
func NewGetRecofrienderV1RegistrationsByNetworkParams() *GetRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &GetRecofrienderV1RegistrationsByNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecofrienderV1RegistrationsByNetworkParamsWithTimeout creates a new GetRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecofrienderV1RegistrationsByNetworkParamsWithTimeout(timeout time.Duration) *GetRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &GetRecofrienderV1RegistrationsByNetworkParams{

		timeout: timeout,
	}
}

// NewGetRecofrienderV1RegistrationsByNetworkParamsWithContext creates a new GetRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecofrienderV1RegistrationsByNetworkParamsWithContext(ctx context.Context) *GetRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &GetRecofrienderV1RegistrationsByNetworkParams{

		Context: ctx,
	}
}

// NewGetRecofrienderV1RegistrationsByNetworkParamsWithHTTPClient creates a new GetRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecofrienderV1RegistrationsByNetworkParamsWithHTTPClient(client *http.Client) *GetRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &GetRecofrienderV1RegistrationsByNetworkParams{
		HTTPClient: client,
	}
}

/*GetRecofrienderV1RegistrationsByNetworkParams contains all the parameters to send to the API endpoint
for the get recofriender v1 registrations by network operation typically these are written to a http.Request
*/
type GetRecofrienderV1RegistrationsByNetworkParams struct {

	/*Network*/
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) WithTimeout(timeout time.Duration) *GetRecofrienderV1RegistrationsByNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) WithContext(ctx context.Context) *GetRecofrienderV1RegistrationsByNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) WithHTTPClient(client *http.Client) *GetRecofrienderV1RegistrationsByNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetwork adds the network to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) WithNetwork(network string) *GetRecofrienderV1RegistrationsByNetworkParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the get recofriender v1 registrations by network params
func (o *GetRecofrienderV1RegistrationsByNetworkParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecofrienderV1RegistrationsByNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param network
	if err := r.SetPathParam("network", o.Network); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
