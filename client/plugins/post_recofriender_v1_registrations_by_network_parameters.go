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

// NewPostRecofrienderV1RegistrationsByNetworkParams creates a new PostRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized.
func NewPostRecofrienderV1RegistrationsByNetworkParams() *PostRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &PostRecofrienderV1RegistrationsByNetworkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRecofrienderV1RegistrationsByNetworkParamsWithTimeout creates a new PostRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRecofrienderV1RegistrationsByNetworkParamsWithTimeout(timeout time.Duration) *PostRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &PostRecofrienderV1RegistrationsByNetworkParams{

		timeout: timeout,
	}
}

// NewPostRecofrienderV1RegistrationsByNetworkParamsWithContext creates a new PostRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRecofrienderV1RegistrationsByNetworkParamsWithContext(ctx context.Context) *PostRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &PostRecofrienderV1RegistrationsByNetworkParams{

		Context: ctx,
	}
}

// NewPostRecofrienderV1RegistrationsByNetworkParamsWithHTTPClient creates a new PostRecofrienderV1RegistrationsByNetworkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRecofrienderV1RegistrationsByNetworkParamsWithHTTPClient(client *http.Client) *PostRecofrienderV1RegistrationsByNetworkParams {
	var ()
	return &PostRecofrienderV1RegistrationsByNetworkParams{
		HTTPClient: client,
	}
}

/*PostRecofrienderV1RegistrationsByNetworkParams contains all the parameters to send to the API endpoint
for the post recofriender v1 registrations by network operation typically these are written to a http.Request
*/
type PostRecofrienderV1RegistrationsByNetworkParams struct {

	/*Network*/
	Network string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) WithTimeout(timeout time.Duration) *PostRecofrienderV1RegistrationsByNetworkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) WithContext(ctx context.Context) *PostRecofrienderV1RegistrationsByNetworkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) WithHTTPClient(client *http.Client) *PostRecofrienderV1RegistrationsByNetworkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNetwork adds the network to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) WithNetwork(network string) *PostRecofrienderV1RegistrationsByNetworkParams {
	o.SetNetwork(network)
	return o
}

// SetNetwork adds the network to the post recofriender v1 registrations by network params
func (o *PostRecofrienderV1RegistrationsByNetworkParams) SetNetwork(network string) {
	o.Network = network
}

// WriteToRequest writes these params to a swagger request
func (o *PostRecofrienderV1RegistrationsByNetworkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
