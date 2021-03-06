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

// NewGetLolEsportStreamNotificationsV1StreamURLParams creates a new GetLolEsportStreamNotificationsV1StreamURLParams object
// with the default values initialized.
func NewGetLolEsportStreamNotificationsV1StreamURLParams() *GetLolEsportStreamNotificationsV1StreamURLParams {

	return &GetLolEsportStreamNotificationsV1StreamURLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolEsportStreamNotificationsV1StreamURLParamsWithTimeout creates a new GetLolEsportStreamNotificationsV1StreamURLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolEsportStreamNotificationsV1StreamURLParamsWithTimeout(timeout time.Duration) *GetLolEsportStreamNotificationsV1StreamURLParams {

	return &GetLolEsportStreamNotificationsV1StreamURLParams{

		timeout: timeout,
	}
}

// NewGetLolEsportStreamNotificationsV1StreamURLParamsWithContext creates a new GetLolEsportStreamNotificationsV1StreamURLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolEsportStreamNotificationsV1StreamURLParamsWithContext(ctx context.Context) *GetLolEsportStreamNotificationsV1StreamURLParams {

	return &GetLolEsportStreamNotificationsV1StreamURLParams{

		Context: ctx,
	}
}

// NewGetLolEsportStreamNotificationsV1StreamURLParamsWithHTTPClient creates a new GetLolEsportStreamNotificationsV1StreamURLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolEsportStreamNotificationsV1StreamURLParamsWithHTTPClient(client *http.Client) *GetLolEsportStreamNotificationsV1StreamURLParams {

	return &GetLolEsportStreamNotificationsV1StreamURLParams{
		HTTPClient: client,
	}
}

/*GetLolEsportStreamNotificationsV1StreamURLParams contains all the parameters to send to the API endpoint
for the get lol esport stream notifications v1 stream Url operation typically these are written to a http.Request
*/
type GetLolEsportStreamNotificationsV1StreamURLParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol esport stream notifications v1 stream Url params
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) WithTimeout(timeout time.Duration) *GetLolEsportStreamNotificationsV1StreamURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol esport stream notifications v1 stream Url params
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol esport stream notifications v1 stream Url params
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) WithContext(ctx context.Context) *GetLolEsportStreamNotificationsV1StreamURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol esport stream notifications v1 stream Url params
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol esport stream notifications v1 stream Url params
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) WithHTTPClient(client *http.Client) *GetLolEsportStreamNotificationsV1StreamURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol esport stream notifications v1 stream Url params
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolEsportStreamNotificationsV1StreamURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
