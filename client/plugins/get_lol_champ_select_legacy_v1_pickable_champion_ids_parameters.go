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

// NewGetLolChampSelectLegacyV1PickableChampionIdsParams creates a new GetLolChampSelectLegacyV1PickableChampionIdsParams object
// with the default values initialized.
func NewGetLolChampSelectLegacyV1PickableChampionIdsParams() *GetLolChampSelectLegacyV1PickableChampionIdsParams {

	return &GetLolChampSelectLegacyV1PickableChampionIdsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectLegacyV1PickableChampionIdsParamsWithTimeout creates a new GetLolChampSelectLegacyV1PickableChampionIdsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectLegacyV1PickableChampionIdsParamsWithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1PickableChampionIdsParams {

	return &GetLolChampSelectLegacyV1PickableChampionIdsParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectLegacyV1PickableChampionIdsParamsWithContext creates a new GetLolChampSelectLegacyV1PickableChampionIdsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectLegacyV1PickableChampionIdsParamsWithContext(ctx context.Context) *GetLolChampSelectLegacyV1PickableChampionIdsParams {

	return &GetLolChampSelectLegacyV1PickableChampionIdsParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectLegacyV1PickableChampionIdsParamsWithHTTPClient creates a new GetLolChampSelectLegacyV1PickableChampionIdsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectLegacyV1PickableChampionIdsParamsWithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1PickableChampionIdsParams {

	return &GetLolChampSelectLegacyV1PickableChampionIdsParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectLegacyV1PickableChampionIdsParams contains all the parameters to send to the API endpoint
for the get lol champ select legacy v1 pickable champion ids operation typically these are written to a http.Request
*/
type GetLolChampSelectLegacyV1PickableChampionIdsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select legacy v1 pickable champion ids params
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) WithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1PickableChampionIdsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select legacy v1 pickable champion ids params
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select legacy v1 pickable champion ids params
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) WithContext(ctx context.Context) *GetLolChampSelectLegacyV1PickableChampionIdsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select legacy v1 pickable champion ids params
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select legacy v1 pickable champion ids params
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) WithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1PickableChampionIdsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select legacy v1 pickable champion ids params
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectLegacyV1PickableChampionIdsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
