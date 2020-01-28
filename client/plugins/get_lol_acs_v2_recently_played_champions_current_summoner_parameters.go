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

// NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams creates a new GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams object
// with the default values initialized.
func NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams() *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	var ()
	return &GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParamsWithTimeout creates a new GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParamsWithTimeout(timeout time.Duration) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	var ()
	return &GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams{

		timeout: timeout,
	}
}

// NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParamsWithContext creates a new GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParamsWithContext(ctx context.Context) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	var ()
	return &GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams{

		Context: ctx,
	}
}

// NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParamsWithHTTPClient creates a new GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParamsWithHTTPClient(client *http.Client) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	var ()
	return &GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams{
		HTTPClient: client,
	}
}

/*GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams contains all the parameters to send to the API endpoint
for the get lol acs v2 recently played champions current summoner operation typically these are written to a http.Request
*/
type GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams struct {

	/*Force*/
	Force bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) WithTimeout(timeout time.Duration) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) WithContext(ctx context.Context) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) WithHTTPClient(client *http.Client) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) WithForce(force bool) *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the get lol acs v2 recently played champions current summoner params
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) SetForce(force bool) {
	o.Force = force
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolAcsV2RecentlyPlayedChampionsCurrentSummonerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param force
	qrForce := o.Force
	qForce := swag.FormatBool(qrForce)
	if qForce != "" {
		if err := r.SetQueryParam("force", qForce); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}