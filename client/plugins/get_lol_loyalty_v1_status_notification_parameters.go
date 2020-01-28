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

// NewGetLolLoyaltyV1StatusNotificationParams creates a new GetLolLoyaltyV1StatusNotificationParams object
// with the default values initialized.
func NewGetLolLoyaltyV1StatusNotificationParams() *GetLolLoyaltyV1StatusNotificationParams {

	return &GetLolLoyaltyV1StatusNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLoyaltyV1StatusNotificationParamsWithTimeout creates a new GetLolLoyaltyV1StatusNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLoyaltyV1StatusNotificationParamsWithTimeout(timeout time.Duration) *GetLolLoyaltyV1StatusNotificationParams {

	return &GetLolLoyaltyV1StatusNotificationParams{

		timeout: timeout,
	}
}

// NewGetLolLoyaltyV1StatusNotificationParamsWithContext creates a new GetLolLoyaltyV1StatusNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLoyaltyV1StatusNotificationParamsWithContext(ctx context.Context) *GetLolLoyaltyV1StatusNotificationParams {

	return &GetLolLoyaltyV1StatusNotificationParams{

		Context: ctx,
	}
}

// NewGetLolLoyaltyV1StatusNotificationParamsWithHTTPClient creates a new GetLolLoyaltyV1StatusNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLoyaltyV1StatusNotificationParamsWithHTTPClient(client *http.Client) *GetLolLoyaltyV1StatusNotificationParams {

	return &GetLolLoyaltyV1StatusNotificationParams{
		HTTPClient: client,
	}
}

/*GetLolLoyaltyV1StatusNotificationParams contains all the parameters to send to the API endpoint
for the get lol loyalty v1 status notification operation typically these are written to a http.Request
*/
type GetLolLoyaltyV1StatusNotificationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol loyalty v1 status notification params
func (o *GetLolLoyaltyV1StatusNotificationParams) WithTimeout(timeout time.Duration) *GetLolLoyaltyV1StatusNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol loyalty v1 status notification params
func (o *GetLolLoyaltyV1StatusNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol loyalty v1 status notification params
func (o *GetLolLoyaltyV1StatusNotificationParams) WithContext(ctx context.Context) *GetLolLoyaltyV1StatusNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol loyalty v1 status notification params
func (o *GetLolLoyaltyV1StatusNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol loyalty v1 status notification params
func (o *GetLolLoyaltyV1StatusNotificationParams) WithHTTPClient(client *http.Client) *GetLolLoyaltyV1StatusNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol loyalty v1 status notification params
func (o *GetLolLoyaltyV1StatusNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLoyaltyV1StatusNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
