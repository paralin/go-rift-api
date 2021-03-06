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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostLolLootV1NewPlayerCheckDoneByNewValueParams creates a new PostLolLootV1NewPlayerCheckDoneByNewValueParams object
// with the default values initialized.
func NewPostLolLootV1NewPlayerCheckDoneByNewValueParams() *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	var ()
	return &PostLolLootV1NewPlayerCheckDoneByNewValueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLolLootV1NewPlayerCheckDoneByNewValueParamsWithTimeout creates a new PostLolLootV1NewPlayerCheckDoneByNewValueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLolLootV1NewPlayerCheckDoneByNewValueParamsWithTimeout(timeout time.Duration) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	var ()
	return &PostLolLootV1NewPlayerCheckDoneByNewValueParams{

		timeout: timeout,
	}
}

// NewPostLolLootV1NewPlayerCheckDoneByNewValueParamsWithContext creates a new PostLolLootV1NewPlayerCheckDoneByNewValueParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLolLootV1NewPlayerCheckDoneByNewValueParamsWithContext(ctx context.Context) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	var ()
	return &PostLolLootV1NewPlayerCheckDoneByNewValueParams{

		Context: ctx,
	}
}

// NewPostLolLootV1NewPlayerCheckDoneByNewValueParamsWithHTTPClient creates a new PostLolLootV1NewPlayerCheckDoneByNewValueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLolLootV1NewPlayerCheckDoneByNewValueParamsWithHTTPClient(client *http.Client) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	var ()
	return &PostLolLootV1NewPlayerCheckDoneByNewValueParams{
		HTTPClient: client,
	}
}

/*PostLolLootV1NewPlayerCheckDoneByNewValueParams contains all the parameters to send to the API endpoint
for the post lol loot v1 new player check done by new value operation typically these are written to a http.Request
*/
type PostLolLootV1NewPlayerCheckDoneByNewValueParams struct {

	/*NewValue*/
	NewValue bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) WithTimeout(timeout time.Duration) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) WithContext(ctx context.Context) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) WithHTTPClient(client *http.Client) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewValue adds the newValue to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) WithNewValue(newValue bool) *PostLolLootV1NewPlayerCheckDoneByNewValueParams {
	o.SetNewValue(newValue)
	return o
}

// SetNewValue adds the newValue to the post lol loot v1 new player check done by new value params
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) SetNewValue(newValue bool) {
	o.NewValue = newValue
}

// WriteToRequest writes these params to a swagger request
func (o *PostLolLootV1NewPlayerCheckDoneByNewValueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param newValue
	if err := r.SetPathParam("newValue", swag.FormatBool(o.NewValue)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
