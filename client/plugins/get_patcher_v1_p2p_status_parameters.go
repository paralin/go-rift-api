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

// NewGetPatcherV1P2pStatusParams creates a new GetPatcherV1P2pStatusParams object
// with the default values initialized.
func NewGetPatcherV1P2pStatusParams() *GetPatcherV1P2pStatusParams {

	return &GetPatcherV1P2pStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatcherV1P2pStatusParamsWithTimeout creates a new GetPatcherV1P2pStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatcherV1P2pStatusParamsWithTimeout(timeout time.Duration) *GetPatcherV1P2pStatusParams {

	return &GetPatcherV1P2pStatusParams{

		timeout: timeout,
	}
}

// NewGetPatcherV1P2pStatusParamsWithContext creates a new GetPatcherV1P2pStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatcherV1P2pStatusParamsWithContext(ctx context.Context) *GetPatcherV1P2pStatusParams {

	return &GetPatcherV1P2pStatusParams{

		Context: ctx,
	}
}

// NewGetPatcherV1P2pStatusParamsWithHTTPClient creates a new GetPatcherV1P2pStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatcherV1P2pStatusParamsWithHTTPClient(client *http.Client) *GetPatcherV1P2pStatusParams {

	return &GetPatcherV1P2pStatusParams{
		HTTPClient: client,
	}
}

/*GetPatcherV1P2pStatusParams contains all the parameters to send to the API endpoint
for the get patcher v1 p2p status operation typically these are written to a http.Request
*/
type GetPatcherV1P2pStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patcher v1 p2p status params
func (o *GetPatcherV1P2pStatusParams) WithTimeout(timeout time.Duration) *GetPatcherV1P2pStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patcher v1 p2p status params
func (o *GetPatcherV1P2pStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patcher v1 p2p status params
func (o *GetPatcherV1P2pStatusParams) WithContext(ctx context.Context) *GetPatcherV1P2pStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patcher v1 p2p status params
func (o *GetPatcherV1P2pStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patcher v1 p2p status params
func (o *GetPatcherV1P2pStatusParams) WithHTTPClient(client *http.Client) *GetPatcherV1P2pStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patcher v1 p2p status params
func (o *GetPatcherV1P2pStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatcherV1P2pStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}