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

// NewGetLolLootV1PlayerLootByLootIDParams creates a new GetLolLootV1PlayerLootByLootIDParams object
// with the default values initialized.
func NewGetLolLootV1PlayerLootByLootIDParams() *GetLolLootV1PlayerLootByLootIDParams {
	var ()
	return &GetLolLootV1PlayerLootByLootIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLolLootV1PlayerLootByLootIDParamsWithTimeout creates a new GetLolLootV1PlayerLootByLootIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLolLootV1PlayerLootByLootIDParamsWithTimeout(timeout time.Duration) *GetLolLootV1PlayerLootByLootIDParams {
	var ()
	return &GetLolLootV1PlayerLootByLootIDParams{

		timeout: timeout,
	}
}

// NewGetLolLootV1PlayerLootByLootIDParamsWithContext creates a new GetLolLootV1PlayerLootByLootIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLolLootV1PlayerLootByLootIDParamsWithContext(ctx context.Context) *GetLolLootV1PlayerLootByLootIDParams {
	var ()
	return &GetLolLootV1PlayerLootByLootIDParams{

		Context: ctx,
	}
}

// NewGetLolLootV1PlayerLootByLootIDParamsWithHTTPClient creates a new GetLolLootV1PlayerLootByLootIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLolLootV1PlayerLootByLootIDParamsWithHTTPClient(client *http.Client) *GetLolLootV1PlayerLootByLootIDParams {
	var ()
	return &GetLolLootV1PlayerLootByLootIDParams{
		HTTPClient: client,
	}
}

/*GetLolLootV1PlayerLootByLootIDParams contains all the parameters to send to the API endpoint
for the get lol loot v1 player loot by loot Id operation typically these are written to a http.Request
*/
type GetLolLootV1PlayerLootByLootIDParams struct {

	/*LootID*/
	LootID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) WithTimeout(timeout time.Duration) *GetLolLootV1PlayerLootByLootIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) WithContext(ctx context.Context) *GetLolLootV1PlayerLootByLootIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) WithHTTPClient(client *http.Client) *GetLolLootV1PlayerLootByLootIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLootID adds the lootID to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) WithLootID(lootID string) *GetLolLootV1PlayerLootByLootIDParams {
	o.SetLootID(lootID)
	return o
}

// SetLootID adds the lootId to the get lol loot v1 player loot by loot Id params
func (o *GetLolLootV1PlayerLootByLootIDParams) SetLootID(lootID string) {
	o.LootID = lootID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLolLootV1PlayerLootByLootIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param lootId
	if err := r.SetPathParam("lootId", o.LootID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
