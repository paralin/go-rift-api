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

// NewGetLolPlayerMessagingV1NotificationParams creates a new GetLolPlayerMessagingV1NotificationParams object
// with the default values initialized.
func NewGetLolPlayerMessagingV1NotificationParams() *GetLolPlayerMessagingV1NotificationParams {

	return &GetLolPlayerMessagingV1NotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPlayerMessagingV1NotificationParamsWithTimeout creates a new GetLolPlayerMessagingV1NotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPlayerMessagingV1NotificationParamsWithTimeout(timeout time.Duration) *GetLolPlayerMessagingV1NotificationParams {

	return &GetLolPlayerMessagingV1NotificationParams{

		timeout: timeout,
	}
}

// NewGetLolPlayerMessagingV1NotificationParamsWithContext creates a new GetLolPlayerMessagingV1NotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPlayerMessagingV1NotificationParamsWithContext(ctx context.Context) *GetLolPlayerMessagingV1NotificationParams {

	return &GetLolPlayerMessagingV1NotificationParams{

		Context: ctx,
	}
}

// NewGetLolPlayerMessagingV1NotificationParamsWithHTTPClient creates a new GetLolPlayerMessagingV1NotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPlayerMessagingV1NotificationParamsWithHTTPClient(client *http.Client) *GetLolPlayerMessagingV1NotificationParams {

	return &GetLolPlayerMessagingV1NotificationParams{
		HTTPClient: client,
	}
}

/*GetLolPlayerMessagingV1NotificationParams contains all the parameters to send to the API endpoint
for the get lol player messaging v1 notification operation typically these are written to a http.Request
*/
type GetLolPlayerMessagingV1NotificationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol player messaging v1 notification params
func (o *GetLolPlayerMessagingV1NotificationParams) WithTimeout(timeout time.Duration) *GetLolPlayerMessagingV1NotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol player messaging v1 notification params
func (o *GetLolPlayerMessagingV1NotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol player messaging v1 notification params
func (o *GetLolPlayerMessagingV1NotificationParams) WithContext(ctx context.Context) *GetLolPlayerMessagingV1NotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol player messaging v1 notification params
func (o *GetLolPlayerMessagingV1NotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol player messaging v1 notification params
func (o *GetLolPlayerMessagingV1NotificationParams) WithHTTPClient(client *http.Client) *GetLolPlayerMessagingV1NotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol player messaging v1 notification params
func (o *GetLolPlayerMessagingV1NotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPlayerMessagingV1NotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
