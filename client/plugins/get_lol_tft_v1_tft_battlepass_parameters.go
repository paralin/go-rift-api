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

// NewGetLolTftV1TftBattlepassParams creates a new GetLolTftV1TftBattlepassParams object
// with the default values initialized.
func NewGetLolTftV1TftBattlepassParams() *GetLolTftV1TftBattlepassParams {

	return &GetLolTftV1TftBattlepassParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolTftV1TftBattlepassParamsWithTimeout creates a new GetLolTftV1TftBattlepassParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolTftV1TftBattlepassParamsWithTimeout(timeout time.Duration) *GetLolTftV1TftBattlepassParams {

	return &GetLolTftV1TftBattlepassParams{

		timeout: timeout,
	}
}

// NewGetLolTftV1TftBattlepassParamsWithContext creates a new GetLolTftV1TftBattlepassParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolTftV1TftBattlepassParamsWithContext(ctx context.Context) *GetLolTftV1TftBattlepassParams {

	return &GetLolTftV1TftBattlepassParams{

		Context: ctx,
	}
}

// NewGetLolTftV1TftBattlepassParamsWithHTTPClient creates a new GetLolTftV1TftBattlepassParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolTftV1TftBattlepassParamsWithHTTPClient(client *http.Client) *GetLolTftV1TftBattlepassParams {

	return &GetLolTftV1TftBattlepassParams{
		HTTPClient: client,
	}
}

/*GetLolTftV1TftBattlepassParams contains all the parameters to send to the API endpoint
for the get lol tft v1 tft battlepass operation typically these are written to a http.Request
*/
type GetLolTftV1TftBattlepassParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol tft v1 tft battlepass params
func (o *GetLolTftV1TftBattlepassParams) WithTimeout(timeout time.Duration) *GetLolTftV1TftBattlepassParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol tft v1 tft battlepass params
func (o *GetLolTftV1TftBattlepassParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol tft v1 tft battlepass params
func (o *GetLolTftV1TftBattlepassParams) WithContext(ctx context.Context) *GetLolTftV1TftBattlepassParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol tft v1 tft battlepass params
func (o *GetLolTftV1TftBattlepassParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol tft v1 tft battlepass params
func (o *GetLolTftV1TftBattlepassParams) WithHTTPClient(client *http.Client) *GetLolTftV1TftBattlepassParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol tft v1 tft battlepass params
func (o *GetLolTftV1TftBattlepassParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolTftV1TftBattlepassParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
