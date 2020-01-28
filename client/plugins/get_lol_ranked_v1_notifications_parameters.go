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

// NewGetLolRankedV1NotificationsParams creates a new GetLolRankedV1NotificationsParams object
// with the default values initialized.
func NewGetLolRankedV1NotificationsParams() *GetLolRankedV1NotificationsParams {

	return &GetLolRankedV1NotificationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolRankedV1NotificationsParamsWithTimeout creates a new GetLolRankedV1NotificationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolRankedV1NotificationsParamsWithTimeout(timeout time.Duration) *GetLolRankedV1NotificationsParams {

	return &GetLolRankedV1NotificationsParams{

		timeout: timeout,
	}
}

// NewGetLolRankedV1NotificationsParamsWithContext creates a new GetLolRankedV1NotificationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolRankedV1NotificationsParamsWithContext(ctx context.Context) *GetLolRankedV1NotificationsParams {

	return &GetLolRankedV1NotificationsParams{

		Context: ctx,
	}
}

// NewGetLolRankedV1NotificationsParamsWithHTTPClient creates a new GetLolRankedV1NotificationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolRankedV1NotificationsParamsWithHTTPClient(client *http.Client) *GetLolRankedV1NotificationsParams {

	return &GetLolRankedV1NotificationsParams{
		HTTPClient: client,
	}
}

/*GetLolRankedV1NotificationsParams contains all the parameters to send to the API endpoint
for the get lol ranked v1 notifications operation typically these are written to a http.Request
*/
type GetLolRankedV1NotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol ranked v1 notifications params
func (o *GetLolRankedV1NotificationsParams) WithTimeout(timeout time.Duration) *GetLolRankedV1NotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol ranked v1 notifications params
func (o *GetLolRankedV1NotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol ranked v1 notifications params
func (o *GetLolRankedV1NotificationsParams) WithContext(ctx context.Context) *GetLolRankedV1NotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol ranked v1 notifications params
func (o *GetLolRankedV1NotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol ranked v1 notifications params
func (o *GetLolRankedV1NotificationsParams) WithHTTPClient(client *http.Client) *GetLolRankedV1NotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol ranked v1 notifications params
func (o *GetLolRankedV1NotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolRankedV1NotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}