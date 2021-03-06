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

// NewGetLolSummonerV2SummonersPuuidByPuuidParams creates a new GetLolSummonerV2SummonersPuuidByPuuidParams object
// with the default values initialized.
func NewGetLolSummonerV2SummonersPuuidByPuuidParams() *GetLolSummonerV2SummonersPuuidByPuuidParams {
	var ()
	return &GetLolSummonerV2SummonersPuuidByPuuidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSummonerV2SummonersPuuidByPuuidParamsWithTimeout creates a new GetLolSummonerV2SummonersPuuidByPuuidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSummonerV2SummonersPuuidByPuuidParamsWithTimeout(timeout time.Duration) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	var ()
	return &GetLolSummonerV2SummonersPuuidByPuuidParams{

		timeout: timeout,
	}
}

// NewGetLolSummonerV2SummonersPuuidByPuuidParamsWithContext creates a new GetLolSummonerV2SummonersPuuidByPuuidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSummonerV2SummonersPuuidByPuuidParamsWithContext(ctx context.Context) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	var ()
	return &GetLolSummonerV2SummonersPuuidByPuuidParams{

		Context: ctx,
	}
}

// NewGetLolSummonerV2SummonersPuuidByPuuidParamsWithHTTPClient creates a new GetLolSummonerV2SummonersPuuidByPuuidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSummonerV2SummonersPuuidByPuuidParamsWithHTTPClient(client *http.Client) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	var ()
	return &GetLolSummonerV2SummonersPuuidByPuuidParams{
		HTTPClient: client,
	}
}

/*GetLolSummonerV2SummonersPuuidByPuuidParams contains all the parameters to send to the API endpoint
for the get lol summoner v2 summoners puuid by puuid operation typically these are written to a http.Request
*/
type GetLolSummonerV2SummonersPuuidByPuuidParams struct {

	/*Puuid*/
	Puuid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) WithTimeout(timeout time.Duration) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) WithContext(ctx context.Context) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) WithHTTPClient(client *http.Client) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPuuid adds the puuid to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) WithPuuid(puuid string) *GetLolSummonerV2SummonersPuuidByPuuidParams {
	o.SetPuuid(puuid)
	return o
}

// SetPuuid adds the puuid to the get lol summoner v2 summoners puuid by puuid params
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) SetPuuid(puuid string) {
	o.Puuid = puuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSummonerV2SummonersPuuidByPuuidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param puuid
	if err := r.SetPathParam("puuid", o.Puuid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
