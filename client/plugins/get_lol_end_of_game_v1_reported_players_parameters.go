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

// NewGetLolEndOfGameV1ReportedPlayersParams creates a new GetLolEndOfGameV1ReportedPlayersParams object
// with the default values initialized.
func NewGetLolEndOfGameV1ReportedPlayersParams() *GetLolEndOfGameV1ReportedPlayersParams {

	return &GetLolEndOfGameV1ReportedPlayersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolEndOfGameV1ReportedPlayersParamsWithTimeout creates a new GetLolEndOfGameV1ReportedPlayersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolEndOfGameV1ReportedPlayersParamsWithTimeout(timeout time.Duration) *GetLolEndOfGameV1ReportedPlayersParams {

	return &GetLolEndOfGameV1ReportedPlayersParams{

		timeout: timeout,
	}
}

// NewGetLolEndOfGameV1ReportedPlayersParamsWithContext creates a new GetLolEndOfGameV1ReportedPlayersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolEndOfGameV1ReportedPlayersParamsWithContext(ctx context.Context) *GetLolEndOfGameV1ReportedPlayersParams {

	return &GetLolEndOfGameV1ReportedPlayersParams{

		Context: ctx,
	}
}

// NewGetLolEndOfGameV1ReportedPlayersParamsWithHTTPClient creates a new GetLolEndOfGameV1ReportedPlayersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolEndOfGameV1ReportedPlayersParamsWithHTTPClient(client *http.Client) *GetLolEndOfGameV1ReportedPlayersParams {

	return &GetLolEndOfGameV1ReportedPlayersParams{
		HTTPClient: client,
	}
}

/*GetLolEndOfGameV1ReportedPlayersParams contains all the parameters to send to the API endpoint
for the get lol end of game v1 reported players operation typically these are written to a http.Request
*/
type GetLolEndOfGameV1ReportedPlayersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol end of game v1 reported players params
func (o *GetLolEndOfGameV1ReportedPlayersParams) WithTimeout(timeout time.Duration) *GetLolEndOfGameV1ReportedPlayersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol end of game v1 reported players params
func (o *GetLolEndOfGameV1ReportedPlayersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol end of game v1 reported players params
func (o *GetLolEndOfGameV1ReportedPlayersParams) WithContext(ctx context.Context) *GetLolEndOfGameV1ReportedPlayersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol end of game v1 reported players params
func (o *GetLolEndOfGameV1ReportedPlayersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol end of game v1 reported players params
func (o *GetLolEndOfGameV1ReportedPlayersParams) WithHTTPClient(client *http.Client) *GetLolEndOfGameV1ReportedPlayersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol end of game v1 reported players params
func (o *GetLolEndOfGameV1ReportedPlayersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolEndOfGameV1ReportedPlayersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
