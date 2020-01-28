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

// NewGetPatcherV1ProductsByProductIDPathsParams creates a new GetPatcherV1ProductsByProductIDPathsParams object
// with the default values initialized.
func NewGetPatcherV1ProductsByProductIDPathsParams() *GetPatcherV1ProductsByProductIDPathsParams {
	var ()
	return &GetPatcherV1ProductsByProductIDPathsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPatcherV1ProductsByProductIDPathsParamsWithTimeout creates a new GetPatcherV1ProductsByProductIDPathsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPatcherV1ProductsByProductIDPathsParamsWithTimeout(timeout time.Duration) *GetPatcherV1ProductsByProductIDPathsParams {
	var ()
	return &GetPatcherV1ProductsByProductIDPathsParams{

		timeout: timeout,
	}
}

// NewGetPatcherV1ProductsByProductIDPathsParamsWithContext creates a new GetPatcherV1ProductsByProductIDPathsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPatcherV1ProductsByProductIDPathsParamsWithContext(ctx context.Context) *GetPatcherV1ProductsByProductIDPathsParams {
	var ()
	return &GetPatcherV1ProductsByProductIDPathsParams{

		Context: ctx,
	}
}

// NewGetPatcherV1ProductsByProductIDPathsParamsWithHTTPClient creates a new GetPatcherV1ProductsByProductIDPathsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPatcherV1ProductsByProductIDPathsParamsWithHTTPClient(client *http.Client) *GetPatcherV1ProductsByProductIDPathsParams {
	var ()
	return &GetPatcherV1ProductsByProductIDPathsParams{
		HTTPClient: client,
	}
}

/*GetPatcherV1ProductsByProductIDPathsParams contains all the parameters to send to the API endpoint
for the get patcher v1 products by product Id paths operation typically these are written to a http.Request
*/
type GetPatcherV1ProductsByProductIDPathsParams struct {

	/*ProductID*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) WithTimeout(timeout time.Duration) *GetPatcherV1ProductsByProductIDPathsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) WithContext(ctx context.Context) *GetPatcherV1ProductsByProductIDPathsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) WithHTTPClient(client *http.Client) *GetPatcherV1ProductsByProductIDPathsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProductID adds the productID to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) WithProductID(productID string) *GetPatcherV1ProductsByProductIDPathsParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get patcher v1 products by product Id paths params
func (o *GetPatcherV1ProductsByProductIDPathsParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPatcherV1ProductsByProductIDPathsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param product-id
	if err := r.SetPathParam("product-id", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}