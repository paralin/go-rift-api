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

// NewGetLolClashV1HistoryandwinnersParams creates a new GetLolClashV1HistoryandwinnersParams object
// with the default values initialized.
func NewGetLolClashV1HistoryandwinnersParams() *GetLolClashV1HistoryandwinnersParams {

	return &GetLolClashV1HistoryandwinnersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1HistoryandwinnersParamsWithTimeout creates a new GetLolClashV1HistoryandwinnersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1HistoryandwinnersParamsWithTimeout(timeout time.Duration) *GetLolClashV1HistoryandwinnersParams {

	return &GetLolClashV1HistoryandwinnersParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1HistoryandwinnersParamsWithContext creates a new GetLolClashV1HistoryandwinnersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1HistoryandwinnersParamsWithContext(ctx context.Context) *GetLolClashV1HistoryandwinnersParams {

	return &GetLolClashV1HistoryandwinnersParams{

		Context: ctx,
	}
}

// NewGetLolClashV1HistoryandwinnersParamsWithHTTPClient creates a new GetLolClashV1HistoryandwinnersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1HistoryandwinnersParamsWithHTTPClient(client *http.Client) *GetLolClashV1HistoryandwinnersParams {

	return &GetLolClashV1HistoryandwinnersParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1HistoryandwinnersParams contains all the parameters to send to the API endpoint
for the get lol clash v1 historyandwinners operation typically these are written to a http.Request
*/
type GetLolClashV1HistoryandwinnersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 historyandwinners params
func (o *GetLolClashV1HistoryandwinnersParams) WithTimeout(timeout time.Duration) *GetLolClashV1HistoryandwinnersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 historyandwinners params
func (o *GetLolClashV1HistoryandwinnersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 historyandwinners params
func (o *GetLolClashV1HistoryandwinnersParams) WithContext(ctx context.Context) *GetLolClashV1HistoryandwinnersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 historyandwinners params
func (o *GetLolClashV1HistoryandwinnersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 historyandwinners params
func (o *GetLolClashV1HistoryandwinnersParams) WithHTTPClient(client *http.Client) *GetLolClashV1HistoryandwinnersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 historyandwinners params
func (o *GetLolClashV1HistoryandwinnersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1HistoryandwinnersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
