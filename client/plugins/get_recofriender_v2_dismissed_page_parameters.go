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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRecofrienderV2DismissedPageParams creates a new GetRecofrienderV2DismissedPageParams object
// with the default values initialized.
func NewGetRecofrienderV2DismissedPageParams() *GetRecofrienderV2DismissedPageParams {
	var ()
	return &GetRecofrienderV2DismissedPageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecofrienderV2DismissedPageParamsWithTimeout creates a new GetRecofrienderV2DismissedPageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecofrienderV2DismissedPageParamsWithTimeout(timeout time.Duration) *GetRecofrienderV2DismissedPageParams {
	var ()
	return &GetRecofrienderV2DismissedPageParams{

		timeout: timeout,
	}
}

// NewGetRecofrienderV2DismissedPageParamsWithContext creates a new GetRecofrienderV2DismissedPageParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecofrienderV2DismissedPageParamsWithContext(ctx context.Context) *GetRecofrienderV2DismissedPageParams {
	var ()
	return &GetRecofrienderV2DismissedPageParams{

		Context: ctx,
	}
}

// NewGetRecofrienderV2DismissedPageParamsWithHTTPClient creates a new GetRecofrienderV2DismissedPageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecofrienderV2DismissedPageParamsWithHTTPClient(client *http.Client) *GetRecofrienderV2DismissedPageParams {
	var ()
	return &GetRecofrienderV2DismissedPageParams{
		HTTPClient: client,
	}
}

/*GetRecofrienderV2DismissedPageParams contains all the parameters to send to the API endpoint
for the get recofriender v2 dismissed page operation typically these are written to a http.Request
*/
type GetRecofrienderV2DismissedPageParams struct {

	/*Limit*/
	Limit int64
	/*Start*/
	Start int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) WithTimeout(timeout time.Duration) *GetRecofrienderV2DismissedPageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) WithContext(ctx context.Context) *GetRecofrienderV2DismissedPageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) WithHTTPClient(client *http.Client) *GetRecofrienderV2DismissedPageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) WithLimit(limit int64) *GetRecofrienderV2DismissedPageParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) SetLimit(limit int64) {
	o.Limit = limit
}

// WithStart adds the start to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) WithStart(start int64) *GetRecofrienderV2DismissedPageParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the get recofriender v2 dismissed page params
func (o *GetRecofrienderV2DismissedPageParams) SetStart(start int64) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecofrienderV2DismissedPageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param limit
	qrLimit := o.Limit
	qLimit := swag.FormatInt64(qrLimit)
	if qLimit != "" {
		if err := r.SetQueryParam("limit", qLimit); err != nil {
			return err
		}
	}

	// query param start
	qrStart := o.Start
	qStart := swag.FormatInt64(qrStart)
	if qStart != "" {
		if err := r.SetQueryParam("start", qStart); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
