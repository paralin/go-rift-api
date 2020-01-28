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

// NewGetLolContentTargetingV1ProtectedFiltersParams creates a new GetLolContentTargetingV1ProtectedFiltersParams object
// with the default values initialized.
func NewGetLolContentTargetingV1ProtectedFiltersParams() *GetLolContentTargetingV1ProtectedFiltersParams {

	return &GetLolContentTargetingV1ProtectedFiltersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolContentTargetingV1ProtectedFiltersParamsWithTimeout creates a new GetLolContentTargetingV1ProtectedFiltersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolContentTargetingV1ProtectedFiltersParamsWithTimeout(timeout time.Duration) *GetLolContentTargetingV1ProtectedFiltersParams {

	return &GetLolContentTargetingV1ProtectedFiltersParams{

		timeout: timeout,
	}
}

// NewGetLolContentTargetingV1ProtectedFiltersParamsWithContext creates a new GetLolContentTargetingV1ProtectedFiltersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolContentTargetingV1ProtectedFiltersParamsWithContext(ctx context.Context) *GetLolContentTargetingV1ProtectedFiltersParams {

	return &GetLolContentTargetingV1ProtectedFiltersParams{

		Context: ctx,
	}
}

// NewGetLolContentTargetingV1ProtectedFiltersParamsWithHTTPClient creates a new GetLolContentTargetingV1ProtectedFiltersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolContentTargetingV1ProtectedFiltersParamsWithHTTPClient(client *http.Client) *GetLolContentTargetingV1ProtectedFiltersParams {

	return &GetLolContentTargetingV1ProtectedFiltersParams{
		HTTPClient: client,
	}
}

/*GetLolContentTargetingV1ProtectedFiltersParams contains all the parameters to send to the API endpoint
for the get lol content targeting v1 protected filters operation typically these are written to a http.Request
*/
type GetLolContentTargetingV1ProtectedFiltersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol content targeting v1 protected filters params
func (o *GetLolContentTargetingV1ProtectedFiltersParams) WithTimeout(timeout time.Duration) *GetLolContentTargetingV1ProtectedFiltersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol content targeting v1 protected filters params
func (o *GetLolContentTargetingV1ProtectedFiltersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol content targeting v1 protected filters params
func (o *GetLolContentTargetingV1ProtectedFiltersParams) WithContext(ctx context.Context) *GetLolContentTargetingV1ProtectedFiltersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol content targeting v1 protected filters params
func (o *GetLolContentTargetingV1ProtectedFiltersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol content targeting v1 protected filters params
func (o *GetLolContentTargetingV1ProtectedFiltersParams) WithHTTPClient(client *http.Client) *GetLolContentTargetingV1ProtectedFiltersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol content targeting v1 protected filters params
func (o *GetLolContentTargetingV1ProtectedFiltersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolContentTargetingV1ProtectedFiltersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
