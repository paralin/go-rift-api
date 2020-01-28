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

// NewPostLolHighlightsV1HighlightsParams creates a new PostLolHighlightsV1HighlightsParams object
// with the default values initialized.
func NewPostLolHighlightsV1HighlightsParams() *PostLolHighlightsV1HighlightsParams {

	return &PostLolHighlightsV1HighlightsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolHighlightsV1HighlightsParamsWithTimeout creates a new PostLolHighlightsV1HighlightsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolHighlightsV1HighlightsParamsWithTimeout(timeout time.Duration) *PostLolHighlightsV1HighlightsParams {

	return &PostLolHighlightsV1HighlightsParams{

		timeout: timeout,
	}
}

// NewPostLolHighlightsV1HighlightsParamsWithContext creates a new PostLolHighlightsV1HighlightsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolHighlightsV1HighlightsParamsWithContext(ctx context.Context) *PostLolHighlightsV1HighlightsParams {

	return &PostLolHighlightsV1HighlightsParams{

		Context: ctx,
	}
}

// NewPostLolHighlightsV1HighlightsParamsWithHTTPClient creates a new PostLolHighlightsV1HighlightsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolHighlightsV1HighlightsParamsWithHTTPClient(client *http.Client) *PostLolHighlightsV1HighlightsParams {

	return &PostLolHighlightsV1HighlightsParams{
		HTTPClient: client,
	}
}

/*PostLolHighlightsV1HighlightsParams contains all the parameters to send to the API endpoint
for the post lol highlights v1 highlights operation typically these are written to a http.Request
*/
type PostLolHighlightsV1HighlightsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol highlights v1 highlights params
func (o *PostLolHighlightsV1HighlightsParams) WithTimeout(timeout time.Duration) *PostLolHighlightsV1HighlightsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol highlights v1 highlights params
func (o *PostLolHighlightsV1HighlightsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol highlights v1 highlights params
func (o *PostLolHighlightsV1HighlightsParams) WithContext(ctx context.Context) *PostLolHighlightsV1HighlightsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol highlights v1 highlights params
func (o *PostLolHighlightsV1HighlightsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol highlights v1 highlights params
func (o *PostLolHighlightsV1HighlightsParams) WithHTTPClient(client *http.Client) *PostLolHighlightsV1HighlightsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol highlights v1 highlights params
func (o *PostLolHighlightsV1HighlightsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolHighlightsV1HighlightsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}