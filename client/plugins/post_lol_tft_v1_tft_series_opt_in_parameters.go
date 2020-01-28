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

// NewPostLolTftV1TftSeriesOptInParams creates a new PostLolTftV1TftSeriesOptInParams object
// with the default values initialized.
func NewPostLolTftV1TftSeriesOptInParams() *PostLolTftV1TftSeriesOptInParams {

	return &PostLolTftV1TftSeriesOptInParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolTftV1TftSeriesOptInParamsWithTimeout creates a new PostLolTftV1TftSeriesOptInParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolTftV1TftSeriesOptInParamsWithTimeout(timeout time.Duration) *PostLolTftV1TftSeriesOptInParams {

	return &PostLolTftV1TftSeriesOptInParams{

		timeout: timeout,
	}
}

// NewPostLolTftV1TftSeriesOptInParamsWithContext creates a new PostLolTftV1TftSeriesOptInParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolTftV1TftSeriesOptInParamsWithContext(ctx context.Context) *PostLolTftV1TftSeriesOptInParams {

	return &PostLolTftV1TftSeriesOptInParams{

		Context: ctx,
	}
}

// NewPostLolTftV1TftSeriesOptInParamsWithHTTPClient creates a new PostLolTftV1TftSeriesOptInParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolTftV1TftSeriesOptInParamsWithHTTPClient(client *http.Client) *PostLolTftV1TftSeriesOptInParams {

	return &PostLolTftV1TftSeriesOptInParams{
		HTTPClient: client,
	}
}

/*PostLolTftV1TftSeriesOptInParams contains all the parameters to send to the API endpoint
for the post lol tft v1 tft series opt in operation typically these are written to a http.Request
*/
type PostLolTftV1TftSeriesOptInParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol tft v1 tft series opt in params
func (o *PostLolTftV1TftSeriesOptInParams) WithTimeout(timeout time.Duration) *PostLolTftV1TftSeriesOptInParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol tft v1 tft series opt in params
func (o *PostLolTftV1TftSeriesOptInParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol tft v1 tft series opt in params
func (o *PostLolTftV1TftSeriesOptInParams) WithContext(ctx context.Context) *PostLolTftV1TftSeriesOptInParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol tft v1 tft series opt in params
func (o *PostLolTftV1TftSeriesOptInParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol tft v1 tft series opt in params
func (o *PostLolTftV1TftSeriesOptInParams) WithHTTPClient(client *http.Client) *PostLolTftV1TftSeriesOptInParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol tft v1 tft series opt in params
func (o *PostLolTftV1TftSeriesOptInParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolTftV1TftSeriesOptInParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
