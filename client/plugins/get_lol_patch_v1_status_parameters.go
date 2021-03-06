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

// NewGetLolPatchV1StatusParams creates a new GetLolPatchV1StatusParams object
// with the default values initialized.
func NewGetLolPatchV1StatusParams() *GetLolPatchV1StatusParams {

	return &GetLolPatchV1StatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPatchV1StatusParamsWithTimeout creates a new GetLolPatchV1StatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPatchV1StatusParamsWithTimeout(timeout time.Duration) *GetLolPatchV1StatusParams {

	return &GetLolPatchV1StatusParams{

		timeout: timeout,
	}
}

// NewGetLolPatchV1StatusParamsWithContext creates a new GetLolPatchV1StatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPatchV1StatusParamsWithContext(ctx context.Context) *GetLolPatchV1StatusParams {

	return &GetLolPatchV1StatusParams{

		Context: ctx,
	}
}

// NewGetLolPatchV1StatusParamsWithHTTPClient creates a new GetLolPatchV1StatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPatchV1StatusParamsWithHTTPClient(client *http.Client) *GetLolPatchV1StatusParams {

	return &GetLolPatchV1StatusParams{
		HTTPClient: client,
	}
}

/*GetLolPatchV1StatusParams contains all the parameters to send to the API endpoint
for the get lol patch v1 status operation typically these are written to a http.Request
*/
type GetLolPatchV1StatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol patch v1 status params
func (o *GetLolPatchV1StatusParams) WithTimeout(timeout time.Duration) *GetLolPatchV1StatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol patch v1 status params
func (o *GetLolPatchV1StatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol patch v1 status params
func (o *GetLolPatchV1StatusParams) WithContext(ctx context.Context) *GetLolPatchV1StatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol patch v1 status params
func (o *GetLolPatchV1StatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol patch v1 status params
func (o *GetLolPatchV1StatusParams) WithHTTPClient(client *http.Client) *GetLolPatchV1StatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol patch v1 status params
func (o *GetLolPatchV1StatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPatchV1StatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
