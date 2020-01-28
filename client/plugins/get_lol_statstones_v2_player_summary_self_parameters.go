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

// NewGetLolStatstonesV2PlayerSummarySelfParams creates a new GetLolStatstonesV2PlayerSummarySelfParams object
// with the default values initialized.
func NewGetLolStatstonesV2PlayerSummarySelfParams() *GetLolStatstonesV2PlayerSummarySelfParams {

	return &GetLolStatstonesV2PlayerSummarySelfParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolStatstonesV2PlayerSummarySelfParamsWithTimeout creates a new GetLolStatstonesV2PlayerSummarySelfParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolStatstonesV2PlayerSummarySelfParamsWithTimeout(timeout time.Duration) *GetLolStatstonesV2PlayerSummarySelfParams {

	return &GetLolStatstonesV2PlayerSummarySelfParams{

		timeout: timeout,
	}
}

// NewGetLolStatstonesV2PlayerSummarySelfParamsWithContext creates a new GetLolStatstonesV2PlayerSummarySelfParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolStatstonesV2PlayerSummarySelfParamsWithContext(ctx context.Context) *GetLolStatstonesV2PlayerSummarySelfParams {

	return &GetLolStatstonesV2PlayerSummarySelfParams{

		Context: ctx,
	}
}

// NewGetLolStatstonesV2PlayerSummarySelfParamsWithHTTPClient creates a new GetLolStatstonesV2PlayerSummarySelfParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolStatstonesV2PlayerSummarySelfParamsWithHTTPClient(client *http.Client) *GetLolStatstonesV2PlayerSummarySelfParams {

	return &GetLolStatstonesV2PlayerSummarySelfParams{
		HTTPClient: client,
	}
}

/*GetLolStatstonesV2PlayerSummarySelfParams contains all the parameters to send to the API endpoint
for the get lol statstones v2 player summary self operation typically these are written to a http.Request
*/
type GetLolStatstonesV2PlayerSummarySelfParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol statstones v2 player summary self params
func (o *GetLolStatstonesV2PlayerSummarySelfParams) WithTimeout(timeout time.Duration) *GetLolStatstonesV2PlayerSummarySelfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol statstones v2 player summary self params
func (o *GetLolStatstonesV2PlayerSummarySelfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol statstones v2 player summary self params
func (o *GetLolStatstonesV2PlayerSummarySelfParams) WithContext(ctx context.Context) *GetLolStatstonesV2PlayerSummarySelfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol statstones v2 player summary self params
func (o *GetLolStatstonesV2PlayerSummarySelfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol statstones v2 player summary self params
func (o *GetLolStatstonesV2PlayerSummarySelfParams) WithHTTPClient(client *http.Client) *GetLolStatstonesV2PlayerSummarySelfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol statstones v2 player summary self params
func (o *GetLolStatstonesV2PlayerSummarySelfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolStatstonesV2PlayerSummarySelfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}