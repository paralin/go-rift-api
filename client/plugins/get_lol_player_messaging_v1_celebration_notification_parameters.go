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

// NewGetLolPlayerMessagingV1CelebrationNotificationParams creates a new GetLolPlayerMessagingV1CelebrationNotificationParams object
// with the default values initialized.
func NewGetLolPlayerMessagingV1CelebrationNotificationParams() *GetLolPlayerMessagingV1CelebrationNotificationParams {

	return &GetLolPlayerMessagingV1CelebrationNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPlayerMessagingV1CelebrationNotificationParamsWithTimeout creates a new GetLolPlayerMessagingV1CelebrationNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPlayerMessagingV1CelebrationNotificationParamsWithTimeout(timeout time.Duration) *GetLolPlayerMessagingV1CelebrationNotificationParams {

	return &GetLolPlayerMessagingV1CelebrationNotificationParams{

		timeout: timeout,
	}
}

// NewGetLolPlayerMessagingV1CelebrationNotificationParamsWithContext creates a new GetLolPlayerMessagingV1CelebrationNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPlayerMessagingV1CelebrationNotificationParamsWithContext(ctx context.Context) *GetLolPlayerMessagingV1CelebrationNotificationParams {

	return &GetLolPlayerMessagingV1CelebrationNotificationParams{

		Context: ctx,
	}
}

// NewGetLolPlayerMessagingV1CelebrationNotificationParamsWithHTTPClient creates a new GetLolPlayerMessagingV1CelebrationNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPlayerMessagingV1CelebrationNotificationParamsWithHTTPClient(client *http.Client) *GetLolPlayerMessagingV1CelebrationNotificationParams {

	return &GetLolPlayerMessagingV1CelebrationNotificationParams{
		HTTPClient: client,
	}
}

/*GetLolPlayerMessagingV1CelebrationNotificationParams contains all the parameters to send to the API endpoint
for the get lol player messaging v1 celebration notification operation typically these are written to a http.Request
*/
type GetLolPlayerMessagingV1CelebrationNotificationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol player messaging v1 celebration notification params
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) WithTimeout(timeout time.Duration) *GetLolPlayerMessagingV1CelebrationNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol player messaging v1 celebration notification params
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol player messaging v1 celebration notification params
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) WithContext(ctx context.Context) *GetLolPlayerMessagingV1CelebrationNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol player messaging v1 celebration notification params
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol player messaging v1 celebration notification params
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) WithHTTPClient(client *http.Client) *GetLolPlayerMessagingV1CelebrationNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol player messaging v1 celebration notification params
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPlayerMessagingV1CelebrationNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
