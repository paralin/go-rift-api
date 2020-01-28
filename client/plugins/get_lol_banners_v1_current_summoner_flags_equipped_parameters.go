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

// NewGetLolBannersV1CurrentSummonerFlagsEquippedParams creates a new GetLolBannersV1CurrentSummonerFlagsEquippedParams object
// with the default values initialized.
func NewGetLolBannersV1CurrentSummonerFlagsEquippedParams() *GetLolBannersV1CurrentSummonerFlagsEquippedParams {

	return &GetLolBannersV1CurrentSummonerFlagsEquippedParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolBannersV1CurrentSummonerFlagsEquippedParamsWithTimeout creates a new GetLolBannersV1CurrentSummonerFlagsEquippedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolBannersV1CurrentSummonerFlagsEquippedParamsWithTimeout(timeout time.Duration) *GetLolBannersV1CurrentSummonerFlagsEquippedParams {

	return &GetLolBannersV1CurrentSummonerFlagsEquippedParams{

		timeout: timeout,
	}
}

// NewGetLolBannersV1CurrentSummonerFlagsEquippedParamsWithContext creates a new GetLolBannersV1CurrentSummonerFlagsEquippedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolBannersV1CurrentSummonerFlagsEquippedParamsWithContext(ctx context.Context) *GetLolBannersV1CurrentSummonerFlagsEquippedParams {

	return &GetLolBannersV1CurrentSummonerFlagsEquippedParams{

		Context: ctx,
	}
}

// NewGetLolBannersV1CurrentSummonerFlagsEquippedParamsWithHTTPClient creates a new GetLolBannersV1CurrentSummonerFlagsEquippedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolBannersV1CurrentSummonerFlagsEquippedParamsWithHTTPClient(client *http.Client) *GetLolBannersV1CurrentSummonerFlagsEquippedParams {

	return &GetLolBannersV1CurrentSummonerFlagsEquippedParams{
		HTTPClient: client,
	}
}

/*GetLolBannersV1CurrentSummonerFlagsEquippedParams contains all the parameters to send to the API endpoint
for the get lol banners v1 current summoner flags equipped operation typically these are written to a http.Request
*/
type GetLolBannersV1CurrentSummonerFlagsEquippedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol banners v1 current summoner flags equipped params
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) WithTimeout(timeout time.Duration) *GetLolBannersV1CurrentSummonerFlagsEquippedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol banners v1 current summoner flags equipped params
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol banners v1 current summoner flags equipped params
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) WithContext(ctx context.Context) *GetLolBannersV1CurrentSummonerFlagsEquippedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol banners v1 current summoner flags equipped params
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol banners v1 current summoner flags equipped params
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) WithHTTPClient(client *http.Client) *GetLolBannersV1CurrentSummonerFlagsEquippedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol banners v1 current summoner flags equipped params
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolBannersV1CurrentSummonerFlagsEquippedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}