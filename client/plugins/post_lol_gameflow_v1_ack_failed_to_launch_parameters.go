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

// NewPostLolGameflowV1AckFailedToLaunchParams creates a new PostLolGameflowV1AckFailedToLaunchParams object
// with the default values initialized.
func NewPostLolGameflowV1AckFailedToLaunchParams() *PostLolGameflowV1AckFailedToLaunchParams {

	return &PostLolGameflowV1AckFailedToLaunchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolGameflowV1AckFailedToLaunchParamsWithTimeout creates a new PostLolGameflowV1AckFailedToLaunchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolGameflowV1AckFailedToLaunchParamsWithTimeout(timeout time.Duration) *PostLolGameflowV1AckFailedToLaunchParams {

	return &PostLolGameflowV1AckFailedToLaunchParams{

		timeout: timeout,
	}
}

// NewPostLolGameflowV1AckFailedToLaunchParamsWithContext creates a new PostLolGameflowV1AckFailedToLaunchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolGameflowV1AckFailedToLaunchParamsWithContext(ctx context.Context) *PostLolGameflowV1AckFailedToLaunchParams {

	return &PostLolGameflowV1AckFailedToLaunchParams{

		Context: ctx,
	}
}

// NewPostLolGameflowV1AckFailedToLaunchParamsWithHTTPClient creates a new PostLolGameflowV1AckFailedToLaunchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolGameflowV1AckFailedToLaunchParamsWithHTTPClient(client *http.Client) *PostLolGameflowV1AckFailedToLaunchParams {

	return &PostLolGameflowV1AckFailedToLaunchParams{
		HTTPClient: client,
	}
}

/*PostLolGameflowV1AckFailedToLaunchParams contains all the parameters to send to the API endpoint
for the post lol gameflow v1 ack failed to launch operation typically these are written to a http.Request
*/
type PostLolGameflowV1AckFailedToLaunchParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol gameflow v1 ack failed to launch params
func (o *PostLolGameflowV1AckFailedToLaunchParams) WithTimeout(timeout time.Duration) *PostLolGameflowV1AckFailedToLaunchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol gameflow v1 ack failed to launch params
func (o *PostLolGameflowV1AckFailedToLaunchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol gameflow v1 ack failed to launch params
func (o *PostLolGameflowV1AckFailedToLaunchParams) WithContext(ctx context.Context) *PostLolGameflowV1AckFailedToLaunchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol gameflow v1 ack failed to launch params
func (o *PostLolGameflowV1AckFailedToLaunchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol gameflow v1 ack failed to launch params
func (o *PostLolGameflowV1AckFailedToLaunchParams) WithHTTPClient(client *http.Client) *PostLolGameflowV1AckFailedToLaunchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol gameflow v1 ack failed to launch params
func (o *PostLolGameflowV1AckFailedToLaunchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolGameflowV1AckFailedToLaunchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
