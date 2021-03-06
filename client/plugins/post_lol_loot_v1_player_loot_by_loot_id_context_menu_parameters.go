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

// NewPostLolLootV1PlayerLootByLootIDContextMenuParams creates a new PostLolLootV1PlayerLootByLootIDContextMenuParams object
// with the default values initialized.
func NewPostLolLootV1PlayerLootByLootIDContextMenuParams() *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	var ()
	return &PostLolLootV1PlayerLootByLootIDContextMenuParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLootV1PlayerLootByLootIDContextMenuParamsWithTimeout creates a new PostLolLootV1PlayerLootByLootIDContextMenuParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLootV1PlayerLootByLootIDContextMenuParamsWithTimeout(timeout time.Duration) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	var ()
	return &PostLolLootV1PlayerLootByLootIDContextMenuParams{

		timeout: timeout,
	}
}

// NewPostLolLootV1PlayerLootByLootIDContextMenuParamsWithContext creates a new PostLolLootV1PlayerLootByLootIDContextMenuParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLootV1PlayerLootByLootIDContextMenuParamsWithContext(ctx context.Context) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	var ()
	return &PostLolLootV1PlayerLootByLootIDContextMenuParams{

		Context: ctx,
	}
}

// NewPostLolLootV1PlayerLootByLootIDContextMenuParamsWithHTTPClient creates a new PostLolLootV1PlayerLootByLootIDContextMenuParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLootV1PlayerLootByLootIDContextMenuParamsWithHTTPClient(client *http.Client) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	var ()
	return &PostLolLootV1PlayerLootByLootIDContextMenuParams{
		HTTPClient: client,
	}
}

/*PostLolLootV1PlayerLootByLootIDContextMenuParams contains all the parameters to send to the API endpoint
for the post lol loot v1 player loot by loot Id context menu operation typically these are written to a http.Request
*/
type PostLolLootV1PlayerLootByLootIDContextMenuParams struct {

	/*LootID*/
	LootID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) WithTimeout(timeout time.Duration) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) WithContext(ctx context.Context) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) WithHTTPClient(client *http.Client) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLootID adds the lootID to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) WithLootID(lootID string) *PostLolLootV1PlayerLootByLootIDContextMenuParams {
	o.SetLootID(lootID)
	return o
}

// SetLootID adds the lootId to the post lol loot v1 player loot by loot Id context menu params
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) SetLootID(lootID string) {
	o.LootID = lootID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLootV1PlayerLootByLootIDContextMenuParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
