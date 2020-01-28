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

// NewPostLolClashV1EventsParams creates a new PostLolClashV1EventsParams object
// with the default values initialized.
func NewPostLolClashV1EventsParams() *PostLolClashV1EventsParams {
	var ()
	return &PostLolClashV1EventsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolClashV1EventsParamsWithTimeout creates a new PostLolClashV1EventsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolClashV1EventsParamsWithTimeout(timeout time.Duration) *PostLolClashV1EventsParams {
	var ()
	return &PostLolClashV1EventsParams{

		timeout: timeout,
	}
}

// NewPostLolClashV1EventsParamsWithContext creates a new PostLolClashV1EventsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolClashV1EventsParamsWithContext(ctx context.Context) *PostLolClashV1EventsParams {
	var ()
	return &PostLolClashV1EventsParams{

		Context: ctx,
	}
}

// NewPostLolClashV1EventsParamsWithHTTPClient creates a new PostLolClashV1EventsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolClashV1EventsParamsWithHTTPClient(client *http.Client) *PostLolClashV1EventsParams {
	var ()
	return &PostLolClashV1EventsParams{
		HTTPClient: client,
	}
}

/*PostLolClashV1EventsParams contains all the parameters to send to the API endpoint
for the post lol clash v1 events operation typically these are written to a http.Request
*/
type PostLolClashV1EventsParams struct {

	/*Uuids*/
	Uuids []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) WithTimeout(timeout time.Duration) *PostLolClashV1EventsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) WithContext(ctx context.Context) *PostLolClashV1EventsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) WithHTTPClient(client *http.Client) *PostLolClashV1EventsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUuids adds the uuids to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) WithUuids(uuids []string) *PostLolClashV1EventsParams {
	o.SetUuids(uuids)
	return o
}

// SetUuids adds the uuids to the post lol clash v1 events params
func (o *PostLolClashV1EventsParams) SetUuids(uuids []string) {
	o.Uuids = uuids
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolClashV1EventsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Uuids != nil {
		if err := r.SetBodyParam(o.Uuids); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
