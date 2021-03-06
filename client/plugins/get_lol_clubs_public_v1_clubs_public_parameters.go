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

// NewGetLolClubsPublicV1ClubsPublicParams creates a new GetLolClubsPublicV1ClubsPublicParams object
// with the default values initialized.
func NewGetLolClubsPublicV1ClubsPublicParams() *GetLolClubsPublicV1ClubsPublicParams {
	var ()
	return &GetLolClubsPublicV1ClubsPublicParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolClubsPublicV1ClubsPublicParamsWithTimeout creates a new GetLolClubsPublicV1ClubsPublicParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolClubsPublicV1ClubsPublicParamsWithTimeout(timeout time.Duration) *GetLolClubsPublicV1ClubsPublicParams {
	var ()
	return &GetLolClubsPublicV1ClubsPublicParams{

		timeout: timeout,
	}
}

// NewGetLolClubsPublicV1ClubsPublicParamsWithContext creates a new GetLolClubsPublicV1ClubsPublicParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolClubsPublicV1ClubsPublicParamsWithContext(ctx context.Context) *GetLolClubsPublicV1ClubsPublicParams {
	var ()
	return &GetLolClubsPublicV1ClubsPublicParams{

		Context: ctx,
	}
}

// NewGetLolClubsPublicV1ClubsPublicParamsWithHTTPClient creates a new GetLolClubsPublicV1ClubsPublicParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolClubsPublicV1ClubsPublicParamsWithHTTPClient(client *http.Client) *GetLolClubsPublicV1ClubsPublicParams {
	var ()
	return &GetLolClubsPublicV1ClubsPublicParams{
		HTTPClient: client,
	}
}

/*GetLolClubsPublicV1ClubsPublicParams contains all the parameters to send to the API endpoint
for the get lol clubs public v1 clubs public operation typically these are written to a http.Request
*/
type GetLolClubsPublicV1ClubsPublicParams struct {

	/*SummonerNames*/
	SummonerNames string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) WithTimeout(timeout time.Duration) *GetLolClubsPublicV1ClubsPublicParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) WithContext(ctx context.Context) *GetLolClubsPublicV1ClubsPublicParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) WithHTTPClient(client *http.Client) *GetLolClubsPublicV1ClubsPublicParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSummonerNames adds the summonerNames to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) WithSummonerNames(summonerNames string) *GetLolClubsPublicV1ClubsPublicParams {
	o.SetSummonerNames(summonerNames)
	return o
}

// SetSummonerNames adds the summonerNames to the get lol clubs public v1 clubs public params
func (o *GetLolClubsPublicV1ClubsPublicParams) SetSummonerNames(summonerNames string) {
	o.SummonerNames = summonerNames
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolClubsPublicV1ClubsPublicParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param summonerNames
	qrSummonerNames := o.SummonerNames
	qSummonerNames := qrSummonerNames
	if qSummonerNames != "" {
		if err := r.SetQueryParam("summonerNames", qSummonerNames); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
