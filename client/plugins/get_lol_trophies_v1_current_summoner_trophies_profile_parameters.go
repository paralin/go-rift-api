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

// NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParams creates a new GetLolTrophiesV1CurrentSummonerTrophiesProfileParams object
// with the default values initialized.
func NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParams() *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {

	return &GetLolTrophiesV1CurrentSummonerTrophiesProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParamsWithTimeout creates a new GetLolTrophiesV1CurrentSummonerTrophiesProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParamsWithTimeout(timeout time.Duration) *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {

	return &GetLolTrophiesV1CurrentSummonerTrophiesProfileParams{

		timeout: timeout,
	}
}

// NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParamsWithContext creates a new GetLolTrophiesV1CurrentSummonerTrophiesProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParamsWithContext(ctx context.Context) *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {

	return &GetLolTrophiesV1CurrentSummonerTrophiesProfileParams{

		Context: ctx,
	}
}

// NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParamsWithHTTPClient creates a new GetLolTrophiesV1CurrentSummonerTrophiesProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolTrophiesV1CurrentSummonerTrophiesProfileParamsWithHTTPClient(client *http.Client) *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {

	return &GetLolTrophiesV1CurrentSummonerTrophiesProfileParams{
		HTTPClient: client,
	}
}

/*GetLolTrophiesV1CurrentSummonerTrophiesProfileParams contains all the parameters to send to the API endpoint
for the get lol trophies v1 current summoner trophies profile operation typically these are written to a http.Request
*/
type GetLolTrophiesV1CurrentSummonerTrophiesProfileParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol trophies v1 current summoner trophies profile params
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) WithTimeout(timeout time.Duration) *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol trophies v1 current summoner trophies profile params
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol trophies v1 current summoner trophies profile params
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) WithContext(ctx context.Context) *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol trophies v1 current summoner trophies profile params
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol trophies v1 current summoner trophies profile params
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) WithHTTPClient(client *http.Client) *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol trophies v1 current summoner trophies profile params
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolTrophiesV1CurrentSummonerTrophiesProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}