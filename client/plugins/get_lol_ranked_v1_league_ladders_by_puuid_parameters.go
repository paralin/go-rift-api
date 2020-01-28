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

// NewGetLolRankedV1LeagueLaddersByPuuidParams creates a new GetLolRankedV1LeagueLaddersByPuuidParams object
// with the default values initialized.
func NewGetLolRankedV1LeagueLaddersByPuuidParams() *GetLolRankedV1LeagueLaddersByPuuidParams {
	var ()
	return &GetLolRankedV1LeagueLaddersByPuuidParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolRankedV1LeagueLaddersByPuuidParamsWithTimeout creates a new GetLolRankedV1LeagueLaddersByPuuidParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolRankedV1LeagueLaddersByPuuidParamsWithTimeout(timeout time.Duration) *GetLolRankedV1LeagueLaddersByPuuidParams {
	var ()
	return &GetLolRankedV1LeagueLaddersByPuuidParams{

		timeout: timeout,
	}
}

// NewGetLolRankedV1LeagueLaddersByPuuidParamsWithContext creates a new GetLolRankedV1LeagueLaddersByPuuidParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolRankedV1LeagueLaddersByPuuidParamsWithContext(ctx context.Context) *GetLolRankedV1LeagueLaddersByPuuidParams {
	var ()
	return &GetLolRankedV1LeagueLaddersByPuuidParams{

		Context: ctx,
	}
}

// NewGetLolRankedV1LeagueLaddersByPuuidParamsWithHTTPClient creates a new GetLolRankedV1LeagueLaddersByPuuidParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolRankedV1LeagueLaddersByPuuidParamsWithHTTPClient(client *http.Client) *GetLolRankedV1LeagueLaddersByPuuidParams {
	var ()
	return &GetLolRankedV1LeagueLaddersByPuuidParams{
		HTTPClient: client,
	}
}

/*GetLolRankedV1LeagueLaddersByPuuidParams contains all the parameters to send to the API endpoint
for the get lol ranked v1 league ladders by puuid operation typically these are written to a http.Request
*/
type GetLolRankedV1LeagueLaddersByPuuidParams struct {

	/*Puuid*/
	Puuid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) WithTimeout(timeout time.Duration) *GetLolRankedV1LeagueLaddersByPuuidParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) WithContext(ctx context.Context) *GetLolRankedV1LeagueLaddersByPuuidParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) WithHTTPClient(client *http.Client) *GetLolRankedV1LeagueLaddersByPuuidParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPuuid adds the puuid to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) WithPuuid(puuid string) *GetLolRankedV1LeagueLaddersByPuuidParams {
	o.SetPuuid(puuid)
	return o
}

// SetPuuid adds the puuid to the get lol ranked v1 league ladders by puuid params
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) SetPuuid(puuid string) {
	o.Puuid = puuid
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolRankedV1LeagueLaddersByPuuidParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
