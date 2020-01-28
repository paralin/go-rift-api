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

// NewGetLolClashV1NotificationsParams creates a new GetLolClashV1NotificationsParams object
// with the default values initialized.
func NewGetLolClashV1NotificationsParams() *GetLolClashV1NotificationsParams {

	return &GetLolClashV1NotificationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1NotificationsParamsWithTimeout creates a new GetLolClashV1NotificationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1NotificationsParamsWithTimeout(timeout time.Duration) *GetLolClashV1NotificationsParams {

	return &GetLolClashV1NotificationsParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1NotificationsParamsWithContext creates a new GetLolClashV1NotificationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1NotificationsParamsWithContext(ctx context.Context) *GetLolClashV1NotificationsParams {

	return &GetLolClashV1NotificationsParams{

		Context: ctx,
	}
}

// NewGetLolClashV1NotificationsParamsWithHTTPClient creates a new GetLolClashV1NotificationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1NotificationsParamsWithHTTPClient(client *http.Client) *GetLolClashV1NotificationsParams {

	return &GetLolClashV1NotificationsParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1NotificationsParams contains all the parameters to send to the API endpoint
for the get lol clash v1 notifications operation typically these are written to a http.Request
*/
type GetLolClashV1NotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 notifications params
func (o *GetLolClashV1NotificationsParams) WithTimeout(timeout time.Duration) *GetLolClashV1NotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 notifications params
func (o *GetLolClashV1NotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 notifications params
func (o *GetLolClashV1NotificationsParams) WithContext(ctx context.Context) *GetLolClashV1NotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 notifications params
func (o *GetLolClashV1NotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 notifications params
func (o *GetLolClashV1NotificationsParams) WithHTTPClient(client *http.Client) *GetLolClashV1NotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 notifications params
func (o *GetLolClashV1NotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1NotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}