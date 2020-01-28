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

// NewGetLolChampSelectV1BannableChampionIdsParams creates a new GetLolChampSelectV1BannableChampionIdsParams object
// with the default values initialized.
func NewGetLolChampSelectV1BannableChampionIdsParams() *GetLolChampSelectV1BannableChampionIdsParams {

	return &GetLolChampSelectV1BannableChampionIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectV1BannableChampionIdsParamsWithTimeout creates a new GetLolChampSelectV1BannableChampionIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectV1BannableChampionIdsParamsWithTimeout(timeout time.Duration) *GetLolChampSelectV1BannableChampionIdsParams {

	return &GetLolChampSelectV1BannableChampionIdsParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectV1BannableChampionIdsParamsWithContext creates a new GetLolChampSelectV1BannableChampionIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectV1BannableChampionIdsParamsWithContext(ctx context.Context) *GetLolChampSelectV1BannableChampionIdsParams {

	return &GetLolChampSelectV1BannableChampionIdsParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectV1BannableChampionIdsParamsWithHTTPClient creates a new GetLolChampSelectV1BannableChampionIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectV1BannableChampionIdsParamsWithHTTPClient(client *http.Client) *GetLolChampSelectV1BannableChampionIdsParams {

	return &GetLolChampSelectV1BannableChampionIdsParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectV1BannableChampionIdsParams contains all the parameters to send to the API endpoint
for the get lol champ select v1 bannable champion ids operation typically these are written to a http.Request
*/
type GetLolChampSelectV1BannableChampionIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select v1 bannable champion ids params
func (o *GetLolChampSelectV1BannableChampionIdsParams) WithTimeout(timeout time.Duration) *GetLolChampSelectV1BannableChampionIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select v1 bannable champion ids params
func (o *GetLolChampSelectV1BannableChampionIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select v1 bannable champion ids params
func (o *GetLolChampSelectV1BannableChampionIdsParams) WithContext(ctx context.Context) *GetLolChampSelectV1BannableChampionIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select v1 bannable champion ids params
func (o *GetLolChampSelectV1BannableChampionIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select v1 bannable champion ids params
func (o *GetLolChampSelectV1BannableChampionIdsParams) WithHTTPClient(client *http.Client) *GetLolChampSelectV1BannableChampionIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select v1 bannable champion ids params
func (o *GetLolChampSelectV1BannableChampionIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectV1BannableChampionIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
