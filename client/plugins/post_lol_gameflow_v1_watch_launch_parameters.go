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

// NewPostLolGameflowV1WatchLaunchParams creates a new PostLolGameflowV1WatchLaunchParams object
// with the default values initialized.
func NewPostLolGameflowV1WatchLaunchParams() *PostLolGameflowV1WatchLaunchParams {
	var ()
	return &PostLolGameflowV1WatchLaunchParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolGameflowV1WatchLaunchParamsWithTimeout creates a new PostLolGameflowV1WatchLaunchParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolGameflowV1WatchLaunchParamsWithTimeout(timeout time.Duration) *PostLolGameflowV1WatchLaunchParams {
	var ()
	return &PostLolGameflowV1WatchLaunchParams{

		timeout: timeout,
	}
}

// NewPostLolGameflowV1WatchLaunchParamsWithContext creates a new PostLolGameflowV1WatchLaunchParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolGameflowV1WatchLaunchParamsWithContext(ctx context.Context) *PostLolGameflowV1WatchLaunchParams {
	var ()
	return &PostLolGameflowV1WatchLaunchParams{

		Context: ctx,
	}
}

// NewPostLolGameflowV1WatchLaunchParamsWithHTTPClient creates a new PostLolGameflowV1WatchLaunchParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolGameflowV1WatchLaunchParamsWithHTTPClient(client *http.Client) *PostLolGameflowV1WatchLaunchParams {
	var ()
	return &PostLolGameflowV1WatchLaunchParams{
		HTTPClient: client,
	}
}

/*PostLolGameflowV1WatchLaunchParams contains all the parameters to send to the API endpoint
for the post lol gameflow v1 watch launch operation typically these are written to a http.Request
*/
type PostLolGameflowV1WatchLaunchParams struct {

	/*Args*/
	Args []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) WithTimeout(timeout time.Duration) *PostLolGameflowV1WatchLaunchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) WithContext(ctx context.Context) *PostLolGameflowV1WatchLaunchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) WithHTTPClient(client *http.Client) *PostLolGameflowV1WatchLaunchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithArgs adds the args to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) WithArgs(args []string) *PostLolGameflowV1WatchLaunchParams {
	o.SetArgs(args)
	return o
}

// SetArgs adds the args to the post lol gameflow v1 watch launch params
func (o *PostLolGameflowV1WatchLaunchParams) SetArgs(args []string) {
	o.Args = args
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolGameflowV1WatchLaunchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Args != nil {
		if err := r.SetBodyParam(o.Args); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
