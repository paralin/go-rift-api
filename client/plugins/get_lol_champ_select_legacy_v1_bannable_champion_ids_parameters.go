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

// NewGetLolChampSelectLegacyV1BannableChampionIdsParams creates a new GetLolChampSelectLegacyV1BannableChampionIdsParams object
// with the default values initialized.
func NewGetLolChampSelectLegacyV1BannableChampionIdsParams() *GetLolChampSelectLegacyV1BannableChampionIdsParams {

	return &GetLolChampSelectLegacyV1BannableChampionIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectLegacyV1BannableChampionIdsParamsWithTimeout creates a new GetLolChampSelectLegacyV1BannableChampionIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectLegacyV1BannableChampionIdsParamsWithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1BannableChampionIdsParams {

	return &GetLolChampSelectLegacyV1BannableChampionIdsParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectLegacyV1BannableChampionIdsParamsWithContext creates a new GetLolChampSelectLegacyV1BannableChampionIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectLegacyV1BannableChampionIdsParamsWithContext(ctx context.Context) *GetLolChampSelectLegacyV1BannableChampionIdsParams {

	return &GetLolChampSelectLegacyV1BannableChampionIdsParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectLegacyV1BannableChampionIdsParamsWithHTTPClient creates a new GetLolChampSelectLegacyV1BannableChampionIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectLegacyV1BannableChampionIdsParamsWithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1BannableChampionIdsParams {

	return &GetLolChampSelectLegacyV1BannableChampionIdsParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectLegacyV1BannableChampionIdsParams contains all the parameters to send to the API endpoint
for the get lol champ select legacy v1 bannable champion ids operation typically these are written to a http.Request
*/
type GetLolChampSelectLegacyV1BannableChampionIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select legacy v1 bannable champion ids params
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) WithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1BannableChampionIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select legacy v1 bannable champion ids params
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select legacy v1 bannable champion ids params
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) WithContext(ctx context.Context) *GetLolChampSelectLegacyV1BannableChampionIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select legacy v1 bannable champion ids params
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select legacy v1 bannable champion ids params
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) WithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1BannableChampionIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select legacy v1 bannable champion ids params
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectLegacyV1BannableChampionIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}