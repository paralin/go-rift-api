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

// NewGetLolChampSelectV1PickableChampionIdsParams creates a new GetLolChampSelectV1PickableChampionIdsParams object
// with the default values initialized.
func NewGetLolChampSelectV1PickableChampionIdsParams() *GetLolChampSelectV1PickableChampionIdsParams {

	return &GetLolChampSelectV1PickableChampionIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectV1PickableChampionIdsParamsWithTimeout creates a new GetLolChampSelectV1PickableChampionIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectV1PickableChampionIdsParamsWithTimeout(timeout time.Duration) *GetLolChampSelectV1PickableChampionIdsParams {

	return &GetLolChampSelectV1PickableChampionIdsParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectV1PickableChampionIdsParamsWithContext creates a new GetLolChampSelectV1PickableChampionIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectV1PickableChampionIdsParamsWithContext(ctx context.Context) *GetLolChampSelectV1PickableChampionIdsParams {

	return &GetLolChampSelectV1PickableChampionIdsParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectV1PickableChampionIdsParamsWithHTTPClient creates a new GetLolChampSelectV1PickableChampionIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectV1PickableChampionIdsParamsWithHTTPClient(client *http.Client) *GetLolChampSelectV1PickableChampionIdsParams {

	return &GetLolChampSelectV1PickableChampionIdsParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectV1PickableChampionIdsParams contains all the parameters to send to the API endpoint
for the get lol champ select v1 pickable champion ids operation typically these are written to a http.Request
*/
type GetLolChampSelectV1PickableChampionIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select v1 pickable champion ids params
func (o *GetLolChampSelectV1PickableChampionIdsParams) WithTimeout(timeout time.Duration) *GetLolChampSelectV1PickableChampionIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select v1 pickable champion ids params
func (o *GetLolChampSelectV1PickableChampionIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select v1 pickable champion ids params
func (o *GetLolChampSelectV1PickableChampionIdsParams) WithContext(ctx context.Context) *GetLolChampSelectV1PickableChampionIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select v1 pickable champion ids params
func (o *GetLolChampSelectV1PickableChampionIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select v1 pickable champion ids params
func (o *GetLolChampSelectV1PickableChampionIdsParams) WithHTTPClient(client *http.Client) *GetLolChampSelectV1PickableChampionIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select v1 pickable champion ids params
func (o *GetLolChampSelectV1PickableChampionIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectV1PickableChampionIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}