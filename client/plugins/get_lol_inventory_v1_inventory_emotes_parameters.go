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

// NewGetLolInventoryV1InventoryEmotesParams creates a new GetLolInventoryV1InventoryEmotesParams object
// with the default values initialized.
func NewGetLolInventoryV1InventoryEmotesParams() *GetLolInventoryV1InventoryEmotesParams {

	return &GetLolInventoryV1InventoryEmotesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolInventoryV1InventoryEmotesParamsWithTimeout creates a new GetLolInventoryV1InventoryEmotesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolInventoryV1InventoryEmotesParamsWithTimeout(timeout time.Duration) *GetLolInventoryV1InventoryEmotesParams {

	return &GetLolInventoryV1InventoryEmotesParams{

		timeout: timeout,
	}
}

// NewGetLolInventoryV1InventoryEmotesParamsWithContext creates a new GetLolInventoryV1InventoryEmotesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolInventoryV1InventoryEmotesParamsWithContext(ctx context.Context) *GetLolInventoryV1InventoryEmotesParams {

	return &GetLolInventoryV1InventoryEmotesParams{

		Context: ctx,
	}
}

// NewGetLolInventoryV1InventoryEmotesParamsWithHTTPClient creates a new GetLolInventoryV1InventoryEmotesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolInventoryV1InventoryEmotesParamsWithHTTPClient(client *http.Client) *GetLolInventoryV1InventoryEmotesParams {

	return &GetLolInventoryV1InventoryEmotesParams{
		HTTPClient: client,
	}
}

/*GetLolInventoryV1InventoryEmotesParams contains all the parameters to send to the API endpoint
for the get lol inventory v1 inventory emotes operation typically these are written to a http.Request
*/
type GetLolInventoryV1InventoryEmotesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol inventory v1 inventory emotes params
func (o *GetLolInventoryV1InventoryEmotesParams) WithTimeout(timeout time.Duration) *GetLolInventoryV1InventoryEmotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol inventory v1 inventory emotes params
func (o *GetLolInventoryV1InventoryEmotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol inventory v1 inventory emotes params
func (o *GetLolInventoryV1InventoryEmotesParams) WithContext(ctx context.Context) *GetLolInventoryV1InventoryEmotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol inventory v1 inventory emotes params
func (o *GetLolInventoryV1InventoryEmotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol inventory v1 inventory emotes params
func (o *GetLolInventoryV1InventoryEmotesParams) WithHTTPClient(client *http.Client) *GetLolInventoryV1InventoryEmotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol inventory v1 inventory emotes params
func (o *GetLolInventoryV1InventoryEmotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolInventoryV1InventoryEmotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
