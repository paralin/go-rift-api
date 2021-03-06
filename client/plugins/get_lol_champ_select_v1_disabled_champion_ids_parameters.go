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

// NewGetLolChampSelectV1DisabledChampionIdsParams creates a new GetLolChampSelectV1DisabledChampionIdsParams object
// with the default values initialized.
func NewGetLolChampSelectV1DisabledChampionIdsParams() *GetLolChampSelectV1DisabledChampionIdsParams {

	return &GetLolChampSelectV1DisabledChampionIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectV1DisabledChampionIdsParamsWithTimeout creates a new GetLolChampSelectV1DisabledChampionIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectV1DisabledChampionIdsParamsWithTimeout(timeout time.Duration) *GetLolChampSelectV1DisabledChampionIdsParams {

	return &GetLolChampSelectV1DisabledChampionIdsParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectV1DisabledChampionIdsParamsWithContext creates a new GetLolChampSelectV1DisabledChampionIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectV1DisabledChampionIdsParamsWithContext(ctx context.Context) *GetLolChampSelectV1DisabledChampionIdsParams {

	return &GetLolChampSelectV1DisabledChampionIdsParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectV1DisabledChampionIdsParamsWithHTTPClient creates a new GetLolChampSelectV1DisabledChampionIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectV1DisabledChampionIdsParamsWithHTTPClient(client *http.Client) *GetLolChampSelectV1DisabledChampionIdsParams {

	return &GetLolChampSelectV1DisabledChampionIdsParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectV1DisabledChampionIdsParams contains all the parameters to send to the API endpoint
for the get lol champ select v1 disabled champion ids operation typically these are written to a http.Request
*/
type GetLolChampSelectV1DisabledChampionIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select v1 disabled champion ids params
func (o *GetLolChampSelectV1DisabledChampionIdsParams) WithTimeout(timeout time.Duration) *GetLolChampSelectV1DisabledChampionIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select v1 disabled champion ids params
func (o *GetLolChampSelectV1DisabledChampionIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select v1 disabled champion ids params
func (o *GetLolChampSelectV1DisabledChampionIdsParams) WithContext(ctx context.Context) *GetLolChampSelectV1DisabledChampionIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select v1 disabled champion ids params
func (o *GetLolChampSelectV1DisabledChampionIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select v1 disabled champion ids params
func (o *GetLolChampSelectV1DisabledChampionIdsParams) WithHTTPClient(client *http.Client) *GetLolChampSelectV1DisabledChampionIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select v1 disabled champion ids params
func (o *GetLolChampSelectV1DisabledChampionIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectV1DisabledChampionIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
