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

// NewPostLolClashV1VoiceParams creates a new PostLolClashV1VoiceParams object
// with the default values initialized.
func NewPostLolClashV1VoiceParams() *PostLolClashV1VoiceParams {

	return &PostLolClashV1VoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1VoiceParamsWithTimeout creates a new PostLolClashV1VoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1VoiceParamsWithTimeout(timeout time.Duration) *PostLolClashV1VoiceParams {

	return &PostLolClashV1VoiceParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1VoiceParamsWithContext creates a new PostLolClashV1VoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1VoiceParamsWithContext(ctx context.Context) *PostLolClashV1VoiceParams {

	return &PostLolClashV1VoiceParams{

		Context: ctx,
	}
}

// NewPostLolClashV1VoiceParamsWithHTTPClient creates a new PostLolClashV1VoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1VoiceParamsWithHTTPClient(client *http.Client) *PostLolClashV1VoiceParams {

	return &PostLolClashV1VoiceParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1VoiceParams contains all the parameters to send to the API endpoint
for the post lol clash v1 voice operation typically these are written to a http.Request
*/
type PostLolClashV1VoiceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 voice params
func (o *PostLolClashV1VoiceParams) WithTimeout(timeout time.Duration) *PostLolClashV1VoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 voice params
func (o *PostLolClashV1VoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 voice params
func (o *PostLolClashV1VoiceParams) WithContext(ctx context.Context) *PostLolClashV1VoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 voice params
func (o *PostLolClashV1VoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 voice params
func (o *PostLolClashV1VoiceParams) WithHTTPClient(client *http.Client) *PostLolClashV1VoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 voice params
func (o *PostLolClashV1VoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1VoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}