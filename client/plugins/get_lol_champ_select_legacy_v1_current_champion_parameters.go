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

// NewGetLolChampSelectLegacyV1CurrentChampionParams creates a new GetLolChampSelectLegacyV1CurrentChampionParams object
// with the default values initialized.
func NewGetLolChampSelectLegacyV1CurrentChampionParams() *GetLolChampSelectLegacyV1CurrentChampionParams {

	return &GetLolChampSelectLegacyV1CurrentChampionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectLegacyV1CurrentChampionParamsWithTimeout creates a new GetLolChampSelectLegacyV1CurrentChampionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectLegacyV1CurrentChampionParamsWithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1CurrentChampionParams {

	return &GetLolChampSelectLegacyV1CurrentChampionParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectLegacyV1CurrentChampionParamsWithContext creates a new GetLolChampSelectLegacyV1CurrentChampionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectLegacyV1CurrentChampionParamsWithContext(ctx context.Context) *GetLolChampSelectLegacyV1CurrentChampionParams {

	return &GetLolChampSelectLegacyV1CurrentChampionParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectLegacyV1CurrentChampionParamsWithHTTPClient creates a new GetLolChampSelectLegacyV1CurrentChampionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectLegacyV1CurrentChampionParamsWithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1CurrentChampionParams {

	return &GetLolChampSelectLegacyV1CurrentChampionParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectLegacyV1CurrentChampionParams contains all the parameters to send to the API endpoint
for the get lol champ select legacy v1 current champion operation typically these are written to a http.Request
*/
type GetLolChampSelectLegacyV1CurrentChampionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select legacy v1 current champion params
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) WithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1CurrentChampionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select legacy v1 current champion params
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select legacy v1 current champion params
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) WithContext(ctx context.Context) *GetLolChampSelectLegacyV1CurrentChampionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select legacy v1 current champion params
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select legacy v1 current champion params
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) WithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1CurrentChampionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select legacy v1 current champion params
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectLegacyV1CurrentChampionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
