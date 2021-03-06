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

// NewGetLolGameSettingsV1ReadyParams creates a new GetLolGameSettingsV1ReadyParams object
// with the default values initialized.
func NewGetLolGameSettingsV1ReadyParams() *GetLolGameSettingsV1ReadyParams {

	return &GetLolGameSettingsV1ReadyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolGameSettingsV1ReadyParamsWithTimeout creates a new GetLolGameSettingsV1ReadyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolGameSettingsV1ReadyParamsWithTimeout(timeout time.Duration) *GetLolGameSettingsV1ReadyParams {

	return &GetLolGameSettingsV1ReadyParams{

		timeout: timeout,
	}
}

// NewGetLolGameSettingsV1ReadyParamsWithContext creates a new GetLolGameSettingsV1ReadyParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolGameSettingsV1ReadyParamsWithContext(ctx context.Context) *GetLolGameSettingsV1ReadyParams {

	return &GetLolGameSettingsV1ReadyParams{

		Context: ctx,
	}
}

// NewGetLolGameSettingsV1ReadyParamsWithHTTPClient creates a new GetLolGameSettingsV1ReadyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolGameSettingsV1ReadyParamsWithHTTPClient(client *http.Client) *GetLolGameSettingsV1ReadyParams {

	return &GetLolGameSettingsV1ReadyParams{
		HTTPClient: client,
	}
}

/*GetLolGameSettingsV1ReadyParams contains all the parameters to send to the API endpoint
for the get lol game settings v1 ready operation typically these are written to a http.Request
*/
type GetLolGameSettingsV1ReadyParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol game settings v1 ready params
func (o *GetLolGameSettingsV1ReadyParams) WithTimeout(timeout time.Duration) *GetLolGameSettingsV1ReadyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol game settings v1 ready params
func (o *GetLolGameSettingsV1ReadyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol game settings v1 ready params
func (o *GetLolGameSettingsV1ReadyParams) WithContext(ctx context.Context) *GetLolGameSettingsV1ReadyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol game settings v1 ready params
func (o *GetLolGameSettingsV1ReadyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol game settings v1 ready params
func (o *GetLolGameSettingsV1ReadyParams) WithHTTPClient(client *http.Client) *GetLolGameSettingsV1ReadyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol game settings v1 ready params
func (o *GetLolGameSettingsV1ReadyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolGameSettingsV1ReadyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
