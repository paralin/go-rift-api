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

// NewGetLolSummonerV1CurrentSummonerJwtParams creates a new GetLolSummonerV1CurrentSummonerJwtParams object
// with the default values initialized.
func NewGetLolSummonerV1CurrentSummonerJwtParams() *GetLolSummonerV1CurrentSummonerJwtParams {

	return &GetLolSummonerV1CurrentSummonerJwtParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSummonerV1CurrentSummonerJwtParamsWithTimeout creates a new GetLolSummonerV1CurrentSummonerJwtParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSummonerV1CurrentSummonerJwtParamsWithTimeout(timeout time.Duration) *GetLolSummonerV1CurrentSummonerJwtParams {

	return &GetLolSummonerV1CurrentSummonerJwtParams{

		timeout: timeout,
	}
}

// NewGetLolSummonerV1CurrentSummonerJwtParamsWithContext creates a new GetLolSummonerV1CurrentSummonerJwtParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSummonerV1CurrentSummonerJwtParamsWithContext(ctx context.Context) *GetLolSummonerV1CurrentSummonerJwtParams {

	return &GetLolSummonerV1CurrentSummonerJwtParams{

		Context: ctx,
	}
}

// NewGetLolSummonerV1CurrentSummonerJwtParamsWithHTTPClient creates a new GetLolSummonerV1CurrentSummonerJwtParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSummonerV1CurrentSummonerJwtParamsWithHTTPClient(client *http.Client) *GetLolSummonerV1CurrentSummonerJwtParams {

	return &GetLolSummonerV1CurrentSummonerJwtParams{
		HTTPClient: client,
	}
}

/*GetLolSummonerV1CurrentSummonerJwtParams contains all the parameters to send to the API endpoint
for the get lol summoner v1 current summoner jwt operation typically these are written to a http.Request
*/
type GetLolSummonerV1CurrentSummonerJwtParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol summoner v1 current summoner jwt params
func (o *GetLolSummonerV1CurrentSummonerJwtParams) WithTimeout(timeout time.Duration) *GetLolSummonerV1CurrentSummonerJwtParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol summoner v1 current summoner jwt params
func (o *GetLolSummonerV1CurrentSummonerJwtParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol summoner v1 current summoner jwt params
func (o *GetLolSummonerV1CurrentSummonerJwtParams) WithContext(ctx context.Context) *GetLolSummonerV1CurrentSummonerJwtParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol summoner v1 current summoner jwt params
func (o *GetLolSummonerV1CurrentSummonerJwtParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol summoner v1 current summoner jwt params
func (o *GetLolSummonerV1CurrentSummonerJwtParams) WithHTTPClient(client *http.Client) *GetLolSummonerV1CurrentSummonerJwtParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol summoner v1 current summoner jwt params
func (o *GetLolSummonerV1CurrentSummonerJwtParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSummonerV1CurrentSummonerJwtParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
