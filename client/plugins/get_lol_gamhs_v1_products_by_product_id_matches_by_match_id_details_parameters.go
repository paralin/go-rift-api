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

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams object
// with the default values initialized.
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams() *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParamsWithTimeout creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParamsWithTimeout(timeout time.Duration) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams{

		timeout: timeout,
	}
}

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParamsWithContext creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParamsWithContext(ctx context.Context) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams{

		Context: ctx,
	}
}

// NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParamsWithHTTPClient creates a new GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParamsWithHTTPClient(client *http.Client) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	var ()
	return &GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams{
		HTTPClient: client,
	}
}

/*GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams contains all the parameters to send to the API endpoint
for the get lol gamhs v1 products by product Id matches by match Id details operation typically these are written to a http.Request
*/
type GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams struct {

	/*MatchID*/
	MatchID string
	/*ProductID*/
	ProductID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) WithTimeout(timeout time.Duration) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) WithContext(ctx context.Context) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) WithHTTPClient(client *http.Client) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMatchID adds the matchID to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) WithMatchID(matchID string) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	o.SetMatchID(matchID)
	return o
}

// SetMatchID adds the matchId to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) SetMatchID(matchID string) {
	o.MatchID = matchID
}

// WithProductID adds the productID to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) WithProductID(productID string) *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams {
	o.SetProductID(productID)
	return o
}

// SetProductID adds the productId to the get lol gamhs v1 products by product Id matches by match Id details params
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) SetProductID(productID string) {
	o.ProductID = productID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolGamhsV1ProductsByProductIDMatchesByMatchIDDetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
