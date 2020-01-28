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

// NewGetLolReplaysV1RoflsPathParams creates a new GetLolReplaysV1RoflsPathParams object
// with the default values initialized.
func NewGetLolReplaysV1RoflsPathParams() *GetLolReplaysV1RoflsPathParams {

	return &GetLolReplaysV1RoflsPathParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolReplaysV1RoflsPathParamsWithTimeout creates a new GetLolReplaysV1RoflsPathParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolReplaysV1RoflsPathParamsWithTimeout(timeout time.Duration) *GetLolReplaysV1RoflsPathParams {

	return &GetLolReplaysV1RoflsPathParams{

		timeout: timeout,
	}
}

// NewGetLolReplaysV1RoflsPathParamsWithContext creates a new GetLolReplaysV1RoflsPathParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolReplaysV1RoflsPathParamsWithContext(ctx context.Context) *GetLolReplaysV1RoflsPathParams {

	return &GetLolReplaysV1RoflsPathParams{

		Context: ctx,
	}
}

// NewGetLolReplaysV1RoflsPathParamsWithHTTPClient creates a new GetLolReplaysV1RoflsPathParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolReplaysV1RoflsPathParamsWithHTTPClient(client *http.Client) *GetLolReplaysV1RoflsPathParams {

	return &GetLolReplaysV1RoflsPathParams{
		HTTPClient: client,
	}
}

/*GetLolReplaysV1RoflsPathParams contains all the parameters to send to the API endpoint
for the get lol replays v1 rofls path operation typically these are written to a http.Request
*/
type GetLolReplaysV1RoflsPathParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol replays v1 rofls path params
func (o *GetLolReplaysV1RoflsPathParams) WithTimeout(timeout time.Duration) *GetLolReplaysV1RoflsPathParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol replays v1 rofls path params
func (o *GetLolReplaysV1RoflsPathParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol replays v1 rofls path params
func (o *GetLolReplaysV1RoflsPathParams) WithContext(ctx context.Context) *GetLolReplaysV1RoflsPathParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol replays v1 rofls path params
func (o *GetLolReplaysV1RoflsPathParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol replays v1 rofls path params
func (o *GetLolReplaysV1RoflsPathParams) WithHTTPClient(client *http.Client) *GetLolReplaysV1RoflsPathParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol replays v1 rofls path params
func (o *GetLolReplaysV1RoflsPathParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolReplaysV1RoflsPathParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
