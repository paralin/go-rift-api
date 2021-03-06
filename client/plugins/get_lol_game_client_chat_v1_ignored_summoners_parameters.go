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

// NewGetLolGameClientChatV1IgnoredSummonersParams creates a new GetLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized.
func NewGetLolGameClientChatV1IgnoredSummonersParams() *GetLolGameClientChatV1IgnoredSummonersParams {

	return &GetLolGameClientChatV1IgnoredSummonersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolGameClientChatV1IgnoredSummonersParamsWithTimeout creates a new GetLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolGameClientChatV1IgnoredSummonersParamsWithTimeout(timeout time.Duration) *GetLolGameClientChatV1IgnoredSummonersParams {

	return &GetLolGameClientChatV1IgnoredSummonersParams{

		timeout: timeout,
	}
}

// NewGetLolGameClientChatV1IgnoredSummonersParamsWithContext creates a new GetLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolGameClientChatV1IgnoredSummonersParamsWithContext(ctx context.Context) *GetLolGameClientChatV1IgnoredSummonersParams {

	return &GetLolGameClientChatV1IgnoredSummonersParams{

		Context: ctx,
	}
}

// NewGetLolGameClientChatV1IgnoredSummonersParamsWithHTTPClient creates a new GetLolGameClientChatV1IgnoredSummonersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolGameClientChatV1IgnoredSummonersParamsWithHTTPClient(client *http.Client) *GetLolGameClientChatV1IgnoredSummonersParams {

	return &GetLolGameClientChatV1IgnoredSummonersParams{
		HTTPClient: client,
	}
}

/*GetLolGameClientChatV1IgnoredSummonersParams contains all the parameters to send to the API endpoint
for the get lol game client chat v1 ignored summoners operation typically these are written to a http.Request
*/
type GetLolGameClientChatV1IgnoredSummonersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol game client chat v1 ignored summoners params
func (o *GetLolGameClientChatV1IgnoredSummonersParams) WithTimeout(timeout time.Duration) *GetLolGameClientChatV1IgnoredSummonersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol game client chat v1 ignored summoners params
func (o *GetLolGameClientChatV1IgnoredSummonersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol game client chat v1 ignored summoners params
func (o *GetLolGameClientChatV1IgnoredSummonersParams) WithContext(ctx context.Context) *GetLolGameClientChatV1IgnoredSummonersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol game client chat v1 ignored summoners params
func (o *GetLolGameClientChatV1IgnoredSummonersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol game client chat v1 ignored summoners params
func (o *GetLolGameClientChatV1IgnoredSummonersParams) WithHTTPClient(client *http.Client) *GetLolGameClientChatV1IgnoredSummonersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol game client chat v1 ignored summoners params
func (o *GetLolGameClientChatV1IgnoredSummonersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolGameClientChatV1IgnoredSummonersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
