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

// NewPostLolAccountVerificationV1InvalidateParams creates a new PostLolAccountVerificationV1InvalidateParams object
// with the default values initialized.
func NewPostLolAccountVerificationV1InvalidateParams() *PostLolAccountVerificationV1InvalidateParams {

	return &PostLolAccountVerificationV1InvalidateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolAccountVerificationV1InvalidateParamsWithTimeout creates a new PostLolAccountVerificationV1InvalidateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolAccountVerificationV1InvalidateParamsWithTimeout(timeout time.Duration) *PostLolAccountVerificationV1InvalidateParams {

	return &PostLolAccountVerificationV1InvalidateParams{

		timeout: timeout,
	}
}

// NewPostLolAccountVerificationV1InvalidateParamsWithContext creates a new PostLolAccountVerificationV1InvalidateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolAccountVerificationV1InvalidateParamsWithContext(ctx context.Context) *PostLolAccountVerificationV1InvalidateParams {

	return &PostLolAccountVerificationV1InvalidateParams{

		Context: ctx,
	}
}

// NewPostLolAccountVerificationV1InvalidateParamsWithHTTPClient creates a new PostLolAccountVerificationV1InvalidateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolAccountVerificationV1InvalidateParamsWithHTTPClient(client *http.Client) *PostLolAccountVerificationV1InvalidateParams {

	return &PostLolAccountVerificationV1InvalidateParams{
		HTTPClient: client,
	}
}

/*PostLolAccountVerificationV1InvalidateParams contains all the parameters to send to the API endpoint
for the post lol account verification v1 invalidate operation typically these are written to a http.Request
*/
type PostLolAccountVerificationV1InvalidateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol account verification v1 invalidate params
func (o *PostLolAccountVerificationV1InvalidateParams) WithTimeout(timeout time.Duration) *PostLolAccountVerificationV1InvalidateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol account verification v1 invalidate params
func (o *PostLolAccountVerificationV1InvalidateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol account verification v1 invalidate params
func (o *PostLolAccountVerificationV1InvalidateParams) WithContext(ctx context.Context) *PostLolAccountVerificationV1InvalidateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol account verification v1 invalidate params
func (o *PostLolAccountVerificationV1InvalidateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol account verification v1 invalidate params
func (o *PostLolAccountVerificationV1InvalidateParams) WithHTTPClient(client *http.Client) *PostLolAccountVerificationV1InvalidateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol account verification v1 invalidate params
func (o *PostLolAccountVerificationV1InvalidateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolAccountVerificationV1InvalidateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
