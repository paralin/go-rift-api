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

// NewGetLolPurchaseWidgetV1ConfigurationParams creates a new GetLolPurchaseWidgetV1ConfigurationParams object
// with the default values initialized.
func NewGetLolPurchaseWidgetV1ConfigurationParams() *GetLolPurchaseWidgetV1ConfigurationParams {

	return &GetLolPurchaseWidgetV1ConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolPurchaseWidgetV1ConfigurationParamsWithTimeout creates a new GetLolPurchaseWidgetV1ConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolPurchaseWidgetV1ConfigurationParamsWithTimeout(timeout time.Duration) *GetLolPurchaseWidgetV1ConfigurationParams {

	return &GetLolPurchaseWidgetV1ConfigurationParams{

		timeout: timeout,
	}
}

// NewGetLolPurchaseWidgetV1ConfigurationParamsWithContext creates a new GetLolPurchaseWidgetV1ConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolPurchaseWidgetV1ConfigurationParamsWithContext(ctx context.Context) *GetLolPurchaseWidgetV1ConfigurationParams {

	return &GetLolPurchaseWidgetV1ConfigurationParams{

		Context: ctx,
	}
}

// NewGetLolPurchaseWidgetV1ConfigurationParamsWithHTTPClient creates a new GetLolPurchaseWidgetV1ConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolPurchaseWidgetV1ConfigurationParamsWithHTTPClient(client *http.Client) *GetLolPurchaseWidgetV1ConfigurationParams {

	return &GetLolPurchaseWidgetV1ConfigurationParams{
		HTTPClient: client,
	}
}

/*GetLolPurchaseWidgetV1ConfigurationParams contains all the parameters to send to the API endpoint
for the get lol purchase widget v1 configuration operation typically these are written to a http.Request
*/
type GetLolPurchaseWidgetV1ConfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol purchase widget v1 configuration params
func (o *GetLolPurchaseWidgetV1ConfigurationParams) WithTimeout(timeout time.Duration) *GetLolPurchaseWidgetV1ConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol purchase widget v1 configuration params
func (o *GetLolPurchaseWidgetV1ConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol purchase widget v1 configuration params
func (o *GetLolPurchaseWidgetV1ConfigurationParams) WithContext(ctx context.Context) *GetLolPurchaseWidgetV1ConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol purchase widget v1 configuration params
func (o *GetLolPurchaseWidgetV1ConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol purchase widget v1 configuration params
func (o *GetLolPurchaseWidgetV1ConfigurationParams) WithHTTPClient(client *http.Client) *GetLolPurchaseWidgetV1ConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol purchase widget v1 configuration params
func (o *GetLolPurchaseWidgetV1ConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolPurchaseWidgetV1ConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
