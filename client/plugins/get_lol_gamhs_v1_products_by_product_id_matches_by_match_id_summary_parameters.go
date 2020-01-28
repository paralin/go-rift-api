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

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams object
// with the default values initialized.
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams() *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParamsWithTimeout creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParamsWithTimeout(timeout time.Duration) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams{

		timeout: timeout,
	}
}

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParamsWithContext creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParamsWithContext(ctx context.Context) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams{

		Context: ctx,
	}
}

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParamsWithHTTPClient creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParamsWithHTTPClient(client *http.Client) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams{
		HTTPClient: client,
	}
}

/*GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams contains all the parameters to send to the API endpoint
for the get lol gamhs v1 products by product Id matches by match Id summary operation typically these are written to a http.Request
*/
type GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams struct {

	/*MatchID*/
	MatchID string
	/*ProductID*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) WithTimeout(timeout time.Duration) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) WithContext(ctx context.Context) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) WithHTTPClient(client *http.Client) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMatchID adds the matchID to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) WithMatchID(matchID string) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	o.SetMatchID(matchID)
	return o
}

// SetMatchID adds the matchId to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) SetMatchID(matchID string) {
	o.MatchID = matchID
}

// WithProductID adds the productID to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) WithProductID(productID string) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get lol gamhs v1 products by product Id matches by match Id summary params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param matchId
	if err := r.SetPathParam("matchId", o.MatchID); err != nil {
		return err
	}

	// path param productId
	if err := r.SetPathParam("productId", o.ProductID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
