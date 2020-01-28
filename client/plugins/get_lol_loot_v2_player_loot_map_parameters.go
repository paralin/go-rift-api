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

// NewGetLolLootV2PlayerLootMapParams creates a new GetLolLootV2PlayerLootMapParams object
// with the default values initialized.
func NewGetLolLootV2PlayerLootMapParams() *GetLolLootV2PlayerLootMapParams {

	return &GetLolLootV2PlayerLootMapParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLootV2PlayerLootMapParamsWithTimeout creates a new GetLolLootV2PlayerLootMapParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLootV2PlayerLootMapParamsWithTimeout(timeout time.Duration) *GetLolLootV2PlayerLootMapParams {

	return &GetLolLootV2PlayerLootMapParams{

		timeout: timeout,
	}
}

// NewGetLolLootV2PlayerLootMapParamsWithContext creates a new GetLolLootV2PlayerLootMapParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLootV2PlayerLootMapParamsWithContext(ctx context.Context) *GetLolLootV2PlayerLootMapParams {

	return &GetLolLootV2PlayerLootMapParams{

		Context: ctx,
	}
}

// NewGetLolLootV2PlayerLootMapParamsWithHTTPClient creates a new GetLolLootV2PlayerLootMapParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLootV2PlayerLootMapParamsWithHTTPClient(client *http.Client) *GetLolLootV2PlayerLootMapParams {

	return &GetLolLootV2PlayerLootMapParams{
		HTTPClient: client,
	}
}

/*GetLolLootV2PlayerLootMapParams contains all the parameters to send to the API endpoint
for the get lol loot v2 player loot map operation typically these are written to a http.Request
*/
type GetLolLootV2PlayerLootMapParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol loot v2 player loot map params
func (o *GetLolLootV2PlayerLootMapParams) WithTimeout(timeout time.Duration) *GetLolLootV2PlayerLootMapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol loot v2 player loot map params
func (o *GetLolLootV2PlayerLootMapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol loot v2 player loot map params
func (o *GetLolLootV2PlayerLootMapParams) WithContext(ctx context.Context) *GetLolLootV2PlayerLootMapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol loot v2 player loot map params
func (o *GetLolLootV2PlayerLootMapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol loot v2 player loot map params
func (o *GetLolLootV2PlayerLootMapParams) WithHTTPClient(client *http.Client) *GetLolLootV2PlayerLootMapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol loot v2 player loot map params
func (o *GetLolLootV2PlayerLootMapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLootV2PlayerLootMapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}