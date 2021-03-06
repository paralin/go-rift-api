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

// NewGetLolPlayerBehaviorV1ConfigParams creates a new GetLolPlayerBehaviorV1ConfigParams object
// with the default values initialized.
func NewGetLolPlayerBehaviorV1ConfigParams() *GetLolPlayerBehaviorV1ConfigParams {

	return &GetLolPlayerBehaviorV1ConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPlayerBehaviorV1ConfigParamsWithTimeout creates a new GetLolPlayerBehaviorV1ConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPlayerBehaviorV1ConfigParamsWithTimeout(timeout time.Duration) *GetLolPlayerBehaviorV1ConfigParams {

	return &GetLolPlayerBehaviorV1ConfigParams{

		timeout: timeout,
	}
}

// NewGetLolPlayerBehaviorV1ConfigParamsWithContext creates a new GetLolPlayerBehaviorV1ConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPlayerBehaviorV1ConfigParamsWithContext(ctx context.Context) *GetLolPlayerBehaviorV1ConfigParams {

	return &GetLolPlayerBehaviorV1ConfigParams{

		Context: ctx,
	}
}

// NewGetLolPlayerBehaviorV1ConfigParamsWithHTTPClient creates a new GetLolPlayerBehaviorV1ConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPlayerBehaviorV1ConfigParamsWithHTTPClient(client *http.Client) *GetLolPlayerBehaviorV1ConfigParams {

	return &GetLolPlayerBehaviorV1ConfigParams{
		HTTPClient: client,
	}
}

/*GetLolPlayerBehaviorV1ConfigParams contains all the parameters to send to the API endpoint
for the get lol player behavior v1 config operation typically these are written to a http.Request
*/
type GetLolPlayerBehaviorV1ConfigParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol player behavior v1 config params
func (o *GetLolPlayerBehaviorV1ConfigParams) WithTimeout(timeout time.Duration) *GetLolPlayerBehaviorV1ConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol player behavior v1 config params
func (o *GetLolPlayerBehaviorV1ConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol player behavior v1 config params
func (o *GetLolPlayerBehaviorV1ConfigParams) WithContext(ctx context.Context) *GetLolPlayerBehaviorV1ConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol player behavior v1 config params
func (o *GetLolPlayerBehaviorV1ConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol player behavior v1 config params
func (o *GetLolPlayerBehaviorV1ConfigParams) WithHTTPClient(client *http.Client) *GetLolPlayerBehaviorV1ConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol player behavior v1 config params
func (o *GetLolPlayerBehaviorV1ConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPlayerBehaviorV1ConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
