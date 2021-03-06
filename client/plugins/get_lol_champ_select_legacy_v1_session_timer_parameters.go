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

// NewGetLolChampSelectLegacyV1SessionTimerParams creates a new GetLolChampSelectLegacyV1SessionTimerParams object
// with the default values initialized.
func NewGetLolChampSelectLegacyV1SessionTimerParams() *GetLolChampSelectLegacyV1SessionTimerParams {

	return &GetLolChampSelectLegacyV1SessionTimerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolChampSelectLegacyV1SessionTimerParamsWithTimeout creates a new GetLolChampSelectLegacyV1SessionTimerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolChampSelectLegacyV1SessionTimerParamsWithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1SessionTimerParams {

	return &GetLolChampSelectLegacyV1SessionTimerParams{

		timeout: timeout,
	}
}

// NewGetLolChampSelectLegacyV1SessionTimerParamsWithContext creates a new GetLolChampSelectLegacyV1SessionTimerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolChampSelectLegacyV1SessionTimerParamsWithContext(ctx context.Context) *GetLolChampSelectLegacyV1SessionTimerParams {

	return &GetLolChampSelectLegacyV1SessionTimerParams{

		Context: ctx,
	}
}

// NewGetLolChampSelectLegacyV1SessionTimerParamsWithHTTPClient creates a new GetLolChampSelectLegacyV1SessionTimerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolChampSelectLegacyV1SessionTimerParamsWithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1SessionTimerParams {

	return &GetLolChampSelectLegacyV1SessionTimerParams{
		HTTPClient: client,
	}
}

/*GetLolChampSelectLegacyV1SessionTimerParams contains all the parameters to send to the API endpoint
for the get lol champ select legacy v1 session timer operation typically these are written to a http.Request
*/
type GetLolChampSelectLegacyV1SessionTimerParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol champ select legacy v1 session timer params
func (o *GetLolChampSelectLegacyV1SessionTimerParams) WithTimeout(timeout time.Duration) *GetLolChampSelectLegacyV1SessionTimerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol champ select legacy v1 session timer params
func (o *GetLolChampSelectLegacyV1SessionTimerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol champ select legacy v1 session timer params
func (o *GetLolChampSelectLegacyV1SessionTimerParams) WithContext(ctx context.Context) *GetLolChampSelectLegacyV1SessionTimerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol champ select legacy v1 session timer params
func (o *GetLolChampSelectLegacyV1SessionTimerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol champ select legacy v1 session timer params
func (o *GetLolChampSelectLegacyV1SessionTimerParams) WithHTTPClient(client *http.Client) *GetLolChampSelectLegacyV1SessionTimerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol champ select legacy v1 session timer params
func (o *GetLolChampSelectLegacyV1SessionTimerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolChampSelectLegacyV1SessionTimerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
