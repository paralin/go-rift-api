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

// NewGetLolPlayerBehaviorV1ReformCardParams creates a new GetLolPlayerBehaviorV1ReformCardParams object
// with the default values initialized.
func NewGetLolPlayerBehaviorV1ReformCardParams() *GetLolPlayerBehaviorV1ReformCardParams {

	return &GetLolPlayerBehaviorV1ReformCardParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPlayerBehaviorV1ReformCardParamsWithTimeout creates a new GetLolPlayerBehaviorV1ReformCardParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPlayerBehaviorV1ReformCardParamsWithTimeout(timeout time.Duration) *GetLolPlayerBehaviorV1ReformCardParams {

	return &GetLolPlayerBehaviorV1ReformCardParams{

		timeout: timeout,
	}
}

// NewGetLolPlayerBehaviorV1ReformCardParamsWithContext creates a new GetLolPlayerBehaviorV1ReformCardParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPlayerBehaviorV1ReformCardParamsWithContext(ctx context.Context) *GetLolPlayerBehaviorV1ReformCardParams {

	return &GetLolPlayerBehaviorV1ReformCardParams{

		Context: ctx,
	}
}

// NewGetLolPlayerBehaviorV1ReformCardParamsWithHTTPClient creates a new GetLolPlayerBehaviorV1ReformCardParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPlayerBehaviorV1ReformCardParamsWithHTTPClient(client *http.Client) *GetLolPlayerBehaviorV1ReformCardParams {

	return &GetLolPlayerBehaviorV1ReformCardParams{
		HTTPClient: client,
	}
}

/*GetLolPlayerBehaviorV1ReformCardParams contains all the parameters to send to the API endpoint
for the get lol player behavior v1 reform card operation typically these are written to a http.Request
*/
type GetLolPlayerBehaviorV1ReformCardParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol player behavior v1 reform card params
func (o *GetLolPlayerBehaviorV1ReformCardParams) WithTimeout(timeout time.Duration) *GetLolPlayerBehaviorV1ReformCardParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol player behavior v1 reform card params
func (o *GetLolPlayerBehaviorV1ReformCardParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol player behavior v1 reform card params
func (o *GetLolPlayerBehaviorV1ReformCardParams) WithContext(ctx context.Context) *GetLolPlayerBehaviorV1ReformCardParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol player behavior v1 reform card params
func (o *GetLolPlayerBehaviorV1ReformCardParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol player behavior v1 reform card params
func (o *GetLolPlayerBehaviorV1ReformCardParams) WithHTTPClient(client *http.Client) *GetLolPlayerBehaviorV1ReformCardParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol player behavior v1 reform card params
func (o *GetLolPlayerBehaviorV1ReformCardParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPlayerBehaviorV1ReformCardParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
