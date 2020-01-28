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

// NewGetLolSummonerV1SummonersByIDParams creates a new GetLolSummonerV1SummonersByIDParams object
// with the default values initialized.
func NewGetLolSummonerV1SummonersByIDParams() *GetLolSummonerV1SummonersByIDParams {
	var ()
	return &GetLolSummonerV1SummonersByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolSummonerV1SummonersByIDParamsWithTimeout creates a new GetLolSummonerV1SummonersByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolSummonerV1SummonersByIDParamsWithTimeout(timeout time.Duration) *GetLolSummonerV1SummonersByIDParams {
	var ()
	return &GetLolSummonerV1SummonersByIDParams{

		timeout: timeout,
	}
}

// NewGetLolSummonerV1SummonersByIDParamsWithContext creates a new GetLolSummonerV1SummonersByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolSummonerV1SummonersByIDParamsWithContext(ctx context.Context) *GetLolSummonerV1SummonersByIDParams {
	var ()
	return &GetLolSummonerV1SummonersByIDParams{

		Context: ctx,
	}
}

// NewGetLolSummonerV1SummonersByIDParamsWithHTTPClient creates a new GetLolSummonerV1SummonersByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolSummonerV1SummonersByIDParamsWithHTTPClient(client *http.Client) *GetLolSummonerV1SummonersByIDParams {
	var ()
	return &GetLolSummonerV1SummonersByIDParams{
		HTTPClient: client,
	}
}

/*GetLolSummonerV1SummonersByIDParams contains all the parameters to send to the API endpoint
for the get lol summoner v1 summoners by Id operation typically these are written to a http.Request
*/
type GetLolSummonerV1SummonersByIDParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) WithTimeout(timeout time.Duration) *GetLolSummonerV1SummonersByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) WithContext(ctx context.Context) *GetLolSummonerV1SummonersByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) WithHTTPClient(client *http.Client) *GetLolSummonerV1SummonersByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) WithID(id int64) *GetLolSummonerV1SummonersByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get lol summoner v1 summoners by Id params
func (o *GetLolSummonerV1SummonersByIDParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolSummonerV1SummonersByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
