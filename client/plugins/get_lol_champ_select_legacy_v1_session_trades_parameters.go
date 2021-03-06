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

// NewGetLolChampSelectLegacyV1SessionTradesParams creates a new GetLolChampSelectLegacyV1SessionTradesParams object
// with the default values initialized.
func NewGetLolChampSelectLegacyV1SessionTradesParams() *GetLolChampSelectLegacyV1SessionTradesParams {

	return &GetLolChampSelectLegacyV1SessionTradesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectLegacyV1SessionTradesParamsWithTimeout creates a new GetLolChampSelectLegacyV1SessionTradesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectLegacyV1SessionTradesParamsWithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1SessionTradesParams {

	return &GetLolChampSelectLegacyV1SessionTradesParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectLegacyV1SessionTradesParamsWithContext creates a new GetLolChampSelectLegacyV1SessionTradesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectLegacyV1SessionTradesParamsWithContext(ctx context.Context) *GetLolChampSelectLegacyV1SessionTradesParams {

	return &GetLolChampSelectLegacyV1SessionTradesParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectLegacyV1SessionTradesParamsWithHTTPClient creates a new GetLolChampSelectLegacyV1SessionTradesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectLegacyV1SessionTradesParamsWithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1SessionTradesParams {

	return &GetLolChampSelectLegacyV1SessionTradesParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectLegacyV1SessionTradesParams contains all the parameters to send to the API endpoint
for the get lol champ select legacy v1 session trades operation typically these are written to a http.Request
*/
type GetLolChampSelectLegacyV1SessionTradesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select legacy v1 session trades params
func (o *GetLolChampSelectLegacyV1SessionTradesParams) WithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1SessionTradesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select legacy v1 session trades params
func (o *GetLolChampSelectLegacyV1SessionTradesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select legacy v1 session trades params
func (o *GetLolChampSelectLegacyV1SessionTradesParams) WithContext(ctx context.Context) *GetLolChampSelectLegacyV1SessionTradesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select legacy v1 session trades params
func (o *GetLolChampSelectLegacyV1SessionTradesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select legacy v1 session trades params
func (o *GetLolChampSelectLegacyV1SessionTradesParams) WithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1SessionTradesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select legacy v1 session trades params
func (o *GetLolChampSelectLegacyV1SessionTradesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectLegacyV1SessionTradesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
