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

// NewGetLolTencentQtV1UIStatesParams creates a new GetLolTencentQtV1UIStatesParams object
// with the default values initialized.
func NewGetLolTencentQtV1UIStatesParams() *GetLolTencentQtV1UIStatesParams {

	return &GetLolTencentQtV1UIStatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolTencentQtV1UIStatesParamsWithTimeout creates a new GetLolTencentQtV1UIStatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolTencentQtV1UIStatesParamsWithTimeout(timeout time.Duration) *GetLolTencentQtV1UIStatesParams {

	return &GetLolTencentQtV1UIStatesParams{

		timeout: timeout,
	}
}

// NewGetLolTencentQtV1UIStatesParamsWithContext creates a new GetLolTencentQtV1UIStatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolTencentQtV1UIStatesParamsWithContext(ctx context.Context) *GetLolTencentQtV1UIStatesParams {

	return &GetLolTencentQtV1UIStatesParams{

		Context: ctx,
	}
}

// NewGetLolTencentQtV1UIStatesParamsWithHTTPClient creates a new GetLolTencentQtV1UIStatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolTencentQtV1UIStatesParamsWithHTTPClient(client *http.Client) *GetLolTencentQtV1UIStatesParams {

	return &GetLolTencentQtV1UIStatesParams{
		HTTPClient: client,
	}
}

/*GetLolTencentQtV1UIStatesParams contains all the parameters to send to the API endpoint
for the get lol tencent qt v1 Ui states operation typically these are written to a http.Request
*/
type GetLolTencentQtV1UIStatesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol tencent qt v1 Ui states params
func (o *GetLolTencentQtV1UIStatesParams) WithTimeout(timeout time.Duration) *GetLolTencentQtV1UIStatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol tencent qt v1 Ui states params
func (o *GetLolTencentQtV1UIStatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol tencent qt v1 Ui states params
func (o *GetLolTencentQtV1UIStatesParams) WithContext(ctx context.Context) *GetLolTencentQtV1UIStatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol tencent qt v1 Ui states params
func (o *GetLolTencentQtV1UIStatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol tencent qt v1 Ui states params
func (o *GetLolTencentQtV1UIStatesParams) WithHTTPClient(client *http.Client) *GetLolTencentQtV1UIStatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol tencent qt v1 Ui states params
func (o *GetLolTencentQtV1UIStatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolTencentQtV1UIStatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}