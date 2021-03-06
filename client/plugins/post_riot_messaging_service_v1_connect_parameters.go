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

// NewPostRiotMessagingServiceV1ConnectParams creates a new PostRiotMessagingServiceV1ConnectParams object
// with the default values initialized.
func NewPostRiotMessagingServiceV1ConnectParams() *PostRiotMessagingServiceV1ConnectParams {
	var ()
	return &PostRiotMessagingServiceV1ConnectParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRiotMessagingServiceV1ConnectParamsWithTimeout creates a new PostRiotMessagingServiceV1ConnectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRiotMessagingServiceV1ConnectParamsWithTimeout(timeout time.Duration) *PostRiotMessagingServiceV1ConnectParams {
	var ()
	return &PostRiotMessagingServiceV1ConnectParams{

		timeout: timeout,
	}
}

// NewPostRiotMessagingServiceV1ConnectParamsWithContext creates a new PostRiotMessagingServiceV1ConnectParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRiotMessagingServiceV1ConnectParamsWithContext(ctx context.Context) *PostRiotMessagingServiceV1ConnectParams {
	var ()
	return &PostRiotMessagingServiceV1ConnectParams{

		Context: ctx,
	}
}

// NewPostRiotMessagingServiceV1ConnectParamsWithHTTPClient creates a new PostRiotMessagingServiceV1ConnectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRiotMessagingServiceV1ConnectParamsWithHTTPClient(client *http.Client) *PostRiotMessagingServiceV1ConnectParams {
	var ()
	return &PostRiotMessagingServiceV1ConnectParams{
		HTTPClient: client,
	}
}

/*PostRiotMessagingServiceV1ConnectParams contains all the parameters to send to the API endpoint
for the post riot messaging service v1 connect operation typically these are written to a http.Request
*/
type PostRiotMessagingServiceV1ConnectParams struct {

	/*IDToken*/
	IDToken string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) WithTimeout(timeout time.Duration) *PostRiotMessagingServiceV1ConnectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) WithContext(ctx context.Context) *PostRiotMessagingServiceV1ConnectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) WithHTTPClient(client *http.Client) *PostRiotMessagingServiceV1ConnectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIDToken adds the iDToken to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) WithIDToken(iDToken string) *PostRiotMessagingServiceV1ConnectParams {
	o.SetIDToken(iDToken)
	return o
}

// SetIDToken adds the idToken to the post riot messaging service v1 connect params
func (o *PostRiotMessagingServiceV1ConnectParams) SetIDToken(iDToken string) {
	o.IDToken = iDToken
}

// WriteToRequest writes these params to a swagger request
func (o *PostRiotMessagingServiceV1ConnectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.IDToken); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
