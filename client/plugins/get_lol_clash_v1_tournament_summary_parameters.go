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

// NewGetLolClashV1TournamentSummaryParams creates a new GetLolClashV1TournamentSummaryParams object
// with the default values initialized.
func NewGetLolClashV1TournamentSummaryParams() *GetLolClashV1TournamentSummaryParams {

	return &GetLolClashV1TournamentSummaryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClashV1TournamentSummaryParamsWithTimeout creates a new GetLolClashV1TournamentSummaryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClashV1TournamentSummaryParamsWithTimeout(timeout time.Duration) *GetLolClashV1TournamentSummaryParams {

	return &GetLolClashV1TournamentSummaryParams{

		timeout: timeout,
	}
}

// NewGetLolClashV1TournamentSummaryParamsWithContext creates a new GetLolClashV1TournamentSummaryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClashV1TournamentSummaryParamsWithContext(ctx context.Context) *GetLolClashV1TournamentSummaryParams {

	return &GetLolClashV1TournamentSummaryParams{

		Context: ctx,
	}
}

// NewGetLolClashV1TournamentSummaryParamsWithHTTPClient creates a new GetLolClashV1TournamentSummaryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClashV1TournamentSummaryParamsWithHTTPClient(client *http.Client) *GetLolClashV1TournamentSummaryParams {

	return &GetLolClashV1TournamentSummaryParams{
		HTTPClient: client,
	}
}

/*GetLolClashV1TournamentSummaryParams contains all the parameters to send to the API endpoint
for the get lol clash v1 tournament summary operation typically these are written to a http.Request
*/
type GetLolClashV1TournamentSummaryParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clash v1 tournament summary params
func (o *GetLolClashV1TournamentSummaryParams) WithTimeout(timeout time.Duration) *GetLolClashV1TournamentSummaryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clash v1 tournament summary params
func (o *GetLolClashV1TournamentSummaryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clash v1 tournament summary params
func (o *GetLolClashV1TournamentSummaryParams) WithContext(ctx context.Context) *GetLolClashV1TournamentSummaryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clash v1 tournament summary params
func (o *GetLolClashV1TournamentSummaryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clash v1 tournament summary params
func (o *GetLolClashV1TournamentSummaryParams) WithHTTPClient(client *http.Client) *GetLolClashV1TournamentSummaryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clash v1 tournament summary params
func (o *GetLolClashV1TournamentSummaryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClashV1TournamentSummaryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
