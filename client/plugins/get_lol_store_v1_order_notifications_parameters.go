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

// NewGetLolStoreV1OrderNotificationsParams creates a new GetLolStoreV1OrderNotificationsParams object
// with the default values initialized.
func NewGetLolStoreV1OrderNotificationsParams() *GetLolStoreV1OrderNotificationsParams {

	return &GetLolStoreV1OrderNotificationsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolStoreV1OrderNotificationsParamsWithTimeout creates a new GetLolStoreV1OrderNotificationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolStoreV1OrderNotificationsParamsWithTimeout(timeout time.Duration) *GetLolStoreV1OrderNotificationsParams {

	return &GetLolStoreV1OrderNotificationsParams{

		timeout: timeout,
	}
}

// NewGetLolStoreV1OrderNotificationsParamsWithContext creates a new GetLolStoreV1OrderNotificationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolStoreV1OrderNotificationsParamsWithContext(ctx context.Context) *GetLolStoreV1OrderNotificationsParams {

	return &GetLolStoreV1OrderNotificationsParams{

		Context: ctx,
	}
}

// NewGetLolStoreV1OrderNotificationsParamsWithHTTPClient creates a new GetLolStoreV1OrderNotificationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolStoreV1OrderNotificationsParamsWithHTTPClient(client *http.Client) *GetLolStoreV1OrderNotificationsParams {

	return &GetLolStoreV1OrderNotificationsParams{
		HTTPClient: client,
	}
}

/*GetLolStoreV1OrderNotificationsParams contains all the parameters to send to the API endpoint
for the get lol store v1 order notifications operation typically these are written to a http.Request
*/
type GetLolStoreV1OrderNotificationsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol store v1 order notifications params
func (o *GetLolStoreV1OrderNotificationsParams) WithTimeout(timeout time.Duration) *GetLolStoreV1OrderNotificationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol store v1 order notifications params
func (o *GetLolStoreV1OrderNotificationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol store v1 order notifications params
func (o *GetLolStoreV1OrderNotificationsParams) WithContext(ctx context.Context) *GetLolStoreV1OrderNotificationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol store v1 order notifications params
func (o *GetLolStoreV1OrderNotificationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol store v1 order notifications params
func (o *GetLolStoreV1OrderNotificationsParams) WithHTTPClient(client *http.Client) *GetLolStoreV1OrderNotificationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol store v1 order notifications params
func (o *GetLolStoreV1OrderNotificationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolStoreV1OrderNotificationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}